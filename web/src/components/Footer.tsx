"use client";

import { useLanguage } from "@/context/LanguageContext";
import { GithubIcon } from "./GithubIcon";
import { Terminal, ExternalLink } from "lucide-react";

export function Footer() {
  const { t } = useLanguage();

  return (
    <footer className="border-t border-slate-800 bg-slate-950/90 text-slate-400">
      <div className="mx-auto max-w-6xl px-4 sm:px-6 py-12 sm:py-16">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8 mb-12">
          {/* Col 1: Brand & Tagline */}
          <div className="md:col-span-2 space-y-4">
            <a href="#" className="flex items-center gap-2.5 transition hover:opacity-90">
              <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-sky-500/10 border border-sky-500/30 text-sky-400 shadow-sm shadow-sky-500/20">
                <Terminal className="h-4 w-4" />
              </div>
              <span className="font-semibold text-base tracking-tight text-white flex items-center gap-2">
                Port Detective
                <span className="rounded-full bg-sky-500/10 px-2 py-0.5 text-[10px] font-mono font-medium text-sky-400 border border-sky-500/20">
                  v1.1.2
                </span>
              </span>
            </a>
            <p className="text-sm text-slate-400 max-w-sm leading-relaxed">
              {t.footer.tagline}
            </p>
            <div className="flex items-center gap-3 pt-2">
              <a
                href="https://github.com/Khoa180806/Port_Detective"
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center gap-1.5 rounded-lg border border-slate-800 bg-slate-900 px-3 py-1.5 text-xs text-slate-300 hover:text-white hover:border-slate-700 transition"
              >
                <GithubIcon className="h-3.5 w-3.5" />
                <span>Khoa180806/Port_Detective</span>
              </a>
            </div>
          </div>

          {/* Col 2: Documentation Links */}
          <div className="space-y-3">
            <h4 className="text-xs font-semibold uppercase tracking-wider text-slate-200">
              {t.footer.resources}
            </h4>
            <ul className="space-y-2 text-sm">
              <li>
                <a
                  href="https://github.com/Khoa180806/Port_Detective/blob/master/docs/02-technical/architecture.md"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="hover:text-white transition flex items-center gap-1"
                >
                  <span>{t.footer.architecture}</span>
                  <ExternalLink className="h-3 w-3 opacity-60" />
                </a>
              </li>
              <li>
                <a
                  href="https://github.com/Khoa180806/Port_Detective/blob/master/docs/02-technical/api-reference.md"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="hover:text-white transition flex items-center gap-1"
                >
                  <span>{t.footer.cliRef}</span>
                  <ExternalLink className="h-3 w-3 opacity-60" />
                </a>
              </li>
              <li>
                <a
                  href="https://github.com/Khoa180806/Port_Detective/blob/master/docs/01-overview/getting-started.md"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="hover:text-white transition flex items-center gap-1"
                >
                  <span>{t.footer.installGuide}</span>
                  <ExternalLink className="h-3 w-3 opacity-60" />
                </a>
              </li>
              <li>
                <a
                  href="https://github.com/Khoa180806/Port_Detective/blob/master/docs/01-overview/development-workflow.md"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="hover:text-white transition flex items-center gap-1"
                >
                  <span>{t.footer.contributing}</span>
                  <ExternalLink className="h-3 w-3 opacity-60" />
                </a>
              </li>
            </ul>
          </div>

          {/* Col 3: Community & Releases */}
          <div className="space-y-3">
            <h4 className="text-xs font-semibold uppercase tracking-wider text-slate-200">
              {t.footer.community}
            </h4>
            <ul className="space-y-2 text-sm">
              <li>
                <a
                  href="https://github.com/Khoa180806/Port_Detective/releases/tag/v1.1.2"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="hover:text-white transition"
                >
                  {t.footer.releases}
                </a>
              </li>
              <li>
                <a
                  href="https://github.com/Khoa180806/Port_Detective/issues"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="hover:text-white transition"
                >
                  {t.footer.issues}
                </a>
              </li>
              <li>
                <a
                  href="https://github.com/Khoa180806/Port_Detective/blob/master/LICENSE"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="hover:text-white transition"
                >
                  {t.footer.mitLicense}
                </a>
              </li>
            </ul>
          </div>
        </div>

        {/* Bottom copyright line */}
        <div className="border-t border-slate-800/80 pt-8 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs text-slate-500">
          <p>{t.footer.copyright}</p>
          <p>{t.footer.builtWith}</p>
        </div>
      </div>
    </footer>
  );
}
