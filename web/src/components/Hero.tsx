import { Terminal, ShieldCheck, Zap, ArrowRight } from "lucide-react";
import { GithubIcon } from "./GithubIcon";
import { InstallTabs } from "./InstallTabs";

export function Hero() {
  return (
    <section className="relative pt-12 pb-16 sm:pt-20 sm:pb-24 overflow-hidden">
      {/* Background radial gradient glow */}
      <div className="absolute top-1/4 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[350px] bg-sky-500/10 blur-[130px] rounded-full pointer-events-none -z-10" />

      <div className="mx-auto max-w-5xl px-4 sm:px-6 text-center">
        {/* Release / Announcement pill */}
        <div className="inline-flex items-center gap-2 rounded-full border border-sky-500/30 bg-sky-500/10 px-3.5 py-1 text-xs font-medium text-sky-300 backdrop-blur-md mb-8 hover:bg-sky-500/15 transition shadow-sm shadow-sky-500/10">
          <span className="flex h-2 w-2 rounded-full bg-sky-400 animate-pulse" />
          <span>Port Detective v1.1.1 Released</span>
          <span className="text-slate-500">·</span>
          <a
            href="https://github.com/Khoa180806/Port_Detective/releases/latest"
            target="_blank"
            rel="noopener noreferrer"
            className="flex items-center gap-1 hover:text-white transition-colors"
          >
            Changelog <ArrowRight className="h-3 w-3" />
          </a>
        </div>

        {/* Main Headline */}
        <h1 className="text-4xl sm:text-6xl font-extrabold tracking-tight text-white max-w-4xl mx-auto leading-[1.15]">
          Stop wrestling with <code className="text-slate-400 font-mono text-[0.85em] bg-slate-900 border border-slate-800 px-2 py-0.5 rounded-lg">netstat</code>.
          <br />
          <span className="bg-gradient-to-r from-sky-400 via-cyan-300 to-indigo-400 bg-clip-text text-transparent">
            Investigate & kill ports in 1 second.
          </span>
        </h1>

        {/* Subtitle */}
        <p className="mt-6 text-base sm:text-lg text-slate-300 max-w-2xl mx-auto leading-relaxed">
          A lightning-fast, native Go CLI tool to inspect zombie processes, safely terminate port-hogging servers with confirmation prompts, and scan port ranges concurrently.
        </p>

        {/* Value badges */}
        <div className="mt-8 flex flex-wrap items-center justify-center gap-3 sm:gap-6 text-xs sm:text-sm text-slate-300">
          <div className="flex items-center gap-1.5 rounded-md bg-slate-900/60 border border-slate-800/80 px-2.5 py-1">
            <Zap className="h-3.5 w-3.5 text-amber-400" />
            <span>Sub-50ms execution</span>
          </div>
          <div className="flex items-center gap-1.5 rounded-md bg-slate-900/60 border border-slate-800/80 px-2.5 py-1">
            <ShieldCheck className="h-3.5 w-3.5 text-emerald-400" />
            <span>Interactive Safe Kill [y/N]</span>
          </div>
          <div className="flex items-center gap-1.5 rounded-md bg-slate-900/60 border border-slate-800/80 px-2.5 py-1">
            <Terminal className="h-3.5 w-3.5 text-sky-400" />
            <span>Windows, macOS & Linux</span>
          </div>
        </div>

        {/* Quick Install Tabs Component */}
        <div id="install" className="mt-10 sm:mt-12 scroll-mt-24">
          <InstallTabs />
        </div>

        {/* Alternative CTAs */}
        <div className="mt-6 flex flex-wrap items-center justify-center gap-4 text-xs sm:text-sm text-slate-400">
          <span>Or explore source code:</span>
          <a
            href="https://github.com/Khoa180806/Port_Detective"
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex items-center gap-1.5 text-slate-300 hover:text-white underline underline-offset-4 decoration-slate-700 hover:decoration-sky-400 transition"
          >
            <GithubIcon className="h-3.5 w-3.5" />
            <span>GitHub Repository (Khoa180806/Port_Detective)</span>
          </a>
        </div>
      </div>
    </section>
  );
}
