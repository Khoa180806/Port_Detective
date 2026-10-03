const fs = require('fs');

const html = fs.readFileSync('.archify/pd-arch/architecture.html', 'utf8');

// 1. Get SVG element
const svgMatch = html.match(/<svg[\s\S]*?<\/svg>/i);
if (!svgMatch) {
  console.error('No SVG found');
  process.exit(1);
}
let svg = svgMatch[0];

// Check xmlns
if (!svg.includes('xmlns="http://www.w3.org/2000/svg"')) {
  svg = svg.replace('<svg ', '<svg xmlns="http://www.w3.org/2000/svg" ');
}

// 2. Extract styles from HTML
const styleMatches = html.match(/<style[^>]*>([\s\S]*?)<\/style>/gi) || [];
// We need the fonts and the diagram styles
let stylesCombined = '';
for (const s of styleMatches) {
  // Strip <style...> and </style>
  const content = s.replace(/^<style[^>]*>/i, '').replace(/<\/style>$/i, '');
  stylesCombined += content + '\n';
}

// Add dark theme variables explicitly mapped to :root inside SVG so it doesn't depend on html data-theme
const svgThemeOverrides = `
:root {
  --bg: #090d16;
  --grid: #1e293b;
  --canvas-dot: rgba(148, 163, 184, 0.16);
  --text: #f8fafc;
  --text-muted: #94a3b8;
  --text-dim: #64748b;
  --mask: #090d16;
  --panel: #0f172a;
  --panel-border: rgba(148, 163, 184, 0.2);
  --selection: rgba(56, 189, 248, 0.2);
  --accent: #38bdf8;
  --accent-emphasis: #0284c7;
  --accent-muted: rgba(56, 189, 248, 0.15);
  font-family: 'JetBrains Mono', ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
}
svg {
  background-color: #090d16;
  border-radius: 8px;
}
`;

// Insert <style> inside <defs> of SVG
const defsIndex = svg.indexOf('<defs>');
if (defsIndex !== -1) {
  const insertPos = defsIndex + '<defs>'.length;
  svg = svg.slice(0, insertPos) + '\n<style>\n' + stylesCombined + '\n' + svgThemeOverrides + '\n</style>\n' + svg.slice(insertPos);
} else {
  // insert before first child
  svg = svg.replace('>', '><defs><style>' + stylesCombined + '\n' + svgThemeOverrides + '</style></defs>');
}

fs.writeFileSync('assets/architecture.svg', svg);
console.log('Successfully wrote self-contained assets/architecture.svg, size:', svg.length);
