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
Port 9999 is free.
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

1. **Interactive Prompt:** By default, `pd kill` identifies the process and asks for confirmation (`Are you sure? [y/N]`) before sending a kill signal.
2. **Dry Run Mode:** With `--dry-run`, you can inspect what would be killed without making any changes to running services.

#### Examples

**Interactive mode:**
```bash
$ pd kill 8080
Port 8080 is occupied by process 'node.exe' (PID: 14280).
Do you want to kill this process? [y/N]: y
Successfully terminated process 'node.exe' (PID: 14280).
```

**Non-interactive (force kill):**
```bash
$ pd kill 8080 --force
Successfully terminated process 'node.exe' (PID: 14280).
```

**Dry run:**
```bash
$ pd kill 8080 --dry-run
[DRY-RUN] Would terminate process 'node.exe' (PID: 14280) on port 8080. No action taken.
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
Scanning ports 3000 to 3005...

Port  Status  PID    Process
3000  BUSY    18204  node.exe
3001  FREE    -      -
3002  FREE    -      -
3003  BUSY    9142   docker-proxy
3004  FREE    -      -
3005  FREE    -      -
```

**JSON scan output:**
```bash
$ pd scan 3000-3002 --json
[
  {
    "port": 3000,
    "status": "BUSY",
    "process": {
      "pid": 18204,
      "name": "node.exe",
      "command": "node.exe",
      "port": 3000,
      "protocol": "tcp"
    }
  },
  {
    "port": 3001,
    "status": "FREE",
    "process": null
  },
  {
    "port": 3002,
    "status": "FREE",
    "process": null
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
