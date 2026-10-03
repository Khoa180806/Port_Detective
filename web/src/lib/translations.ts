export type Locale = "en" | "vi";

export const translations = {
  en: {
    // Nav
    nav: {
      features: "Features",
      demo: "Demo",
      benchmark: "Benchmark",
      docs: "Docs",
      install: "Install",
    },
    // Hero
    hero: {
      releaseBadge: "Port Detective v1.1.1 Released",
      changelog: "Changelog",
      titleLine1: "Stop wrestling with",
      titleLine2: "Investigate & kill ports in 1 second.",
      subtitle:
        "A lightning-fast, native Go CLI tool to inspect zombie processes, safely terminate port-hogging servers with confirmation prompts, and scan port ranges concurrently.",
      sub50ms: "Sub-50ms execution",
      safeKill: "Interactive Safe Kill [y/N]",
      crossPlatform: "Windows, macOS & Linux",
      exploreSource: "Or explore source code:",
      repoText: "GitHub Repository (Khoa180806/Port_Detective)",
    },
    // Install Tabs
    install: {
      windows: "Windows (PowerShell)",
      unix: "Linux & macOS",
      go: "Go Install",
      winComment: "# Run in PowerShell — Auto detects arch & configures PATH",
      unixComment: "# Run in Terminal — Zero-config one-liner installer",
      goComment: "# Build & install binary directly to your GOPATH/bin",
      verify: "Verify:",
      readyImmediately: "Ready immediately after install",
      copied: "Copied!",
    },
    // Terminal Preview
    demo: {
      badge: "Interactive CLI Preview",
      title: "See Port Detective in Action",
      subtitle: "Click any command below to test the syntax and see genuine terminal output.",
      replay: "Replay",
      copy: "Copy",
      copied: "Copied",
      executing: "Executing native Go routine...",
      infoPrefix: "Info:",
      exitCode: "Exit Code: 0",
      checkDesc: "Instantly identify which application or zombie process is holding port 8080.",
      killDesc: "Safely terminate port-blocking processes with interactive confirmation prompts.",
      scanDesc: "Scan port ranges concurrently in milliseconds using a Go worker pool.",
      jsonDesc: "Clean JSON output designed for headless automation, CI/CD, and jq parsing.",
    },
    // Features
    features: {
      badge: "Capabilities",
      title: "Engineered for Modern Developers",
      subtitle:
        "Everything you need to troubleshoot, free up ports, and streamline your local development environment.",
      f1_title: "Sub-50ms Execution",
      f1_desc:
        "Built with Go and compiled to a lightweight standalone binary. Zero interpreter startup delay, no NodeJS or Python runtime overhead.",
      f1_badge: "Native Performance",
      f2_title: "Safe by Default",
      f2_desc:
        "Prevents catastrophic accidental kills. Displays full process details with an interactive [y/N] prompt and supports --dry-run simulation.",
      f2_badge: "Human Guardrails",
      f3_title: "Unified Cross-Platform",
      f3_desc:
        "One syntax across Windows, macOS, and Linux. No more memorizing netstat -ano, lsof -i, or taskkill commands when switching machines.",
      f3_badge: "Zero Friction",
      f4_title: "Concurrent Port Scanner",
      f4_desc:
        "High-throughput goroutine worker pool scans ranges of up to 5,000 ports in milliseconds. Detect open ports and bound services instantly.",
      f4_badge: "Goroutines",
      f5_title: "Machine-Readable JSON",
      f5_desc:
        "Pass the --json flag on any command for strict, clean JSON output. Built specifically for CI/CD pipelines, shell scripts, and jq filtering.",
      f5_badge: "Automation Ready",
      f6_title: "Native Bilingual CLI",
      f6_desc:
        "Native support for English and Vietnamese out of the box. Switch dynamically via --lang vi or configure system-wide with environment variables.",
      f6_badge: "i18n Built-in",
    },
    // Benchmark
    benchmark: {
      badge: "Developer Productivity",
      title: "1 Second vs 30 Seconds",
      subtitle:
        "Compare the friction of traditional multi-step OS terminal commands against the streamlined workflow of Port Detective.",
      tradTitle: "Traditional OS Commands",
      tradFriction: "~25–30s Friction",
      tradStep1: "Remember & type netstat -ano | findstr :8080",
      tradStep2: "Squint across terminal columns to identify PID (e.g. 14280)",
      tradStep3: "Query process name via tasklist | findstr 14280 or Task Manager",
      tradStep4: "Execute destructive kill: taskkill /F /PID 14280",
      tradStep5: "Risk terminating critical system services due to typing error",
      tradSummary: "❌ High cognitive load, error-prone PID matching, breaks developer flow state.",
      pdTitle: "Port Detective",
      pdTime: "⚡ ~1s Total",
      pdBadge: "Recommended",
      pdCheckDesc: "Instantly prints PID, Process Name, Protocol, and Command line in high-contrast color.",
      pdKillDesc: "Safely asks confirmation with full process context, or bypass with -f.",
      pdSummary: "One unified command across Windows, macOS, and Linux. Zero friction.",
      stat1: "~48ms",
      stat1_label: "pd check runtime",
      stat2: "20x",
      stat2_label: "Faster developer workflow",
      stat3: "5–8 MB",
      stat3_label: "Standalone binary size",
      stat4: "0 dep",
      stat4_label: "External dependencies",
    },
    // Footer
    footer: {
      tagline:
        "A lightning-fast, cross-platform CLI tool to investigate and terminate processes occupying network ports.",
      resources: "Resources",
      architecture: "Technical Architecture",
      cliRef: "CLI Reference Manual",
      installGuide: "Installation Guide",
      contributing: "Contributing Guide",
      community: "Community",
      githubRepo: "GitHub Repository",
      releases: "Releases (v1.1.1)",
      issues: "Report an Issue",
      mitLicense: "MIT License",
      builtWith: "Built with Go. Designed for engineers worldwide.",
      copyright: "Port Detective. Released under the MIT License.",
    },
  },
  vi: {
    // Nav
    nav: {
      features: "Tính năng",
      demo: "Dùng thử",
      benchmark: "So sánh",
      docs: "Tài liệu",
      install: "Cài đặt",
    },
    // Hero
    hero: {
      releaseBadge: "Đã phát hành Port Detective v1.1.1",
      changelog: "Nhật ký thay đổi",
      titleLine1: "Không còn vật lộn với",
      titleLine2: "Tra cứu & giải phóng cổng trong 1 giây.",
      subtitle:
        "Công cụ dòng lệnh (CLI) siêu tốc viết bằng Go: tìm tiến trình chiếm cổng, tắt an toàn có hỏi xác nhận và quét dải cổng đồng thời cực nhanh.",
      sub50ms: "Thực thi dưới 50ms",
      safeKill: "Xác nhận an toàn [y/N]",
      crossPlatform: "Windows, macOS & Linux",
      exploreSource: "Hoặc xem mã nguồn:",
      repoText: "Kho mã nguồn GitHub (Khoa180806/Port_Detective)",
    },
    // Install Tabs
    install: {
      windows: "Windows (PowerShell)",
      unix: "Linux & macOS",
      go: "Cài qua Go",
      winComment: "# Chạy trong PowerShell — Tự nhận arch & cấu hình PATH",
      unixComment: "# Chạy trong Terminal — Cài đặt tự động 1 câu lệnh",
      goComment: "# Tải & build binary trực tiếp vào GOPATH/bin",
      verify: "Kiểm tra:",
      readyImmediately: "Dùng được ngay sau khi cài",
      copied: "Đã sao chép!",
    },
    // Terminal Preview
    demo: {
      badge: "Mô phỏng CLI Tương tác",
      title: "Trải nghiệm Port Detective",
      subtitle: "Nhấp vào bất kỳ lệnh nào dưới đây để xem kết quả terminal chân thực.",
      replay: "Chạy lại",
      copy: "Sao chép",
      copied: "Đã chép",
      executing: "Đang thực thi Go routine...",
      infoPrefix: "Ghi chú:",
      exitCode: "Mã thoát: 0",
      checkDesc: "Xác định ngay ứng dụng hoặc tiến trình nền nào đang giữ cổng 8080.",
      killDesc: "Dừng tiến trình chiếm cổng an toàn với hộp thoại xác nhận tương tác.",
      scanDesc: "Quét dải cổng đồng thời trong vài mili-giây bằng Goroutines worker pool.",
      jsonDesc: "Xuất dữ liệu chuẩn JSON cho CI/CD, script tự động hóa và lọc qua jq.",
    },
    // Features
    features: {
      badge: "Sức mạnh cốt lõi",
      title: "Thiết kế cho Lập trình viên Hiện đại",
      subtitle:
        "Mọi thứ bạn cần để chẩn đoán, giải phóng cổng mạng và tối ưu môi trường lập trình local.",
      f1_title: "Thực thi siêu tốc (< 50ms)",
      f1_desc:
        "Viết bằng Go, biên dịch thẳng ra mã máy native. Khởi động 0ms, không phụ thuộc NodeJS hay Python.",
      f1_badge: "Hiệu năng Native",
      f2_title: "An toàn là trên hết",
      f2_desc:
        "Ngăn ngừa tắt nhầm service quan trọng. Hiện rõ tên app, PID và yêu cầu gõ [y/N] xác nhận trước khi kill.",
      f2_badge: "Cơ chế bảo vệ",
      f3_title: "Đa nền tảng đồng nhất",
      f3_desc:
        "Chung một cú pháp trên Windows, macOS và Linux. Không cần nhớ netstat, lsof hay taskkill riêng lẻ.",
      f3_badge: "Tiện lợi tối đa",
      f4_title: "Quét cổng đa luồng",
      f4_desc:
        "Worker pool tận dụng Go Goroutines quét tới 5.000 port cùng lúc trong chớp mắt.",
      f4_badge: "Goroutines",
      f5_title: "Hỗ trợ chuẩn JSON",
      f5_desc:
        "Cờ --json trên mọi câu lệnh xuất kết quả JSON chuẩn xác, sẵn sàng tích hợp CI/CD và bash script.",
      f5_badge: "Tự động hóa",
      f6_title: "Hỗ trợ Song ngữ gốc",
      f6_desc:
        "Tích hợp sẵn tiếng Anh và tiếng Việt. Chuyển đổi linh hoạt qua cờ --lang vi hoặc biến môi trường.",
      f6_badge: "Song ngữ EN/VI",
    },
    // Benchmark
    benchmark: {
      badge: "Năng suất lập trình",
      title: "1 Giây so với 30 Giây",
      subtitle:
        "So sánh quy trình thao tác thủ công phức tạp của hệ điều hành với sự mượt mà của Port Detective.",
      tradTitle: "Cách gõ lệnh truyền thống",
      tradFriction: "~25–30s Thao tác",
      tradStep1: "Nhớ cú pháp & gõ netstat -ano | findstr :8080",
      tradStep2: "Căng mắt dò tìm số PID ở cột cuối (vd: 14280)",
      tradStep3: "Tra cứu tên app qua tasklist | findstr 14280 hoặc mở Task Manager",
      tradStep4: "Gõ lệnh tắt cưỡng bức: taskkill /F /PID 14280",
      tradStep5: "Rủi ro gõ nhầm PID gây tắt nhầm service hệ thống quan trọng",
      tradSummary: "❌ Tốn thời gian, dễ nhầm lẫn số PID, làm gián đoạn luồng tập trung khi code.",
      pdTitle: "Port Detective",
      pdTime: "⚡ ~1s Tổng cộng",
      pdBadge: "Khuyên dùng",
      pdCheckDesc: "In ngay PID, Tên tiến trình, Giao thức và Lệnh thực thi với màu sắc nổi bật.",
      pdKillDesc: "Hỏi xác nhận kèm đầy đủ thông tin app, hoặc tắt nhanh với -f.",
      pdSummary: "Một lệnh duy nhất cho Windows, macOS và Linux. Không độ trễ.",
      stat1: "~48ms",
      stat1_label: "Thời gian chạy pd check",
      stat2: "20x",
      stat2_label: "Nhanh hơn trong công việc",
      stat3: "5–8 MB",
      stat3_label: "Kích thước file binary",
      stat4: "0 dep",
      stat4_label: "Phụ thuộc thư viện ngoài",
    },
    // Footer
    footer: {
      tagline:
        "Công cụ dòng lệnh (CLI) siêu tốc, đa nền tảng giúp tra cứu và giải phóng các cổng mạng bị chiếm dụng.",
      resources: "Tài liệu kỹ thuật",
      architecture: "Kiến trúc hệ thống",
      cliRef: "Cẩm nang tra cứu CLI",
      installGuide: "Hướng dẫn cài đặt",
      contributing: "Hướng dẫn đóng góp",
      community: "Cộng đồng",
      githubRepo: "Kho mã nguồn GitHub",
      releases: "Bản phát hành (v1.1.1)",
      issues: "Báo cáo lỗi (Issues)",
      mitLicense: "Giấy phép MIT",
      builtWith: "Viết bằng Go. Dành cho lập trình viên toàn cầu.",
      copyright: "Port Detective. Phát hành dưới giấy phép MIT License.",
    },
  },
};
