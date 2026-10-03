import type { Metadata, Viewport } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const viewport: Viewport = {
  themeColor: "#090d16",
  width: "device-width",
  initialScale: 1,
};

export const metadata: Metadata = {
  metadataBase: new URL("https://port-detective.vercel.app"),
  title: "Port Detective (pd) — Lightning-Fast Cross-Platform Port & Process Investigation CLI",
  description:
    "A lightning-fast, native Go CLI tool to investigate, kill, and scan processes occupying network ports across Windows, macOS, and Linux. Sub-50ms execution, safe interactive confirmation, and zero external runtime dependencies.",
  keywords: [
    "port detective",
    "port-detective",
    "pd cli",
    "port scanner",
    "kill port",
    "listen EADDRINUSE",
    "netstat alternative",
    "lsof alternative",
    "golang cli",
    "cross-platform devtool",
  ],
  authors: [{ name: "Khoa180806", url: "https://github.com/Khoa180806" }],
  creator: "Khoa180806",
  icons: {
    icon: "/icon.svg",
    shortcut: "/icon.svg",
    apple: "/icon.svg",
  },
  openGraph: {
    type: "website",
    locale: "en_US",
    alternateLocale: "vi_VN",
    url: "https://port-detective.vercel.app",
    siteName: "Port Detective",
    title: "Port Detective (pd) — Investigate & Kill Port Conflicts in 1 Second",
    description:
      "Stop wrestling with netstat and convoluted commands. Free up occupied ports in 1 second with a unified cross-platform Go CLI.",
  },
  twitter: {
    card: "summary_large_image",
    title: "Port Detective (pd) — Cross-Platform Port Investigator CLI",
    description:
      "A lightning-fast Go CLI tool to investigate and kill processes occupying network ports. Sub-50ms execution, safe confirmation, zero runtime dependencies.",
  },
  robots: {
    index: true,
    follow: true,
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html
      lang="en"
      className={`${geistSans.variable} ${geistMono.variable} h-full antialiased`}
    >
      <body className="min-h-full flex flex-col">{children}</body>
    </html>
  );
}
