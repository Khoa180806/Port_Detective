package lookup

import (
	"errors"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

// ErrNotSupported được trả về khi thao tác không được hỗ trợ trên hệ điều hành hiện tại.
var ErrNotSupported = errors.New("không được hỗ trợ trên hệ điều hành này")

// PortLookupStrategy định nghĩa giao ước chung (contract) cho việc tương tác
// với các tiến trình mạng. Mỗi hệ điều hành (Windows, Linux, macOS) sẽ
// cung cấp một implementation riêng cho interface này.
type PortLookupStrategy interface {
	// FindProcessByPort trả về danh sách các tiến trình đang chiếm giữ một port cụ thể.
	FindProcessByPort(port int) ([]process.ProcessInfo, error)

	// KillProcess buộc dừng một tiến trình dựa vào Process ID (PID).
	KillProcess(pid int) error
}

// OSStrategy lưu trữ implementation hiện tại của hệ điều hành đang chạy.
// Biến này sẽ được gán trong hàm init() của các file OS cụ thể (windows.go, linux.go, ...).
var OSStrategy PortLookupStrategy

// FindProcessByPort là hàm tiện ích (dispatch) chuyển tiếp lệnh gọi đến strategy của OS hiện tại.
func FindProcessByPort(port int) ([]process.ProcessInfo, error) {
	if OSStrategy == nil {
		return nil, ErrNotSupported
	}
	return OSStrategy.FindProcessByPort(port)
}

// KillProcess là hàm tiện ích (dispatch) chuyển tiếp lệnh gọi đến strategy của OS hiện tại.
func KillProcess(pid int) error {
	if OSStrategy == nil {
		return ErrNotSupported
	}
	return OSStrategy.KillProcess(pid)
}
