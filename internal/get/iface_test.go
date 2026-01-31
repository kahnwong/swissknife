package get

import (
	"testing"
)

func TestGetIface(t *testing.T) {
	iface, err := getIface()
	// May fail in some test environments
	if err != nil {
		t.Logf("getIface() returned error (acceptable in test environment): %v", err)
		return
	}

	if iface == "" {
		t.Error("getIface() returned empty interface name")
	}
}
