//go:build darwin
// +build darwin

package lookup

import (
	"reflect"
	"testing"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

func TestParseDarwinLsofOutput(t *testing.T) {
	sampleOutput := `COMMAND   PID  USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
java    12345 admin  113u  IPv6 0x1234567890abcdef      0t0  TCP *:8080 (LISTEN)
node    54321 admin   12u  IPv4 0xabcdef1234567890      0t0  UDP localhost:8080 
`
	procs := parseDarwinLsofOutput(sampleOutput, 8080)

	if len(procs) != 2 {
		t.Fatalf("Mong đợi 2 process, nhận được: %d", len(procs))
	}

	expected1 := process.ProcessInfo{
		PID:      12345,
		Name:     "java",
		Command:  "java",
		Port:     8080,
		Protocol: "tcp",
	}
	if !reflect.DeepEqual(procs[0], expected1) {
		t.Errorf("Process 1 không khớp.\nNhận được: %+v\nMong đợi: %+v", procs[0], expected1)
	}

	expected2 := process.ProcessInfo{
		PID:      54321,
		Name:     "node",
		Command:  "node",
		Port:     8080,
		Protocol: "udp",
	}
	if !reflect.DeepEqual(procs[1], expected2) {
		t.Errorf("Process 2 không khớp.\nNhận được: %+v\nMong đợi: %+v", procs[1], expected2)
	}
}

func TestParseDarwinLsofOutput_Empty(t *testing.T) {
	// Header only
	sampleOutput := `COMMAND   PID  USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME`
	procs := parseDarwinLsofOutput(sampleOutput, 8080)

	if len(procs) != 0 {
		t.Fatalf("Mong đợi 0 process, nhận được: %d", len(procs))
	}
}
