# Development Workflow

This document outlines the branching strategy, commit conventions, cross-platform build validation, and quality gates for Port Detective.

---

## 1. Branching Strategy

- **`master` Branch**: Production-ready code. All commits on `master` must pass unit tests, static vet checks, and cross-platform compilation.
- **Feature Branches**: Branch from `master` using the convention:
  - `feat/<feature-name>` for new capabilities
  - `fix/<bug-name>` for bug fixes
  - `refactor/<target>` for structural code refactoring
  - `docs/<topic>` for documentation updates

---

## 2. Commit Message Convention

Commits must follow the **Conventional Commits** specification:

```text
<type>(<scope>): <short description>
```

### Supported Commit Types

| Type | Description | Example |
| :--- | :--- | :--- |
| `feat:` | New feature or capability | `feat(scan): add concurrency flag` |
| `fix:` | Bug fix in lookup or output | `fix(windows): resolve process name truncation` |
| `refactor:` | Code restructuring without feature change | `refactor(lookup): split socket parsing helper` |
| `test:` | Adding or modifying tests | `test(output): add test for empty slice format` |
| `docs:` | Documentation changes only | `docs(overview): update tech stack table` |
| `style:` | Formatting or styling adjustments | `style: apply fatih/color for CLI output` |
| `chore:` | Build scripts, dependencies, CI config | `chore(release): bump version to 1.1.1` |

---

## 3. Testing Approach & Verification

We adhere strictly to test-driven development (TDD) and multi-platform compilation checks:

### 3.1 Unit Testing
Tests reside adjacent to their source files (`*_test.go`).
```bash
go test -v ./...
```

### 3.2 Cross-Platform Compilation Check
Because Port Detective relies on OS-specific build tags (`windows.go`, `linux.go`, `darwin.go`), any changes touching `internal/lookup` must be verified across all target platforms:

```bash
# Windows (x86_64)
GOOS=windows GOARCH=amd64 go build -o /dev/null ./cmd/pd

# Linux (x86_64)
GOOS=linux GOARCH=amd64 go build -o /dev/null ./cmd/pd

# macOS (Apple Silicon arm64)
GOOS=darwin GOARCH=arm64 go build -o /dev/null ./cmd/pd
```
*(On Windows cmd/powershell, replace `/dev/null` with `NUL`)*.

---

## 4. Pre-Merge Checklist

Before merging a branch into `master`:

- [ ] All unit tests pass: `go test -v ./...`
- [ ] Code is formatted: `go fmt ./...`
- [ ] Static analysis passes: `go vet ./...`
- [ ] Cross-compilation passes for Windows, Linux, and macOS.
- [ ] Frontend builds cleanly with zero lint warnings (if `web/` was modified): `cd web && npm run lint && npm run build`.
*(source: [docs/contributing.md](file:///d:/Project/PortDetective/docs/contributing.md))*
