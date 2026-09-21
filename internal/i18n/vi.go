package i18n

func init() {
	RegisterMessages("vi", viMessages)
}

var viMessages = MessageMap{
	// Root command
	"root.short": "Port Detective - CLI tool giúp tra cứu và quản lý port",
	"root.long": `Port Detective là công cụ dòng lệnh (CLI) cross-platform được viết bằng Go.
Công cụ giúp bạn nhanh chóng tìm ra process nào đang chiếm giữ port cụ thể
và hỗ trợ dừng (kill) process đó nếu cần thiết.

Hỗ trợ các hệ điều hành: Windows, macOS, Linux.`,

	// Flags
	"flag.json":     "Hiển thị kết quả dưới định dạng JSON",
	"flag.force":    "Kill process không cần xác nhận (bỏ qua y/N)",
	"flag.dry_run":  "Chỉ hiển thị các process sẽ bị kill mà không thực thi",
	"flag.lang":     "Chỉ định ngôn ngữ hiển thị (en, vi)",

	// Check command
	"check.short":                "Kiểm tra process đang chạy trên một port",
	"check.long":                 "Kiểm tra và hiển thị thông tin về các process đang chiếm giữ hoặc lắng nghe trên port được chỉ định.",
	"check.error.invalid_port":   "Lỗi: Port phải là số nguyên từ 1 đến 65535.",
	"check.error.permission":     "Lỗi: Không đủ quyền truy cập để lấy thông tin chi tiết.",
	"check.error.permission_hint": "Gợi ý: Hãy thử chạy lại lệnh với quyền Administrator (hoặc sudo trên Linux/macOS).",
	"check.error.lookup":         "Lỗi khi tra cứu port: %v",
	"check.not_found":            "Không tìm thấy process nào đang chạy trên port %d.",

	// Kill command
	"kill.short":                 "Dừng (kill) các process đang chiếm giữ port",
	"kill.long": `Kiểm tra xem port nào đang bị chiếm giữ và dừng (kill)
tất cả các process đang sử dụng port đó.

Mặc định lệnh này sẽ hiển thị thông tin process và hỏi xác nhận [y/N] trước khi kill.
Dùng cờ --force (-f) để bỏ qua bước xác nhận.`,
	"kill.error.invalid_port":    "Lỗi: Port phải là số nguyên từ 1 đến 65535.",
	"kill.error.lookup":          "Lỗi khi tra cứu port: %v",
	"kill.not_found":             "Không tìm thấy process nào đang chạy trên port %d.",
	"kill.warning":               "CẢNH BÁO: Phát hiện các process sau đang chiếm dụng port:",
	"kill.dry_run":               "[DRY RUN] Sẽ không dừng bất kỳ process nào.",
	"kill.confirm":               "Bạn có chắc chắn muốn KILL tất cả các process trên không? [y/N]: ",
	"kill.cancelled":             "Đã hủy thao tác.",
	"kill.killing":               "Đang kill PID %d (%s)... ",
	"kill.failed.permission":     "THẤT BẠI: Không đủ quyền (Permission Denied). Cần quyền Admin/sudo.",
	"kill.failed":                "THẤT BẠI: %v",
	"kill.success":               "THÀNH CÔNG",

	// Scan command
	"scan.short":                 "Quét một dải port để tìm các process đang hoạt động",
	"scan.long": `Lệnh scan giúp bạn kiểm tra hàng loạt port trong một dải cụ thể.
Ví dụ: port-detective scan 3000-3010
Dải port hợp lệ là từ 1 đến 65535. Để bảo vệ hệ thống, giới hạn tối đa 5000 port mỗi lần quét.`,
	"scan.error.format":          "Lỗi: Sai định dạng dải port. Vui lòng dùng định dạng <start>-<end> (vd: 3000-3010).",
	"scan.error.range":           "Lỗi: Dải port không hợp lệ. (1 <= start <= end <= 65535).",
	"scan.error.too_large":       "Lỗi: Khoảng quét quá lớn (tối đa 5000 port mỗi lần) để tránh quá tải hệ thống.",
	"scan.scanning":              "Đang quét dải port từ %d đến %d...",
	"scan.no_results":            "Không có process nào đang chiếm dụng dải port này.",
	"scan.permission_warn":       "\n[Chú ý] Quá trình quét bị từ chối truy cập (Permission Denied) ở một số port.",
	"scan.permission_hint":       "Gợi ý: Hãy chạy lại công cụ với quyền Administrator/sudo để có danh sách đầy đủ nhất.",

	// Output format
	"output.occupied_by":         "Port %s đang bị chiếm bởi:\n",
	"output.found_processes":     "Tìm thấy %d process:\n",
	"output.no_processes":        "Không tìm thấy process nào.",
	"output.no_info":             "Không có thông tin tiến trình.",

	// Internal errors
	"error.permission_denied":    "không đủ quyền truy cập (thử chạy công cụ với quyền Administrator hoặc sudo)",
	"error.process_not_found":    "không tìm thấy process",
	"error.kill_failed":          "không thể kill process",
	"error.not_supported":        "không được hỗ trợ trên hệ điều hành này",
	"error.netstat_failed":       "lỗi khi chạy netstat: %v",
	"error.lsof_failed":          "lỗi khi chạy lsof: %v",
	"error.kill_pid_failed":      "không thể kill PID %d: %s",
	"error.lsof_procfs_failed":   "cả lsof và procfs đều thất bại: %v",
}
