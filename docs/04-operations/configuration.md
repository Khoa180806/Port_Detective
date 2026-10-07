# Configuration Reference

This document catalogs all environment variables, runtime flags, and configuration parameters supported by Port Detective.

---

## 1. Environment Variables

| Variable Name | Type | Default Value | Description |
| :--- | :---: | :---: | :--- |
| `PORT_DETECTIVE_LANG` | `string` | `"en"` | Global language preference for CLI output. Supported values: `en` (English), `vi` (Vietnamese). |

### Precedence Hierarchy for Localization:
1. Explicit CLI flag: `--lang vi` / `--lang en`
2. Environment variable: `PORT_DETECTIVE_LANG`
3. Operating system environment: `$LANG` / `$LC_ALL`
4. Fallback default: `en`
*(source: [cmd/root.go](file:///d:/Project/PortDetective/cmd/root.go), [internal/i18n/i18n.go](file:///d:/Project/PortDetective/internal/i18n/i18n.go))*

---

## 2. CLI Runtime Flags

### Global Flags

| Flag | Shorthand | Type | Default | Description |
| :--- | :---: | :---: | :---: | :--- |
| `--lang` | `-l` | `string` | `"en"` | Override language dictionary for the current execution. |
| `--help` | `-h` | `bool` | `false` | Display help synopsis and command list. |
| `--version` | `-v` | `bool` | `false` | Display compiled version of the CLI. |

### Command Flags

| Command | Flag | Shorthand | Type | Default | Description |
| :--- | :--- | :---: | :---: | :---: | :--- |
| `pd check` | `--json` | `-j` | `bool` | `false` | Serialize process records to structured JSON. |
| `pd kill` | `--force` | `-f` | `bool` | `false` | Terminate process immediately without interactive confirmation. |
| `pd kill` | `--dry-run` | `-d` | `bool` | `false` | Simulate process kill without sending OS termination signals. |
| `pd scan` | `--json` | `-j` | `bool` | `false` | Output discovered active ports in JSON format. |

*(source: [cmd/check.go](file:///d:/Project/PortDetective/cmd/check.go), [cmd/kill.go](file:///d:/Project/PortDetective/cmd/kill.go), [cmd/scan.go](file:///d:/Project/PortDetective/cmd/scan.go))*

---

## 3. Web Application Configuration (`web/`)

The web frontend operates with zero runtime database configuration.

### Static Build Configurations:
- **Framework**: Next.js 16 (App Router)
- **Port**: Defaults to `3000` during local development (`npm run dev`).
- **Security Headers**: Defined in `web/vercel.json` (`X-Content-Type-Options: nosniff`, `X-Frame-Options: DENY`, `X-XSS-Protection: 1; mode=block`).
*(source: [web/vercel.json](file:///d:/Project/PortDetective/web/vercel.json))*
