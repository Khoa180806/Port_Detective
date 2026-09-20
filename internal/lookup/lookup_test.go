package lookup_test

import (
	"errors"
	"testing"

	"github.com/Khoa180806/Port_Detective/internal/lookup"
	"github.com/Khoa180806/Port_Detective/internal/process"
)

// mockStrategy là một đối tượng giả lập (mock) cho PortLookupStrategy
type mockStrategy struct{}

func (m *mockStrategy) FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	if port == 8080 {
		return []process.ProcessInfo{
			{PID: 1234, Name: "mock.exe", Port: 8080, Protocol: "tcp"},
		}, nil
	}
	return nil, errors.New("port không tìm thấy")
}

func (m *mockStrategy) KillProcess(pid int) error {
	if pid == 1234 {
		return nil
	}
	return errors.New("pid không tồn tại")
}

func TestDispatch_FindProcessByPort(t *testing.T) {
	// Giữ lại state cũ để hoàn trả sau khi test
	oldStrategy := lookup.OSStrategy
	defer func() { lookup.OSStrategy = oldStrategy }()

	// Gán strategy thành mock
	lookup.OSStrategy = &mockStrategy{}

	// Test case thành công
	procs, err := lookup.FindProcessByPort(8080)
	if err != nil {
		t.Fatalf("Lỗi không mong đợi: %v", err)
	}
	if len(procs) != 1 || procs[0].PID != 1234 {
		t.Errorf("Kết quả trả về không khớp mong đợi: %+v", procs)
	}

	// Test case thất bại
	_, err = lookup.FindProcessByPort(9999)
	if err == nil {
		t.Errorf("Mong đợi một lỗi nhưng lại không có lỗi")
	}
}

func TestDispatch_KillProcess(t *testing.T) {
	// Giữ lại state cũ
	oldStrategy := lookup.OSStrategy
	defer func() { lookup.OSStrategy = oldStrategy }()

	lookup.OSStrategy = &mockStrategy{}

	// Test case thành công
	err := lookup.KillProcess(1234)
	if err != nil {
		t.Fatalf("Lỗi không mong đợi khi kill pid hợp lệ: %v", err)
	}

	// Test case thất bại
	err = lookup.KillProcess(9999)
	if err == nil {
		t.Errorf("Mong đợi một lỗi khi kill pid không hợp lệ")
	}
}

func TestDispatch_NilStrategy(t *testing.T) {
	oldStrategy := lookup.OSStrategy
	defer func() { lookup.OSStrategy = oldStrategy }()

	// Set strategy = nil để test lỗi ErrNotSupported
	lookup.OSStrategy = nil

	_, err := lookup.FindProcessByPort(8080)
	if !errors.Is(err, lookup.ErrNotSupported) {
		t.Errorf("Mong đợi ErrNotSupported, nhận được: %v", err)
	}

	err = lookup.KillProcess(1234)
	if !errors.Is(err, lookup.ErrNotSupported) {
		t.Errorf("Mong đợi ErrNotSupported, nhận được: %v", err)
	}
}
