# 🗺️ Lộ trình phát triển & Định hướng tương lai — Port Detective

<div align="center">

[![Port Detective Documentation](https://img.shields.io/badge/docs-roadmap-blue?style=flat-square)](../README.md)
[![Status](https://img.shields.io/badge/status-planning-orange?style=flat-square)](../README.md)
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen?style=flat-square)](./contributing.md)

</div>

Tài liệu này tổng hợp các ý tưởng, đề xuất tính năng và kế hoạch phát triển chiến lược cho **Port Detective (`pd`)** trong các phiên bản tiếp theo, nhằm đưa công cụ từ một tiện ích CLI cá nhân trở thành một công cụ mã nguồn mở phổ biến trong cộng đồng lập trình viên.

---

## 📑 Mục lục lộ trình

- [Giai đoạn 1: Phân phối & Quản lý gói cài đặt (Package Managers)](#-giai-đoạn-1-phân-phối--quản-lý-gói-cài-đặt-package-managers)
- [Giai đoạn 2: Tính năng CLI & Trải nghiệm tương tác (Feature Enhancements)](#-giai-đoạn-2-tính-năng-cli--trải-nghiệm-tương-tác-feature-enhancements)
- [Giai đoạn 3: Hệ sinh thái & Tích hợp IDE (Ecosystem & Integrations)](#-giai-đoạn-3-hệ-sinh-thái--tích-hợp-ide-ecosystem--integrations)
- [Giai đoạn 4: Thư viện Go SDK (`pkg/detective`)](#-giai-đoạn-4-thư-viện-go-sdk-pkgdetective)
- [Giai đoạn 5: Tự động hóa CI/CD & Lan tỏa cộng đồng (Community & CI/CD)](#-giai-đoạn-5-tự-động-hóa-cicd--lan-tỏa-cộng-đồng-community--cicd)

---

## 📦 Giai đoạn 1: Phân phối & Quản lý gói cài đặt (Package Managers)

Mục tiêu: Giúp người dùng cài đặt và cập nhật `pd` bằng các trình quản lý gói quen thuộc trên mọi nền tảng thay vì chỉ dùng script curl/powershell.

- [ ] **Homebrew (`brew install port-detective`)**:
  - Tạo Homebrew Tap repository (`Khoa180806/homebrew-tap`).
  - Hỗ trợ macOS (Apple Silicon + Intel) và Linux.
- [ ] **Windows Package Managers**:
  - **Scoop**: Đưa manifest vào Scoop Main/Extras bucket (`scoop install pd`).
  - **Chocolatey** / **Winget**: Cung cấp package cài đặt native cho Windows.
- [ ] **Linux Repositories**:
  - **AUR (Arch User Repository)** cho người dùng Arch Linux / Manjaro.
  - Gói `.deb` (Debian/Ubuntu) và `.rpm` (Fedora/RHEL).
- [ ] **Docker Hub / GHCR Container**:
  - Cung cấp image container siêu nhẹ (`ghcr.io/khoa180806/port-detective:latest`).
  - Phù hợp để debug network trong môi trường container: `docker run --net=host --rm port-detective check 8080`.

---

## 🚀 Giai đoạn 2: Tính năng CLI & Trải nghiệm tương tác (Feature Enhancements)

Mục tiêu: Mở rộng khả năng điều tra sâu tiến trình và mang lại trải nghiệm tương tác trực quan ngay trong terminal.

- [ ] **Dashboard TUI tương tác thời gian thực (`pd top` / `pd monitor`)**:
  - Tích hợp framework terminal hiện đại như `charmbracelet/bubbletea`.
  - Hiển thị bảng live-monitoring các port đang mở, tự động cập nhật theo chu kỳ.
  - Cho phép dùng phím mũi tên `↑/↓` để duyệt danh sách và phím tắt `k` để kill tiến trình trực tiếp.
- [ ] **Chế độ theo dõi & Tự động giải phóng port (`pd watch <port>`)**:
  - Giám sát một hoặc nhiều port chỉ định (ví dụ port dev `3000`, `8080`).
  - Khi phát hiện server crash để lại tiến trình zombie, tự động phát cảnh báo hoặc giải phóng port ngay tức khắc.
- [ ] **Thông tin chi tiết tiến trình (Deep Process Insights)**:
  - Bổ sung thông tin:
    - **CWD (Current Working Directory)**: Thư mục chứa mã nguồn của tiến trình đang chạy (giúp biết ngay PID thuộc project nào).
    - **Uptime / Start Time**: Thời gian tiến trình đã chạy.
    - **Memory / CPU Usage**: Mức độ tiêu thụ tài nguyên hệ thống của PID.
- [ ] **Hỗ trợ hủy toàn bộ nhánh tiến trình (`pd kill <port> --tree`)**:
  - Hỗ trợ kill đệ quy toàn bộ Process Tree (cả process cha và các process con sinh ra từ `npm`, `nodemon`, `vite`, `gunicorn`).

---

## 🧩 Giai đoạn 3: Hệ sinh thái & Tích hợp IDE (Ecosystem & Integrations)

Mục tiêu: Đưa Port Detective vào nơi lập trình viên làm việc hàng ngày.

- [ ] **VS Code Extension (`Port Detective for VS Code`)**:
  - Hiển thị các port đang active ngay trên thanh trạng thái (Status Bar) hoặc bảng điều khiển riêng ở Activity Bar.
  - Nút bấm 1-click **"Free Port"** để giải phóng port bị kẹt mà không cần mở terminal.
- [ ] **Raycast / Alfred Extension (macOS)**:
  - Cho phép gọi `pd <port>` trực tiếp từ launcher nhanh trên macOS.

---

## 🐹 Giai đoạn 4: Thư viện Go SDK (`pkg/detective`)

Mục tiêu: Cung cấp API chuẩn cho các nhà phát triển khác nhúng logic kiểm tra port vào ứng dụng Go của họ.

- [ ] Tách core logic kiểm tra/kill port ra package `pkg/detective` tách biệt khỏi presentation layer của Cobra.
- [ ] Cung cấp Go documentation (`pkg.go.dev/github.com/Khoa180806/Port_Detective/pkg/detective`).
- [ ] Viết ví dụ mẫu tích hợp vào test suites (ví dụ: tự động dọn dẹp port trước khi chạy integration test).

---

## 📣 Giai đoạn 5: Tự động hóa CI/CD & Lan tỏa cộng đồng (Community & CI/CD)

- [ ] **Tích hợp GoReleaser**:
  - Tự động hóa quy trình build matrix đa nền tảng, tạo changelog, tính mã băm checksum và upload release assets khi push git tag.
  - Tự động update Homebrew formula khi có release mới.
- [ ] **GitHub Community Templates**:
  - Issue templates: Bug report, Feature request.
  - Pull request template với checklist rõ ràng.
- [ ] **Bài viết & Chia sẻ kỹ thuật**:
  - Viết bài chia sẻ kỹ thuật trên Dev.to / Hashnode: *"How we built a zero-dependency cross-platform port detective in Go"*.
  - Giới thiệu trên các cộng đồng lập trình: Reddit (`r/golang`, `r/webdev`), Facebook Golang Vietnam, J2TEAM.

---

## 🧭 Điều hướng tài liệu

- [🏠 Trang chủ & README](../README.md)
- [📐 Kiến trúc hệ thống](./architecture.md)
- [📖 Cẩm nang tra cứu CLI](./cli-reference.md)
- [📦 Hướng dẫn cài đặt](./installation.md)
- [🤝 Hướng dẫn đóng góp](./contributing.md)
