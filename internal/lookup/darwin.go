//go:build darwin
// +build darwin

package lookup

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

func init() {
	OSStrategy = &DarwinStrategy{}
}

// DarwinStrategy thực thi PortLookupStrategy cho môi trường macOS
type DarwinStrategy struct{}

func (s *DarwinStrategy) FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	cmd := exec.Command("lsof", "-i", fmt.Sprintf(":%d", port), "-P", "-n")
	output, err := cmd.Output()
	if err != nil {
		// lsof trả về exit code 1 nếu không tìm thấy tiến trình nào,
		// ta cần bắt trường hợp này để trả về mảng rỗng thay vì báo lỗi hệ thống.
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return nil, nil
		}
		return nil, fmt.Errorf("lỗi khi chạy lsof: %v", err)
	}

	return parseDarwinLsofOutput(string(output), port), nil
}

func (s *DarwinStrategy) KillProcess(pid int) error {
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("không thể kill PID %d: %s", pid, stderr.String())
	}
	return nil
}

// parseDarwinLsofOutput phân tích cú pháp kết quả trả về từ lệnh lsof trên macOS
func parseDarwinLsofOutput(output string, port int) []process.ProcessInfo {
	var results []process.ProcessInfo
	seen := make(map[int]bool)
	lines := strings.Split(output, "\n")
	
	for i, line := range lines {
		if i == 0 || strings.TrimSpace(line) == "" {
			continue // Bỏ qua dòng header (COMMAND PID USER ...)
		}
		
		fields := strings.Fields(line)
		if len(fields) < 8 {
			continue
		}

		// Field 1 là PID
		pidStr := fields[1]
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			continue
		}

		if seen[pid] {
			continue // Tránh add duplicate process
		}

		// Field 0 là Command/Name
		command := fields[0]
		
		// Protocol thường nằm ở cột NAME hoặc TYPE.
		// VD macOS: TCP *:8080 (LISTEN) hoặc UDP *:*
		protocol := "tcp"
		lowerLine := strings.ToLower(line)
		if strings.Contains(lowerLine, "udp") {
			protocol = "udp"
		}

		results = append(results, process.ProcessInfo{
			PID:      pid,
			Name:     command,
			Command:  command, // lsof không cung cấp full command line parameters
			Port:     port,
			Protocol: protocol,
		})
		seen[pid] = true
	}
	
	return results
}
