package i18n

func init() {
	RegisterMessages("vi", viMessages)
}

var viMessages = MessageMap{
	// Root command
	"root.short": "Port Detective - CLI tool giúp tra cứu và quản lý các cổng mạng (ports)",
	"root.long": `Port Detective là một công cụ dòng lệnh (CLI) cross-platform được viết bằng Go.
Công cụ giúp bạn nhanh chóng tìm ra tiến trình (process) nào đang chiếm giữ một cổng mạng (port) cụ thể
và hỗ trợ buộc dừng (kill) tiến trình đó nếu cần thiết.

Hỗ trợ các hệ điều hành: Windows, macOS, Linux.`,

	// Flags
	"flag.json":     "Hiển thị kết quả dưới định dạng JSON",
	"flag.force":    "Buộc dừng tiến trình không cần xác nhận (Bỏ qua y/N)",
	"flag.dry_run":  "Chỉ hiển thị các tiến trình sẽ bị kill mà không thực thi",
	"flag.lang":     "Chỉ định ngôn ngữ hiển thị (en, vi)",

	// Check command
	"check.short":                "Kiểm tra tiến trình đang chạy trên một port",
	"check.long":                 "Kiểm tra và hiển thị thông tin về các tiến trình (processes) đang chiếm giữ hoặc lắng nghe trên cổng mạng (port) được chỉ định.",
	"check.error.invalid_port":   "Lỗi: Port phải là một số nguyên từ 1 đến 65535.",
	"check.error.permission":     "Lỗi: Không đủ quyền truy cập để lấy thông tin chi tiết.",
	"check.error.permission_hint": "Gợi ý: Hãy thử chạy lại lệnh dưới quyền Administrator (hoặc sudo trên Linux/macOS).",
	"check.error.lookup":         "Lỗi khi tra cứu: %v",
	"check.not_found":            "Không tìm thấy tiến trình nào đang chạy trên port %d.",

	// Kill command
	"kill.short":                 "Buộc dừng các tiến trình đang chiếm giữ port",
	"kill.long": `Kiểm tra xem cổng mạng (port) nào đang bị chiếm giữ và buộc dừng (kill)
tất cả các tiến trình (processes) đang sử dụng cổng mạng đó.

Mặc định lệnh này sẽ hiển thị thông tin tiến trình và hỏi bạn (y/N) trước khi kill.
Bạn có thể dùng cờ --force (-f) để bỏ qua câu hỏi xác nhận.`,
	"kill.error.invalid_port":    "Lỗi: Port phải là một số nguyên từ 1 đến 65535.",
	"kill.error.lookup":          "Lỗi khi tra cứu: %v",
	"kill.not_found":             "Không tìm thấy tiến trình nào đang chạy trên port %d.",
	"kill.warning":               "CẢNH BÁO: Phát hiện các tiến trình sau đang chiếm dụng port:",
	"kill.dry_run":               "[DRY RUN] Sẽ không thực hiện lệnh kill nào.",
	"kill.confirm":               "Bạn có chắc chắn muốn KILL tất cả các tiến trình trên không? [y/N]: ",
	"kill.cancelled":             "Đã hủy thao tác.",
	"kill.killing":               "Đang kill PID %d (%s)... ",
	"kill.failed.permission":     "THẤT BẠI: Không đủ quyền (Permission Denied). Cần quyền Admin/sudo.",
	"kill.failed":                "THẤT BẠI: %v",
	"kill.success":               "THÀNH CÔNG",

	// Scan command
	"scan.short":                 "Quét một dải cổng mạng (port range) để tìm các tiến trình đang hoạt động",
	"scan.long": `Lệnh scan giúp bạn kiểm tra hàng loạt cổng mạng trong một dải cụ thể.
Ví dụ: port-detective scan 3000-3010
Dải cổng hợp lệ là từ 1 đến 65535. Để bảo vệ hệ thống, giới hạn tối đa 5000 cổng mỗi lần quét.`,
	"scan.error.format":          "Lỗi: Sai định dạng dải port. Vui lòng sử dụng định dạng <start>-<end> (vd: 3000-3010).",
	"scan.error.range":           "Lỗi: Dải port không hợp lệ. (1 <= start <= end <= 65535).",
	"scan.error.too_large":       "Lỗi: Khoảng quét quá lớn (tối đa 5000 port mỗi lần) để tránh quá tải hệ thống.",
	"scan.scanning":              "Đang quét dải port từ %d đến %d...",
	"scan.no_results":            "Tuyệt vời! Không có tiến trình nào đang chiếm dụng dải port này.",
	"scan.permission_warn":       "\n[Chú ý] Quá trình quét bị từ chối truy cập (Permission Denied) ở một số port.",
	"scan.permission_hint":       "Gợi ý: Hãy chạy lại công cụ dưới quyền Administrator/sudo để có danh sách đầy đủ nhất.",

	// Output format
	"output.occupied_by":         "Port %s is occupied by:\n",
	"output.found_processes":     "Found %d processes:\n",
	"output.no_processes":        "No processes found.",
	"output.no_info":             "Không có thông tin tiến trình khả dụng.",

	// Internal errors
	"error.permission_denied":    "không đủ quyền truy cập (thử chạy công cụ với quyền Administrator hoặc sudo)",
	"error.process_not_found":    "không tìm thấy tiến trình",
	"error.kill_failed":          "không thể buộc dừng tiến trình",
	"error.not_supported":        "không được hỗ trợ trên hệ điều hành này",
	"error.netstat_failed":       "lỗi khi chạy netstat: %v",
	"error.lsof_failed":          "lỗi khi chạy lsof: %v",
	"error.kill_pid_failed":      "không thể kill PID %d: %s",
	"error.lsof_procfs_failed":   "cả lsof và procfs đều thất bại: %v",
}
