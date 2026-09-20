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
		t.Fatalf("Không thể marshal ProcessInfo sang JSON: %v", err)
	}

	expectedJSON := `{"pid":12345,"name":"node.exe","command":"node server.js","port":8080,"protocol":"tcp"}`
	if string(data) != expectedJSON {
		t.Errorf("Kết quả JSON không đúng định dạng mong đợi.\nNhận được: %s\nMong đợi:  %s", string(data), expectedJSON)
	}
}

func TestProcessInfo_JSONUnmarshal(t *testing.T) {
	inputJSON := `{"pid":54321,"name":"postgres","command":"postgres -D /data","port":5432,"protocol":"tcp"}`

	var proc process.ProcessInfo
	err := json.Unmarshal([]byte(inputJSON), &proc)
	if err != nil {
		t.Fatalf("Không thể unmarshal JSON thành ProcessInfo: %v", err)
	}

	expected := process.ProcessInfo{
		PID:      54321,
		Name:     "postgres",
		Command:  "postgres -D /data",
		Port:     5432,
		Protocol: "tcp",
	}

	if !reflect.DeepEqual(proc, expected) {
		t.Errorf("Dữ liệu sau khi unmarshal không khớp.\nNhận được: %+v\nMong đợi:  %+v", proc, expected)
	}
}

func TestProcessInfo_ZeroValues(t *testing.T) {
	var proc process.ProcessInfo
	data, err := json.Marshal(proc)
	if err != nil {
		t.Fatalf("Lỗi khi marshal struct rỗng: %v", err)
	}

	expectedJSON := `{"pid":0,"name":"","command":"","port":0,"protocol":""}`
	if string(data) != expectedJSON {
		t.Errorf("Kết quả JSON cho struct rỗng không khớp.\nNhận được: %s\nMong đợi:  %s", string(data), expectedJSON)
	}
}
