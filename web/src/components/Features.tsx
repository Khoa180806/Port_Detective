import { Zap, ShieldCheck, Cpu, Layers, Braces, Globe } from "lucide-react";

interface Feature {
  icon: typeof Zap;
  title: string;
  description: string;
  badge: string;
  badgeColor: string;
}

const FEATURES: Feature[] = [
  {
    icon: Zap,
    title: "Sub-50ms Execution",
    description:
      "Built with Go and compiled to a lightweight standalone binary. Zero interpreter startup delay, no NodeJS or Python runtime overhead.",
    badge: "Native Performance",
    badgeColor: "text-amber-400 bg-amber-500/10 border-amber-500/20",
  },
  {
    icon: ShieldCheck,
    title: "Safe by Default",
    description:
      "Prevents catastrophic accidental kills. Displays full process details with an interactive [y/N] prompt and supports --dry-run simulation.",
    badge: "Human Guardrails",
    badgeColor: "text-emerald-400 bg-emerald-500/10 border-emerald-500/20",
  },
  {
    icon: Cpu,
    title: "Unified Cross-Platform",
    description:
      "One syntax across Windows, macOS, and Linux. No more memorizing netstat -ano, lsof -i, or taskkill commands when switching machines.",
    badge: "Zero Friction",
    badgeColor: "text-sky-400 bg-sky-500/10 border-sky-500/20",
  },
  {
    icon: Layers,
    title: "Concurrent Port Scanner",
    description:
      "High-throughput goroutine worker pool scans ranges of up to 5,000 ports in milliseconds. Detect open ports and bound services instantly.",
    badge: "Goroutines",
    badgeColor: "text-cyan-400 bg-cyan-500/10 border-cyan-500/20",
  },
  {
    icon: Braces,
    title: "Machine-Readable JSON",
    description:
      "Pass the --json flag on any command for strict, clean JSON output. Built specifically for CI/CD pipelines, shell scripts, and jq filtering.",
    badge: "Automation Ready",
    badgeColor: "text-purple-400 bg-purple-500/10 border-purple-500/20",
  },
  {
    icon: Globe,
    title: "Native Bilingual CLI",
    description:
      "Native support for English and Vietnamese out of the box. Switch dynamically via --lang vi or configure system-wide with environment variables.",
    badge: "i18n Built-in",
    badgeColor: "text-rose-400 bg-rose-500/10 border-rose-500/20",
  },
];

export function Features() {
  return (
    <section id="features" className="relative py-16 sm:py-24 scroll-mt-16">
      <div className="mx-auto max-w-6xl px-4 sm:px-6">
        {/* Section Header */}
        <div className="text-center max-w-2xl mx-auto mb-12 sm:mb-16">
          <div className="inline-flex items-center gap-1.5 rounded-full border border-sky-500/20 bg-sky-500/10 px-3 py-1 text-xs font-medium text-sky-400 mb-3">
            <span>Capabilities</span>
          </div>
          <h2 className="text-3xl sm:text-4xl font-bold tracking-tight text-white">
            Engineered for Modern Developers
          </h2>
          <p className="mt-3 text-sm sm:text-base text-slate-400">
            Everything you need to troubleshoot, free up ports, and streamline your local development environment.
          </p>
        </div>

        {/* 3x2 Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {FEATURES.map((feat, index) => {
            const Icon = feat.icon;
            return (
              <div
                key={index}
                className="group relative rounded-xl border border-slate-800 bg-slate-900/50 p-6 backdrop-blur-sm transition-all hover:border-slate-700 hover:bg-slate-900/80 hover:shadow-xl hover:shadow-sky-500/5"
              >
                <div className="flex items-center justify-between mb-4">
                  <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-slate-800 border border-slate-700/80 text-sky-400 group-hover:scale-105 group-hover:text-sky-300 transition-transform">
                    <Icon className="h-5 w-5" />
                  </div>
                  <span
                    className={`rounded-full border px-2.5 py-0.5 text-[11px] font-mono font-medium ${feat.badgeColor}`}
                  >
                    {feat.badge}
                  </span>
                </div>

                <h3 className="text-base sm:text-lg font-semibold text-white mb-2 group-hover:text-sky-300 transition-colors">
                  {feat.title}
                </h3>
                <p className="text-sm text-slate-400 leading-relaxed">
                  {feat.description}
                </p>
              </div>
            );
          })}
        </div>
      </div>
    </section>
  );
}
