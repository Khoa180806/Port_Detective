# 🤝 Contributing Guide — Port Detective

<div align="center">

[![Port Detective Documentation](https://img.shields.io/badge/docs-contributing-blue?style=flat-square)](../README.md)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=flat-square)](https://conventionalcommits.org)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/Khoa180806/Port_Detective/pulls)

</div>

Thank you for your interest in contributing to **Port Detective (`pd`)**! We appreciate community contributions, feature ideas, bug fixes, and documentation improvements.

---

## 🛠️ Development Setup

### Prerequisites
- **Go**: Version **1.21** or later installed.
- **Git**: Configured with your developer identity.

### Clone & Download Dependencies
```bash
git clone https://github.com/Khoa180806/Port_Detective.git
cd Port_Detective
go mod download
```

---

## 🧪 Testing & Verification Workflow

### 1. Run Unit Tests
We adhere strictly to test-driven and regression testing practices. Ensure all tests pass before proposing any pull requests:

```bash
go test -v ./...
```

### 2. Format & Linting Check
Ensure all code conforms to standard Go idioms:

```bash
# Format source files
go fmt ./...

# Static vet analysis
go vet ./...
```

### 3. Cross-Platform Compilation Check
Because Port Detective relies on OS-specific build tags (`windows.go`, `linux.go`, `darwin.go`), always verify that your changes compile successfully across all supported platforms:

```bash
# Windows (amd64)
GOOS=windows GOARCH=amd64 go build -o bin/pd-windows.exe ./cmd/pd

# Linux (amd64)
GOOS=linux GOARCH=amd64 go build -o bin/pd-linux ./cmd/pd

# macOS (Apple Silicon arm64)
GOOS=darwin GOARCH=arm64 go build -o bin/pd-darwin ./cmd/pd
```

---

## 📝 Commit Conventions

All commits must follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

| Prefix | Type | Example |
| :--- | :--- | :--- |
| `feat:` | New feature or capability | `feat(scan): add concurrency limit flag` |
| `fix:` | Bug fix | `fix(windows): resolve process name truncation in tasklist` |
| `refactor:` | Code restructuring without feature change | `refactor(lookup): split socket parsing logic into helper` |
| `test:` | Adding or modifying unit / integration tests | `test(check): add test cases for unoccupied port responses` |
| `docs:` | Documentation changes only | `docs(readme): add troubleshooting section` |
| `chore:` | Build tooling, CI workflows, dependencies | `chore(ci): update release action workflow` |

> [!TIP]
> Keep commits atomic: make focused commits for each individual change rather than large monolithic diffs.

---

## 🔀 Pull Request (PR) Checklist

Before submitting a Pull Request, please ensure you have completed the following checklist:

1. [ ] Created a descriptive branch from `master` (`git checkout -b feat/your-feature-name`).
2. [ ] Added accompanying unit tests for any new logic or bug fixes.
3. [ ] All unit tests pass: `go test -v ./...`.
4. [ ] Code is formatted with `go fmt ./...` and passes `go vet ./...`.
5. [ ] Verified cross-platform builds (`GOOS=windows`, `GOOS=linux`, `GOOS=darwin`).
6. [ ] Followed Conventional Commits in your commit messages.

---

## 🧭 Navigation

- [📐 System Architecture](./architecture.md)
- [📖 CLI Reference Guide](./cli-reference.md)
- [📦 Installation Guide](./installation.md)
- [🏠 Project Root & README](../README.md)
