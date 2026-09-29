# 🤝 Contributing Guide

Thank you for your interest in contributing to **Port Detective (`pd`)**! We welcome community contributions, bug reports, and suggestions. This guide covers how to set up your development environment, run tests, and submit high-quality pull requests.

---

## 🛠️ Development Setup

### Prerequisites
- **Go**: Version **1.21** or later installed.
- **Git**: Configured on your system.

### Getting the Code
```bash
git clone https://github.com/Khoa180806/Port_Detective.git
cd Port_Detective
go mod download
```

---

## 🧪 Development & Testing Workflow

### 1. Run Unit Tests
We advocate for test-driven development (TDD). Ensure all tests pass before making any changes:

```bash
go test -v ./...
```

### 2. Static Analysis & Linting
Ensure code adheres to idiomatic Go conventions:

```bash
# Format code
go fmt ./...

# Static analysis
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

## 📝 Commit Guidelines

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

- `feat:` A new feature or capability (e.g., `feat: add scan command for port ranges`)
- `fix:` A bug fix (e.g., `fix: handle edge case when netstat PID is 0`)
- `refactor:` Code restructuring that neither fixes a bug nor adds a feature
- `test:` Adding missing tests or correcting existing tests
- `docs:` Documentation changes only
- `chore:` Changes to build process, CI workflows, or auxiliary tooling

Keep commits atomic: make small, focused commits with concise, descriptive messages.

---

## 🔀 Submitting a Pull Request (PR)

1. Fork the repository and create a feature branch from `master`:
   ```bash
   git checkout -b feat/your-feature-name
   ```
2. Write clean, idiomatic Go code with accompanying unit tests.
3. Run `go test -v ./...` and verify cross-compilation passes.
4. Commit your changes following Conventional Commits.
5. Push your branch to your fork:
   ```bash
   git push origin feat/your-feature-name
   ```
6. Open a Pull Request against `master`. Provide a clear summary of your changes, the rationale behind them, and verification steps.
