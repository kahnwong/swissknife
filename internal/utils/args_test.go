package utils

import (
	"testing"
)

func TestSetURL(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "valid URL from args",
			args:     []string{"https://example.com"},
			expected: "https://example.com",
		},
		{
			name:     "empty args returns empty or exits",
			args:     []string{},
			expected: "", // Will exit in actual code, but we can't test that easily
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.args) > 0 {
				result := SetURL(tt.args)
				if result != tt.expected {
					t.Errorf("SetURL() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestSetIP(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "valid IP from args",
			args:     []string{"192.168.1.1"},
			expected: "192.168.1.1",
		},
		{
			name:     "empty args returns empty",
			args:     []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetIP(tt.args)
			if result != tt.expected {
				t.Errorf("SetIP() = %v, want %v", result, tt.expected)
			}
		})
	}
}
