# Deployment & Release Engineering

This document outlines the build, release, and deployment pipelines for the Port Detective CLI tool and the Next.js web application.

---

## 1. CLI Release Pipeline (GoReleaser + GitHub Actions)

The CLI tool utilizes **GoReleaser** executed inside a GitHub Actions runner upon Git tag pushes.

### Workflow Configuration: `.github/workflows/release.yml`

```yaml
name: goreleaser

on:
  push:
    tags:
      - '*'

permissions:
  contents: write

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v6
        with:
          distribution: goreleaser
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

### Release Artifacts Matrix

When a tag (e.g. `v1.1.2`) is pushed, the workflow generates compiled standalone binaries packaged into archives with SHA-256 checksums:

| OS | Architecture | Archive Format | Output Binary |
| :--- | :--- | :--- | :--- |
| **Windows** | `x86_64`, `arm64` | `.zip` | `pd.exe` |
| **Linux** | `x86_64`, `arm64` | `.tar.gz` | `pd` |
| **macOS** | `x86_64`, `arm64` | `.tar.gz` | `pd` |

*(source: [.github/workflows/release.yml](file:///d:/Project/PortDetective/.github/workflows/release.yml))*

---

## 2. Automated Installer Distribution

Pre-built binaries are distributed via automated shell scripts hosted directly in the repository:

- **Linux & macOS (`scripts/install.sh`)**:
  - Fetches the latest release tag metadata from the GitHub API.
  - Detects host OS (`uname -s`) and CPU architecture (`uname -m`).
  - Downloads, verifies, and installs `pd` into `~/.local/bin` or `/usr/local/bin`.
- **Windows PowerShell (`scripts/install.ps1`)**:
  - Detects CPU architecture (`AMD64` vs `ARM64`).
  - Extracts `pd.exe` to `$HOME\AppData\Local\PortDetective\bin`.
  - Automatically updates the User `PATH` environment variable.
*(source: [scripts/install.sh](file:///d:/Project/PortDetective/scripts/install.sh), [scripts/install.ps1](file:///d:/Project/PortDetective/scripts/install.ps1))*

---

## 3. Web Showcase Deployment (Vercel)

The web documentation and interactive showcase (`web/`) is hosted on the Vercel Edge Network.

### Configuration: `web/vercel.json`

```json
{
  "framework": "nextjs",
  "cleanUrls": true,
  "headers": [
    {
      "source": "/(.*)",
      "headers": [
        { "key": "X-Content-Type-Options", "value": "nosniff" },
        { "key": "X-Frame-Options", "value": "DENY" },
        { "key": "X-XSS-Protection", "value": "1; mode=block" }
      ]
    }
  ]
}
```

### Build & Deploy Command
- **Build**: `next build` (generates 100% static SSG pages).
- **Production URL**: `https://port-detective.vercel.app`
*(source: [web/vercel.json](file:///d:/Project/PortDetective/web/vercel.json))*
