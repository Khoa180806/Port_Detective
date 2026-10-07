# Business Flow

This document details the operational business workflows and sequence diagrams for Port Detective's primary CLI commands.

---

## 1. Port Investigation Flow (`pd check <port>`)

Executes port discovery, validates port bounds, invokes the platform strategy, and formats the result.

```mermaid
sequenceDiagram
    autonumber
    actor Dev as Developer / Script
    participant Root as Root & i18n Hook (cmd/root.go)
    participant Check as CheckCmd (cmd/check.go)
    participant Dispatch as Dispatcher (internal/lookup)
    participant OS as OS Strategy (windows/linux/darwin)
    participant Output as Formatter (internal/output)

    Dev->>Root: pd check 8080 [--json] [--lang vi]
    Root->>Root: Set language dictionary
    Root->>Check: RunE(args)
    Check->>Check: Validate port bounds (1 - 65535)
    Check->>Dispatch: FindProcessByPort(8080)
    Dispatch->>OS: Execute native lookup (netstat / lsof / procfs)
    OS-->>Dispatch: []ProcessInfo
    Dispatch-->>Check: Process records / nil
    alt Port is occupied
        Check->>Output: FormatText(processes) / FormatJSON(processes)
        Output-->>Dev: Render results (exit code 0)
    else Port is available
        Check-->>Dev: Print unoccupied message (exit code 1)
    end
```
*(source: [cmd/check.go](file:///d:/Project/PortDetective/cmd/check.go), [internal/lookup/lookup.go](file:///d:/Project/PortDetective/internal/lookup/lookup.go))*

---

## 2. Process Termination Flow (`pd kill <port>`)

Terminates target processes with safety verification, dry-run simulation, and PID protection.

```mermaid
sequenceDiagram
    autonumber
    actor Dev as Developer / CI Script
    participant Kill as KillCmd (cmd/kill.go)
    participant Dispatch as Dispatcher (internal/lookup)
    participant OS as OS Strategy (windows/linux/darwin)

    Dev->>Kill: pd kill 8080 [--force] [--dry-run]
    Kill->>Dispatch: FindProcessByPort(8080)
    Dispatch->>OS: Locate active processes
    OS-->>Dispatch: []ProcessInfo
    Dispatch-->>Kill: Return process list
    
    alt --dry-run enabled
        Kill-->>Dev: Display target processes & [DRY RUN] notice (exit code 0)
    else Interactive Mode (no --force)
        Kill-->>Dev: Prompt confirmation [y/N]
        Dev->>Kill: Response ("y" / "n")
        alt User confirms ("y")
            Kill->>Dispatch: KillProcess(PID)
            Dispatch->>OS: Signal termination (taskkill / SIGKILL)
            OS-->>Kill: Success / Error
            Kill-->>Dev: Print termination summary (exit code 0)
        else User aborts ("n")
            Kill-->>Dev: Print cancellation notice (exit code 1)
        end
    else Automated Mode (--force)
        Kill->>Dispatch: KillProcess(PID)
        Dispatch->>OS: Terminate process immediately
        OS-->>Kill: Success
        Kill-->>Dev: Print success message (exit code 0)
    end
```
*(source: [cmd/kill.go](file:///d:/Project/PortDetective/cmd/kill.go), [internal/lookup/lookup.go](file:///d:/Project/PortDetective/internal/lookup/lookup.go))*

---

## 3. Concurrent Range Scan Flow (`pd scan <start>-<end>`)

Orchestrates multi-threaded port inspection across a designated range.

```mermaid
sequenceDiagram
    autonumber
    actor Dev as Developer / Tool
    participant Scan as ScanCmd (cmd/scan.go)
    participant Pool as Worker Pool (Goroutines)
    participant Net as TCP Socket Dialer
    participant Dispatch as Strategy Dispatcher

    Dev->>Scan: pd scan 3000-3010 [--json]
    Scan->>Scan: Parse & validate range syntax
    Scan->>Pool: Spawn 100 worker goroutines
    loop For each port in range
        Scan->>Pool: Send port to buffered channel
        Pool->>Net: DialTimeout(tcp, port, 50ms)
        alt Port open
            Pool->>Dispatch: FindProcessByPort(port)
            Dispatch-->>Pool: Return ProcessInfo
            Pool->>Scan: Append to thread-safe result list
        else Port closed
            Pool-->>Pool: Discard
        end
    end
    Scan->>Dev: Output aggregated table / JSON array (exit code 0)
```
*(source: [cmd/scan.go](file:///d:/Project/PortDetective/cmd/scan.go))*
