import { CheckCircle2, XCircle, Clock, Zap } from "lucide-react";

export function Benchmark() {
  return (
    <section id="benchmark" className="relative py-16 sm:py-24 border-t border-slate-800/80 scroll-mt-16 bg-slate-950/60">
      <div className="mx-auto max-w-5xl px-4 sm:px-6">
        {/* Section Header */}
        <div className="text-center max-w-2xl mx-auto mb-12 sm:mb-16">
          <div className="inline-flex items-center gap-1.5 rounded-full border border-amber-500/20 bg-amber-500/10 px-3 py-1 text-xs font-medium text-amber-400 mb-3">
            <Clock className="h-3.5 w-3.5" />
            <span>Developer Productivity</span>
          </div>
          <h2 className="text-3xl sm:text-4xl font-bold tracking-tight text-white">
            1 Second vs 30 Seconds
          </h2>
          <p className="mt-3 text-sm sm:text-base text-slate-400">
            Compare the friction of traditional multi-step OS terminal commands against the streamlined workflow of Port Detective.
          </p>
        </div>

        {/* Side-by-Side Comparison Cards */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8 items-stretch">
          {/* Traditional Way */}
          <div className="flex flex-col rounded-xl border border-rose-900/30 bg-rose-950/10 p-6 sm:p-8 backdrop-blur-sm">
            <div className="flex items-center justify-between pb-6 border-b border-rose-900/30">
              <div className="flex items-center gap-2 text-rose-400">
                <XCircle className="h-5 w-5" />
                <h3 className="font-semibold text-lg">Traditional OS Commands</h3>
              </div>
              <span className="rounded-full bg-rose-500/10 border border-rose-500/20 px-2.5 py-1 text-xs font-mono font-medium text-rose-400">
                ~25–30s Friction
              </span>
            </div>

            <ol className="mt-6 space-y-4 text-xs sm:text-sm text-slate-400 font-mono flex-1">
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">1.</span>
                <span>
                  Remember & type <code className="text-rose-300">netstat -ano | findstr :8080</code>
                </span>
              </li>
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">2.</span>
                <span>Squint across terminal columns to identify PID (e.g. 14280)</span>
              </li>
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">3.</span>
                <span>
                  Query process name via <code className="text-rose-300">tasklist | findstr 14280</code> or Task Manager
                </span>
              </li>
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">4.</span>
                <span>
                  Execute destructive kill: <code className="text-rose-300">taskkill /F /PID 14280</code>
                </span>
              </li>
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">5.</span>
                <span>Risk terminating critical system services due to typing error</span>
              </li>
            </ol>

            <div className="mt-8 pt-4 border-t border-rose-900/20 text-xs text-rose-300/80 font-sans">
              ❌ High cognitive load, error-prone PID matching, breaks developer flow state.
            </div>
          </div>

          {/* Port Detective Way */}
          <div className="flex flex-col rounded-xl border border-sky-500/40 bg-sky-950/20 p-6 sm:p-8 backdrop-blur-sm relative shadow-xl shadow-sky-500/5">
            <div className="absolute -top-3 right-6 rounded-full bg-sky-500 px-3 py-0.5 text-[11px] font-sans font-bold text-slate-950 uppercase tracking-wider shadow-sm">
              Recommended
            </div>

            <div className="flex items-center justify-between pb-6 border-b border-sky-800/40">
              <div className="flex items-center gap-2 text-sky-400">
                <CheckCircle2 className="h-5 w-5" />
                <h3 className="font-semibold text-lg text-white">Port Detective</h3>
              </div>
              <span className="rounded-full bg-sky-500/10 border border-sky-500/30 px-2.5 py-1 text-xs font-mono font-medium text-sky-400">
                ⚡ ~1s Total
              </span>
            </div>

            <div className="mt-6 space-y-4 text-xs sm:text-sm text-slate-300 font-mono flex-1">
              <div className="rounded-lg bg-slate-900/90 border border-slate-800 p-3.5 space-y-2">
                <div className="flex items-center gap-2 text-sky-300 font-semibold">
                  <span>$ pd check 8080</span>
                </div>
                <p className="text-[11px] font-sans text-slate-400">
                  Instantly prints PID, Process Name, Protocol, and Command line in high-contrast color.
                </p>
              </div>

              <div className="rounded-lg bg-slate-900/90 border border-slate-800 p-3.5 space-y-2">
                <div className="flex items-center gap-2 text-sky-300 font-semibold">
                  <span>$ pd kill 8080</span>
                </div>
                <p className="text-[11px] font-sans text-slate-400">
                  Safely asks confirmation with full process context, or bypass with <code className="text-sky-300">-f</code>.
                </p>
              </div>
            </div>

            <div className="mt-8 pt-4 border-t border-sky-800/30 text-xs text-sky-300 font-sans flex items-center gap-1.5">
              <Zap className="h-4 w-4 text-amber-400 shrink-0" />
              <span>One unified command across Windows, macOS, and Linux. Zero friction.</span>
            </div>
          </div>
        </div>

        {/* Benchmark Metric Grid */}
        <div className="mt-12 grid grid-cols-2 sm:grid-cols-4 gap-4 text-center">
          <div className="rounded-xl border border-slate-800 bg-slate-900/40 p-4">
            <div className="text-2xl sm:text-3xl font-extrabold text-white font-mono">~48ms</div>
            <div className="text-xs text-slate-400 mt-1">pd check runtime</div>
          </div>
          <div className="rounded-xl border border-slate-800 bg-slate-900/40 p-4">
            <div className="text-2xl sm:text-3xl font-extrabold text-emerald-400 font-mono">20x</div>
            <div className="text-xs text-slate-400 mt-1">Faster developer workflow</div>
          </div>
          <div className="rounded-xl border border-slate-800 bg-slate-900/40 p-4">
            <div className="text-2xl sm:text-3xl font-extrabold text-sky-400 font-mono">5–8 MB</div>
            <div className="text-xs text-slate-400 mt-1">Standalone binary size</div>
          </div>
          <div className="rounded-xl border border-slate-800 bg-slate-900/40 p-4">
            <div className="text-2xl sm:text-3xl font-extrabold text-purple-400 font-mono">0 dep</div>
            <div className="text-xs text-slate-400 mt-1">External dependencies</div>
          </div>
        </div>
      </div>
    </section>
  );
}
