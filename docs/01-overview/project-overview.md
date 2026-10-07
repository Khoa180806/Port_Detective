# Project Overview

This document introduces Port Detective, outlining the developer problem it solves, target users, core features, and technology stack.

---

## 1. Problem Statement

Modern software engineers frequently encounter socket conflict errors when launching local development servers:

```text
Error: listen EADDRINUSE: address already in use :::8080
bind: address already in use
```

Investigating and resolving these conflicts typically requires remembering OS-specific terminal commands:
- **Windows**: `netstat -ano | findstr :8080` followed by `tasklist /FI "PID eq <pid>"` and `taskkill /F /PID <pid>`
- **Linux / macOS**: `lsof -i :8080` followed by `kill -9 <pid>`

These manual steps are slow, error-prone, and inconsistent across operating systems. **Port Detective (`pd`)** provides a unified, cross-platform CLI tool to query, preview, and terminate port-holding processes in a single command.
*(source: [README.md](file:///d:/Project/PortDetective/README.md))*

---

## 2. Target Users

- **Full-Stack & Backend Developers**: Running multiple local servers (Node.js, Go, Python, Docker, Spring Boot) who encounter port collision errors.
- **DevOps & Platform Engineers**: Debugging container socket listeners in stripped Linux Docker containers.
- **CI/CD Pipeline Authors**: Requiring deterministic exit codes and `--json` serialization to verify port availability in automated test runners.

---

## 3. Key Capabilities

- **Unified CLI Syntax**: Consistent command signatures (`check`, `kill`, `scan`) across Windows, macOS, and Linux.
- **Instant Execution**: Native compiled Go binary executing in under 10 milliseconds with zero external runtime dependencies.
- **Safety by Default**: Interactive confirmation prompt (`[y/N]`) before killing processes, with `--dry-run` preview and OS critical PID protection (PID 0, 4, launchd).
- **High-Throughput Scanner**: Concurrent Goroutine worker pool scanning hundreds of ports in milliseconds.
- **Machine-Parseable Output**: Full `--json` flag support across all commands for script piping and `jq` automation.
- **Built-in Bilingual Support**: Instant runtime language switching between English and Vietnamese (`--lang vi` / `PORT_DETECTIVE_LANG=vi`).
*(source: [cmd/check.go](file:///d:/Project/PortDetective/cmd/check.go), [cmd/kill.go](file:///d:/Project/PortDetective/cmd/kill.go), [cmd/scan.go](file:///d:/Project/PortDetective/cmd/scan.go))*

<div align="center">
  <img src="../assets/screenshots/cli-check-kill.png" alt="Terminal output of check and kill operations" width="85%" />
  <p><em>Figure 1: Inspecting occupied ports and executing safe dry-run preview before termination.</em></p>
</div>

<div align="center">
  <img src="../assets/screenshots/cli-scan-range.png" alt="Terminal output of high-throughput range scanner" width="85%" />
  <p><em>Figure 2: Concurrently scanning contiguous port ranges using Goroutine worker pools.</em></p>
</div>

---

## 4. Technology Stack

### Core CLI Tool
- **Language**: Go 1.22+
- **CLI Framework**: `github.com/spf13/cobra` v1.10.2
- **Color Formatting**: `github.com/fatih/color` v1.19.0
- **Cross-Platform Abstraction**: Go compile-time build tags (`//go:build`)
- **Release Automation**: GoReleaser + GitHub Actions
*(source: [go.mod](file:///d:/Project/PortDetective/go.mod))*

### Web Showcase & Documentation
- **Framework**: Next.js 16 (App Router, SSG static export)
- **Styling**: Tailwind CSS v4
- **Icons**: Lucide React
- **Hosting**: Vercel Edge Network
*(source: [web/package.json](file:///d:/Project/PortDetective/web/package.json))*
