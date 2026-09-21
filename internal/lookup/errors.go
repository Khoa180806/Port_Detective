package lookup

import (
	"errors"
	"strings"
)

var (
	// ErrPermissionDenied indicates insufficient permissions to access process details or terminate a process.
	ErrPermissionDenied = errors.New("permission denied (try running with Administrator or sudo privileges)")

	// ErrProcessNotFound indicates no matching process was found for the specified criteria.
	ErrProcessNotFound = errors.New("process not found")

	// ErrKillFailed indicates process termination failed.
	ErrKillFailed = errors.New("failed to kill process")
)

// CheckPermissionError analyzes error output and maps to ErrPermissionDenied if permission issues are detected.
func CheckPermissionError(err error, stderr string) error {
	if err == nil {
		return nil
	}

	errStr := strings.ToLower(err.Error() + " " + stderr)
	if strings.Contains(errStr, "access is denied") ||
		strings.Contains(errStr, "permission denied") ||
		strings.Contains(errStr, "operation not permitted") {
		return ErrPermissionDenied
	}

	return err
}
