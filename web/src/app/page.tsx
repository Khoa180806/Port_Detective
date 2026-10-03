import { Navbar } from "@/components/Navbar";
import { Hero } from "@/components/Hero";
import { TerminalPreview } from "@/components/TerminalPreview";

export default function Home() {
  return (
    <div className="min-h-screen flex flex-col bg-slate-950 text-slate-100 selection:bg-sky-500 selection:text-slate-950">
      <Navbar />
      <main className="flex-1">
        <Hero />
        <TerminalPreview />
      </main>
    </div>
  );
}
