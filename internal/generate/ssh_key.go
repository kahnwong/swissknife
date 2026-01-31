package generate

import (
	"crypto"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

// helpers
func writeStringToFile(filePath string, data string, permission os.FileMode) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}

	if _, err = file.WriteString(data); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filePath, err)
	}

	if err = file.Chmod(permission); err != nil {
		return fmt.Errorf("failed to change file permissions: %w", err)
	}

	return nil
}

func returnKeyPath(fileName string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	keyPath := filepath.Join(currentDir, fileName)

	return keyPath, nil
}

// main
func generateSSHKeyEDSA() (string, string, error) {
	// Generate a new Ed25519 private key
	//// If rand is nil, crypto/rand.Reader will be used
	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate ed25519 key: %w", err)
	}

	// public key
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return "", "", fmt.Errorf("failed to create public key: %w", err)
	}
	publicKeyString := fmt.Sprintf("ssh-ed25519 %s", base64.StdEncoding.EncodeToString(publicKey.Marshal()))

	// private key
	//// p stands for pem
	p, err := ssh.MarshalPrivateKey(crypto.PrivateKey(private), "")
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal private key: %w", err)
	}
	privateKeyPem := pem.EncodeToMemory(p)
	privateKeyString := string(privateKeyPem)

	return publicKeyString, privateKeyString, nil
}

func SSHKey(args []string) error {
	//init
	if len(args) == 0 {
		return fmt.Errorf("please specify key name")
	}

	// main
	publicKeyString, privateKeyString, err := generateSSHKeyEDSA()
	if err != nil {
		return err
	}

	// write key to file
	publicKeyFilename := fmt.Sprintf("%s.pub", args[0])
	if err = writeStringToFile(publicKeyFilename, publicKeyString, 0644); err != nil {
		return err
	}

	privateKeyFilename := args[0]
	if err = writeStringToFile(privateKeyFilename, privateKeyString, 0600); err != nil {
		return err
	}

	keyPath, err := returnKeyPath(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("SSH key created at: %s\n", keyPath)
	return nil
}
