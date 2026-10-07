# Product Roadmap & Evolution

This document tracks the product evolution of Port Detective, capturing completed capabilities, current operational status, and concrete future milestones.

---

## 1. Completed Capabilities (Shipped)

- [x] **Core CLI Subcommands**: `pd check <port>`, `pd kill <port>`, `pd scan <range>`.
- [x] **Safety Controls**: Interactive confirmation `[y/N]`, non-destructive `--dry-run`, automated `--force`.
- [x] **Automation & CI/CD**: Strict `--json` output across all commands, deterministic POSIX exit codes (0, 1, 2, 3, 4).
- [x] **Cross-Platform Engine**: Go build tags for Windows (`netstat`), Linux (`lsof` + `/proc/net/tcp` inode fallback), and macOS (`lsof`).
- [x] **Native Bilingual Support**: Runtime language switching between English and Vietnamese (`--lang vi`, `PORT_DETECTIVE_LANG=vi`).
- [x] **Zero-Config Installers**: Automated one-line scripts (`scripts/install.sh`, `scripts/install.ps1`).
- [x] **Automated Releases**: GitHub Actions workflow powered by GoReleaser (`.github/workflows/release.yml`).
- [x] **Web Showcase & Documentation**: Next.js 16 landing page with interactive terminal simulator and port reference table deployed on Vercel.
*(source: Git history and codebase implementation)*

---

## 2. In Progress / Maintenance

- [x] **Documentation Refactoring**: Restructuring technical and product documentation into standardized modular directories (`docs/01-overview/` through `docs/07-notes/`).
- [ ] **Cross-Platform Test Coverage**: Expanding automated end-to-end integration tests on native macOS and Linux runners.
*(source: Active development branch)*

---

## 3. Concrete Future Milestones

### Milestone 1: Package Manager Distribution
- [ ] **Homebrew Formula**: Create official Homebrew tap for macOS and Linux (`brew install port-detective`).
- [ ] **Windows Package Managers**: Publish manifest to Scoop Main bucket and Windows Package Manager (Winget).
- [ ] **Linux Repositories**: Provide Arch User Repository (AUR) PKGBUILD and native `.deb` / `.rpm` release assets.
- [ ] **Container Image**: Publish official micro-container to GitHub Container Registry (`ghcr.io/khoa180806/port-detective:latest`).

### Milestone 2: Advanced CLI Features
- [ ] **Process Tree Termination (`--tree` / `-t`)**: Recursively terminate parent and child processes spawned by process managers (e.g., `npm`, `nodemon`, `vite`).
- [ ] **Live Interactive TUI (`pd top` / `pd monitor`)**: Terminal dashboard built with `charmbracelet/bubbletea` for live socket monitoring and 1-key process termination.
- [ ] **Watch & Auto-Release Mode (`pd watch <port>`)**: Background watcher that detects crashed zombie processes on development ports and auto-clears them.
- [ ] **Process Deep Inspection**: Display Working Directory (CWD), Uptime, and Memory/CPU usage metrics for identified PIDs.

### Milestone 3: Ecosystem Extensions
- [ ] **VS Code Extension**: Status bar indicator showing active local development ports with 1-click "Free Port" action.
- [ ] **Go Public SDK (`pkg/detective`)**: Extract core socket query engine into a standalone public Go package for programmatic integration in third-party test suites.
*(source: Code analysis and [docs/roadmap.md](file:///d:/Project/PortDetective/docs/roadmap.md))*
