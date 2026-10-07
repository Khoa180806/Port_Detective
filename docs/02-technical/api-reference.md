# CLI Command & Public Interface Reference

This document defines the public command-line interface, syntax specifications, machine JSON schemas, and POSIX exit codes of Port Detective.

---

## 1. Global Commands & Syntax

```bash
pd [command] [flags]
```

### Global Flags

| Flag | Shorthand | Type | Default | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--lang` | `-l` | `string` | `"en"` | Output language code: `en` (English) or `vi` (Vietnamese). |
| `--help` | `-h` | `bool` | `false` | Display command usage documentation. |
| `--version` | `-v` | `bool` | `false` | Output compiled binary version. |

*(source: [cmd/root.go](file:///d:/Project/PortDetective/cmd/root.go))*

---

## 2. Command Specifications

### 2.1 `pd check <port>`
Inspects active listeners on a specific network port.

- **Argument**: `<port>` (Integer between `1` and `65535`).
- **Flags**:
  - `--json` (`-j`): Serialize results to structured JSON.
- **Example Usage**:
  ```bash
  pd check 8080
  pd check 8080 --json
  ```
*(source: [cmd/check.go](file:///d:/Project/PortDetective/cmd/check.go))*

### 2.2 `pd kill <port>`
Terminates processes bound to a designated network port.

- **Argument**: `<port>` (Integer between `1` and `65535`).
- **Flags**:
  - `--force` (`-f`): Bypass interactive confirmation prompt (`[y/N]`).
  - `--dry-run` (`-d`): Preview target processes without killing.
- **Example Usage**:
  ```bash
  pd kill 8080
  pd kill 8080 --force
  pd kill 8080 --dry-run
  ```
*(source: [cmd/kill.go](file:///d:/Project/PortDetective/cmd/kill.go))*

### 2.3 `pd scan <range>`
Concurrently scans a contiguous range of ports.

- **Argument**: `<range>` (Formatted as `<start>-<end>`, e.g., `3000-3005`).
- **Flags**:
  - `--json` (`-j`): Output scan results in JSON format.
- **Example Usage**:
  ```bash
  pd scan 3000-3005
  pd scan 3000-3005 --json
  ```
*(source: [cmd/scan.go](file:///d:/Project/PortDetective/cmd/scan.go))*

---

## 3. Machine-Readable JSON Schema

When executed with `--json`, output is emitted to `stdout` as a JSON array of `ProcessInfo` objects:

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

### JSON Property Types

| Field | Type | Description |
| :--- | :--- | :--- |
| `pid` | `integer` | Process Identifier assigned by the host operating system. |
| `name` | `string` | Binary / Executable name (e.g. `node.exe`, `docker-proxy`). |
| `command` | `string` | Full execution command or resolved executable path. |
| `port` | `integer` | Target port number (`1` - `65535`). |
| `protocol` | `string` | Socket transport layer protocol (`"tcp"` or `"udp"`). |

*(source: [internal/process/process.go](file:///d:/Project/PortDetective/internal/process/process.go), [internal/output/json.go](file:///d:/Project/PortDetective/internal/output/json.go))*

---

## 4. POSIX Exit Codes

Port Detective returns deterministic exit codes for shell scripting and CI/CD pipelines:

| Exit Code | Classification | Trigger Condition |
| :---: | :--- | :--- |
| `0` | **Success** | Process identified (`check`), process killed (`kill`), or range scan completed. |
| `1` | **Port Unoccupied / Aborted** | Port is available (no process bound), or interactive kill aborted with `n`. |
| `2` | **Permission Denied** | Insufficient user elevation (requires root / Administrator privileges). |
| `3` | **Invalid Arguments** | Port out of range (`1-65535`), or malformed range format. |
| `4` | **System Failure** | Operating system command execution or procfs parse failure. |

*(source: [cmd/check.go](file:///d:/Project/PortDetective/cmd/check.go), [cmd/kill.go](file:///d:/Project/PortDetective/cmd/kill.go), [internal/lookup/errors.go](file:///d:/Project/PortDetective/internal/lookup/errors.go))*
