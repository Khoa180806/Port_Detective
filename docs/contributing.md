# 🤝 Hướng dẫn đóng góp (Contributing Guide)

Cảm ơn bạn đã quan tâm đến việc đóng góp cho **Port Detective**! Dưới đây là các hướng dẫn giúp bạn thiết lập môi trường phát triển và gửi các đóng góp chất lượng.

---

## 🛠️ Chuẩn bị môi trường

1. **Cài đặt Go**: Yêu cầu Go phiên bản **1.21** trở lên.
2. **Clone repository**:
   ```bash
   git clone https://github.com/Khoa180806/Port_Detective.git
   cd Port_Detective
   ```
3. **Cài đặt dependencies**:
   ```bash
   go mod download
   ```

---

## 🧪 Quy trình phát triển & Kiểm thử

### 1. Chạy Unit Test
Chúng tôi khuyến khích phát triển theo hướng kiểm thử (TDD). Chạy toàn bộ test suite bằng lệnh:
```bash
go test -v ./...
```

### 2. Kiểm tra tĩnh (Lint & Vet)
Đảm bảo mã nguồn tuân thủ các quy chuẩn idiomatic của Go:
```bash
go vet ./...
```

### 3. Kiểm thử Biên dịch Đa nền tảng (Cross-compile Check)
Vì dự án dùng Go Build Tags cho từng hệ điều hành, hãy kiểm tra biên dịch cho cả 3 nền tảng trước khi tạo PR:

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o bin/port-detective-windows.exe .

# Linux
GOOS=linux GOARCH=amd64 go build -o bin/port-detective-linux .

# macOS (Darwin)
GOOS=darwin GOARCH=arm64 go build -o bin/port-detective-darwin .
```

---

## 📝 Quy chuẩn Commit (Commit Message Guidelines)

Chúng tôi áp dụng định dạng **Conventional Commits**:

- `feat:` Thêm tính năng mới (ví dụ: `feat: add scan command for port ranges`)
- `fix:` Sửa lỗi (ví dụ: `fix: handle edge case when netstat PID is 0`)
- `refactor:` Tái cấu trúc mã nguồn mà không thay đổi hành vi logic
- `test:` Bổ sung hoặc cập nhật unit test
- `docs:` Thay đổi hoặc bổ sung tài liệu
- `chore:` Các thay đổi phụ trợ như cấu hình build, gitignore...

---

## 🔀 Quy trình gửi Pull Request (PR)

1. Fork repository về tài khoản cá nhân.
2. Tạo nhánh tính năng mới từ `master` hoặc `main` (ví dụ: `git checkout -b feat/support-udp-ports`).
3. Thực hiện thay đổi, chia nhỏ thành các commit có ý nghĩa rõ ràng.
4. Đảm bảo mọi test case đều pass và mã nguồn được định dạng bằng `gofmt`.
5. Tạo Pull Request và mô tả chi tiết những gì bạn đã làm cùng các bước kiểm thử thực tế.
