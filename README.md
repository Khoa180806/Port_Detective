# 🔍 Port Detective

> **Port Detective** là công cụ dòng lệnh (CLI) hiệu năng cao, viết bằng **Go (Golang)**, giúp lập trình viên nhanh chóng tra cứu và giải phóng (kill) các tiến trình (process) đang chiếm dụng port trên máy local.

---

## 🌟 Tính năng nổi bật

- ⚡ **Siêu nhanh & Nhẹ**: Khởi động tức thì, tiêu tốn cực ít tài nguyên, biên dịch thành 1 file thực thi (binary) duy nhất không cần cài đặt runtime.
- 💻 **Đa nền tảng (Cross-platform)**: Hoạt động mượt mà trên **Windows**, **macOS**, và **Linux** thông qua cơ chế Go Build Tags (Strategy Pattern).
- 🛡️ **Giải phóng port an toàn**: Lệnh `kill` luôn hiển thị chi tiết tiến trình (PID, tên, lệnh khởi chạy) và yêu cầu xác nhận trước khi thực hiện, tránh kill nhầm tiến trình quan trọng.
- 🎨 **Giao diện trực quan**: Hỗ trợ hiển thị dạng bảng có màu sắc sinh động trên terminal, tự động tắt màu khi xuất dữ liệu qua đường ống (pipe/redirect).
- 🤖 **Thân thiện với Tự động hóa**: Hỗ trợ cờ `--json` xuất dữ liệu chuẩn JSON, dễ dàng tích hợp vào shell script, CI/CD pipeline, hoặc các công cụ giám sát.

---

## 🚀 Cài đặt

### Cách 1: Sử dụng `go install` (Yêu cầu Go 1.21+)

```bash
go install github.com/Khoa180806/Port_Detective@latest
```

### Cách 2: Tải bản phát hành (Pre-built Binary)

Tải file thực thi tương ứng với hệ điều hành của bạn từ trang [GitHub Releases](https://github.com/Khoa180806/Port_Detective/releases).

### Cách 3: Build từ mã nguồn

```bash
git clone https://github.com/Khoa180806/Port_Detective.git
cd Port_Detective
go build -o port-detective .
```

---

## 📖 Hướng dẫn sử dụng

### 1. Kiểm tra tiến trình đang chiếm port (`check`)

Kiểm tra xem port cụ thể có đang bị chiếm dụng hay không:

```bash
# Kiểm tra port 8080
port-detective check 8080

# Xuất kết quả dạng JSON
port-detective check 8080 --json
```

**Ví dụ đầu ra (Text):**
```text
Port 8080 is occupied by:
  PID:      12345
  Process:  node.exe
  Command:  node server.js
  Protocol: tcp
```

**Ví dụ đầu ra (JSON):**
```json
[
  {
    "pid": 12345,
    "name": "node.exe",
    "command": "node server.js",
    "port": 8080,
    "protocol": "tcp"
  }
]
```

**Quy ước mã thoát (Exit codes):**
- `0`: Tìm thấy tiến trình đang giữ port.
- `1`: Port đang khả dụng (free / không có tiến trình nào chiếm).
- `2`: Lỗi hệ thống (port không hợp lệ, không đủ quyền truy cập...).

---

### 2. Dừng tiến trình đang chiếm port (`kill`)

Giải phóng port bằng cách kết thúc tiến trình đang chiếm giữ:

```bash
# Hiển thị thông tin tiến trình và hỏi xác nhận [y/N] trước khi dừng
port-detective kill 8080

# Dừng tiến trình ngay lập tức, bỏ qua bước xác nhận
port-detective kill 8080 --force

# Xem trước (Dry Run) những tiến trình sẽ bị tiêu diệt mà không thực sự kill
port-detective kill 8080 --dry-run
```

---

### 3. Quét dải port (`scan` - Đang phát triển / Task 15)

Quét một khoảng port để xem những port nào đang mở:

```bash
port-detective scan 3000-3010
port-detective scan 8000-8080 --json
```

---

## 🏗️ Kiến trúc dự án

Dự án được tổ chức phân tầng rõ ràng theo chuẩn Go idiomatic:

```text
port-detective/
├── cmd/                        # Định nghĩa lệnh CLI (Cobra Framework)
│   ├── root.go                 # Lệnh gốc & cấu hình toàn cục
│   ├── check.go                # Lệnh kiểm tra port (`check`)
│   ├── kill.go                 # Lệnh giải phóng port (`kill`)
│   └── scan.go                 # Lệnh quét dải port (`scan`)
├── internal/                   # Logic nghiệp vụ nội bộ (không public)
│   ├── lookup/                 # OS Strategy Pattern tra cứu port
│   │   ├── lookup.go           # Interface PortLookupStrategy
│   │   ├── windows.go          # Xử lý trên Windows (netstat + tasklist)
│   │   ├── linux.go            # Xử lý trên Linux (lsof / procfs)
│   │   └── darwin.go           # Xử lý trên macOS (lsof)
│   ├── process/                # Data model định nghĩa ProcessInfo
│   └── output/                 # Định dạng dữ liệu (Text/Color & JSON)
├── docs/                       # Tài liệu kỹ thuật chuyên sâu & ADRs
└── main.go                     # Điểm khởi chạy ứng dụng
```

Xem chi tiết trong:
- [Kiến trúc hệ thống](docs/architecture.md)
- [Bản kế hoạch chi tiết (Original Plan)](docs/port-detective-plan.md)
- [Quyết định kiến trúc (ADR-001: Build Tags Strategy)](docs/decisions/0001-strategy-pattern-via-build-tags.md)

---

## 🤝 Đóng góp phát triển

Mọi ý kiến đóng góp, báo lỗi (issue) và đề xuất tính năng (pull request) đều được hoan nghênh! Vui lòng đọc qua [Hướng dẫn đóng góp](docs/contributing.md) trước khi tạo pull request.

---

## 📄 Bản quyền (License)

Dự án được phân phối dưới giấy phép [MIT License](LICENSE).
