# 📖 CLI Reference — Port Detective

Comprehensive command-line usage and parameter reference for **Port Detective (`pd`)**.

---

## 📌 Global Synopsis

```bash
pd [command] [flags]
```

### Global Flags

| Flag | Shorthand | Type | Default | Description |
|---|---|---|---|---|
| `--lang` | `-l` | string | `"en"` | Language code for CLI output (`en` for English, `vi` for Vietnamese). |
| `--help` | `-h` | bool | `false` | Display help and usage information for any command. |
| `--version` | `-v` | bool | `false` | Print the current version of Port Detective. |

> [!TIP]
> You can also configure the language globally using the `PORT_DETECTIVE_LANG` environment variable:
> ```bash
> export PORT_DETECTIVE_LANG=vi    # Linux/macOS
> $env:PORT_DETECTIVE_LANG="vi"    # PowerShell
> ```

---

## 🔍 Commands

### 1. `pd check <port>`

Inspect which process or service is currently occupying a specific network port.

```bash
pd check <port> [flags]
```

#### Flags
| Flag | Shorthand | Type | Default | Description |
|---|---|---|---|---|
| `--json` | `-j` | bool | `false` | Output results in valid JSON format. |

#### Examples

**Human-readable output:**
```bash
$ pd check 8080
Port 8080 is occupied by:
  PID:      14280
  Process:  node.exe
  Command:  node server.js
  Protocol: tcp
```

**JSON output (for automation and scripts):**
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

When a port is unoccupied:
```bash
$ pd check 9999
No process found running on port 9999.
```

---

### 2. `pd kill <port>`

Terminate the process currently occupying the specified port.

```bash
pd kill <port> [flags]
```

#### Flags
| Flag | Shorthand | Type | Default | Description |
|---|---|---|---|---|
| `--force` | `-f` | bool | `false` | Bypass interactive confirmation prompt and kill immediately. |
| `--dry-run` | `-d` | bool | `false` | Simulate the kill action without terminating any process. |

#### Safety Safeguards

1. **Interactive Prompt:** By default, `pd kill` identifies all processes occupying the port, prints their details, and asks for confirmation (`Are you sure you want to KILL all processes above? [y/N]: `) before sending a kill signal.
2. **Dry Run Mode:** With `--dry-run`, you can inspect what would be killed without making any changes to running services.

#### Examples

**Interactive mode:**
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

**Non-interactive (force kill):**
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

**Dry run:**
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

Concurrently inspect a port range using an asynchronous worker pool.

```bash
pd scan <start-port>-<end-port> [flags]
```

#### Arguments
- `<range>`: Formatted as `start-end` where both values are between `1` and `65535` (e.g. `3000-3010`).

#### Flags
| Flag | Shorthand | Type | Default | Description |
|---|---|---|---|---|
| `--json` | `-j` | bool | `false` | Output active ports in JSON format. |

#### Examples

**Human-readable scan:**
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

**JSON scan output:**
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

## 🚦 Exit Codes

Port Detective returns standard POSIX exit codes suitable for CI/CD assertions:

| Code | Status | Meaning |
|:---:|---|---|
| `0` | **Success** | Process located (`check`), process terminated (`kill`), or scan completed. |
| `1` | **Port Free / Action Aborted** | Port is free/available, or interactive kill was declined by user. |
| `2` | **System Error / Bad Input** | Port out of range (`1-65535`), invalid syntax, or missing root/Admin permissions. |

**Example in a shell script:**
```bash
pd check 3000 > /dev/null 2>&1
if [ $? -eq 0 ]; then
  echo "Port 3000 is occupied! Freeing port..."
  pd kill 3000 --force
fi
```
