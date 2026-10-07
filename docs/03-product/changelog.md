# Changelog

All notable changes to the Port Detective project are documented in this file, structured by releases and milestones derived from Git history.

---

## [v1.1.1] — 2026-09-21

### Changed
- Bumped version to `1.1.1` across automated installer scripts.
- Refined release archive asset detection via GitHub Releases API.
*(source: Git commits `186ed45`, `1481cb3`)*

---

## [v1.1.0] — 2026-09-21

### Added
- **Zero-Config Installer Scripts**:
  - Added `scripts/install.sh` for one-line automated installation on Linux and macOS.
  - Added `scripts/install.ps1` for one-line automated installation on Windows PowerShell.
- **Canonical Entrypoint**:
  - Moved CLI entrypoint to `cmd/pd/main.go` to support standard `go install github.com/.../cmd/pd@latest`.
- **Media & Architecture Assets**:
  - Added vector SVG and high-resolution PNG architecture diagrams (`assets/architecture.png`).
  - Added interactive demo GIF and terminal screenshot assets.

### Changed
- Removed redundant root `main.go`.
*(source: Git commits `e8ef842`, `a78e932`, `22dd7bc`)*

---

## [v1.0.0] — 2026-09-21

### Added
- **Core CLI Subcommands**:
  - `pd check <port>`: Queries active process metadata occupying a specific port.
  - `pd kill <port>`: Terminates processes on a target port with interactive `[y/N]` confirmation prompts.
  - `--force` / `-f`: Bypasses confirmation prompt for automation.
  - `--dry-run` / `-d`: Previews target processes without terminating.
  - `pd scan <start>-<end>`: Concurrently scans port ranges using a 100-worker Goroutine pool.
  - `--json` / `-j`: Machine-readable JSON output for all subcommands.
- **Cross-Platform Strategy Engine (`internal/lookup`)**:
  - Windows: `netstat -ano` output parser + `tasklist /FO CSV` resolution + `taskkill /F /PID`.
  - Linux: `lsof -i :<port>` with fallback direct `/proc/net/tcp` inode parsing for stripped Docker containers.
  - macOS: Native `lsof -i` query engine.
- **Internationalization Subsystem (`internal/i18n`)**:
  - In-memory bilingual dictionaries for English (`en`) and Vietnamese (`vi`).
  - Runtime language selection via `--lang` flag or `PORT_DETECTIVE_LANG` environment variable.
- **Output Formatter (`internal/output`)**:
  - High-contrast colorized ANSI tables via `fatih/color`.
- **CI/CD Automation**:
  - Configured GoReleaser and GitHub Actions workflow (`.github/workflows/release.yml`).
*(source: Git commits `18b9b43` through `cbad427`)*

---

## Web Showcase & Documentation Milestone — 2026-10-03

### Added
- **Interactive Web Showcase (`web/`)**:
  - Initialized Next.js 16 app with Tailwind CSS v4 and Lucide React.
  - Interactive TerminalPreview mockup simulating `pd check`, `pd kill`, and `pd scan`.
  - Common Port Reference table with filtering and quick copy snippets.
  - Bilingual switcher, SEO metadata, dynamic OpenGraph image, and Vercel edge deployment (`web/vercel.json`).
*(source: Git commits `394b496` through `501b163`)*
