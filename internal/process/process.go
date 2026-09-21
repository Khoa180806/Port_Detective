package process

// ProcessInfo represents information about a process occupying or listening on a network port.
type ProcessInfo struct {
	PID      int    `json:"pid"`
	Name     string `json:"name"`
	Command  string `json:"command"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"` // "tcp" | "udp"
}
