# Port Detective Documentation Hub

Welcome to the official technical and product documentation for **Port Detective (`pd`)**, a lightning-fast, cross-platform CLI tool for investigating and resolving occupied network ports.

---

## 📚 Documentation Index

### 01. Overview
- [Project Overview](01-overview/project-overview.md) — Problem statement, target developer personas, key capabilities, and technology stack.
- [Project Structure](01-overview/project-structure.md) — Repository tree layout and responsibilities of each Go package and web module.
- [Getting Started](01-overview/getting-started.md) — Prerequisites, installation methods, local build instructions, and test commands.
- [Development Workflow](01-overview/development-workflow.md) — Git branching guidelines, Conventional Commits specification, and cross-compilation validation.

### 02. Technical Architecture
- [Architecture](02-technical/architecture.md) — System layers, component boundaries, Go build tag strategy, and concurrency worker pool.
- [Business Flow](02-technical/business-flow.md) — Detailed sequence diagrams for `pd check`, `pd kill`, and `pd scan`.
- [CLI Reference & Public API](02-technical/api-reference.md) — Complete command syntax, options, machine-readable JSON schemas, and POSIX exit codes.
- [Technical Architecture Decisions (ADR)](02-technical/tech-decisions.md) — Architecture decision records detailing context, decisions, alternatives, and trade-offs.

### 03. Product
- [Changelog](03-product/changelog.md) — Detailed version release history and milestones derived from Git history.
- [Product Roadmap](03-product/roadmap.md) — Completed features, current operations, and upcoming package manager / CLI milestones.

### 04. Operations & Configuration
- [Deployment & Release Engineering](04-operations/deployment.md) — GoReleaser matrix, GitHub Actions CI/CD pipeline, and Vercel edge deployment.
- [Configuration Reference](04-operations/configuration.md) — Environment variables (`PORT_DETECTIVE_LANG`), CLI flags, and frontend settings.

### 05. Templates
- [Bug Report Template](05-templates/bug-report.md) — Standardized template for submitting bug reports.
- [ADR Template](05-templates/adr-template.md) — Standardized template for authoring Architecture Decision Records.

### 06. Frontend & UI Design
- [Web UI Architecture Notes](06-design/ui-notes.md) — Next.js 16 landing page composition, simulated terminal preview, and accessibility implementations.

### 07. Notes & Known Constraints
- [Known Issues & Limitations](07-notes/known-issues.md) — Documented OS permission constraints, stripped container caveats, and process tree behavior.
