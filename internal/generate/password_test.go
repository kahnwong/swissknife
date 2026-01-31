package generate

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	password, err := generatePassword()
	if err != nil {
		t.Fatalf("generatePassword() returned error: %v", err)
	}

	if len(password) == 0 {
		t.Error("generatePassword() returned empty password")
	}

	if len(password) != 32 {
		t.Errorf("generatePassword() returned password of length %d, expected 32", len(password))
	}
}

func TestPassword(t *testing.T) {
	// Test that Password() doesn't panic and returns error properly
	err := Password()
	// May fail due to clipboard, but should return error not panic
	if err != nil {
		t.Logf("Password() returned error (expected in test environment): %v", err)
	}
}
