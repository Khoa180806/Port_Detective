package output

import (
	"encoding/json"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

// FormatJSON returns a formatted JSON string for a single process.
func FormatJSON(p *process.ProcessInfo) (string, error) {
	if p == nil {
		return "{}", nil
	}
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FormatJSONMultiple returns a formatted JSON string for a list of processes.
func FormatJSONMultiple(ps []process.ProcessInfo) (string, error) {
	if ps == nil {
		return "[]", nil
	}
	data, err := json.MarshalIndent(ps, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
