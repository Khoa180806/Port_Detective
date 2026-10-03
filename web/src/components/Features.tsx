"use client";

import { Zap, ShieldCheck, Cpu, Layers, Braces, Globe } from "lucide-react";
import { useLanguage } from "@/context/LanguageContext";

export function Features() {
  const { t } = useLanguage();

  const FEATURES = [
    {
      icon: Zap,
      title: t.features.f1_title,
      description: t.features.f1_desc,
      badge: t.features.f1_badge,
      badgeColor: "text-amber-400 bg-amber-500/10 border-amber-500/20",
    },
    {
      icon: ShieldCheck,
      title: t.features.f2_title,
      description: t.features.f2_desc,
      badge: t.features.f2_badge,
      badgeColor: "text-emerald-400 bg-emerald-500/10 border-emerald-500/20",
    },
    {
      icon: Cpu,
      title: t.features.f3_title,
      description: t.features.f3_desc,
      badge: t.features.f3_badge,
      badgeColor: "text-sky-400 bg-sky-500/10 border-sky-500/20",
    },
    {
      icon: Layers,
      title: t.features.f4_title,
      description: t.features.f4_desc,
      badge: t.features.f4_badge,
      badgeColor: "text-cyan-400 bg-cyan-500/10 border-cyan-500/20",
    },
    {
      icon: Braces,
      title: t.features.f5_title,
      description: t.features.f5_desc,
      badge: t.features.f5_badge,
      badgeColor: "text-purple-400 bg-purple-500/10 border-purple-500/20",
    },
    {
      icon: Globe,
      title: t.features.f6_title,
      description: t.features.f6_desc,
      badge: t.features.f6_badge,
      badgeColor: "text-rose-400 bg-rose-500/10 border-rose-500/20",
    },
  ];

  return (
    <section id="features" className="relative py-10 sm:py-14 scroll-mt-16">
      <div className="mx-auto max-w-6xl px-4 sm:px-6">
        {/* Section Header */}
        <div className="text-center max-w-2xl mx-auto mb-12 sm:mb-16">
          <div className="inline-flex items-center gap-1.5 rounded-full border border-sky-500/20 bg-sky-500/10 px-3 py-1 text-xs font-medium text-sky-400 mb-3">
            <span>{t.features.badge}</span>
          </div>
          <h2 className="text-3xl sm:text-4xl font-bold tracking-tight text-white">
            {t.features.title}
          </h2>
          <p className="mt-3 text-sm sm:text-base text-slate-400">
            {t.features.subtitle}
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
