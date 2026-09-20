# Port Detective — Project Plan

> Tài liệu này là bản kế hoạch kỹ thuật đầy đủ cho dự án **Port Detective**, dùng để AI coding agent (hoặc developer khác) có thể đọc và triển khai độc lập mà không cần thêm ngữ cảnh.

## 1. Tổng quan

**Tên dự án:** Port Detective
**Loại:** Developer tool — Command Line Interface (CLI)
**Ngôn ngữ / stack:** Go (Golang)
**Bài toán:** Tìm process nào đang chiếm (occupy) một port cụ thể trên máy local.
**Mục đích:** Giúp developer debug nhanh lỗi "port already in use" khi chạy service ở local, thay vì phải tra thủ công qua `lsof` / `netstat` rồi tự tra PID.

**Người dùng mục tiêu:** Developer (backend, fullstack) làm việc trên Windows / macOS / Linux, cần công cụ nhanh gọn chạy trong terminal.

## 2. Mục tiêu & phạm vi

### Trong phạm vi (in scope)
- Tra cứu process đang giữ 1 port cụ thể, hiển thị PID + tên process + command line.
- Kill process đang giữ port đó (có xác nhận trước khi kill).
- Hỗ trợ cả 3 hệ điều hành: Linux, macOS, Windows.
- Output dạng text đẹp cho người đọc, và dạng JSON cho máy đọc (script/CI).
- Phân phối dưới dạng 1 binary độc lập, không cần cài runtime.

### Ngoài phạm vi (out of scope, giai đoạn đầu)
- Không quản lý port ở tầng network/firewall (không phải security tool).
- Không hỗ trợ máy remote (chỉ chạy local machine).
- Không có GUI — chỉ CLI.
- Không cần đăng ký / license / telemetry.

## 3. Kiến trúc kỹ thuật

### 3.1 Nguyên lý hoạt động
Trên mỗi OS, việc tra cứu process theo port dựa trên lệnh hệ thống có sẵn:
- **Linux:** `lsof -i :<port>` hoặc parse `/proc/net/tcp`
- **macOS:** `lsof -i :<port>` (tương tự Linux)
- **Windows:** `netstat -ano | findstr :<port>` để lấy PID, sau đó `tasklist /FI "PID eq <pid>"` để lấy tên process

Port Detective sẽ gọi các lệnh này qua `os/exec`, parse output, và chuẩn hóa về 1 struct dữ liệu chung.

### 3.2 Design pattern: Strategy Pattern qua Go build tags
Thay vì runtime `if/else` theo OS, dùng Go build constraints để compiler tự chọn đúng file implementation theo OS lúc build:

```go
//go:build linux
package lookup

func FindProcessByPort(port int) (*ProcessInfo, error) { ... }
```
Mỗi OS có 1 file (`linux.go`, `darwin.go`, `windows.go`) cùng implement chung 1 interface `PortLookupStrategy`.

### 3.3 Cấu trúc thư mục

```
port-detective/
├── go.mod
├── go.sum
├── main.go                    // entry point, khởi tạo Cobra root command
├── cmd/
│   ├── root.go                 // root command definition
│   ├── check.go                 // `port-detective check <port>`
│   ├── kill.go                  // `port-detective kill <port>`
│   └── scan.go                  // `port-detective scan <range>` (M5)
├── internal/
│   ├── lookup/
│   │   ├── lookup.go             // interface PortLookupStrategy + dispatch
│   │   ├── linux.go               // //go:build linux
│   │   ├── darwin.go              // //go:build darwin
│   │   └── windows.go             // //go:build windows
│   ├── process/
│   │   └── process.go             // struct ProcessInfo{PID, Name, Cmd, Port, Protocol}
│   └── output/
│       ├── text.go                // format output dạng text đẹp (bảng, màu)
│       └── json.go                // format output dạng JSON
├── README.md
└── .goreleaser.yml              // (M5) config để build binary đa nền tảng
```

### 3.4 Thư viện đề xuất
| Nhu cầu | Thư viện | Lý do |
|---|---|---|
| CLI framework (subcommand, flags, help) | `github.com/spf13/cobra` | Chuẩn công nghiệp, dùng bởi `kubectl`, `docker`, `hugo` |
| Output có màu trong terminal | `github.com/fatih/color` | Đơn giản, nhẹ |
| Build & release binary đa OS | `goreleaser` | Tự động cross-compile Windows/macOS/Linux + tạo GitHub Release |

## 4. Data model

```go
type ProcessInfo struct {
    PID      int    `json:"pid"`
    Name     string `json:"name"`
    Command  string `json:"command"`
    Port     int    `json:"port"`
    Protocol string `json:"protocol"` // "tcp" | "udp"
}
```

## 5. CLI Interface (contract)

```
port-detective check <port> [--json]
    → Tìm và hiển thị process đang chiếm <port>.
    → Exit code 0 nếu tìm thấy, 1 nếu port đang free, 2 nếu lỗi hệ thống.

port-detective kill <port> [--force] [--json]
    → Kill process đang chiếm <port>.
    → Mặc định hỏi xác nhận (y/n) trừ khi có --force.

port-detective scan <start>-<end> [--json]     (M5, optional)
    → Quét dải port, liệt kê port nào đang bận.
```

Output text mẫu (`check`):
```
Port 8080 is occupied by:
  PID:      12345
  Process:  java
  Command:  java -jar app.jar
  Protocol: tcp
```

Output JSON mẫu:
```json
{ "port": 8080, "pid": 12345, "name": "java", "command": "java -jar app.jar", "protocol": "tcp" }
```

## 6. Roadmap theo milestone

| Milestone | Nội dung | Definition of Done |
|---|---|---|
| **M1** | MVP `check` command, 1 OS duy nhất (chọn theo OS dev đang dùng) | Chạy `port-detective check <port>` ra đúng PID + tên process, có test thủ công trên ít nhất 1 port đang mở thật |
| **M2** | `kill` command + xác nhận | Kill đúng process theo PID, có prompt xác nhận, có flag `--force` để bỏ qua prompt |
| **M3** | Cross-platform | Cả 3 file `linux.go`, `darwin.go`, `windows.go` implement cùng interface, build thành công trên cả 3 OS (dùng `GOOS=... go build` để test cross-compile) |
| **M4** | Chuyên nghiệp hóa CLI | Chuyển sang dùng Cobra, có `--help` tự động, có `--json` flag hoạt động đúng cho cả `check` và `kill` |
| **M5** | Nâng cao (optional) | `scan` command hoạt động, có sẵn binary release cho 3 OS qua GitHub Releases (dùng goreleaser) |

## 7. Tiêu chí hoàn thành dự án (Acceptance Criteria)
- [ ] `check` hoạt động đúng trên cả 3 OS
- [ ] `kill` hoạt động đúng, không kill nhầm process khi có nhiều process share port (edge case cần xử lý)
- [ ] Output JSON valid, parse được bằng `jq` hoặc script khác
- [ ] Có README hướng dẫn cài đặt (`go install` hoặc tải binary) và usage
- [ ] Xử lý lỗi rõ ràng: port không tồn tại process nào, port input không hợp lệ, không đủ quyền (permission denied khi kill)

## 8. Rủi ro / điểm cần lưu ý khi triển khai
- Trên Linux, một số lệnh (`lsof`) có thể không có sẵn mặc định trên mọi distro → cần có fallback đọc `/proc/net/tcp`.
- Kill process trên Windows cần quyền Administrator trong một số trường hợp → cần thông báo lỗi rõ ràng thay vì crash.
- Một port có thể có nhiều process cùng listen (SO_REUSEPORT) → cần quyết định UX: hiển thị tất cả hay chỉ 1.

## 9. Ghi chú cho AI agent triển khai
- Ưu tiên viết code Go idiomatic: xử lý error tường minh (`if err != nil`), không dùng panic cho lỗi runtime thông thường.
- Mỗi package trong `internal/` nên có unit test tối thiểu cho phần parse output (vì đây là phần dễ vỡ nhất khi format lệnh OS thay đổi).
- Không cần implement M5 nếu mục tiêu là MVP nhanh cho portfolio — M1-M4 đã đủ để demo và đưa vào CV.
