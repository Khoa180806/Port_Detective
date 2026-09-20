package output

import (
	"encoding/json"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

// FormatJSON trả về chuỗi định dạng JSON cho một tiến trình.
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

// FormatJSONMultiple trả về chuỗi định dạng JSON cho danh sách các tiến trình.
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
