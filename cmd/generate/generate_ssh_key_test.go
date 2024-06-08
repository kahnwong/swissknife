package generate

import (
	"strings"
	"testing"
)

func TestGenerateSSHKey(t *testing.T) {
	publicKey, privateKey, err := generateSSHKeyEDSA()
	if err != nil {
		t.Errorf("generateSSHKeyEDSA() error = %v", err)
	} else {
		// assert public key
		if !strings.HasPrefix(publicKey, "ssh-ed25519") {
			t.Errorf("generateSSHKeyEDSA() publicKey is not a public key")
		}

		// assert private key
		if !strings.HasPrefix(privateKey, "-----BEGIN OPENSSH PRIVATE KEY-----") {
			t.Errorf("generateSSHKeyEDSA() privateKey is not a private key")
		}
	}
}
