package lookup_test

import (
	"errors"
	"testing"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
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
	return nil, errors.New(lookup.ErrProcessNotFoundMessage())
}

func (m *mockStrategy) KillProcess(pid int) error {
	if pid == 1234 {
		return nil
	}
	return errors.New(lookup.ErrKillFailedMessage())
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
	// English tests
	_ = i18n.SetLang("en")
	if msg := lookup.ErrPermissionDeniedMessage(); msg != "permission denied (try running with Administrator or sudo privileges)" {
		t.Errorf("unexpected EN ErrPermissionDeniedMessage: %s", msg)
	}
	if msg := lookup.ErrProcessNotFoundMessage(); msg != "process not found" {
		t.Errorf("unexpected EN ErrProcessNotFoundMessage: %s", msg)
	}
	if msg := lookup.ErrKillFailedMessage(); msg != "failed to kill process" {
		t.Errorf("unexpected EN ErrKillFailedMessage: %s", msg)
	}
	if msg := lookup.ErrNotSupportedMessage(); msg != "operation not supported on this operating system" {
		t.Errorf("unexpected EN ErrNotSupportedMessage: %s", msg)
	}

	// Vietnamese tests
	_ = i18n.SetLang("vi")
	if msg := lookup.ErrPermissionDeniedMessage(); msg != "không đủ quyền truy cập (thử chạy công cụ với quyền Administrator hoặc sudo)" {
		t.Errorf("unexpected VI ErrPermissionDeniedMessage: %s", msg)
	}
	if msg := lookup.ErrProcessNotFoundMessage(); msg != "không tìm thấy process" {
		t.Errorf("unexpected VI ErrProcessNotFoundMessage: %s", msg)
	}
	if msg := lookup.ErrKillFailedMessage(); msg != "không thể kill process" {
		t.Errorf("unexpected VI ErrKillFailedMessage: %s", msg)
	}
	if msg := lookup.ErrNotSupportedMessage(); msg != "không được hỗ trợ trên hệ điều hành này" {
		t.Errorf("unexpected VI ErrNotSupportedMessage: %s", msg)
	}

	// Reset to English
	_ = i18n.SetLang("en")
}
