package get

import (
	"net"
	"testing"
)

func TestGetInternalIP(t *testing.T) {
	ip, err := getInternalIP()
	// May fail in some test environments, but should not panic
	if err != nil {
		t.Logf("getInternalIP() returned error (acceptable in test environment): %v", err)
		return
	}

	if ip != "" {
		// Verify it's a valid IP
		if net.ParseIP(ip) == nil {
			t.Errorf("getInternalIP() returned invalid IP: %s", ip)
		}
	}
}

func TestGetLocalIP(t *testing.T) {
	ip, err := getLocalIP()
	// May fail in some test environments, but should not panic
	if err != nil {
		t.Logf("getLocalIP() returned error (acceptable in test environment): %v", err)
		return
	}

	if ip != "" {
		// Verify it's a valid IP
		if net.ParseIP(ip) == nil {
			t.Errorf("getLocalIP() returned invalid IP: %s", ip)
		}
	}
}

func TestGetPublicIP(t *testing.T) {
	// This requires network access, so we just test it doesn't panic
	_, err := getPublicIP()
	if err != nil {
		t.Logf("getPublicIP() returned error (expected without network): %v", err)
	}
}

func TestGetIPLocation(t *testing.T) {
	// This requires network access, so we just test it doesn't panic
	_, err := getIPLocation("8.8.8.8")
	if err != nil {
		t.Logf("getIPLocation() returned error (expected without network): %v", err)
	}
}
