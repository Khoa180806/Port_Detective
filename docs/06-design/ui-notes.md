# Web UI & Frontend Architecture Notes

This document details the frontend architecture, component composition, accessibility implementations, and styling design tokens for the Port Detective web application (`web/`).

---

## 1. Frontend Technology Stack

- **Framework**: Next.js 16 (App Router with 100% Static Site Generation)
- **UI Library**: React 19
- **CSS Utility Engine**: Tailwind CSS v4
- **Iconography**: `lucide-react`
*(source: [web/package.json](file:///d:/Project/PortDetective/web/package.json))*

---

## 2. Component Hierarchy & Composition

The page layout in `web/app/page.tsx` is structured into focused, single-responsibility components:

```text
web/app/page.tsx
├── <Navbar />                 # Brand logo, GitHub release badge, docs links, and bilingual switcher
├── <Hero />                   # Primary value proposition headline and call-to-action buttons
├── <InstallTabs />            # One-click copy tabbed install commands (Linux/macOS vs. Windows)
├── <TerminalPreview />        # Interactive terminal simulator with tabbed command scenarios
├── <Features />               # Grid of 6 core feature cards with Lucide icons
├── <BenchmarkComparison />    # Performance latency comparison table vs. npx kill-port and netstat
├── <PortReferenceTable />     # Searchable / filterable common developer port reference table
├── <FAQ />                    # Collapsible accordion answering permissions and cross-platform questions
└── <Footer />                 # MIT license declaration, repository links, and author attribution
```
*(source: [web/app/page.tsx](file:///d:/Project/PortDetective/web/app/page.tsx), [web/components/](file:///d:/Project/PortDetective/web/components))*

---

## 3. Interactive Terminal Simulator (`TerminalPreview.tsx`)

The `TerminalPreview` component demonstrates Port Detective's runtime behavior in the browser without requiring a backend shell:

- **State Management**: React state handles active command selection (`pd check`, `pd kill`, `pd scan`).
- **Real Code Parity**: Terminal outputs match the exact character formatting, syntax colors, and indentation produced by the Go binary formatters (`internal/output/text.go`).
- **Lazy Initialization**: Uses lazy state initializers to prevent hydration mismatches between client and server.
*(source: [web/components/TerminalPreview.tsx](file:///d:/Project/PortDetective/web/components/TerminalPreview.tsx))*

---

## 4. Accessibility (A11y) & SEO Architecture

- **Semantic ARIA Attributes**:
  - Tab controls include `role="tab"`, `aria-selected`, and `aria-controls`.
  - FAQ accordion items include `role="region"` and `aria-expanded`.
- **High-Contrast Dark Theme**:
  - Slate and Zinc palette with high contrast ratios meeting WCAG 2.1 AA standards.
- **Dynamic SEO Metadata**:
  - Dynamically generated OpenGraph preview image via `web/app/opengraph-image.tsx`.
  - Configured `robots.txt` and `sitemap.xml` for search engine indexation.
*(source: [web/app/opengraph-image.tsx](file:///d:/Project/PortDetective/web/app/opengraph-image.tsx), [web/app/robots.ts](file:///d:/Project/PortDetective/web/app/robots.ts))*
