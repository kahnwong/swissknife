package generate

import (
	"crypto"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ssh"
)

// helpers
func writeStringToFile(filePath string, data string, permission os.FileMode) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal().Msgf("Failed to create file %s", filePath)
	}

	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal().Msgf("Failed to write to file %s", filePath)
	}

	err = file.Chmod(permission)
	if err != nil {
		log.Fatal().Msgf("Failed to change file permissions")
	}
}

func returnKeyPath(fileName string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal().Msgf("Failed to get current directory")
	}

	keyPath := filepath.Join(currentDir, fileName)
	keyPath = keyPath

	return keyPath
}

// main
func generateSSHKeyEDSA() (string, string) {
	// Generate a new Ed25519 private key
	//// If rand is nil, crypto/rand.Reader will be used
	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal().Msgf("Failed to generate ed25519 key")
	}

	// public key
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		log.Fatal().Msgf("Failed to create public key")
	}
	publicKeyString := fmt.Sprintf("ssh-ed25519 %s", base64.StdEncoding.EncodeToString(publicKey.Marshal()))

	// private key
	//// p stands for pem
	p, err := ssh.MarshalPrivateKey(crypto.PrivateKey(private), "")
	if err != nil {
		log.Fatal().Msgf("Failed to marshal private key")
	}
	privateKeyPem := pem.EncodeToMemory(p)
	privateKeyString := string(privateKeyPem)

	return publicKeyString, privateKeyString
}

func SSHKey(args []string) {
	//init
	if len(args) == 0 {
		fmt.Println("Please specify key name")
		os.Exit(1)
	}

	// main
	publicKeyString, privateKeyString := generateSSHKeyEDSA()

	// write key to file
	publicKeyFilename := fmt.Sprintf("%s.pub", args[0])
	writeStringToFile(publicKeyFilename, publicKeyString, 0644)

	privateKeyFilename := args[0]
	writeStringToFile(privateKeyFilename, privateKeyString, 0600)

	fmt.Printf("SSH key created at: %s\n", returnKeyPath(args[0]))
}
