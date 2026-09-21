package lookup

import (
	"errors"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
	"github.com/Khoa180806/Port_Detective/internal/process"
)

// ErrNotSupported is returned when an operation is not supported on the current operating system.
var ErrNotSupported = errors.New("operation not supported")

// ErrNotSupportedMessage returns the localized error message for unsupported operations.
func ErrNotSupportedMessage() string {
	return i18n.T("error.not_supported")
}

// PortLookupStrategy defines the contract for interacting with network processes on different operating systems.
type PortLookupStrategy interface {
	// FindProcessByPort returns a list of processes occupying the specified port.
	FindProcessByPort(port int) ([]process.ProcessInfo, error)

	// KillProcess terminates a process by its Process ID (PID).
	KillProcess(pid int) error
}

// OSStrategy holds the active OS-specific PortLookupStrategy implementation.
var OSStrategy PortLookupStrategy

// FindProcessByPort forwards the lookup request to the active OS strategy.
func FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	if OSStrategy == nil {
		return nil, ErrNotSupported
	}
	return OSStrategy.FindProcessByPort(port)
}

// KillProcess forwards the termination request to the active OS strategy.
func KillProcess(pid int) error {
	if OSStrategy == nil {
		return ErrNotSupported
	}
	return OSStrategy.KillProcess(pid)
}
