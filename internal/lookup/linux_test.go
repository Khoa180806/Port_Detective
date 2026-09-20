//go:build linux
// +build linux

package lookup

import (
	"reflect"
	"testing"

	"github.com/Khoa180806/Port_Detective/internal/process"
)

func TestParseLsofOutput(t *testing.T) {
	sampleOutput := `COMMAND     PID   USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
node      12345   root   18u  IPv6 123456      0t0  TCP *:8080 (LISTEN)
postgres  54321  admin   12u  IPv4 987654      0t0  UDP localhost:8080 
`
	procs := parseLsofOutput(sampleOutput, 8080)

	if len(procs) != 2 {
		t.Fatalf("Mong đợi 2 process, nhận được: %d", len(procs))
	}

	expected1 := process.ProcessInfo{
		PID:      12345,
		Name:     "node",
		Command:  "node",
		Port:     8080,
		Protocol: "tcp",
	}
	if !reflect.DeepEqual(procs[0], expected1) {
		t.Errorf("Process 1 không khớp.\nNhận được: %+v\nMong đợi: %+v", procs[0], expected1)
	}

	expected2 := process.ProcessInfo{
		PID:      54321,
		Name:     "postgres",
		Command:  "postgres",
		Port:     8080,
		Protocol: "udp",
	}
	if !reflect.DeepEqual(procs[1], expected2) {
		t.Errorf("Process 2 không khớp.\nNhận được: %+v\nMong đợi: %+v", procs[1], expected2)
	}
}

func TestParseInodesFromProcNet(t *testing.T) {
	// 1F90 hex = 8080 decimal
	sampleProcNet := `  sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode                                                     
   0: 00000000:1F90 00000000:0000 0A 00000000:00000000 00:00000000 00000000  1000        0 99999 1 0000000000000000 100 0 0 10 0
   1: 0100007F:0035 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 88888 1 0000000000000000 100 0 0 10 0
`
	inodes := parseInodesFromProcNet(sampleProcNet, 8080)
	
	if len(inodes) != 1 {
		t.Fatalf("Mong đợi 1 inode, nhận được: %d", len(inodes))
	}
	if inodes[0] != "99999" {
		t.Errorf("Inode không khớp. Nhận được: %s, mong đợi: 99999", inodes[0])
	}
}
