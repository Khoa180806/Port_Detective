//go:build windows
// +build windows

package lookup

import (
	"reflect"
	"testing"
)

func TestParseNetstatForPort(t *testing.T) {
	sampleOutput := `
Active Connections

  Proto  Local Address          Foreign Address        State           PID
  TCP    0.0.0.0:135            0.0.0.0:0              LISTENING       1552
  TCP    0.0.0.0:8080           0.0.0.0:0              LISTENING       1234
  TCP    [::]:8080              [::]:0                 LISTENING       1234
  UDP    0.0.0.0:8080           *:*                                    5678
  TCP    127.0.0.1:5432         0.0.0.0:0              LISTENING       9999
`

	pids, protocols := parseNetstatForPort(sampleOutput, 8080)

	expectedPids := []int{1234, 5678}
	expectedProtocols := []string{"tcp", "udp"}

	if !reflect.DeepEqual(pids, expectedPids) {
		t.Errorf("PIDs mismatch. Got: %v, expected: %v", pids, expectedPids)
	}

	if !reflect.DeepEqual(protocols, expectedProtocols) {
		t.Errorf("Protocols mismatch. Got: %v, expected: %v", protocols, expectedProtocols)
	}
}

func TestParseNetstatForPort_NotFound(t *testing.T) {
	sampleOutput := `
Active Connections

  Proto  Local Address          Foreign Address        State           PID
  TCP    0.0.0.0:135            0.0.0.0:0              LISTENING       1552
`
	pids, protocols := parseNetstatForPort(sampleOutput, 8080)

	if len(pids) != 0 || len(protocols) != 0 {
		t.Errorf("expected empty slices when port is not found, got: %v, %v", pids, protocols)
	}
}
