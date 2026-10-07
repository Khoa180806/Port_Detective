import { ImageResponse } from "next/og";

export const alt = "Port Detective — Cross-Platform Port Investigator CLI";
export const size = {
  width: 1200,
  height: 630,
};
export const contentType = "image/png";

export default async function Image() {
  return new ImageResponse(
    (
      <div
        style={{
          height: "100%",
          width: "100%",
          display: "flex",
          flexDirection: "column",
          justifyContent: "space-between",
          padding: "80px",
          backgroundColor: "#090d16",
          color: "#f1f5f9",
          fontFamily: "sans-serif",
          position: "relative",
        }}
      >
        {/* Top Header Badge */}
        <div style={{ display: "flex", alignItems: "center", gap: "16px" }}>
          <div
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              width: "56px",
              height: "56px",
              borderRadius: "14px",
              backgroundColor: "rgba(14, 165, 233, 0.15)",
              border: "1px solid rgba(14, 165, 233, 0.4)",
              color: "#38bdf8",
              fontSize: "28px",
              fontWeight: "bold",
            }}
          >
            &gt;_
          </div>
          <span style={{ fontSize: "32px", fontWeight: "bold", color: "#ffffff" }}>
            Port Detective
          </span>
          <span
            style={{
              display: "flex",
              padding: "4px 12px",
              borderRadius: "9999px",
              fontSize: "16px",
              fontWeight: "600",
              backgroundColor: "rgba(14, 165, 233, 0.15)",
              color: "#38bdf8",
              border: "1px solid rgba(14, 165, 233, 0.3)",
            }}
          >
            v1.1.2
          </span>
        </div>

        {/* Center Tagline */}
        <div style={{ display: "flex", flexDirection: "column", gap: "20px" }}>
          <div
            style={{
              display: "flex",
              flexDirection: "column",
              fontSize: "60px",
              fontWeight: 800,
              lineHeight: 1.15,
              color: "#ffffff",
              letterSpacing: "-0.02em",
            }}
          >
            <span>Stop wrestling with netstat.</span>
            <span style={{ color: "#38bdf8" }}>
              Kill port conflicts in 1 second.
            </span>
          </div>

          <div
            style={{
              display: "flex",
              fontSize: "24px",
              color: "#94a3b8",
              maxWidth: "900px",
              lineHeight: 1.5,
            }}
          >
            A lightning-fast, native Go CLI tool to investigate, kill, and scan network ports across Windows, macOS, and Linux.
          </div>
        </div>

        {/* Bottom Feature Badges */}
        <div style={{ display: "flex", alignItems: "center", gap: "24px", fontSize: "18px", color: "#cbd5e1" }}>
          <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
            <span style={{ color: "#38bdf8" }}>⚡</span>
            <span>Sub-50ms execution</span>
          </div>
          <span style={{ color: "#475569" }}>·</span>
          <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
            <span style={{ color: "#34d399" }}>🛡️</span>
            <span>Interactive Safe Kill [y/N]</span>
          </div>
          <span style={{ color: "#475569" }}>·</span>
          <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
            <span style={{ color: "#a78bfa" }}>🌍</span>
            <span>Cross-Platform (Win / Mac / Linux)</span>
          </div>
        </div>
      </div>
    ),
    {
      ...size,
    }
  );
}
