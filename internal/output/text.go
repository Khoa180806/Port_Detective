package output

import (
	"fmt"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/process"
	"github.com/fatih/color"
)

var (
	colorPort     = color.New(color.FgCyan, color.Bold)
	colorPID      = color.New(color.FgRed, color.Bold)
	colorProcess  = color.New(color.FgYellow)
	colorCommand  = color.New(color.FgHiBlack) // Xám đậm
	colorProtocol = color.New(color.FgGreen)
)

// FormatText trả về chuỗi định dạng văn bản trực quan cho một tiến trình, có màu sắc.
func FormatText(p *process.ProcessInfo) string {
	if p == nil {
		return "No process information available."
	}

	var sb strings.Builder
	portStr := colorPort.Sprintf("%d", p.Port)
	sb.WriteString(fmt.Sprintf("Port %s is occupied by:\n", portStr))
	
	sb.WriteString(fmt.Sprintf("  PID:      %s\n", colorPID.Sprintf("%d", p.PID)))
	sb.WriteString(fmt.Sprintf("  Process:  %s\n", colorProcess.Sprintf("%s", p.Name)))
	sb.WriteString(fmt.Sprintf("  Command:  %s\n", colorCommand.Sprintf("%s", p.Command)))
	sb.WriteString(fmt.Sprintf("  Protocol: %s", colorProtocol.Sprintf("%s", p.Protocol)))

	return sb.String()
}

// FormatTextMultiple trả về chuỗi định dạng văn bản có màu sắc cho danh sách các tiến trình.
func FormatTextMultiple(ps []process.ProcessInfo) string {
	if len(ps) == 0 {
		return "No processes found."
	}

	if len(ps) == 1 {
		return FormatText(&ps[0])
	}

	var sb strings.Builder
	sb.WriteString(color.New(color.FgHiWhite, color.Bold).Sprintf("Found %d processes:\n", len(ps)))

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
