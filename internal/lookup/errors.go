package lookup

import (
	"errors"
	"strings"
)

var (
	// ErrPermissionDenied được trả về khi người dùng không đủ quyền truy cập (vd: cần Admin/root).
	ErrPermissionDenied = errors.New("không đủ quyền truy cập (thử chạy công cụ với quyền Administrator hoặc sudo)")

	// ErrProcessNotFound được trả về khi không tìm thấy tiến trình nào.
	ErrProcessNotFound = errors.New("không tìm thấy tiến trình")

	// ErrKillFailed được trả về khi thao tác kill process gặp lỗi.
	ErrKillFailed = errors.New("không thể buộc dừng tiến trình")
)

// CheckPermissionError là một hàm trợ giúp để phân tích thông báo lỗi thô (từ stderr hoặc os) 
// và map nó về ErrPermissionDenied nếu thấy dấu hiệu "permission" hoặc "access is denied".
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
