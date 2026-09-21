package lookup_test

import (
	"errors"
	"testing"

	"github.com/Khoa180806/Port_Detective/internal/lookup"
	"github.com/Khoa180806/Port_Detective/internal/process"
)

// mockStrategy is a mock implementation of PortLookupStrategy for testing.
type mockStrategy struct{}

func (m *mockStrategy) FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	if port == 8080 {
		return []process.ProcessInfo{
			{PID: 1234, Name: "mock.exe", Port: 8080, Protocol: "tcp"},
		}, nil
	}
	return nil, errors.New("port not found")
}

func (m *mockStrategy) KillProcess(pid int) error {
	if pid == 1234 {
		return nil
	}
	return errors.New("pid not found")
}

func TestDispatch_FindProcessByPort(t *testing.T) {
	// Preserve old strategy and restore after test completion
	oldStrategy := lookup.OSStrategy
	defer func() { lookup.OSStrategy = oldStrategy }()

	lookup.OSStrategy = &mockStrategy{}

	// Successful lookup
	procs, err := lookup.FindProcessByPort(8080)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(procs) != 1 || procs[0].PID != 1234 {
		t.Errorf("unexpected process result: %+v", procs)
	}

	// Failed lookup
	_, err = lookup.FindProcessByPort(9999)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDispatch_KillProcess(t *testing.T) {
	oldStrategy := lookup.OSStrategy
	defer func() { lookup.OSStrategy = oldStrategy }()

	lookup.OSStrategy = &mockStrategy{}

	// Successful kill
	err := lookup.KillProcess(1234)
	if err != nil {
		t.Fatalf("unexpected error killing valid PID: %v", err)
	}

	// Failed kill
	err = lookup.KillProcess(9999)
	if err == nil {
		t.Errorf("expected error killing invalid PID, got nil")
	}
}

func TestDispatch_NilStrategy(t *testing.T) {
	oldStrategy := lookup.OSStrategy
	defer func() { lookup.OSStrategy = oldStrategy }()

	// Set strategy to nil to verify ErrNotSupported
	lookup.OSStrategy = nil

	_, err := lookup.FindProcessByPort(8080)
	if !errors.Is(err, lookup.ErrNotSupported) {
		t.Errorf("expected ErrNotSupported, got: %v", err)
	}

	err = lookup.KillProcess(1234)
	if !errors.Is(err, lookup.ErrNotSupported) {
		t.Errorf("expected ErrNotSupported, got: %v", err)
	}
}

func TestLocalizedErrorMessages(t *testing.T) {
	if msg := lookup.ErrPermissionDeniedMessage(); msg == "" {
		t.Errorf("expected non-empty ErrPermissionDeniedMessage")
	}
	if msg := lookup.ErrProcessNotFoundMessage(); msg == "" {
		t.Errorf("expected non-empty ErrProcessNotFoundMessage")
	}
	if msg := lookup.ErrKillFailedMessage(); msg == "" {
		t.Errorf("expected non-empty ErrKillFailedMessage")
	}
	if msg := lookup.ErrNotSupportedMessage(); msg == "" {
		t.Errorf("expected non-empty ErrNotSupportedMessage")
	}
}
