package lookup

import (
	"errors"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
)

var (
	// ErrPermissionDenied indicates insufficient permissions to access process details or terminate a process.
	ErrPermissionDenied = errors.New("permission denied")

	// ErrProcessNotFound indicates no matching process was found for the specified criteria.
	ErrProcessNotFound = errors.New("process not found")

	// ErrKillFailed indicates process termination failed.
	ErrKillFailed = errors.New("failed to kill process")
)

// ErrPermissionDeniedMessage returns the localized error message for permission denied.
func ErrPermissionDeniedMessage() string {
	return i18n.T("error.permission_denied")
}

// ErrProcessNotFoundMessage returns the localized error message for process not found.
func ErrProcessNotFoundMessage() string {
	return i18n.T("error.process_not_found")
}

// ErrKillFailedMessage returns the localized error message for kill failed.
func ErrKillFailedMessage() string {
	return i18n.T("error.kill_failed")
}

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
