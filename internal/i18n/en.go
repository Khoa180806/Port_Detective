package i18n

func init() {
	RegisterMessages("en", enMessages)
}

var enMessages = MessageMap{
	// Root command
	"root.short": "Port Detective - Cross-platform CLI tool to investigate and manage network ports",
	"root.long": `Port Detective is a cross-platform command-line interface (CLI) written in Go.
It helps you quickly identify which process is occupying a specific network port,
inspect process details, and optionally terminate (kill) the process.

Supported operating systems: Windows, macOS, Linux.`,

	// Flags
	"flag.json":     "Output result in JSON format",
	"flag.force":    "Force kill process without interactive confirmation",
	"flag.dry_run":  "Preview processes that would be terminated without executing kill",
	"flag.lang":     "Specify language for CLI output (en, vi)",

	// Check command
	"check.short":                "Check the process listening on a specified port",
	"check.long":                 "Inspect and display details of the process currently occupying or listening on the specified port.",
	"check.error.invalid_port":   "Error: Port must be an integer between 1 and 65535.",
	"check.error.permission":     "Error: Permission denied while accessing process information.",
	"check.error.permission_hint": "Hint: Try running the command with Administrator privileges (or sudo on Linux/macOS).",
	"check.error.lookup":         "Error looking up port: %v",
	"check.not_found":            "No process found running on port %d.",

	// Kill command
	"kill.short":                 "Terminate processes occupying a specified port",
	"kill.long": `Check which processes are occupying the specified network port and terminate them.

By default, this command displays process information and asks for confirmation [y/N] before killing.
Use --force (-f) to skip confirmation.`,
	"kill.error.invalid_port":    "Error: Port must be an integer between 1 and 65535.",
	"kill.error.lookup":          "Error looking up port: %v",
	"kill.not_found":             "No process found running on port %d.",
	"kill.warning":               "WARNING: Detected process(es) occupying port:",
	"kill.dry_run":               "[DRY RUN] No processes will be terminated.",
	"kill.confirm":               "Are you sure you want to KILL all processes above? [y/N]: ",
	"kill.cancelled":             "Operation cancelled.",
	"kill.killing":               "Terminating PID %d (%s)... ",
	"kill.failed.permission":     "FAILED: Permission denied. Administrator or root privileges required.",
	"kill.failed":                "FAILED: %v",
	"kill.success":               "SUCCESS",

	// Scan command
	"scan.short":                 "Scan a range of network ports to find active processes",
	"scan.long": `The scan command inspects multiple ports within a specified range.
Example: port-detective scan 3000-3010
Valid port range is from 1 to 65535. To protect system performance, range is limited to 5000 ports per scan.`,
	"scan.error.format":          "Error: Invalid port range format. Use <start>-<end> (e.g. 3000-3010).",
	"scan.error.range":           "Error: Invalid port range. Ensure 1 <= start <= end <= 65535.",
	"scan.error.too_large":       "Error: Port range too large (max 5000 ports per scan) to avoid system overload.",
	"scan.scanning":              "Scanning ports from %d to %d...",
	"scan.no_results":            "Great! No processes are occupying this port range.",
	"scan.permission_warn":       "\n[Notice] Access was denied (Permission Denied) on some ports during the scan.",
	"scan.permission_hint":       "Hint: Re-run with Administrator/sudo privileges for comprehensive results.",

	// Output format
	"output.occupied_by":         "Port %s is occupied by:\n",
	"output.found_processes":     "Found %d processes:\n",
	"output.no_processes":        "No processes found.",

	// Internal errors
	"error.permission_denied":    "permission denied (try running with Administrator or sudo privileges)",
	"error.process_not_found":    "process not found",
	"error.kill_failed":          "failed to kill process",
	"error.not_supported":        "operation not supported on this operating system",
	"error.netstat_failed":       "error running netstat: %v",
	"error.lsof_failed":          "error running lsof: %v",
	"error.kill_pid_failed":      "failed to kill PID %d: %s",
	"error.lsof_procfs_failed":   "both lsof and procfs failed: %v",
}
