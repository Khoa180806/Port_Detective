package process_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

func TestProcessInfo_JSONMarshal(t *testing.T) {
	proc := process.ProcessInfo{
		PID:      12345,
		Name:     "node.exe",
		Command:  "node server.js",
		Port:     8080,
		Protocol: "tcp",
	}

	data, err := json.Marshal(proc)
	if err != nil {
		t.Fatalf("failed to marshal ProcessInfo to JSON: %v", err)
	}

	expectedJSON := `{"pid":12345,"name":"node.exe","command":"node server.js","port":8080,"protocol":"tcp"}`
	if string(data) != expectedJSON {
		t.Errorf("JSON output does not match expected format.\nGot:      %s\nExpected: %s", string(data), expectedJSON)
	}
}

func TestProcessInfo_JSONUnmarshal(t *testing.T) {
	inputJSON := `{"pid":54321,"name":"postgres","command":"postgres -D /data","port":5432,"protocol":"tcp"}`

	var proc process.ProcessInfo
	err := json.Unmarshal([]byte(inputJSON), &proc)
	if err != nil {
		t.Fatalf("failed to unmarshal JSON to ProcessInfo: %v", err)
	}

	expected := process.ProcessInfo{
		PID:      54321,
		Name:     "postgres",
		Command:  "postgres -D /data",
		Port:     5432,
		Protocol: "tcp",
	}

	if !reflect.DeepEqual(proc, expected) {
		t.Errorf("unmarshaled data mismatch.\nGot:      %+v\nExpected: %+v", proc, expected)
	}
}

func TestProcessInfo_ZeroValues(t *testing.T) {
	var proc process.ProcessInfo
	data, err := json.Marshal(proc)
	if err != nil {
		t.Fatalf("failed to marshal empty struct: %v", err)
	}

	expectedJSON := `{"pid":0,"name":"","command":"","port":0,"protocol":""}`
	if string(data) != expectedJSON {
		t.Errorf("JSON output for empty struct mismatch.\nGot:      %s\nExpected: %s", string(data), expectedJSON)
	}
}
