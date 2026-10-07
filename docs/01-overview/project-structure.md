# Codebase Project Structure

This document outlines the file tree and responsibilities of each package in the Port Detective repository.

---

## 1. Directory Tree Overview

```text
port-detective/
├── .github/
│   └── workflows/
│       └── release.yml          # GitHub Actions workflow for GoReleaser automation
├── cmd/
│   ├── pd/
│   │   └── main.go              # Primary application entrypoint (compiles to 'pd' binary)
│   ├── root.go                  # Cobra root command definition & i18n initialization hook
│   ├── check.go                 # 'pd check <port>' subcommand
│   ├── kill.go                  # 'pd kill <port>' subcommand with safety confirmation
│   └── scan.go                  # 'pd scan <range>' subcommand with worker pool concurrency
├── docs/                        # Complete technical and product documentation
│   └── assets/                  # Media assets (screenshots, terminal gif, architecture diagram)
│       ├── screenshots/         # Terminal CLI screenshots and execution demo GIF
│       └── diagrams/            # Architecture diagrams (PNG, vector SVG)
├── internal/
│   ├── i18n/                    # In-memory dictionary maps (en.go, vi.go, i18n.go)
│   ├── lookup/                  # Platform strategy implementations using Go build tags
│   │   ├── lookup.go            # PortLookupStrategy interface and dispatch logic
│   │   ├── windows.go           # Windows strategy (netstat + tasklist + taskkill)
│   │   ├── linux.go             # Linux strategy (lsof + /proc/net/tcp inode fallback)
│   │   ├── darwin.go            # macOS strategy (native lsof -i)
│   │   └── errors.go            # Custom domain error wrappers & permission helpers
│   ├── output/                  # Terminal table formatting and strict JSON serializers
│   │   ├── text.go              # Human-readable ANSI colorized tables
│   │   └── json.go              # Machine-readable JSON array serialization
│   └── process/                 # Core domain entity (ProcessInfo struct)
├── scripts/
│   ├── install.sh               # Zero-config automated installer for Linux / macOS
│   ├── install.ps1              # Zero-config automated installer for Windows PowerShell
│   └── build_all.bat            # Cross-compilation helper script for all target platforms
├── web/                         # Next.js 16 Web Landing Page and interactive terminal showcase
│   ├── app/                     # Next.js App Router (layout, page, opengraph-image)
│   ├── components/              # Modular UI components (Hero, Terminal, Features, FAQ, etc.)
│   ├── public/                  # Public web assets (favicons, manifest)
│   └── vercel.json              # Vercel deployment configuration
├── go.mod                       # Go module dependency manifest (Go 1.22+)
└── go.sum                       # Cryptographic checksums of Go dependencies
```

---

## 2. Package Roles and Boundaries

| Package / Module | Responsibility | Key Files |
| :--- | :--- | :--- |
| **`cmd/pd`** | Main entrypoint. Calls `cmd.Execute()`. | `main.go` |
| **`cmd`** | CLI command presentation. Flag definition, validation, argument parsing, exit codes. | `root.go`, `check.go`, `kill.go`, `scan.go` |
| **`internal/process`** | Core domain model. Defines `ProcessInfo` data contract. | `process.go` |
| **`internal/lookup`** | Operating system abstraction layer. Resolves ports to PIDs and kills processes. | `lookup.go`, `windows.go`, `linux.go`, `darwin.go` |
| **`internal/output`** | Presentation layer. Formats `ProcessInfo` into colorized terminal text or JSON. | `text.go`, `json.go` |
| **`internal/i18n`** | Localization subsystem. Manages thread-safe key-value dictionaries. | `i18n.go`, `en.go`, `vi.go` |
| **`web`** | Marketing landing page and interactive simulated terminal showcase. | `app/page.tsx`, `components/*` |
*(source: Codebase root directory inspection)*
