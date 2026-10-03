"use client";

import { CheckCircle2, XCircle, Clock, Zap } from "lucide-react";
import { useLanguage } from "@/context/LanguageContext";

export function Benchmark() {
  const { t } = useLanguage();

  return (
    <section id="benchmark" className="relative py-10 sm:py-14 border-t border-slate-800/80 scroll-mt-16 bg-slate-950/60">
      <div className="mx-auto max-w-5xl px-4 sm:px-6">
        {/* Section Header */}
        <div className="text-center max-w-2xl mx-auto mb-12 sm:mb-16">
          <div className="inline-flex items-center gap-1.5 rounded-full border border-amber-500/20 bg-amber-500/10 px-3 py-1 text-xs font-medium text-amber-400 mb-3">
            <Clock className="h-3.5 w-3.5" />
            <span>{t.benchmark.badge}</span>
          </div>
          <h2 className="text-3xl sm:text-4xl font-bold tracking-tight text-white">
            {t.benchmark.title}
          </h2>
          <p className="mt-3 text-sm sm:text-base text-slate-400">
            {t.benchmark.subtitle}
          </p>
        </div>

        {/* Side-by-Side Comparison Cards */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8 items-stretch">
          {/* Traditional Way */}
          <div className="flex flex-col rounded-xl border border-rose-900/30 bg-rose-950/10 p-6 sm:p-8 backdrop-blur-sm">
            <div className="flex items-center justify-between pb-6 border-b border-rose-900/30">
              <div className="flex items-center gap-2 text-rose-400">
                <XCircle className="h-5 w-5" />
                <h3 className="font-semibold text-lg">{t.benchmark.tradTitle}</h3>
              </div>
              <span className="rounded-full bg-rose-500/10 border border-rose-500/20 px-2.5 py-1 text-xs font-mono font-medium text-rose-400">
                {t.benchmark.tradFriction}
              </span>
            </div>

            <ol className="mt-6 space-y-4 text-xs sm:text-sm text-slate-400 font-mono flex-1">
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">1.</span>
                <span>{t.benchmark.tradStep1}</span>
              </li>
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">2.</span>
                <span>{t.benchmark.tradStep2}</span>
              </li>
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">3.</span>
                <span>{t.benchmark.tradStep3}</span>
              </li>
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">4.</span>
                <span>{t.benchmark.tradStep4}</span>
              </li>
              <li className="flex items-start gap-2.5">
                <span className="text-rose-500 font-bold shrink-0">5.</span>
                <span>{t.benchmark.tradStep5}</span>
              </li>
            </ol>

            <div className="mt-8 pt-4 border-t border-rose-900/20 text-xs text-rose-300/80 font-sans">
              {t.benchmark.tradSummary}
            </div>
          </div>

          {/* Port Detective Way */}
          <div className="flex flex-col rounded-xl border border-sky-500/40 bg-sky-950/20 p-6 sm:p-8 backdrop-blur-sm relative shadow-xl shadow-sky-500/5">
            <div className="absolute -top-3 right-6 rounded-full bg-sky-500 px-3 py-0.5 text-[11px] font-sans font-bold text-slate-950 uppercase tracking-wider shadow-sm">
              {t.benchmark.pdBadge}
            </div>

            <div className="flex items-center justify-between pb-6 border-b border-sky-800/40">
              <div className="flex items-center gap-2 text-sky-400">
                <CheckCircle2 className="h-5 w-5" />
                <h3 className="font-semibold text-lg text-white">{t.benchmark.pdTitle}</h3>
              </div>
              <span className="rounded-full bg-sky-500/10 border border-sky-500/30 px-2.5 py-1 text-xs font-mono font-medium text-sky-400">
                {t.benchmark.pdTime}
              </span>
            </div>

            <div className="mt-6 space-y-4 text-xs sm:text-sm text-slate-300 font-mono flex-1">
              <div className="rounded-lg bg-slate-900/90 border border-slate-800 p-3.5 space-y-2">
                <div className="flex items-center gap-2 text-sky-300 font-semibold">
                  <span>$ pd check 8080</span>
                </div>
                <p className="text-[11px] font-sans text-slate-400">
                  {t.benchmark.pdCheckDesc}
                </p>
              </div>

              <div className="rounded-lg bg-slate-900/90 border border-slate-800 p-3.5 space-y-2">
                <div className="flex items-center gap-2 text-sky-300 font-semibold">
                  <span>$ pd kill 8080</span>
                </div>
                <p className="text-[11px] font-sans text-slate-400">
                  {t.benchmark.pdKillDesc}
                </p>
              </div>
            </div>

            <div className="mt-8 pt-4 border-t border-sky-800/30 text-xs text-sky-300 font-sans flex items-center gap-1.5">
              <Zap className="h-4 w-4 text-amber-400 shrink-0" />
              <span>{t.benchmark.pdSummary}</span>
            </div>
          </div>
        </div>

        {/* Benchmark Metric Grid */}
        <div className="mt-12 grid grid-cols-2 sm:grid-cols-4 gap-4 text-center">
          <div className="rounded-xl border border-slate-800 bg-slate-900/40 p-4">
            <div className="text-2xl sm:text-3xl font-extrabold text-white font-mono">{t.benchmark.stat1}</div>
            <div className="text-xs text-slate-400 mt-1">{t.benchmark.stat1_label}</div>
          </div>
          <div className="rounded-xl border border-slate-800 bg-slate-900/40 p-4">
            <div className="text-2xl sm:text-3xl font-extrabold text-emerald-400 font-mono">{t.benchmark.stat2}</div>
            <div className="text-xs text-slate-400 mt-1">{t.benchmark.stat2_label}</div>
          </div>
          <div className="rounded-xl border border-slate-800 bg-slate-900/40 p-4">
            <div className="text-2xl sm:text-3xl font-extrabold text-sky-400 font-mono">{t.benchmark.stat3}</div>
            <div className="text-xs text-slate-400 mt-1">{t.benchmark.stat3_label}</div>
          </div>
          <div className="rounded-xl border border-slate-800 bg-slate-900/40 p-4">
            <div className="text-2xl sm:text-3xl font-extrabold text-purple-400 font-mono">{t.benchmark.stat4}</div>
            <div className="text-xs text-slate-400 mt-1">{t.benchmark.stat4_label}</div>
          </div>
        </div>
      </div>
    </section>
  );
}
