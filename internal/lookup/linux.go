//go:build linux
// +build linux

package lookup

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

func init() {
	OSStrategy = &LinuxStrategy{}
}

// LinuxStrategy thực thi PortLookupStrategy cho môi trường Linux
type LinuxStrategy struct{}

func (s *LinuxStrategy) FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	// Ưu tiên 1: Sử dụng lsof
	procs, err := findViaLsof(port)
	if err == nil && len(procs) > 0 {
		return procs, nil
	}

	// Ưu tiên 2: Fallback parse trực tiếp /proc (không cần cài thêm tool)
	procs, err = findViaProcFS(port)
	if err != nil {
		return nil, fmt.Errorf("cả lsof và procfs đều thất bại: %v", err)
	}

	return procs, nil
}

func (s *LinuxStrategy) KillProcess(pid int) error {
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("không thể kill PID %d: %s", pid, stderr.String())
	}
	return nil
}

func findViaLsof(port int) ([]process.ProcessInfo, error) {
	cmd := exec.Command("lsof", "-i", fmt.Sprintf(":%d", port), "-P", "-n")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseLsofOutput(string(output), port), nil
}

func parseLsofOutput(output string, port int) []process.ProcessInfo {
	var results []process.ProcessInfo
	seen := make(map[int]bool)
	lines := strings.Split(output, "\n")
	
	for i, line := range lines {
		if i == 0 || strings.TrimSpace(line) == "" {
			continue // Bỏ qua header
		}
		
		fields := strings.Fields(line)
		if len(fields) < 9 {
			continue
		}

		pidStr := fields[1]
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			continue
		}

		if seen[pid] {
			continue
		}

		command := fields[0]
		protocol := "tcp"
		if strings.Contains(strings.ToLower(fields[7]), "udp") {
			protocol = "udp"
		}

		results = append(results, process.ProcessInfo{
			PID:      pid,
			Name:     command,
			Command:  command, // lsof cắt bớt command line, fallback sau nếu cần
			Port:     port,
			Protocol: protocol,
		})
		seen[pid] = true
	}
	
	return results
}

func findViaProcFS(port int) ([]process.ProcessInfo, error) {
	// Lấy danh sách inode đang listen trên port này (TCP v4 và v6)
	inodes := getInodesForPort(port, "/proc/net/tcp")
	inodes = append(inodes, getInodesForPort(port, "/proc/net/tcp6")...)
	inodes = append(inodes, getInodesForPort(port, "/proc/net/udp")...)
	inodes = append(inodes, getInodesForPort(port, "/proc/net/udp6")...)

	if len(inodes) == 0 {
		return nil, nil
	}

	inodeSet := make(map[string]bool)
	for _, in := range inodes {
		inodeSet[in] = true
	}

	var results []process.ProcessInfo

	// Duyệt qua tất cả các tiến trình trong /proc
	dirs, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		pid, err := strconv.Atoi(d.Name())
		if err != nil {
			continue // Không phải thư mục PID
		}

		fdDir := filepath.Join("/proc", d.Name(), "fd")
		fds, err := os.ReadDir(fdDir)
		if err != nil {
			continue // Permission denied (vd: process của user khác)
		}

		for _, fd := range fds {
			link, err := os.Readlink(filepath.Join(fdDir, fd.Name()))
			if err != nil {
				continue
			}

			// Kiểm tra xem link có dạng socket:[inode] không
			if strings.HasPrefix(link, "socket:[") && strings.HasSuffix(link, "]") {
				inode := link[8 : len(link)-1]
				if inodeSet[inode] {
					// Tìm thấy!
					name, cmdline := getProcDetails(pid)
					results = append(results, process.ProcessInfo{
						PID:      pid,
						Name:     name,
						Command:  cmdline,
						Port:     port,
						Protocol: "unknown", // Có thể tra cứu lại file proc để biết chính xác, tạm để unknown hoặc gộp với logic inodes
					})
					// Xóa inode để tránh add duplicate nếu process có nhiều FD vào cùng socket
					delete(inodeSet, inode)
				}
			}
		}
	}

	return results, nil
}

// getInodesForPort đọc file /proc/net/[protocol] để tìm các inode liên kết với port cụ thể.
func getInodesForPort(port int, procFile string) []string {
	content, err := os.ReadFile(procFile)
	if err != nil {
		return nil
	}
	return parseInodesFromProcNet(string(content), port)
}

func parseInodesFromProcNet(content string, port int) []string {
	var inodes []string
	hexPort := fmt.Sprintf("%04X", port)
	lines := strings.Split(content, "\n")
	
	for i, line := range lines {
		if i == 0 || strings.TrimSpace(line) == "" {
			continue
		}
		
		fields := strings.Fields(line)
		if len(fields) < 10 {
			continue
		}

		// local_address là trường thứ 2, vd: 00000000:1F90
		localAddr := fields[1]
		parts := strings.Split(localAddr, ":")
		if len(parts) != 2 {
			continue
		}

		if parts[1] == hexPort {
			// Inode là trường thứ 10
			inodes = append(inodes, fields[9])
		}
	}
	return inodes
}

// getProcDetails lấy Name và Command Line từ /proc/<pid>/comm và /proc/<pid>/cmdline.
func getProcDetails(pid int) (name, cmdline string) {
	// Đọc comm
	if b, err := os.ReadFile(fmt.Sprintf("/proc/%d/comm", pid)); err == nil {
		name = strings.TrimSpace(string(b))
	}

	// Đọc cmdline (null-terminated)
	if b, err := os.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid)); err == nil {
		cmdline = strings.ReplaceAll(string(b), "\x00", " ")
		cmdline = strings.TrimSpace(cmdline)
	}

	if cmdline == "" {
		cmdline = name
	}
	
	return
}
