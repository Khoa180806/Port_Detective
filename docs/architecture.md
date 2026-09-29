# 📐 Technical Architecture — Port Detective

This document provides an in-depth view of the architecture design, data models, and core components of **Port Detective (`pd`)**.

---

## 1. Design Principles

1. **Performance & Instant Startup:** Near-zero overhead. Compiles to a single lightweight native binary with no heavy runtimes (no Node.js, Python, or JVM required).
2. **Clean Cross-Platform Abstraction:** Utilizes compile-time **Go Build Tags** (Strategy Pattern) rather than bulky runtime branching, keeping each target platform's binary lean and purpose-built.
3. **Robust Separation of Concerns:** Distinct separation between the CLI user presentation layer (`cmd/`), operating system lookup & lifecycle controls (`internal/lookup/`), localization (`internal/i18n/`), and output formatters (`internal/output/`).
4. **Safety by Default:** Destructive operations (killing processes) are guarded with interactive confirmation prompts, `--dry-run` inspection, and system process protection.

---

## 2. System Component Diagram

```
┌────────────────────────────────────────────────────────┐
│                      cmd/pd/main.go                    │
└───────────────────────────┬────────────────────────────┘
                            │ Bootstrap
                            ▼
┌────────────────────────────────────────────────────────┐
│                         cmd/                           │
│   (root.go, check.go, kill.go, scan.go - Cobra CLI)    │
└──────────────┬──────────────────────────┬──────────────┘
               │                          │
    Lookup / Kill requests         Pass ProcessInfo
               │                          │
               ▼                          ▼
┌──────────────────────────┐   ┌──────────────────────────┐
│     internal/lookup/     │   │     internal/output/     │
│ (PortLookupStrategy)     │   │                          │
│                          │   │  - text.go (table, color)│
│ ├── windows.go           │   │  - json.go (machine JSON)│
│ ├── linux.go             │   └──────────────────────────┘
│ └── darwin.go            │                  ▲
└──────────────┬───────────┘                  │
               │ Return                       │
               ▼                              │
┌─────────────────────────────────────────────┴──────────┐
│                   internal/process/                    │
│             (ProcessInfo Struct Data Model)            │
└────────────────────────────────────────────────────────┘
```

---

## 3. Core Components

### 3.1. Data Model (`internal/process`)
`ProcessInfo` is the canonical data structure bridging lookup strategies and output renderers:

```go
type ProcessInfo struct {
    PID      int    `json:"pid"`
    Name     string `json:"name"`
    Command  string `json:"command"`
    Port     int    `json:"port"`
    Protocol string `json:"protocol"` // "tcp" | "udp"
}
```

### 3.2. Cross-Platform Strategy via Go Build Tags (`internal/lookup`)

Rather than relying on runtime conditional checks (`if runtime.GOOS == "windows"`), Port Detective implements compile-time build tags:

```go
type PortLookupStrategy interface {
    FindProcessByPort(port int) ([]process.ProcessInfo, error)
    KillProcess(pid int) error
}
```

- **Windows (`windows.go` with `//go:build windows`):**
  - Executes `netstat -ano` to extract listening PIDs for the target port.
  - Queries `tasklist /FO CSV /FI "PID eq <pid>"` to resolve the process name and command.
  - Uses `taskkill /F /PID <pid>` for termination.

- **Linux (`linux.go` with `//go:build linux`):**
  - Primary strategy: Leverages `lsof -i :<port> -sTCP:LISTEN -P -n`.
  - Kernel procfs fallback: If `lsof` is not installed (e.g. minimal Docker containers), it parses `/proc/net/tcp`, `/proc/net/tcp6`, `/proc/net/udp`, and reads `/proc/<pid>/fd/` socket inodes directly.
  - Terminates processes cleanly using `kill -9 <pid>`.

- **macOS (`darwin.go` with `//go:build darwin`):**
  - Utilizes native macOS `lsof -i :<port> -P -n`.
  - Uses POSIX kill for process termination.

### 3.3. Output Layer (`internal/output`)

Supports two output modes:
1. **Human-Readable Text Format:**
   - Visual terminal formatting using `github.com/fatih/color`.
   - Colored highlights for PIDs, Process Names, and Ports.
2. **Machine-Readable JSON:**
   - Strict JSON serialization activated via `--json`.
   - Designed for headless CI/CD pipelines, shell scripts, and parsing with tools like `jq`.

### 3.4. Localization Engine (`internal/i18n`)

Port Detective includes a lightweight key-value translation subsystem:
- Built-in English (`en`) and Vietnamese (`vi`) language packs.
- Language resolution hierarchy:
  1. Explicit CLI flag: `--lang vi` / `--lang en`
  2. Environment variable: `PORT_DETECTIVE_LANG=vi`
  3. System locale detection: fallback to English if unsupported.

---

## 4. Edge Cases & Resilience

| Scenario | Behavior & Resolution |
|---|---|
| **Port is available / Unoccupied** | Returns an informative message; exits with code `1` (Port free / no match). |
| **Multiple processes bound to 1 port** | Displays all matching processes (`SO_REUSEPORT`). In interactive kill mode, lists each process for confirmation. |
| **Permission Denied** | Detects access restrictions and suggests elevating privileges (`Run as Administrator` on Windows or `sudo` on Linux/macOS). Exits with code `2`. |
| **Invalid Port Input** | Pre-execution validation verifies port bounds (`1`–`65535`). Prevents invalid system calls. |
| **Target Process Exited Before Kill** | Gracefully handles non-existent PIDs without crashing. |
