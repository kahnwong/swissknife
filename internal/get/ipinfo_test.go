package get

import (
	"testing"
)

func TestGetIPInfo(t *testing.T) {
	// This requires network access, so we just test it doesn't panic
	_, err := getIPInfo("8.8.8.8")
	if err != nil {
		t.Logf("getIPInfo() returned error (expected without network): %v", err)
	}
}
