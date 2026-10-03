"use client";

import { useLanguage } from "@/context/LanguageContext";
import { Globe } from "lucide-react";

export function LanguageSwitcher() {
  const { locale, setLocale } = useLanguage();

  return (
    <div className="flex items-center rounded-lg border border-slate-700 bg-slate-900/90 p-0.5 text-xs font-medium text-slate-300">
      <div className="flex items-center px-1.5 text-slate-500">
        <Globe className="h-3.5 w-3.5" />
      </div>
      <button
        onClick={() => setLocale("en")}
        aria-label="Switch language to English"
        className={`rounded-md px-2 py-1 transition-all ${
          locale === "en"
            ? "bg-sky-500 text-slate-950 font-bold shadow-sm"
            : "text-slate-400 hover:text-white"
        }`}
      >
        EN
      </button>
      <button
        onClick={() => setLocale("vi")}
        aria-label="Chuyển sang Tiếng Việt"
        className={`rounded-md px-2 py-1 transition-all ${
          locale === "vi"
            ? "bg-sky-500 text-slate-950 font-bold shadow-sm"
            : "text-slate-400 hover:text-white"
        }`}
      >
        VI
      </button>
    </div>
  );
}
