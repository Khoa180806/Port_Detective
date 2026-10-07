# Known Issues & Technical Limitations

This document tracks identified technical limitations, edge cases, and architectural debt in Port Detective based on actual codebase constraints.

---

## 1. Operating System Permission Boundaries

### Windows Administrator Privileges
- **Issue**: Standard (unprivileged) Windows users can query sockets via `netstat -ano`, but terminating elevated or system services via `taskkill /F /PID <pid>` fails with `Access is denied`.
- **Handling in Code**: Handled cleanly with POSIX Exit Code `2` and an informative hint directing users to "Run terminal as Administrator".
*(source: [internal/lookup/windows.go](file:///d:/Project/PortDetective/internal/lookup/windows.go), [internal/lookup/errors.go](file:///d:/Project/PortDetective/internal/lookup/errors.go))*

### Linux / macOS Root Privileges
- **Issue**: Standard users can only inspect socket inodes and terminate processes owned by their own UID. Querying or killing root-owned system daemons (e.g., Docker, Nginx, PostgreSQL, systemd) returns permission denied.
- **Handling in Code**: Yields Exit Code `2` with a hint suggesting `sudo pd ...`.
*(source: [internal/lookup/linux.go](file:///d:/Project/PortDetective/internal/lookup/linux.go), [internal/lookup/darwin.go](file:///d:/Project/PortDetective/internal/lookup/darwin.go))*

---

## 2. Platform-Specific Query Constraints

### Windows `tasklist` Name Truncation
- **Issue**: On older Windows builds or when running `tasklist` with specific buffer sizes, executable names longer than 25 characters can sometimes be truncated.
- **Handling in Code**: Uses `/FO CSV` to obtain raw unpadded string tokens, minimizing truncation.
*(source: [internal/lookup/windows.go](file:///d:/Project/PortDetective/internal/lookup/windows.go))*

### Stripped Linux Container procfs Access
- **Issue**: In heavily sandboxed container environments where `/proc` is masked or mounted with strict `hidepid=2` restrictions, Port Detective cannot correlate socket inodes to process PIDs without root access.
- **Handling in Code**: Returns an explicit `PermissionDeniedError` instead of crashing.
*(source: [internal/lookup/linux.go](file:///d:/Project/PortDetective/internal/lookup/linux.go))*

---

## 3. Process Tree Cascades

- **Limitation**: `pd kill <port>` terminates the specific process bound to the socket. When server processes are managed by parent wrappers (e.g., `npm run dev` spawning `node`), terminating the child process may cause the parent wrapper to exit or hang.
- **Status**: Tracked in [Product Roadmap](../03-product/roadmap.md) under Milestone 2 (`--tree` flag for recursive process tree termination).
