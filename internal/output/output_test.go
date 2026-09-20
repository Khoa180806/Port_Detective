package output_test

import (
	"strings"
	"testing"

	"github.com/Khoa180806/Port_Detective/internal/output"
	"github.com/Khoa180806/Port_Detective/internal/process"
)

var sampleProcess = process.ProcessInfo{
	PID:      12345,
	Name:     "java",
	Command:  "java -jar app.jar",
	Port:     8080,
	Protocol: "tcp",
}

func TestFormatJSON(t *testing.T) {
	out, err := output.FormatJSON(&sampleProcess)
	if err != nil {
		t.Fatalf("Lỗi FormatJSON: %v", err)
	}

	if !strings.Contains(out, `"pid": 12345`) {
		t.Errorf("JSON output không chứa PID mong đợi. Output: %s", out)
	}
	if !strings.Contains(out, `"name": "java"`) {
		t.Errorf("JSON output không chứa Name mong đợi. Output: %s", out)
	}
}

func TestFormatJSON_Nil(t *testing.T) {
	out, err := output.FormatJSON(nil)
	if err != nil {
		t.Fatalf("Lỗi FormatJSON với nil: %v", err)
	}
	if out != "{}" {
		t.Errorf("JSON output với nil không đúng. Mong đợi '{}', nhận được: %s", out)
	}
}

func TestFormatJSONMultiple(t *testing.T) {
	ps := []process.ProcessInfo{sampleProcess, {PID: 9999, Name: "node", Port: 8080, Protocol: "tcp"}}
	out, err := output.FormatJSONMultiple(ps)
	if err != nil {
		t.Fatalf("Lỗi FormatJSONMultiple: %v", err)
	}

	if !strings.Contains(out, `"pid": 12345`) || !strings.Contains(out, `"pid": 9999`) {
		t.Errorf("JSON output không chứa đủ thông tin các tiến trình. Output: %s", out)
	}
}

func TestFormatJSONMultiple_Nil(t *testing.T) {
	out, err := output.FormatJSONMultiple(nil)
	if err != nil {
		t.Fatalf("Lỗi FormatJSONMultiple với nil: %v", err)
	}
	if out != "[]" {
		t.Errorf("JSON output với nil không đúng. Mong đợi '[]', nhận được: %s", out)
	}
}

func TestFormatText(t *testing.T) {
	out := output.FormatText(&sampleProcess)
	
	expectedLines := []string{
		"Port 8080 is occupied by:",
		"PID:      12345",
		"Process:  java",
		"Command:  java -jar app.jar",
		"Protocol: tcp",
	}

	for _, line := range expectedLines {
		if !strings.Contains(out, line) {
			t.Errorf("Text output thiếu dòng mong đợi: %s\nOutput hiện tại:\n%s", line, out)
		}
	}
}

func TestFormatText_Nil(t *testing.T) {
	out := output.FormatText(nil)
	if !strings.Contains(out, "No process information available") {
		t.Errorf("Text output với nil không đúng. Nhận được: %s", out)
	}
}

func TestFormatTextMultiple(t *testing.T) {
	ps := []process.ProcessInfo{
		sampleProcess,
		{PID: 9999, Name: "node", Command: "node server.js", Port: 8080, Protocol: "tcp"},
	}

	out := output.FormatTextMultiple(ps)
	if !strings.Contains(out, "Found 2 processes") {
		t.Errorf("Text output cho nhiều tiến trình không có tiêu đề đúng. Nhận được:\n%s", out)
	}
	if !strings.Contains(out, "PID:      9999") || !strings.Contains(out, "Process:  node") {
		t.Errorf("Text output thiếu thông tin tiến trình thứ 2. Nhận được:\n%s", out)
	}
}

func TestFormatTextMultiple_Empty(t *testing.T) {
	out := output.FormatTextMultiple([]process.ProcessInfo{})
	if !strings.Contains(out, "No processes found") {
		t.Errorf("Text output với mảng rỗng không đúng. Nhận được: %s", out)
	}
}
