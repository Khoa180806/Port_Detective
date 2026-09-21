//go:build windows
// +build windows

package lookup

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
	"github.com/Khoa180806/Port_Detective/internal/process"
)

func init() {
	OSStrategy = &WindowsStrategy{}
}

// WindowsStrategy implements PortLookupStrategy for Windows systems.
type WindowsStrategy struct{}

func (s *WindowsStrategy) FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	cmd := exec.Command("netstat", "-ano")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("%s", i18n.Tr("error.netstat_failed", err))
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
			Command:  name, // Default to process name if full command line cannot be retrieved
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
		err = CheckPermissionError(err, stderr.String())
		if errors.Is(err, ErrPermissionDenied) {
			return err
		}
		return fmt.Errorf("%s", i18n.Tr("error.kill_pid_failed", pid, stderr.String()))
	}
	return nil
}

// parseNetstatForPort parses netstat -ano output to find PIDs and protocols corresponding to targetPort.
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

		// Local address format: IP:PORT (e.g., 0.0.0.0:8080 or [::1]:8080)
		portSuffix := fmt.Sprintf(":%d", targetPort)
		if !strings.HasSuffix(localAddr, portSuffix) {
			continue
		}

		// For TCP, PID is in the last column (index 4). For UDP (no state column), PID is in column index 3.
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

// getProcessName retrieves process name from PID using tasklist.
func getProcessName(pid int) string {
	cmd := exec.Command("tasklist", "/FO", "CSV", "/NH", "/FI", fmt.Sprintf("PID eq %d", pid))
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	outStr := strings.TrimSpace(string(output))
	if outStr == "" || strings.HasPrefix(outStr, "INFO:") {
		return "" // Not found
	}

	// CSV format: "svchost.exe","1552","Services","0","48,092 K"
	r := csv.NewReader(strings.NewReader(outStr))
	record, err := r.Read()
	if err == nil && len(record) > 0 {
		return record[0]
	}

	return ""
}
