# Port Detective — Implementation Plan

> Tài liệu này mô tả kế hoạch triển khai chi tiết cho dự án Port Detective.
> Bất kỳ AI agent hoặc developer nào cũng có thể đọc file này cùng `todo.md` để biết tiến độ và bắt đầu làm việc.

## Tổng quan

Xây dựng CLI tool viết bằng Go giúp developer tra cứu và kill process đang chiếm port trên máy local. Hỗ trợ Windows, macOS, Linux. Output dạng text đẹp (có màu) hoặc JSON.

## Quyết định kiến trúc

- **Strategy Pattern qua Go Build Tags**: Mỗi OS có file riêng (`linux.go`, `darwin.go`, `windows.go`) cùng implement chung interface. Compiler tự chọn đúng file lúc build — không cần runtime `if/else`.
- **Cobra CLI Framework**: Chuẩn công nghiệp cho Go CLI. Hỗ trợ subcommand, flags, auto-generated help.
- **Internal packages**: Tất cả logic trong `internal/` — 3 package: `lookup`, `process`, `output`.
- **Incremental delivery**: Build từ đáy lên — data model → OS lookup → output format → CLI commands → polish.

## Cấu trúc thư mục mục tiêu

```
port-detective/
├── go.mod
├── go.sum
├── main.go                     // entry point
├── cmd/
│   ├── root.go                 // root command (Cobra)
│   ├── check.go                // `port-detective check <port>`
│   ├── kill.go                 // `port-detective kill <port>`
│   └── scan.go                 // `port-detective scan <range>` (optional)
├── internal/
│   ├── lookup/
│   │   ├── lookup.go           // interface PortLookupStrategy
│   │   ├── linux.go            // //go:build linux
│   │   ├── darwin.go           // //go:build darwin
│   │   └── windows.go          // //go:build windows
│   ├── process/
│   │   └── process.go          // struct ProcessInfo
│   └── output/
│       ├── text.go             // format text đẹp (có màu)
│       └── json.go             // format JSON
├── tasks/
│   ├── plan.md                 // (file này)
│   └── todo.md                 // task tracker
└── README.md
```

## Thư viện sử dụng

| Nhu cầu | Thư viện | Lý do |
|---|---|---|
| CLI framework | `github.com/spf13/cobra` | Chuẩn công nghiệp, dùng bởi kubectl, docker, hugo |
| Output có màu | `github.com/fatih/color` | Đơn giản, nhẹ |
| Build & release | `goreleaser` | Cross-compile + GitHub Release |

## Các Phase triển khai

### Phase 1: Foundation
Khởi tạo Go module, data model `ProcessInfo`, output formatter (text + JSON).

### Phase 2: Core Lookup
Interface `PortLookupStrategy`, implement cho Windows (`netstat` + `tasklist`), Linux (`lsof` / `/proc/net/tcp`), macOS (`lsof`).

### Phase 3: CLI Commands
Tích hợp Cobra, xây dựng `check` command và `kill` command.

### Phase 4: Polish & Cross-platform
Thêm màu output, verify cross-compile, chuẩn hóa error handling.

### Phase 5: Documentation & Release
README hoàn chỉnh, `scan` command (optional), GoReleaser config.

## Rủi ro & Biện pháp

| Rủi ro | Biện pháp |
|---|---|
| `netstat` output format khác giữa Windows versions | Parse linh hoạt, unit test với nhiều sample |
| `lsof` không có trên một số Linux distro | Fallback đọc `/proc/net/tcp` |
| Kill process cần quyền Admin trên Windows | Thông báo lỗi rõ ràng |
| Nhiều process cùng listen 1 port (SO_REUSEPORT) | Trả về danh sách, để user chọn |
