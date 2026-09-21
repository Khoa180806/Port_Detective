package output

import (
	"fmt"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
	"github.com/Khoa180806/Port_Detective/internal/process"
	"github.com/fatih/color"
)

var (
	colorPort     = color.New(color.FgCyan, color.Bold)
	colorPID      = color.New(color.FgRed, color.Bold)
	colorProcess  = color.New(color.FgYellow)
	colorCommand  = color.New(color.FgHiBlack) // Dark gray
	colorProtocol = color.New(color.FgGreen)
)

// FormatText returns a formatted colored string for a single process.
func FormatText(p *process.ProcessInfo) string {
	if p == nil {
		return i18n.T("output.no_info")
	}

	var sb strings.Builder
	portStr := colorPort.Sprintf("%d", p.Port)
	sb.WriteString(i18n.Tr("output.occupied_by", portStr))

	sb.WriteString(fmt.Sprintf("  PID:      %s\n", colorPID.Sprintf("%d", p.PID)))
	sb.WriteString(fmt.Sprintf("  Process:  %s\n", colorProcess.Sprintf("%s", p.Name)))
	sb.WriteString(fmt.Sprintf("  Command:  %s\n", colorCommand.Sprintf("%s", p.Command)))
	sb.WriteString(fmt.Sprintf("  Protocol: %s", colorProtocol.Sprintf("%s", p.Protocol)))

	return sb.String()
}

// FormatTextMultiple returns a formatted colored string for multiple processes.
func FormatTextMultiple(ps []process.ProcessInfo) string {
	if len(ps) == 0 {
		return i18n.T("output.no_processes")
	}

	if len(ps) == 1 {
		return FormatText(&ps[0])
	}

	var sb strings.Builder
	titleText := i18n.Tr("output.found_processes", len(ps))
	sb.WriteString(color.New(color.FgHiWhite, color.Bold).Sprint(titleText))

	for i, p := range ps {
		if i > 0 {
			sb.WriteString(color.HiBlackString("\n--------------------------------------------------\n"))
		}
		sb.WriteString(fmt.Sprintf("  Port:     %s\n", colorPort.Sprintf("%d", p.Port)))
		sb.WriteString(fmt.Sprintf("  PID:      %s\n", colorPID.Sprintf("%d", p.PID)))
		sb.WriteString(fmt.Sprintf("  Process:  %s\n", colorProcess.Sprintf("%s", p.Name)))
		sb.WriteString(fmt.Sprintf("  Command:  %s\n", colorCommand.Sprintf("%s", p.Command)))
		sb.WriteString(fmt.Sprintf("  Protocol: %s", colorProtocol.Sprintf("%s", p.Protocol)))
	}

	return sb.String()
}
