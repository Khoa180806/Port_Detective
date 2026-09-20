package output

import (
	"fmt"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

// FormatText trả về chuỗi định dạng văn bản trực quan cho một tiến trình.
func FormatText(p *process.ProcessInfo) string {
	if p == nil {
		return "No process information available."
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Port %d is occupied by:\n", p.Port))
	sb.WriteString(fmt.Sprintf("  PID:      %d\n", p.PID))
	sb.WriteString(fmt.Sprintf("  Process:  %s\n", p.Name))
	sb.WriteString(fmt.Sprintf("  Command:  %s\n", p.Command))
	sb.WriteString(fmt.Sprintf("  Protocol: %s", p.Protocol))

	return sb.String()
}

// FormatTextMultiple trả về chuỗi định dạng văn bản trực quan cho danh sách các tiến trình.
func FormatTextMultiple(ps []process.ProcessInfo) string {
	if len(ps) == 0 {
		return "No processes found."
	}

	if len(ps) == 1 {
		return FormatText(&ps[0])
	}

	// Trường hợp có nhiều tiến trình cùng chiếm 1 port (SO_REUSEPORT) hoặc list scan
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Found %d processes:\n", len(ps)))

	for i, p := range ps {
		if i > 0 {
			sb.WriteString("\n--------------------------------------------------\n")
		}
		sb.WriteString(fmt.Sprintf("  Port:     %d\n", p.Port))
		sb.WriteString(fmt.Sprintf("  PID:      %d\n", p.PID))
		sb.WriteString(fmt.Sprintf("  Process:  %s\n", p.Name))
		sb.WriteString(fmt.Sprintf("  Command:  %s\n", p.Command))
		sb.WriteString(fmt.Sprintf("  Protocol: %s", p.Protocol))
	}

	return sb.String()
}
