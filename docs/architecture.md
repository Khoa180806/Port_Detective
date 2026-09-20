# 📐 Kiến trúc Kỹ thuật — Port Detective

Tài liệu này cung cấp cái nhìn chi tiết về thiết kế kiến trúc, mô hình dữ liệu và các thành phần cốt lõi của **Port Detective**.

---

## 1. Mục tiêu thiết kế

1. **Hiệu năng & Tốc độ:** Khởi động cực nhanh, parse thông tin tức thời và trả về kết quả cho developer.
2. **Độc lập, không phụ thuộc runtime:** Đóng gói thành 1 file binary duy nhất, không yêu cầu Python, Node.js hay bất kỳ dependency phức tạp nào.
3. **Đa nền tảng (Cross-platform) sạch sẽ:** Sử dụng cơ chế Go Build Tags (Strategy Pattern) để tách biệt mã nguồn của từng hệ điều hành mà không làm phình kích thước binary.
4. **Dễ bảo trì & Mở rộng:** Phân tách rõ rệt giữa tầng giao diện người dùng (CLI), tầng tương tác hệ thống (OS Lookup), và tầng định dạng đầu ra (Output Formatters).

---

## 2. Sơ đồ tương tác các thành phần

```
┌────────────────────────────────────────────────────────┐
│                        main.go                         │
└───────────────────────────┬────────────────────────────┘
                            │ Khởi tạo
                            ▼
┌────────────────────────────────────────────────────────┐
│                         cmd/                           │
│   (root.go, check.go, kill.go, scan.go - Cobra CLI)    │
└──────────────┬──────────────────────────┬──────────────┘
               │                          │
   Yêu cầu tra cứu / kill         Truyền ProcessInfo
               │                          │
               ▼                          ▼
┌──────────────────────────┐   ┌──────────────────────────┐
│     internal/lookup/     │   │     internal/output/     │
│ (PortLookupStrategy)     │   │                          │
│                          │   │  - text.go (bảng, màu)   │
│ ├── windows.go           │   │  - json.go (chuẩn JSON)  │
│ ├── linux.go             │   └──────────────────────────┘
│ └── darwin.go            │                  ▲
└──────────────┬───────────┘                  │
               │ Trả về                       │
               ▼                              │
┌─────────────────────────────────────────────┴──────────┐
│                   internal/process/                    │
│             (ProcessInfo Struct Data Model)            │
└────────────────────────────────────────────────────────┘
```

---

## 3. Các thành phần chính (Key Components)

### 3.1. Data Model (`internal/process`)
`ProcessInfo` là cấu trúc dữ liệu trung tâm làm cầu nối giữa các module:

```go
type ProcessInfo struct {
    PID      int    `json:"pid"`
    Name     string `json:"name"`
    Command  string `json:"command"`
    Port     int    `json:"port"`
    Protocol string `json:"protocol"` // "tcp" | "udp"
}
```

### 3.2. Strategy Pattern qua Go Build Tags (`internal/lookup`)
Thay vì kiểm tra điều kiện lúc chạy (runtime check) như `if runtime.GOOS == "windows"`, dự án sử dụng chỉ dẫn biên dịch của Go (Go Build Tags):

- `windows.go` chứa `//go:build windows`:
  - Thực thi `netstat -ano` để bắt PID theo số port.
  - Sau đó gọi `tasklist /FI "PID eq <pid>"` hoặc WMI để lấy tên chương trình và lệnh thực thi.
  - Sử dụng lệnh `taskkill /F /PID <pid>` khi thực hiện lệnh `kill`.
- `linux.go` chứa `//go:build linux`:
  - Ưu tiên gọi lệnh `lsof -i :<port> -sTCP:LISTEN -P -n`.
  - Fallback: Đọc trực tiếp từ `/proc/net/tcp` và `/proc/<pid>/cmdline` trong môi trường tối giản không cài sẵn `lsof`.
  - Sử dụng syscall `syscall.Kill` (SIGTERM/SIGKILL) cho lệnh `kill`.
- `darwin.go` chứa `//go:build darwin`:
  - Tận dụng `lsof -i :<port> -P -n` trên macOS.
  - Sử dụng syscall POSIX tương thích.

Tất cả các file nền tảng đều hiện thực chung các hàm public trong package `lookup`:
- `FindProcessByPort(port int) ([]process.ProcessInfo, error)`
- `KillProcess(pid int) error`

### 3.3. Tầng định dạng đầu ra (`internal/output`)
Cung cấp hai chế độ hiển thị linh hoạt:
1. **Human-readable (Text):** 
   - Định dạng bảng căn lề đẹp mắt, làm nổi bật PID và tên tiến trình bằng màu sắc (`github.com/fatih/color`).
   - Tự động nhận diện thiết bị đầu ra (`isatty`) để tắt màu khi người dùng pipe hoặc redirect sang file.
2. **Machine-readable (JSON):**
   - Định dạng chuẩn JSON, hỗ trợ cờ `--json` ở mọi subcommand.
   - Thích hợp dùng với `jq`, Python hoặc bash script trong CI/CD.

### 3.4. Tầng điều khiển dòng lệnh (`cmd/`)
Xây dựng trên thư viện chuẩn công nghiệp **Cobra**:
- Quản lý các lệnh con: `check`, `kill`, `scan`.
- Tự động tạo menu trợ giúp (`--help`) và kiểm tra tính hợp lệ của tham số (port từ 1 đến 65535).
- Đảm bảo mã thoát (exit code) chuẩn POSIX.

---

## 4. Xử lý các trường hợp đặc thù (Edge Cases)

| Tình huống | Hướng xử lý |
|---|---|
| **Port không tồn tại tiến trình nào** | Trả về thông báo rõ ràng, thoát với mã lỗi `1` (Port available). |
| **Nhiều tiến trình cùng listen 1 port** | Trả về danh sách tất cả các tiến trình (hỗ trợ các trường hợp `SO_REUSEPORT`). Khi gọi `kill`, người dùng được liệt kê đầy đủ để xác nhận. |
| **Thiếu quyền hạn (Permission Denied)** | Bắt lỗi và thông báo người dùng chạy terminal dưới quyền Administrator (Windows) hoặc dùng `sudo` (Linux/macOS). Thoát mã lỗi `2`. |
| **Đầu vào port không hợp lệ** | Validate port trước khi gọi OS command, từ chối số âm, chữ cái, hoặc port > 65535. |
