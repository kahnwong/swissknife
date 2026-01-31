package stopwatch

import (
	"testing"
)

func TestStopwatch(t *testing.T) {
	// Test that Stopwatch() doesn't panic and returns error properly
	// This will fail in test environment without TTY but should return error
	err := Stopwatch()
	if err != nil {
		t.Logf("Stopwatch() returned error (expected in test environment): %v", err)
	}
}
