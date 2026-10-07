# Bug Report Template

Use this template when submitting a bug report for Port Detective CLI or the Web Landing Page.

---

## Bug Description
A clear and concise description of what the bug is.

## Environment & System Details
- **Port Detective Version**: (Run `pd --version`, e.g., `1.1.1`)
- **Operating System**: (Windows 11 / Ubuntu 22.04 / macOS Sonoma)
- **Architecture**: (x86_64 / arm64 / Apple Silicon)
- **Terminal Shell**: (bash / zsh / PowerShell / Command Prompt)

## Steps to Reproduce
1. Start a mock server on port `...` (e.g., `python -m http.server 8080`)
2. Run command: `pd check 8080`
3. Observe unexpected behavior...

## Expected Behavior
A concise description of what you expected to happen.

## Actual CLI Output
```text
Paste terminal output or error message here...
```

## Additional Context
Add any other context about the problem here (e.g., running inside Docker, user permission levels, administrator elevation).
