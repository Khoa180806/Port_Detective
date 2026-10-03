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
      stat1: "< 50ms",
      stat1_label: "Warm lookup latency",
      stat2: "10-30x",
      stat2_label: "Faster than manual netstat",
      stat3: "~3.7 MB",
      stat3_label: "Single native binary",
      stat4: "Zero",
      stat4_label: "Runtime dependencies required",
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
      copyright: `© ${new Date().getFullYear()} Port Detective. Released under the MIT License.`,
    },
  },
  vi: {
    // Nav
    nav: {
      features: "Tính năng",
      demo: "Demo",
      benchmark: "Benchmark",
      docs: "Tài liệu",
      install: "Cài đặt",
    },
    // Hero
    hero: {
      releaseBadge: "Port Detective v1.1.1 đã ra mắt",
      changelog: "Changelog",
      titleLine1: "Quên đi nỗi ám ảnh",
      titleLine2: "Check và giải phóng port chỉ trong 1 giây.",
      subtitle:
        "Công cụ CLI siêu tốc viết bằng Go: tìm tiến trình chiếm port, kill an toàn có hỏi xác nhận và quét dải port đa luồng cực nhanh.",
      sub50ms: "Khởi động dưới 50ms",
      safeKill: "Xác nhận an toàn [y/N]",
      crossPlatform: "Windows, macOS & Linux",
      exploreSource: "Mã nguồn dự án:",
      repoText: "GitHub Repository (Khoa180806/Port_Detective)",
    },
    // Install Tabs
    install: {
      windows: "Windows (PowerShell)",
      unix: "Linux & macOS",
      go: "Go Install",
      winComment: "# Chạy trong PowerShell — Tự nhận diện kiến trúc & cấu hình PATH",
      unixComment: "# Chạy trong Terminal — Cài đặt tự động không cần cấu hình",
      goComment: "# Biên dịch và cài đặt trực tiếp vào GOPATH/bin",
      verify: "Kiểm tra:",
      readyImmediately: "Dùng được ngay sau khi cài",
      copied: "Đã copy!",
    },
    // Terminal Preview
    demo: {
      badge: "Terminal Demo tương tác",
      title: "Xem Port Detective hoạt động thực tế",
      subtitle: "Bấm vào các lệnh bên dưới để xem output hiển thị chuẩn terminal.",
      replay: "Chạy lại",
      copy: "Copy lệnh",
      copied: "Đã copy",
      executing: "Đang chạy qua Go runtime...",
      infoPrefix: "Chi tiết:",
      exitCode: "Exit Code: 0",
      checkDesc: "Biết ngay ứng dụng hay tiến trình nào đang chiếm giữ port 8080.",
      killDesc: "Kill tiến trình an toàn, luôn hỏi xác nhận [y/N] để tránh tắt nhầm.",
      scanDesc: "Quét đồng thời một dải port chỉ mất vài mili-giây với worker pool.",
      jsonDesc: "Xuất chuẩn JSON phục vụ viết script tự động hóa, CI/CD hoặc lọc bằng jq.",
    },
    // Features
    features: {
      badge: "Điểm nổi bật",
      title: "Tối ưu cho Lập trình viên",
      subtitle:
        "Mọi tính năng cần thiết để xử lý xung đột port khi dev mà không làm đứt mạch làm việc.",
      f1_title: "Tốc độ dưới 50ms",
      f1_desc:
        "Viết bằng Go và build ra binary native độc lập. Khởi động tức thì, không cần môi trường NodeJS hay Python.",
      f1_badge: "Go Native",
      f2_title: "An toàn là trên hết",
      f2_desc:
        "Không bao giờ kill bừa. Luôn hiển thị tên tiến trình, PID và hỏi xác nhận [y/N], có hỗ trợ cờ --dry-run chạy thử.",
      f2_badge: "Safe Guardrails",
      f3_title: "Đa nền tảng đồng nhất",
      f3_desc:
        "Cùng một cú pháp cho cả Windows, macOS và Linux. Không cần nhớ netstat, lsof hay taskkill khi đổi máy làm việc.",
      f3_badge: "Cross-Platform",
      f4_title: "Quét dải port đa luồng",
      f4_desc:
        "Tận dụng sức mạnh Goroutines để quét đồng thời dải port lên tới 5.000 port chỉ trong chớp mắt.",
      f4_badge: "Goroutines",
      f5_title: "Định dạng JSON chuẩn",
      f5_desc:
        "Hỗ trợ cờ --json ở tất cả các lệnh. Dễ dàng tích hợp vào CI/CD pipeline, Makefile hoặc pipe qua jq.",
      f5_badge: "Automation Ready",
      f6_title: "Hỗ trợ Song ngữ gốc",
      f6_desc:
        "Hỗ trợ sẵn tiếng Anh và tiếng Việt ngay trong CLI. Đổi ngôn ngữ linh hoạt qua cờ --lang vi hoặc biến môi trường.",
      f6_badge: "Song ngữ EN/VI",
    },
    // Benchmark
    benchmark: {
      badge: "Đo lường hiệu quả",
      title: "1 Giây thay vì 30 Giây",
      subtitle:
        "So sánh quy trình gõ lệnh thủ công truyền thống và trải nghiệm mượt mà với Port Detective.",
      tradTitle: "Cách gõ lệnh truyền thống",
      tradFriction: "~25–30s thao tác",
      tradStep1: "Gõ lệnh dài dòng: netstat -ano | findstr :8080",
      tradStep2: "Căng mắt nhìn bảng kết quả để tìm số PID (vd: 14280)",
      tradStep3: "Gõ tiếp tasklist | findstr 14280 để xem đó là phần mềm nào",
      tradStep4: "Gõ lệnh kill cưỡng bức: taskkill /F /PID 14280",
      tradStep5: "Dễ gõ nhầm PID dẫn đến tắt nhầm service hệ thống quan trọng",
      tradSummary: "❌ Tốn thời gian, dễ nhầm lẫn số PID, làm gián đoạn luồng tập trung khi đang code.",
      pdTitle: "Port Detective",
      pdTime: "⚡ ~1s là xong",
      pdBadge: "Khuyên dùng",
      pdCheckDesc: "In ngay PID, tên tiến trình, giao thức và lệnh chạy với màu sắc trực quan.",
      pdKillDesc: "Hỏi xác nhận kèm đầy đủ thông tin app, hoặc thêm cờ -f để tắt ngay lập tức.",
      pdSummary: "Một cú pháp duy nhất cho Windows, macOS và Linux. Không độ trễ.",
      stat1: "< 50ms",
      stat1_label: "Độ trễ tra cứu (Warm cache)",
      stat2: "10-30x",
      stat2_label: "Nhanh hơn gõ lệnh netstat",
      stat3: "~3.7 MB",
      stat3_label: "File binary độc lập duy nhất",
      stat4: "0",
      stat4_label: "Không cần runtime phụ thuộc",
    },
    // Footer
    footer: {
      tagline:
        "Công cụ dòng lệnh siêu tốc, đa nền tảng giúp tra cứu và giải phóng port mạng bị chiếm dụng.",
      resources: "Tài liệu",
      architecture: "Kiến trúc kỹ thuật",
      cliRef: "Cẩm nang tra cứu CLI",
      installGuide: "Hướng dẫn cài đặt",
      contributing: "Hướng dẫn đóng góp",
      community: "Cộng đồng",
      githubRepo: "GitHub Repository",
      releases: "Bản phát hành (v1.1.1)",
      issues: "Báo lỗi & Góp ý (Issues)",
      mitLicense: "Giấy phép MIT",
      builtWith: "Viết bằng Go. Dành cho lập trình viên.",
      copyright: `© ${new Date().getFullYear()} Port Detective. Phát hành dưới giấy phép MIT License.`,
    },
  },
};
