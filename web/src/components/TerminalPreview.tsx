"use client";

import { useState } from "react";
import { Terminal, RotateCcw, Copy, Check } from "lucide-react";
import { useLanguage } from "@/context/LanguageContext";

type CommandType = "check" | "kill" | "scan" | "json";

interface CommandDemo {
  id: CommandType;
  title: string;
  command: string;
  description: string;
  outputLines: Array<{
    text: string;
    color?: string;
    bold?: boolean;
    prefix?: string;
  }>;
}

export function TerminalPreview() {
  const [activeTab, setActiveTab] = useState<CommandType>("check");
  const [copied, setCopied] = useState(false);
  const [isSimulating, setIsSimulating] = useState(false);
  const { t } = useLanguage();

  const DEMOS: Record<CommandType, CommandDemo> = {
    check: {
      id: "check",
      title: "pd check",
      command: "pd check 8080",
      description: t.demo.checkDesc,
      outputLines: [
        { text: "Port 8080 is occupied by:", color: "text-slate-200", bold: true },
        { text: "14280", color: "text-rose-400", prefix: "  PID:      ", bold: true },
        { text: "node.exe", color: "text-amber-300", prefix: "  Process:  " },
        { text: "node server.js", color: "text-slate-500", prefix: "  Command:  " },
        { text: "tcp", color: "text-emerald-400", prefix: "  Protocol: " },
      ],
    },
    kill: {
      id: "kill",
      title: "pd kill",
      command: "pd kill 8080",
      description: t.demo.killDesc,
      outputLines: [
        { text: "WARNING: You are about to kill the following process(es):", color: "text-rose-400", bold: true },
        { text: "Port 8080 is occupied by:", color: "text-slate-200" },
        { text: "14280", color: "text-rose-400", prefix: "  PID:      ", bold: true },
        { text: "node.exe", color: "text-amber-300", prefix: "  Process:  " },
        { text: "tcp", color: "text-emerald-400", prefix: "  Protocol: " },
        { text: "", color: "text-transparent" },
        { text: "Are you sure you want to terminate this process? [y/N]: y", color: "text-sky-300", bold: true },
        { text: "✔ Successfully terminated process 'node.exe' (PID: 14280).", color: "text-emerald-400", bold: true },
      ],
    },
    scan: {
      id: "scan",
      title: "pd scan",
      command: "pd scan 3000-3005",
      description: t.demo.scanDesc,
      outputLines: [
        { text: "Scanning ports 3000 to 3005...", color: "text-sky-400" },
        { text: "Found 2 active process(es):", color: "text-slate-100", bold: true },
        { text: "3000", color: "text-cyan-400", prefix: "  Port:     ", bold: true },
        { text: "18204", color: "text-rose-400", prefix: "  PID:      " },
        { text: "node.exe", color: "text-amber-300", prefix: "  Process:  " },
        { text: "tcp", color: "text-emerald-400", prefix: "  Protocol: " },
        { text: "--------------------------------------------------", color: "text-slate-700" },
        { text: "3003", color: "text-cyan-400", prefix: "  Port:     ", bold: true },
        { text: "9142", color: "text-rose-400", prefix: "  PID:      " },
        { text: "docker-proxy", color: "text-amber-300", prefix: "  Process:  " },
        { text: "tcp", color: "text-emerald-400", prefix: "  Protocol: " },
      ],
    },
    json: {
      id: "json",
      title: "pd --json",
      command: "pd check 8080 --json",
      description: t.demo.jsonDesc,
      outputLines: [
        { text: "[", color: "text-slate-400" },
        { text: "  {", color: "text-slate-400" },
        { text: '    "pid": 14280,', color: "text-rose-400 font-semibold" },
        { text: '    "name": "node.exe",', color: "text-amber-300" },
        { text: '    "command": "node server.js",', color: "text-slate-300" },
        { text: '    "port": 8080,', color: "text-cyan-300" },
        { text: '    "protocol": "tcp"', color: "text-emerald-400" },
        { text: "  }", color: "text-slate-400" },
        { text: "]", color: "text-slate-400" },
      ],
    },
  };

  const demo = DEMOS[activeTab];

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(demo.command);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch {
      // Fallback
    }
  };

  const handleRunSimulation = () => {
    setIsSimulating(true);
    setTimeout(() => {
      setIsSimulating(false);
    }, 400);
  };

  return (
    <section id="demo" className="relative pt-4 pb-10 sm:pt-6 sm:pb-12 scroll-mt-20">
      <div className="mx-auto max-w-5xl px-4 sm:px-6">
        {/* Section Header */}
        <div className="text-center max-w-2xl mx-auto mb-10 sm:mb-12">
          <div className="inline-flex items-center gap-2 rounded-full border border-sky-500/20 bg-sky-500/10 px-3 py-1 text-xs font-medium text-sky-400 mb-3">
            <Terminal className="h-3.5 w-3.5" />
            <span>{t.demo.badge}</span>
          </div>
          <h2 className="text-3xl sm:text-4xl font-bold tracking-tight text-white">
            {t.demo.title}
          </h2>
          <p className="mt-3 text-sm sm:text-base text-slate-400">
            {t.demo.subtitle}
          </p>
        </div>

        {/* Command Pill Buttons */}
        <div className="flex flex-wrap items-center justify-center gap-2 sm:gap-3 mb-6">
          {(Object.keys(DEMOS) as CommandType[]).map((key) => {
            const item = DEMOS[key];
            const isActive = activeTab === key;
            return (
              <button
                key={key}
                onClick={() => {
                  setActiveTab(key);
                  setCopied(false);
                  handleRunSimulation();
                }}
                className={`flex items-center gap-2 rounded-lg px-3.5 py-2 text-xs sm:text-sm font-mono font-medium transition-all ${
                  isActive
                    ? "bg-sky-500 text-slate-950 shadow-md shadow-sky-500/20 font-semibold"
                    : "bg-slate-900 border border-slate-800 text-slate-300 hover:bg-slate-800 hover:text-white"
                }`}
              >
                <span>{item.title}</span>
              </button>
            );
          })}
        </div>

        {/* Terminal Window Frame */}
        <div className="rounded-xl border border-slate-800 bg-slate-950 shadow-2xl overflow-hidden backdrop-blur-md">
          {/* Terminal Window Header (macOS style dots) */}
          <div className="flex items-center justify-between border-b border-slate-800/80 bg-slate-900/90 px-4 py-3">
            <div className="flex items-center gap-2">
              <span className="h-3 w-3 rounded-full bg-rose-500/80 inline-block" />
              <span className="h-3 w-3 rounded-full bg-amber-500/80 inline-block" />
              <span className="h-3 w-3 rounded-full bg-emerald-500/80 inline-block" />
              <span className="ml-2 font-mono text-xs text-slate-400 hidden sm:inline">
                terminal — port-detective
              </span>
            </div>

            <div className="flex items-center gap-2">
              <button
                onClick={handleRunSimulation}
                title="Rerun command simulation"
                className="flex items-center gap-1.5 rounded bg-slate-800 px-2 py-1 text-xs text-slate-300 hover:bg-slate-700 hover:text-white transition"
              >
                <RotateCcw className={`h-3 w-3 ${isSimulating ? "animate-spin" : ""}`} />
                <span className="hidden sm:inline">{t.demo.replay}</span>
              </button>
              <button
                onClick={handleCopy}
                title="Copy command"
                className="flex items-center gap-1.5 rounded bg-slate-800 px-2 py-1 text-xs text-slate-300 hover:bg-slate-700 hover:text-white transition"
              >
                {copied ? (
                  <>
                    <Check className="h-3 w-3 text-emerald-400" />
                    <span className="text-emerald-400">{t.demo.copied}</span>
                  </>
                ) : (
                  <>
                    <Copy className="h-3 w-3" />
                    <span className="hidden sm:inline">{t.demo.copy}</span>
                  </>
                )}
              </button>
            </div>
          </div>

          {/* Terminal Content Screen */}
          <div className="p-4 sm:p-6 font-mono text-xs sm:text-sm min-h-[280px] bg-slate-950/95 leading-relaxed overflow-x-auto">
            {/* Input Prompt Line */}
            <div className="flex items-center gap-2 text-slate-100 mb-4 select-none">
              <span className="text-emerald-400 font-bold">user@developer</span>
              <span className="text-slate-600">:</span>
              <span className="text-sky-400">~</span>
              <span className="text-slate-400">$</span>
              <span className="font-semibold text-white ml-1">{demo.command}</span>
            </div>

            {/* Output Lines with simulated loading */}
            {isSimulating ? (
              <div className="py-8 flex items-center justify-center gap-2 text-slate-500 text-xs">
                <span className="inline-block h-2 w-2 rounded-full bg-sky-400 animate-ping" />
                <span>{t.demo.executing}</span>
              </div>
            ) : (
              <div className="space-y-1">
                {demo.outputLines.map((line, index) => (
                  <div key={index} className="flex">
                    {line.prefix && (
                      <span className="text-slate-400 whitespace-pre">{line.prefix}</span>
                    )}
                    <span
                      className={`whitespace-pre ${line.color || "text-slate-300"} ${
                        line.bold ? "font-bold" : ""
                      }`}
                    >
                      {line.text}
                    </span>
                  </div>
                ))}
              </div>
            )}
          </div>

          {/* Terminal Status / Help text */}
          <div className="border-t border-slate-800/80 bg-slate-900/50 px-4 py-2.5 flex items-center justify-between text-xs text-slate-400 font-sans">
            <p className="truncate mr-4">
              <span className="text-slate-500 font-mono mr-1.5">{t.demo.infoPrefix}</span>
              {demo.description}
            </p>
            <div className="flex items-center gap-1.5 font-mono text-[11px] text-slate-500 shrink-0">
              <span className="inline-block h-1.5 w-1.5 rounded-full bg-emerald-400" />
              <span>{t.demo.exitCode}</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
