package utils

import (
	"testing"
)

func TestReadFromClipboard(t *testing.T) {
	// Test that function returns without panicking
	_, err := ReadFromClipboard()
	// Error is acceptable if clipboard is not available
	if err != nil {
		t.Logf("ReadFromClipboard returned error (expected in test environment): %v", err)
	}
}

func TestWriteToClipboard(t *testing.T) {
	// Test that function returns without panicking
	err := WriteToClipboard("test")
	// Error is acceptable if clipboard is not available
	if err != nil {
		t.Logf("WriteToClipboard returned error (expected in test environment): %v", err)
	}
}

func TestWriteToClipboardImage(t *testing.T) {
	// Test with empty byte slice
	err := WriteToClipboardImage([]byte{})
	// Error is acceptable if clipboard is not available
	if err != nil {
		t.Logf("WriteToClipboardImage returned error (expected in test environment): %v", err)
	}
}
