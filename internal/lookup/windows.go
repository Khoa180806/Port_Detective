//go:build windows
// +build windows

package lookup

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

func init() {
	OSStrategy = &WindowsStrategy{}
}

// WindowsStrategy thực thi PortLookupStrategy cho môi trường Windows
type WindowsStrategy struct{}

func (s *WindowsStrategy) FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	cmd := exec.Command("netstat", "-ano")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("lỗi khi chạy netstat: %v", err)
	}

	pids, protocols := parseNetstatForPort(string(output), port)
	if len(pids) == 0 {
		return nil, nil // Port is free
	}

	var results []process.ProcessInfo
	for i, pid := range pids {
		name := getProcessName(pid)
		if name == "" {
			continue // Process might have exited
		}

		results = append(results, process.ProcessInfo{
			PID:      pid,
			Name:     name,
			Command:  name, // Mặc định dùng name làm command nếu không lấy được full command line
			Port:     port,
			Protocol: protocols[i],
		})
	}

	return results, nil
}

func (s *WindowsStrategy) KillProcess(pid int) error {
	cmd := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("không thể kill PID %d: %s", pid, stderr.String())
	}
	return nil
}

// parseNetstatForPort phân tích output của netstat -ano để tìm PID tương ứng với port.
// Trả về danh sách PIDs và danh sách Protocols tương ứng (để xử lý trường hợp nhiều process / protocol).
func parseNetstatForPort(output string, targetPort int) ([]int, []string) {
	var pids []int
	var protocols []string
	seen := make(map[string]bool)

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || (!strings.HasPrefix(line, "TCP") && !strings.HasPrefix(line, "UDP")) {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}

		protocol := strings.ToLower(fields[0])
		localAddr := fields[1]

		// Local address có dạng IP:PORT (vd: 0.0.0.0:8080 hoặc [::1]:8080)
		portSuffix := fmt.Sprintf(":%d", targetPort)
		if !strings.HasSuffix(localAddr, portSuffix) {
			continue
		}

		// Với TCP, PID ở cột cuối (cột 4, do index 4). Với UDP (không có state), PID ở cột 3 (index 3).
		pidStr := fields[len(fields)-1]
		pid, err := strconv.Atoi(pidStr)
		if err != nil || pid == 0 {
			continue
		}

		// Deduplicate
		key := fmt.Sprintf("%d-%s", pid, protocol)
		if !seen[key] {
			seen[key] = true
			pids = append(pids, pid)
			protocols = append(protocols, protocol)
		}
	}
	return pids, protocols
}

// getProcessName lấy tên tiến trình từ PID bằng lệnh tasklist.
func getProcessName(pid int) string {
	cmd := exec.Command("tasklist", "/FO", "CSV", "/NH", "/FI", fmt.Sprintf("PID eq %d", pid))
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	
	outStr := strings.TrimSpace(string(output))
	if outStr == "" || strings.HasPrefix(outStr, "INFO:") {
		return "" // Không tìm thấy
	}

	// Output CSV có dạng: "svchost.exe","1552","Services","0","48,092 K"
	r := csv.NewReader(strings.NewReader(outStr))
	record, err := r.Read()
	if err == nil && len(record) > 0 {
		return record[0]
	}

	return ""
}
