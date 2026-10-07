# Technical Architecture Decisions (ADR)

This document records the foundational architectural decisions, evaluation contexts, considered alternatives, and consequences for Port Detective.

---

## ADR-001: Compile-Time Cross-Platform Strategy via Go Build Tags

### Context
Port querying depends on low-level OS utilities (`netstat`/`tasklist` on Windows, `lsof`/`/proc/net` on Linux, native `lsof` on macOS). Using runtime checks (`if runtime.GOOS == "windows"`) forces every binary to compile dead code from other platforms and can introduce platform-specific compilation errors (e.g. referencing Windows DLL syscalls on Linux).

### Decision
Implement the Strategy Pattern using compile-time Go build tags (`//go:build windows`, `//go:build linux`, `//go:build darwin`). Define a common `PortLookupStrategy` interface in `internal/lookup/lookup.go` that each file implements.

### Alternatives Considered
- **Runtime OS Branching**: Single file with `switch runtime.GOOS`. Rejected due to cross-compilation friction and larger binary size.
- **External CGO Wrapper**: Binding to C socket libraries. Rejected because it breaks portable zero-dependency static cross-compilation.

### Consequences
- **Positive**: Clean separation of platform code; zero runtime branching overhead; binaries compile cleanly with `GOOS=<target> go build`.
- **Negative**: Platform tests must be run on matching host OS or mocked interfaces.
*(source: [internal/lookup/lookup.go](file:///d:/Project/PortDetective/internal/lookup/lookup.go))*

---

## ADR-002: Zero External Runtime Dependency (Native Go Binary)

### Context
Developers need instant feedback when resolving port conflicts (`EADDRINUSE`). Scripting solutions written in Node.js, Python, or shell scripts require pre-installed runtimes, have slow cold-start times (>200ms), or behave inconsistently across shell environments.

### Decision
Build Port Detective as a compiled native Go 1.22+ binary with sub-10ms startup latency and zero runtime dependencies.

### Alternatives Considered
- **Node.js CLI (`npx kill-port`)**: Requires Node.js installed, startup latency is 300-600ms, and cannot inspect stripped Docker containers.
- **Pure Shell Scripts (`.sh` / `.ps1`)**: Requires bash on Windows or PowerShell Core on Linux; error-prone cross-platform parsing.

### Consequences
- **Positive**: Instant execution (<10ms); works in minimal scratch containers; distributed as single standalone binaries.
- **Negative**: Must compile and release separate binaries for each OS/architecture pair.
*(source: [cmd/pd/main.go](file:///d:/Project/PortDetective/cmd/pd/main.go))*

---

## ADR-003: Linux Kernel procfs Inode Parsing Fallback

### Context
In stripped Docker containers (e.g. Alpine, Debian Slim), common networking utilities like `lsof` and `netstat` are frequently missing. Requiring `lsof` causes the CLI tool to fail in containerized environments.

### Decision
Implement a secondary fallback inside `internal/lookup/linux.go` that directly reads `/proc/net/tcp`, `/proc/net/tcp6`, and parses socket inodes under `/proc/<pid>/fd/` when `exec.LookPath("lsof")` returns an error.

### Alternatives Considered
- **Fail with dependency error**: Force users to install `lsof`. Rejected because developers frequently debug minimal container images where installing packages is restricted or inconvenient.

### Consequences
- **Positive**: Port Detective functions inside stripped containers out-of-the-box.
- **Negative**: Additional parsing complexity for hex-encoded IP/port addresses in `/proc/net/tcp`.
*(source: [internal/lookup/linux.go](file:///d:/Project/PortDetective/internal/lookup/linux.go))*

---

## ADR-004: In-Memory Dictionary Maps for Internationalization

### Context
Port Detective supports both English (`en`) and Vietnamese (`vi`). Traditional i18n architectures load `.po`, `.json`, or `.yaml` files from the filesystem at runtime, which requires asset bundling or external file paths.

### Decision
Store localized string dictionaries as compiled in-memory Go maps (`map[string]string`) in `internal/i18n/en.go` and `internal/i18n/vi.go`. Wire language resolution to a global hook in `cmd/root.go`.

### Alternatives Considered
- **`go:embed` external JSON files**: Feasible, but adds JSON parsing overhead on every CLI invocation.
- **External locale directory**: Fragile for single binary distributions if binary is moved to `/usr/local/bin`.

### Consequences
- **Positive**: Zero filesystem I/O on startup; zero parsing latency; single binary remains entirely self-contained.
- **Negative**: Adding new language keys requires modifying Go source files.
*(source: [internal/i18n/i18n.go](file:///d:/Project/PortDetective/internal/i18n/i18n.go))*

---

## ADR-005: Static Next.js SSG for the Web Documentation & Showcase

### Context
The landing page and interactive showcase (`web/`) must have fast load times, cost-effective hosting, and consistent design with the CLI tool.

### Decision
Build the web showcase using Next.js 16 with App Router, Tailwind CSS v4, Lucide icons, and export as a 100% static site (SSG) deployed on Vercel.

### Alternatives Considered
- **Plain HTML/CSS**: Minimalist, but lacks modular component reuse for complex UI elements like the Interactive Terminal Preview and bilingual switcher.
- **Full-stack SSR / Node.js Server**: Unnecessary maintenance and hosting overhead for static marketing pages.

### Consequences
- **Positive**: Instant page loads; zero server-side maintenance; free hosting on Vercel CDN.
- **Negative**: Client-side state used for interactive terminal previews is simulated rather than executing real system commands.
*(source: [web/package.json](file:///d:/Project/PortDetective/web/package.json), [web/vercel.json](file:///d:/Project/PortDetective/web/vercel.json))*
