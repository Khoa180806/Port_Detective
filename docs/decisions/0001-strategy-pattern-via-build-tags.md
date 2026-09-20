# ADR-001: Sử dụng Strategy Pattern qua Go Build Tags cho logic tra cứu hệ điều hành

## Trạng thái (Status)
Đã chấp thuận (Accepted)

## Ngày (Date)
2026-09-20

## Bối cảnh (Context)
Port Detective là một công cụ CLI chạy trên nhiều hệ điều hành (Windows, Linux, macOS). Mỗi hệ điều hành có cách thức quản lý tiến trình và socket mạng khác nhau:
- **Windows** sử dụng các tiện ích như `netstat -ano`, `tasklist`, WMI API hoặc `taskkill`.
- **Linux** sử dụng `lsof`, `/proc/net/tcp`, hoặc POSIX signals.
- **macOS** sử dụng `lsof` với định dạng tham số riêng của Darwin.

Chúng ta cần một cơ chế kiến trúc để phân tách mã nguồn của từng hệ điều hành mà vẫn đảm bảo tính module hóa, dễ kiểm thử và không mang theo code thừa sang binary của nền tảng khác.

## Quyết định (Decision)
Chúng ta quyết định áp dụng **Strategy Pattern** được hiện thực hóa ở cấp độ trình biên dịch thông qua **Go Build Tags** (`//go:build <os>`).

Cụ thể:
1. Package `internal/lookup` sẽ là điểm trung gian cung cấp các hàm public:
   - `FindProcessByPort(port int) ([]process.ProcessInfo, error)`
   - `KillProcess(pid int) error`
2. Mỗi hệ điều hành sẽ có file hiện thực riêng biệt:
   - `windows.go` chứa directive `//go:build windows`
   - `linux.go` chứa directive `//go:build linux`
   - `darwin.go` chứa directive `//go:build darwin`
3. Mỗi file tự đảm nhiệm việc gọi lệnh hoặc API hệ thống tương ứng của OS đó và chuẩn hóa dữ liệu về struct `process.ProcessInfo`.

## Các phương án đã cân nhắc (Alternatives Considered)

### Phương án 1: Runtime `if-else` / `switch runtime.GOOS`
- **Mô tả:** Đặt toàn bộ code vào một file hoặc package và dùng `switch runtime.GOOS` khi hàm được gọi lúc runtime.
- **Lý do loại bỏ:** 
  - Kéo theo các thư viện hoặc lệnh không tương thích giữa các OS vào cùng một binary.
  - Mã nguồn trở nên cồng kềnh, khó đọc và khó viết unit test chuyên biệt cho từng nền tảng.
  - Nguy cơ xảy ra lỗi lúc runtime nếu gọi nhầm các gói đặc thù của OS khác (ví dụ: `golang.org/x/sys/windows`).

### Phương án 2: Sử dụng thư viện CGO hoặc native C bindings
- **Mô tả:** Viết C code hoặc gọi trực tiếp Win32 API / libc socket headers qua CGO.
- **Lý do loại bỏ:** 
  - Phá vỡ tính năng cross-compile đơn giản của Go (`CGO_ENABLED=0`).
  - Đòi hỏi môi trường cài đặt C compiler (GCC/Clang/MSVC) phức tạp cho người đóng góp và pipeline CI/CD.

## Hệ quả & Lợi ích (Consequences)
- **Lợi ích:**
  - Binary cho từng OS chỉ chứa mã nguồn cần thiết cho OS đó (nhẹ hơn, tối ưu hơn).
  - Khả năng cross-compile cực kỳ thuận tiện: chỉ cần đổi `GOOS=linux go build` là có binary cho Linux từ máy Windows mà không cần cấu hình thêm toolchain ngoài.
  - Tách bạch trách nhiệm (Separation of Concerns): sửa đổi hoặc fix bug cho Windows không ảnh hưởng hay gây rủi ro cho Linux/macOS.
- **Thách thức:**
  - Cần viết các bài unit test phân tích dữ liệu giả lập (mocked parser tests) cho cả 3 nền tảng để có thể kiểm thử cross-platform trên một máy phát triển duy nhất.
