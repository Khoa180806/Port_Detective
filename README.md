# 🔍 Port Detective

<div align="center">

<h3>A lightning-fast, cross-platform CLI tool to investigate and terminate processes occupying network ports.</h3>

<p align="center">
  <a href="https://port-detective.vercel.app"><img src="https://img.shields.io/badge/Website-port--detective.vercel.app-0ea5e9?style=flat-square&logo=vercel" alt="Landing Page"></a>
  <a href="https://github.com/Khoa180806/Port_Detective/releases"><img src="https://img.shields.io/github/v/release/Khoa180806/Port_Detective?style=flat-square&color=blue" alt="Release"></a>
  <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go" alt="Go Version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-green?style=flat-square" alt="License"></a>
  <a href="https://github.com/Khoa180806/Port_Detective/actions"><img src="https://img.shields.io/github/actions/workflow/status/Khoa180806/Port_Detective/release.yml?style=flat-square&label=build" alt="Build Status"></a>
</p>

<p align="center">
  🌐 <strong>English</strong> · <a href="README.vi.md">🇻🇳 Tiếng Việt</a> · 🚀 <a href="https://port-detective.vercel.app"><strong>Live Web App</strong></a> · 📖 <a href="docs/cli-reference.md"><strong>Docs</strong></a>
</p>

</div>

---

<div align="center">
  <img src="assets/demo.gif" alt="Port Detective Interactive Demo" width="95%" />
</div>

---

## 📑 Table of Contents

- [⚡ Quickstart in 30 Seconds](#-quickstart-in-30-seconds)
- [🚀 Key Features](#-key-features)
- [📦 Installation Options](#-installation-options)
  - [Option 1: Automated Script (Recommended)](#option-1-automated-script-recommended)
  - [Option 2: Go Toolchain](#option-2-go-toolchain)
  - [Option 3: Pre-compiled GitHub Release Binaries](#option-3-pre-compiled-github-release-binaries)
- [🛠️ Command Showcase](#️-command-showcase)
  - [1. `pd check <port>`](#1-pd-check-port)
  - [2. `pd kill <port>`](#2-pd-kill-port)
  - [3. `pd scan <start-port>-<end-port>`](#3-pd-scan-start-port-end-port)
- [📐 System Architecture](#-system-architecture)
- [📚 Documentation Hub](#-documentation-hub)
- [🚦 POSIX Exit Codes](#-posix-exit-codes)
- [📄 License](#-license)

---

## ⚡ Quickstart in 30 Seconds

Tired of `Error: listen EADDRINUSE: address already in use :::8080`? Resolve it immediately with zero guesswork:

```bash
# 1. Install via automated script
curl -sSfL https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.sh | sh

# 2. Check which process is holding port 8080
pd check 8080

# 3. Kill it safely
pd kill 8080 --force
```

---

## 🚀 Key Features

| Capability | Highlights |
| :--- | :--- |
| **🚀 Native Cross-Platform** | Native strategy engines for Windows (`netstat`/`tasklist`), Linux (`lsof` with kernel `/proc/net` fallback), and macOS (`lsof`). |
| **⚡ Blazing Fast** | Zero runtime dependencies (no Node.js, Python, or JVM required). Compiles to a single lightweight native binary. |
| **🛡️ Safe by Default** | Interactive confirmation prompt (`[y/N]`) before killing, plus `--dry-run` inspection and OS critical PID protection. |
| **🎨 High-Contrast Terminal UI** | Clean colorized tables with distinct highlights for PID, Process Name, Port, and Protocol. |
| **🤖 Machine & CI/CD Ready** | Strict `--json` output across all commands for script pipelines and `jq` automation. |
| **🌐 Bilingual CLI** | Built-in instant switching between English (default) and Vietnamese (`--lang vi` or `PORT_DETECTIVE_LANG=vi`). |
| **🔍 High-Throughput Range Scanner** | Asynchronous goroutine worker pool scanning up to 5,000 ports in milliseconds. |

---

## 📦 Installation Options

### Option 1: Automated Script (Recommended)

Detects OS & CPU architecture, downloads the release binary, verifies checksums, and configures `PATH`:

<table>
<tr>
<td><b>Linux & macOS</b></td>
<td><b>Windows (PowerShell)</b></td>
</tr>
<tr>
<td>

```bash
curl -sSfL https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.sh | sh
```

</td>
<td>

```powershell
iwr -useb https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.ps1 | iex
```

</td>
</tr>
</table>

### Option 2: Go Toolchain
```bash
go install github.com/Khoa180806/Port_Detective/cmd/pd@latest
```

### Option 3: Pre-compiled GitHub Release Binaries
Download directly from [GitHub Releases](https://github.com/Khoa180806/Port_Detective/releases):
- **Windows:** `port-detective_Windows_x86_64.zip` / `arm64.zip`
- **Linux:** `port-detective_Linux_x86_64.tar.gz` / `arm64.tar.gz`
- **macOS:** `port-detective_Darwin_x86_64.tar.gz` / `arm64.tar.gz`

---

## 🛠️ Command Showcase

### 1. `pd check <port>`
Inspect processes bound to a port in human table format or structured JSON:

```bash
pd check 8080
```
```text
Port 8080 is occupied by:
  PID:      14280
  Process:  node.exe
  Command:  node server.js
  Protocol: tcp
```

```bash
pd check 8080 --json
```
```json
[
  {
    "pid": 14280,
    "name": "node.exe",
    "command": "node server.js",
    "port": 8080,
    "protocol": "tcp"
  }
]
```

### 2. `pd kill <port>`
Terminate offending processes interactively or with automation flags:

```bash
# Interactive mode with confirmation prompt:
pd kill 8080

# Force kill without prompt:
pd kill 8080 --force

# Preview process without terminating:
pd kill 8080 --dry-run
```

<div align="center">
  <img src="assets/check_demo.png" alt="Check and Kill Demo" width="90%" />
</div>

### 3. `pd scan <start-port>-<end-port>`
Concurrently scan a range of ports using a worker pool:

```bash
pd scan 3000-3005
```
```text
Scanning ports from 3000 to 3005...
Found 2 processes:
  Port:     3000
  PID:      18204
  Process:  node.exe
  Command:  node.exe
  Protocol: tcp
--------------------------------------------------
  Port:     3003
  PID:      9142
  Process:  docker-proxy
  Command:  docker-proxy
  Protocol: tcp
```

<div align="center">
  <img src="assets/scan_demo.png" alt="Scan Ports Demo" width="90%" />
</div>

---

## 📐 System Architecture

Port Detective leverages Go build tags and the Strategy Pattern for compile-time cross-platform dispatch:

<div align="center">

![Port Detective System Architecture](assets/architecture.svg)

</div>

> [!NOTE]
> For in-depth technical details on Linux kernel `/proc/net/tcp` parsing, Windows `netstat` strategy, and worker pool concurrency, see the [Architecture Guide](docs/architecture.md).

---

## 📚 Documentation Hub

Explore detailed documentation in the [`docs/`](docs/) directory:

| Document | Purpose |
| :--- | :--- |
| [📐 **System Architecture**](docs/architecture.md) | Component design, Go build tag strategy, and concurrency models |
| [📖 **CLI Reference Guide**](docs/cli-reference.md) | Comprehensive command flags, syntax specifications, and exit codes |
| [📦 **Installation Guide**](docs/installation.md) | Package managers, manual binary installs, permissions, and PATH setup |
| [🤝 **Contributing Guide**](docs/contributing.md) | Development setup, unit test execution, and pull request guidelines |

---

## 🚦 POSIX Exit Codes

| Code | Status | Meaning |
| :---: | :--- | :--- |
| `0` | **Success** | Process found, process terminated, or scan finished |
| `1` | **Port Free / Aborted** | Port is available, or interactive kill was declined |
| `2` | **Permission Denied** | Insufficient permissions (requires Administrator or `sudo`) |
| `3` | **Invalid Argument** | Malformed port number or range |
| `4` | **System Failure** | OS lookup mechanism error |

---

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
