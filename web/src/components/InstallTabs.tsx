"use client";

import { useState } from "react";
import { Check, Copy, Terminal } from "lucide-react";
import { useLanguage } from "@/context/LanguageContext";

type Platform = "windows" | "unix" | "go";

export function InstallTabs() {
  const [activeTab, setActiveTab] = useState<Platform>("windows");
  const [copied, setCopied] = useState(false);
  const { t } = useLanguage();

  const INSTALL_OPTIONS = [
    {
      id: "windows" as Platform,
      label: t.install.windows,
      osBadge: "Windows x64 / ARM64",
      command: "iwr -useb https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.ps1 | iex",
      comment: t.install.winComment,
    },
    {
      id: "unix" as Platform,
      label: t.install.unix,
      osBadge: "macOS / Linux",
      command: "curl -sSfL https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.sh | sh",
      comment: t.install.unixComment,
    },
    {
      id: "go" as Platform,
      label: t.install.go,
      osBadge: "Go 1.21+",
      command: "go install github.com/Khoa180806/Port_Detective/cmd/pd@latest",
      comment: t.install.goComment,
    },
  ];

  const activeOption = INSTALL_OPTIONS.find((opt) => opt.id === activeTab) || INSTALL_OPTIONS[0];

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(activeOption.command);
      setCopied(true);
      const timer = setTimeout(() => setCopied(false), 2000);
      return () => clearTimeout(timer);
    } catch {
      // Fallback
    }
  };

  return (
    <div className="w-full max-w-2xl mx-auto rounded-xl border border-slate-800 bg-slate-900/90 shadow-2xl backdrop-blur-md overflow-hidden">
      {/* Tab Selectors */}
      <div className="flex border-b border-slate-800/80 bg-slate-950/70 p-1.5 sm:px-3 sm:py-2 items-center justify-between gap-1 overflow-x-auto">
        <div role="tablist" aria-label="Operating system installation choices" className="flex items-center gap-1 sm:gap-1.5">
          {INSTALL_OPTIONS.map((opt) => (
            <button
              key={opt.id}
              role="tab"
              aria-selected={activeTab === opt.id}
              onClick={() => {
                setActiveTab(opt.id);
                setCopied(false);
              }}
              className={`rounded-lg px-3 py-1.5 text-xs sm:text-sm font-medium transition-all ${
                activeTab === opt.id
                  ? "bg-slate-800 text-sky-400 font-semibold shadow-inner border border-slate-700/60"
                  : "text-slate-400 hover:text-slate-200 hover:bg-slate-900"
              }`}
            >
              {opt.label}
            </button>
          ))}
        </div>
        <span className="hidden sm:inline-block text-[11px] font-mono text-slate-500 uppercase tracking-wider">
          {activeOption.osBadge}
        </span>
      </div>

      {/* Code Display Area */}
      <div className="relative p-4 sm:p-5 font-mono text-xs sm:text-sm">
        <div className="text-slate-500 text-[11px] sm:text-xs mb-2 select-none">
          {activeOption.comment}
        </div>
        <div className="flex items-start justify-between gap-3 text-slate-200 overflow-x-auto pr-12 pb-1">
          <span className="text-sky-500 select-none font-bold mr-1">$</span>
          <span className="flex-1 whitespace-pre-wrap break-all leading-relaxed font-mono">
            {activeOption.command}
          </span>
        </div>

        {/* Copy Button */}
        <button
          onClick={handleCopy}
          aria-label="Copy installation command"
          className="absolute right-3.5 bottom-3.5 rounded-lg border border-slate-700 bg-slate-800/90 p-2 text-slate-300 transition-all hover:bg-slate-700 hover:text-white hover:border-slate-600 focus:outline-none focus:ring-2 focus:ring-sky-500"
        >
          {copied ? (
            <div className="flex items-center gap-1.5 text-emerald-400 text-xs font-sans px-1">
              <Check className="h-4 w-4" />
              <span className="font-medium">{t.install.copied}</span>
            </div>
          ) : (
            <Copy className="h-4 w-4" />
          )}
        </button>
      </div>

      {/* Verification footer */}
      <div className="border-t border-slate-800/80 bg-slate-950/40 px-4 py-2 flex items-center justify-between text-[11px] text-slate-500 font-mono">
        <div className="flex items-center gap-1.5">
          <Terminal className="h-3 w-3 text-sky-500/80" />
          <span>{t.install.verify} <code className="text-slate-400">pd --version</code></span>
        </div>
        <span className="text-emerald-500/90 flex items-center gap-1">
          <span className="inline-block h-1.5 w-1.5 rounded-full bg-emerald-400 animate-pulse"></span>
          {t.install.readyImmediately}
        </span>
      </div>
    </div>
  );
}
