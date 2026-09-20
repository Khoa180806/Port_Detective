package process

// ProcessInfo đại diện cho thông tin của một tiến trình đang chiếm dụng hoặc lắng nghe trên một port.
type ProcessInfo struct {
	PID      int    `json:"pid"`
	Name     string `json:"name"`
	Command  string `json:"command"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"` // "tcp" | "udp"
}
