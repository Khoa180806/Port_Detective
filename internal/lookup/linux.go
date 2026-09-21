//go:build linux
// +build linux

package lookup

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
	"github.com/Khoa180806/Port_Detective/internal/process"
)

func init() {
	OSStrategy = &LinuxStrategy{}
}

// LinuxStrategy implements PortLookupStrategy for Linux systems.
type LinuxStrategy struct{}

func (s *LinuxStrategy) FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	// Priority 1: Use lsof
	procs, err := findViaLsof(port)
	if err == nil && len(procs) > 0 {
		return procs, nil
	}

	// Priority 2: Fallback to parsing /proc directly (zero external tool dependency)
	procs, err = findViaProcFS(port)
	if err != nil {
		err = CheckPermissionError(err, "")
		if errors.Is(err, ErrPermissionDenied) {
			return nil, err
		}
		return nil, fmt.Errorf("%s", i18n.Tr("error.lsof_procfs_failed", err))
	}

	return procs, nil
}

func (s *LinuxStrategy) KillProcess(pid int) error {
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		err = CheckPermissionError(err, stderr.String())
		if errors.Is(err, ErrPermissionDenied) {
			return err
		}
		return fmt.Errorf("%s", i18n.Tr("error.kill_pid_failed", pid, stderr.String()))
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
			continue // Skip header line
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
			Command:  command, // lsof truncates command line
			Port:     port,
			Protocol: protocol,
		})
		seen[pid] = true
	}

	return results
}

func findViaProcFS(port int) ([]process.ProcessInfo, error) {
	// Retrieve listening socket inodes for target port (TCP v4/v6, UDP v4/v6)
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

	// Scan all process directories under /proc
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
			continue // Not a numeric PID directory
		}

		fdDir := filepath.Join("/proc", d.Name(), "fd")
		fds, err := os.ReadDir(fdDir)
		if err != nil {
			continue // Permission denied or process exited
		}

		for _, fd := range fds {
			link, err := os.Readlink(filepath.Join(fdDir, fd.Name()))
			if err != nil {
				continue
			}

			// Check if file descriptor points to socket:[inode]
			if strings.HasPrefix(link, "socket:[") && strings.HasSuffix(link, "]") {
				inode := link[8 : len(link)-1]
				if inodeSet[inode] {
					// Match found
					name, cmdline := getProcDetails(pid)
					results = append(results, process.ProcessInfo{
						PID:      pid,
						Name:     name,
						Command:  cmdline,
						Port:     port,
						Protocol: "unknown",
					})
					// Delete inode to prevent duplicates if multiple FDs reference the same socket
					delete(inodeSet, inode)
				}
			}
		}
	}

	return results, nil
}

// getInodesForPort parses /proc/net/[protocol] to find socket inodes bound to target port.
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

		// local_address is column 2, e.g.: 00000000:1F90
		localAddr := fields[1]
		parts := strings.Split(localAddr, ":")
		if len(parts) != 2 {
			continue
		}

		if parts[1] == hexPort {
			// Inode is column 10 (index 9)
			inodes = append(inodes, fields[9])
		}
	}
	return inodes
}

// getProcDetails reads process name and command line from /proc/<pid>/comm and /proc/<pid>/cmdline.
func getProcDetails(pid int) (name, cmdline string) {
	// Read comm
	if b, err := os.ReadFile(fmt.Sprintf("/proc/%d/comm", pid)); err == nil {
		name = strings.TrimSpace(string(b))
	}

	// Read cmdline (null-delimited)
	if b, err := os.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid)); err == nil {
		cmdline = strings.ReplaceAll(string(b), "\x00", " ")
		cmdline = strings.TrimSpace(cmdline)
	}

	if cmdline == "" {
		cmdline = name
	}

	return
}
