# 📦 Installation Guide — Port Detective

<div align="center">

[![Port Detective Documentation](https://img.shields.io/badge/docs-installation-blue?style=flat-square)](../README.md)
[![Platforms](https://img.shields.io/badge/Platforms-macOS%20%7C%20Linux%20%7C%20Windows-informational?style=flat-square)](../README.md)
[![Release](https://img.shields.io/badge/release-latest-brightgreen?style=flat-square)](https://github.com/Khoa180806/Port_Detective/releases/latest)

</div>

This guide covers all available methods to install **Port Detective (`pd`)** across Windows, macOS, and Linux.

---

## ⚡ Method 1: Automated Script (Recommended)

The automated install scripts detect your operating system and CPU architecture, download the corresponding release binary, verify checksums, and configure your system `PATH`.

### Linux & macOS

Run the following command in your terminal:

```bash
curl -sSfL https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.sh | sh
```

- **Target Destination:** `~/.local/bin` (or `/usr/local/bin` if executed with `sudo`).
- **Supported Architectures:** `x86_64` (Intel/AMD), `arm64` (Apple Silicon M1/M2/M3/M4, ARM64 servers).

### Windows (PowerShell)

Open PowerShell and execute:

```powershell
iwr -useb https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.ps1 | iex
```

- **Target Destination:** `$HOME\AppData\Local\PortDetective\bin\pd.exe`
- **Environment Setup:** Automatically adds the destination directory to your User `PATH` environment variable.

---

## 🐹 Method 2: Via Go Toolchain

If Go (version 1.21 or higher) is installed on your workstation:

```bash
go install github.com/Khoa180806/Port_Detective/cmd/pd@latest
```

> [!NOTE]
> Ensure that your `$GOPATH/bin` (Linux/macOS) or `%USERPROFILE%\go\bin` (Windows) directory is included in your system's `PATH`.

---

## 📥 Method 3: Pre-compiled Binaries (GitHub Releases)

Standalone binaries are published on the official [GitHub Releases](https://github.com/Khoa180806/Port_Detective/releases/latest) page for every release tag:

| Platform | Architecture | Archive Package |
| :--- | :--- | :--- |
| **Windows** | `x86_64` (64-bit Intel/AMD)<br>`arm64` (Qualcomm Snapdragon / Windows ARM) | `port-detective_Windows_x86_64.zip`<br>`port-detective_Windows_arm64.zip` |
| **Linux** | `x86_64` (Standard servers/desktops)<br>`arm64` (Raspberry Pi, AWS Graviton) | `port-detective_Linux_x86_64.tar.gz`<br>`port-detective_Linux_arm64.tar.gz` |
| **macOS** | `x86_64` (Intel Mac)<br>`arm64` (Apple Silicon M-Series) | `port-detective_Darwin_x86_64.tar.gz`<br>`port-detective_Darwin_arm64.tar.gz` |

### Manual Installation Steps:
1. Download the archive for your operating system.
2. Extract the archive contents:
   ```bash
   tar -xzf port-detective_Darwin_arm64.tar.gz
   ```
3. Move the binary into your executable path:
   ```bash
   sudo mv pd /usr/local/bin/pd
   chmod +x /usr/local/bin/pd
   ```

---

## 🔨 Method 4: Build from Source

To compile the latest commit from the repository:

```bash
# Clone the repository
git clone https://github.com/Khoa180806/Port_Detective.git
cd Port_Detective

# Build stripped production binary
go build -ldflags="-s -w" -o pd ./cmd/pd

# Move to system path (Linux / macOS)
sudo mv pd /usr/local/bin/

# On Windows:
# move pd.exe C:\Windows\System32\
```

---

## 🔍 Verifying Your Installation

Confirm that `pd` is recognized by your terminal:

```bash
$ pd --version
pd version 1.0.0

$ pd --help
Port Detective - An instant, cross-platform port & process investigation CLI tool
```

---

## 🛡️ Permission & Privilege Guidelines

Port Detective interacts with low-level OS socket tables and process management facilities:

- **Windows:** Standard users can inspect unprivileged ports. Terminating elevated background services requires launching PowerShell or Terminal as **Administrator**.
- **Linux:** Standard users can inspect processes owned by their UID. Querying or killing system demons (e.g., Docker, Nginx, PostgreSQL, systemd services) requires `sudo`:
  ```bash
  sudo pd check 80
  sudo pd kill 80 --force
  ```
- **macOS:** Standard users can inspect their local processes. For root-owned services, prepend `sudo pd check <port>`.

---

## 🧭 Navigation

- [📐 System Architecture](./architecture.md)
- [📖 CLI Reference Guide](./cli-reference.md)
- [🤝 Contributing Guide](./contributing.md)
- [🏠 Project Root & README](../README.md)
