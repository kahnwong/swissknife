package generate

import (
	"encoding/base64"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key, err := generateKey(48)
	if err != nil {
		t.Fatalf("generateKey() returned error: %v", err)
	}

	if len(key) == 0 {
		t.Error("generateKey() returned empty key")
	}

	// Verify it's valid base64
	_, err = base64.URLEncoding.DecodeString(key)
	if err != nil {
		t.Errorf("generateKey() returned invalid base64: %v", err)
	}
}

func TestGenerateKeyDifferentSizes(t *testing.T) {
	sizes := []int{16, 32, 48, 64}

	for _, size := range sizes {
		t.Run(string(rune(size)), func(t *testing.T) {
			key, err := generateKey(size)
			if err != nil {
				t.Fatalf("generateKey(%d) returned error: %v", size, err)
			}

			if len(key) == 0 {
				t.Errorf("generateKey(%d) returned empty key", size)
			}
		})
	}
}

func TestKey(t *testing.T) {
	// Test that Key() doesn't panic and returns error properly
	err := Key()
	// May fail due to clipboard, but should return error not panic
	if err != nil {
		t.Logf("Key() returned error (expected in test environment): %v", err)
	}
}
