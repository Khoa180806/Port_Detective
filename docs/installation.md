# 📦 Installation Guide — Port Detective

This guide covers all available methods to install **Port Detective (`pd`)** on Windows, macOS, and Linux.

---

## ⚡ Method 1: Automated Script (Recommended)

The automated install scripts download the correct pre-built binary for your operating system and CPU architecture, place it in an appropriate user directory, and ensure it is available on your `PATH`.

### Linux & macOS

Run the following in your terminal:
```bash
curl -sSfL https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.sh | sh
```

- **Target location:** `~/.local/bin` (or `/usr/local/bin` if run with sudo)
- **Supported Architectures:** x86_64, arm64 (Apple Silicon, Raspberry Pi)

### Windows (PowerShell)

Open PowerShell and run:
```powershell
iwr -useb https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.ps1 | iex
```

- **Target location:** `$HOME\AppData\Local\PortDetective\bin\pd.exe`
- **Environment:** Automatically persists into your User `PATH` environment variable.

---

## 🐹 Method 2: Via Go Toolchain

If you have Go (1.21+) installed on your machine:

```bash
go install github.com/Khoa180806/Port_Detective/cmd/pd@latest
```

> [!NOTE]
> Make sure `$GOPATH/bin` (or `%USERPROFILE%\go\bin` on Windows) is in your system's `PATH`.

---

## 📥 Method 3: Pre-compiled Binaries (GitHub Releases)

You can download standalone binaries directly from the [GitHub Releases](https://github.com/Khoa180806/Port_Detective/releases/latest) page.

1. Download the archive corresponding to your operating system:
   - **Windows:** `port-detective_Windows_x86_64.zip` or `port-detective_Windows_arm64.zip`
   - **Linux:** `port-detective_Linux_x86_64.tar.gz` or `port-detective_Linux_arm64.tar.gz`
   - **macOS:** `port-detective_Darwin_x86_64.tar.gz` or `port-detective_Darwin_arm64.tar.gz`
2. Extract the archive.
3. Move the binary (`pd` or `pd.exe`) to a directory in your `PATH` (e.g. `/usr/local/bin` or `C:\Windows\System32`).

---

## 🔨 Method 4: Build from Source

To compile the latest bleeding-edge build from source:

```bash
# Clone the repository
git clone https://github.com/Khoa180806/Port_Detective.git
cd Port_Detective

# Build binary
go build -ldflags="-s -w" -o pd ./cmd/pd

# Move to system path (Linux/macOS)
sudo mv pd /usr/local/bin/

# On Windows:
# move pd.exe C:\Windows\
```

---

## 🔍 Verifying Installation

Verify that `pd` is accessible and working correctly:

```bash
pd --version
pd --help
```

---

## 🛡️ Permission Requirements

Port Detective requires operating system permissions to query network sockets and terminate processes:

- **Windows:** Standard user accounts can inspect most ports. However, to terminate elevated or system processes, run your terminal (PowerShell / Command Prompt) as **Administrator**.
- **Linux:** Standard accounts can inspect sockets owned by the current user. To query or kill system services (e.g. Nginx, Docker, PostgreSQL), prepend commands with `sudo`:
  ```bash
  sudo pd check 80
  sudo pd kill 80 --force
  ```
- **macOS:** Standard users can inspect their own processes. For daemon/root processes, run with `sudo pd check <port>`.
