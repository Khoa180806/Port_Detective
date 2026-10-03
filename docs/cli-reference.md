# 📖 CLI Reference — Port Detective

<div align="center">

[![Port Detective Documentation](https://img.shields.io/badge/docs-cli--reference-blue?style=flat-square)](../README.md)
[![Cobra CLI](https://img.shields.io/badge/CLI-spf13%2Fcobra-brightgreen?style=flat-square)](https://github.com/spf13/cobra)
[![POSIX Compliant](https://img.shields.io/badge/POSIX-Exit%20Codes-orange?style=flat-square)](../README.md)

</div>

Comprehensive syntax, arguments, flag details, and operational examples for the **Port Detective (`pd`)** command-line interface.

---

## 📌 Global Synopsis

```bash
pd [command] [flags]
```

### Global Flags

| Flag | Shorthand | Type | Default | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--lang` | `-l` | `string` | `"en"` | Language code for CLI output (`en` for English, `vi` for Vietnamese). |
| `--help` | `-h` | `bool` | `false` | Display help, usage synopsis, and subcommands. |
| `--version` | `-v` | `bool` | `false` | Print the current compiled version of Port Detective. |

> [!TIP]
> **Persistent Language Setting:** You can configure your language preference globally using the `PORT_DETECTIVE_LANG` environment variable:
> ```bash
> export PORT_DETECTIVE_LANG=vi    # Linux / macOS (bash / zsh)
> $env:PORT_DETECTIVE_LANG="vi"    # Windows PowerShell
> ```

---

## 🔍 Commands Breakdown

### 1. `pd check <port>`

Inspect which process or service is currently occupying a specific network port.

```bash
pd check <port> [flags]
```

#### Arguments
- `<port>` *(required)*: Integer between `1` and `65535`.

#### Flags
| Flag | Shorthand | Type | Default | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--json` | `-j` | `bool` | `false` | Format output as machine-readable JSON array. |

#### Usage Examples

**Human-Readable Table Output:**
```bash
$ pd check 8080
Port 8080 is occupied by:
  PID:      14280
  Process:  node.exe
  Command:  node server.js
  Protocol: tcp
```

**JSON Output (Ideal for scripts and `jq` pipelines):**
```bash
$ pd check 8080 --json
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

**When the port is free:**
```bash
$ pd check 9999
No process found running on port 9999.
```

---

### 2. `pd kill <port>`

Safely terminate the process occupying the specified network port.

```bash
pd kill <port> [flags]
```

#### Arguments
- `<port>` *(required)*: Integer between `1` and `65535`.

#### Flags
| Flag | Shorthand | Type | Default | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--force` | `-f` | `bool` | `false` | Terminate process immediately without interactive confirmation. |
| `--dry-run` | `-d` | `bool` | `false` | Simulate process termination and print preview without killing. |

#### Safety Safeguards
1. **Interactive Prompt:** By default, `pd kill` displays all process metadata and prompts for explicit confirmation (`Are you sure you want to KILL all processes above? [y/N]: `) before issuing any termination signal.
2. **Dry Run Simulation:** Adding `--dry-run` allows you to test operational impact in staging and production without modifying system state.
3. **Critical Process Shield:** Prevents accidental termination of core OS system PIDs (`0`, `4`, or `launchd`).

#### Usage Examples

**Interactive Mode:**
```bash
$ pd kill 8080
WARNING: Detected process(es) occupying port:
Port 8080 is occupied by:
  PID:      14280
  Process:  node.exe
  Command:  node server.js
  Protocol: tcp

Are you sure you want to KILL all processes above? [y/N]: y
Terminating PID 14280 (node.exe)... SUCCESS
```

**Non-Interactive / Automation (`--force`):**
```bash
$ pd kill 8080 --force
WARNING: Detected process(es) occupying port:
Port 8080 is occupied by:
  PID:      14280
  Process:  node.exe
  Command:  node server.js
  Protocol: tcp

Terminating PID 14280 (node.exe)... SUCCESS
```

**Dry-Run Simulation:**
```bash
$ pd kill 8080 --dry-run
WARNING: Detected process(es) occupying port:
Port 8080 is occupied by:
  PID:      14280
  Process:  node.exe
  Command:  node server.js
  Protocol: tcp

[DRY RUN] No processes will be terminated.
```

---

### 3. `pd scan <range>`

Concurrently inspect an inclusive range of ports using a bounded goroutine worker pool.

```bash
pd scan <start-port>-<end-port> [flags]
```

#### Arguments
- `<range>` *(required)*: Formatted as `<start>-<end>` where `1 <= start <= end <= 65535` (e.g. `3000-3005`).

#### Flags
| Flag | Shorthand | Type | Default | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--json` | `-j` | `bool` | `false` | Output discovered active ports in JSON format. |

#### Usage Examples

**Human-Readable Multi-Port Scan:**
```bash
$ pd scan 3000-3005
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

**JSON Output for Port Range:**
```bash
$ pd scan 3000-3005 --json
[
  {
    "pid": 18204,
    "name": "node.exe",
    "command": "node.exe",
    "port": 3000,
    "protocol": "tcp"
  },
  {
    "pid": 9142,
    "name": "docker-proxy",
    "command": "docker-proxy",
    "port": 3003,
    "protocol": "tcp"
  }
]
```

---

## 🚦 POSIX Exit Codes

Port Detective returns standard exit codes designed for deterministic CI/CD and shell scripting integration:

| Exit Code | Classification | Meaning | Scripting Action |
| :---: | :--- | :--- | :--- |
| `0` | **Success** | Process located (`check`), process terminated (`kill`), or scan completed | Proceed safely |
| `1` | **Port Available / Aborted** | Port is free/unoccupied, or interactive kill was aborted by user (`n`) | No action required |
| `2` | **Permission Denied** | Insufficient permissions to inspect or kill target process | Request root / Administrator elevation |
| `3` | **Invalid Arguments** | Port out of range (`1-65535`), malformed port range format | Correct command input |
| `4` | **System Failure** | OS lookup mechanism error | Check system logs |

### Shell Script Automation Pattern

```bash
#!/usr/bin/env bash
set -e

PORT=3000
echo "Checking if port $PORT is free..."

# pd returns 0 if occupied, 1 if free
if pd check $PORT > /dev/null 2>&1; then
    echo "⚠️ Port $PORT is occupied. Freeing port automatically..."
    pd kill $PORT --force
else
    echo "✅ Port $PORT is already free!"
fi

echo "Starting dev server..."
npm run dev
```

---

## 🧭 Navigation

- [📐 System Architecture](./architecture.md)
- [📦 Installation Guide](./installation.md)
- [🤝 Contributing Guide](./contributing.md)
- [🏠 Project Root & README](../README.md)
