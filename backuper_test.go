package main

import (
	"testing"
)

func TestPing(t *testing.T) {
	tests := []struct {
		domain   string
		port     string
		expected bool
	}{
		{"localhost", "631", true},
		{"google.de", "80", true},
		{"bla", "7", false},
	}

	for _, test := range tests {
		result := ping(test.domain, test.port)
		if result != test.expected {
			t.Errorf("ping test failed. input: %s, expected: %t, result: %t", test.domain+":"+test.port, test.expected, result)
		}
	}
}

func TestExecuteCommand(t *testing.T) {
	output := executeCommand("ls")
	if output == "" {
		t.Errorf("executeCommand failed.")
	}
}
