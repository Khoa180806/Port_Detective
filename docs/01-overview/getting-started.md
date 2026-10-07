# Getting Started

This guide details prerequisites, local installation, build commands, and test verification for Port Detective.

---

## 1. Prerequisites

- **Go**: Version `1.21` or later (`1.22+` recommended).
- **Git**: Installed and configured.
- **Node.js** *(Optional, only needed for the `web/` frontend)*: Version `18.17+` or `20+`.

---

## 2. Installation & Quick Setup

### Method A: Automated One-Line Script (End Users)

**Linux & macOS:**
```bash
curl -sSfL https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.sh | sh
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.ps1 | iex
```

### Method B: Via Go Toolchain
```bash
go install github.com/Khoa180806/Port_Detective/cmd/pd@latest
```
*(Ensure `$GOPATH/bin` or `%USERPROFILE%\go\bin` is in your system `PATH`)*.

---

## 3. Building From Source (Developers)

Clone the repository and compile the native binary:

```bash
# Clone the repository
git clone https://github.com/Khoa180806/Port_Detective.git
cd Port_Detective

# Download dependencies
go mod download

# Compile production binary with stripped debug symbols
go build -ldflags="-s -w" -o pd ./cmd/pd

# On Windows:
# go build -ldflags="-s -w" -o pd.exe ./cmd/pd
```

---

## 4. Running the CLI Locally

Test basic execution:

```bash
# Print help
./pd --help

# Check an active port
./pd check 8080

# Check port with JSON output
./pd check 8080 --json

# Run interactive kill
./pd kill 8080

# Run non-destructive dry-run
./pd kill 8080 --dry-run

# Scan port range
./pd scan 3000-3010
```

---

## 5. Running the Automated Test Suite

Run unit and integration tests across all Go packages:

```bash
# Run all tests with verbose output
go test -v ./...

# Run static analysis
go vet ./...
```

---

## 6. Running the Web Showcase (`web/`)

If contributing to the landing page:

```bash
cd web

# Install dependencies
npm install

# Start local Next.js development server
npm run dev

# Run linter and static production build
npm run lint
npm run build
```
*(source: [scripts/install.sh](file:///d:/Project/PortDetective/scripts/install.sh), [scripts/install.ps1](file:///d:/Project/PortDetective/scripts/install.ps1), [web/package.json](file:///d:/Project/PortDetective/web/package.json))*
