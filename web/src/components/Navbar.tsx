"use client";

import { useState } from "react";
import { Terminal, Menu, X, BookOpen, Download } from "lucide-react";
import { GithubIcon } from "./GithubIcon";

export function Navbar() {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  return (
    <header className="sticky top-0 z-50 w-full border-b border-slate-800/80 bg-slate-950/75 backdrop-blur-md">
      <div className="mx-auto flex h-16 max-w-6xl items-center justify-between px-4 sm:px-6">
        {/* Brand Logo */}
        <a href="#" className="flex items-center gap-2.5 transition hover:opacity-90">
          <div className="flex h-9 w-9 items-center justify-center rounded-lg bg-sky-500/10 border border-sky-500/30 text-sky-400 shadow-sm shadow-sky-500/20">
            <Terminal className="h-5 w-5" />
          </div>
          <span className="font-semibold text-lg tracking-tight text-white flex items-center gap-2">
            Port Detective
            <span className="rounded-full bg-sky-500/10 px-2 py-0.5 text-xs font-mono font-medium text-sky-400 border border-sky-500/20">
              v1.1.1
            </span>
          </span>
        </a>

        {/* Desktop Nav Links */}
        <nav className="hidden md:flex items-center gap-7 text-sm font-medium text-slate-300">
          <a href="#features" className="hover:text-white transition-colors">
            Features
          </a>
          <a href="#demo" className="hover:text-white transition-colors">
            Demo
          </a>
          <a href="#benchmark" className="hover:text-white transition-colors">
            Benchmark
          </a>
          <a
            href="https://github.com/Khoa180806/Port_Detective/blob/master/docs/cli-reference.md"
            target="_blank"
            rel="noopener noreferrer"
            className="flex items-center gap-1.5 hover:text-white transition-colors"
          >
            <BookOpen className="h-4 w-4" />
            Docs
          </a>
        </nav>

        {/* Action Buttons */}
        <div className="hidden md:flex items-center gap-3">
          <a
            href="https://github.com/Khoa180806/Port_Detective"
            target="_blank"
            rel="noopener noreferrer"
            className="flex items-center gap-2 rounded-lg border border-slate-700 bg-slate-900/80 px-3.5 py-1.5 text-sm font-medium text-slate-200 transition hover:bg-slate-800 hover:text-white hover:border-slate-600"
          >
            <GithubIcon className="h-4 w-4" />
            <span>GitHub</span>
          </a>
          <a
            href="#install"
            className="flex items-center gap-1.5 rounded-lg bg-sky-500 px-3.5 py-1.5 text-sm font-medium text-slate-950 transition hover:bg-sky-400 font-semibold shadow-sm shadow-sky-500/30"
          >
            <Download className="h-4 w-4" />
            <span>Install</span>
          </a>
        </div>

        {/* Mobile menu toggle */}
        <button
          type="button"
          aria-label="Toggle Navigation Menu"
          onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
          className="md:hidden rounded-lg p-2 text-slate-400 hover:bg-slate-800 hover:text-white"
        >
          {mobileMenuOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
        </button>
      </div>

      {/* Mobile Menu Dropdown */}
      {mobileMenuOpen && (
        <div className="md:hidden border-b border-slate-800 bg-slate-950/95 px-4 pt-3 pb-5 space-y-3">
          <a
            href="#features"
            onClick={() => setMobileMenuOpen(false)}
            className="block py-2 text-sm font-medium text-slate-300 hover:text-white"
          >
            Features
          </a>
          <a
            href="#demo"
            onClick={() => setMobileMenuOpen(false)}
            className="block py-2 text-sm font-medium text-slate-300 hover:text-white"
          >
            Demo
          </a>
          <a
            href="#benchmark"
            onClick={() => setMobileMenuOpen(false)}
            className="block py-2 text-sm font-medium text-slate-300 hover:text-white"
          >
            Benchmark
          </a>
          <a
            href="https://github.com/Khoa180806/Port_Detective/blob/master/docs/cli-reference.md"
            target="_blank"
            rel="noopener noreferrer"
            className="flex items-center gap-2 py-2 text-sm font-medium text-slate-300 hover:text-white"
          >
            <BookOpen className="h-4 w-4" />
            Docs
          </a>
          <div className="pt-2 flex gap-3">
            <a
              href="https://github.com/Khoa180806/Port_Detective"
              target="_blank"
              rel="noopener noreferrer"
              className="flex-1 flex justify-center items-center gap-2 rounded-lg border border-slate-700 bg-slate-900 py-2 text-sm font-medium text-slate-200"
            >
              <GithubIcon className="h-4 w-4" />
              GitHub
            </a>
            <a
              href="#install"
              onClick={() => setMobileMenuOpen(false)}
              className="flex-1 flex justify-center items-center gap-1.5 rounded-lg bg-sky-500 py-2 text-sm font-semibold text-slate-950"
            >
              <Download className="h-4 w-4" />
              Install
            </a>
          </div>
        </div>
      )}
    </header>
  );
}
