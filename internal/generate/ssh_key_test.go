package generate

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateSSHKeyEDSA(t *testing.T) {
	publicKey, privateKey, err := generateSSHKeyEDSA()
	if err != nil {
		t.Fatalf("generateSSHKeyEDSA() returned error: %v", err)
	}

	if len(publicKey) == 0 {
		t.Error("generateSSHKeyEDSA() returned empty public key")
	}

	if len(privateKey) == 0 {
		t.Error("generateSSHKeyEDSA() returned empty private key")
	}

	// Check public key format
	if !strings.HasPrefix(publicKey, "ssh-ed25519 ") {
		t.Error("generateSSHKeyEDSA() public key should start with 'ssh-ed25519 '")
	}

	// Check private key format
	if !strings.Contains(privateKey, "BEGIN OPENSSH PRIVATE KEY") {
		t.Error("generateSSHKeyEDSA() private key should contain PEM header")
	}
}

func TestReturnKeyPath(t *testing.T) {
	path, err := returnKeyPath("test_key")
	if err != nil {
		t.Fatalf("returnKeyPath() returned error: %v", err)
	}

	if len(path) == 0 {
		t.Error("returnKeyPath() returned empty path")
	}

	if !strings.Contains(path, "test_key") {
		t.Error("returnKeyPath() should contain the filename")
	}
}

func TestWriteStringToFile(t *testing.T) {
	tmpFile := "/tmp/test_ssh_key_write"
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Error("error removing tmp file")
		}
	}(tmpFile)

	err := writeStringToFile(tmpFile, "test content", 0600)
	if err != nil {
		t.Fatalf("writeStringToFile() returned error: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		t.Error("writeStringToFile() did not create file")
	}

	// Verify content
	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	if string(content) != "test content" {
		t.Errorf("writeStringToFile() wrote %q, want %q", string(content), "test content")
	}
}

func TestSSHKey(t *testing.T) {
	// Test with no args (should return error)
	err := SSHKey([]string{})
	if err == nil {
		t.Error("SSHKey() with no args should return error")
	}
}
