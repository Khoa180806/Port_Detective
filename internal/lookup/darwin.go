//go:build darwin
// +build darwin

package lookup

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
	"github.com/Khoa180806/Port_Detective/internal/process"
)

func init() {
	OSStrategy = &DarwinStrategy{}
}

// DarwinStrategy implements PortLookupStrategy for macOS systems.
type DarwinStrategy struct{}

func (s *DarwinStrategy) FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	cmd := exec.Command("lsof", "-i", fmt.Sprintf(":%d", port), "-P", "-n")
	output, err := cmd.Output()
	if err != nil {
		// lsof returns exit code 1 when no matching process is found
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return nil, nil
		}

		err = CheckPermissionError(err, "")
		if errors.Is(err, ErrPermissionDenied) {
			return nil, err
		}
		return nil, fmt.Errorf("%s", i18n.Tr("error.lsof_failed", err))
	}

	return parseDarwinLsofOutput(string(output), port), nil
}

func (s *DarwinStrategy) KillProcess(pid int) error {
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

// parseDarwinLsofOutput parses lsof output on macOS into ProcessInfo slices.
func parseDarwinLsofOutput(output string, port int) []process.ProcessInfo {
	var results []process.ProcessInfo
	seen := make(map[int]bool)
	lines := strings.Split(output, "\n")

	for i, line := range lines {
		if i == 0 || strings.TrimSpace(line) == "" {
			continue // Skip header line (COMMAND PID USER ...)
		}

		fields := strings.Fields(line)
		if len(fields) < 8 {
			continue
		}

		// Field 1 is PID
		pidStr := fields[1]
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			continue
		}

		if seen[pid] {
			continue // Avoid duplicates
		}

		// Field 0 is Command/Name
		command := fields[0]

		// Protocol information is in NAME or TYPE column (e.g. TCP *:8080 (LISTEN) or UDP *:*)
		protocol := "tcp"
		lowerLine := strings.ToLower(line)
		if strings.Contains(lowerLine, "udp") {
			protocol = "udp"
		}

		results = append(results, process.ProcessInfo{
			PID:      pid,
			Name:     command,
			Command:  command, // lsof does not provide full command-line arguments
			Port:     port,
			Protocol: protocol,
		})
		seen[pid] = true
	}

	return results
}
