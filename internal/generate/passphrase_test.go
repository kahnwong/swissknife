package generate

import (
	"strings"
	"testing"
)

func TestGeneratePassphrase(t *testing.T) {
	passphrase, err := generatePassphrase()
	if err != nil {
		t.Fatalf("generatePassphrase() returned error: %v", err)
	}

	if len(passphrase) == 0 {
		t.Error("generatePassphrase() returned empty passphrase")
	}

	// Should contain dashes separating words
	if !strings.Contains(passphrase, "-") {
		t.Error("generatePassphrase() should contain dashes between words")
	}

	// Should have 6 words (5 dashes)
	words := strings.Split(passphrase, "-")
	if len(words) != 6 {
		t.Errorf("generatePassphrase() returned %d words, expected 6", len(words))
	}
}

func TestPassphrase(t *testing.T) {
	// Test that Passphrase() doesn't panic and returns error properly
	err := Passphrase()
	// May fail due to clipboard, but should return error not panic
	if err != nil {
		t.Logf("Passphrase() returned error (expected in test environment): %v", err)
	}
}
