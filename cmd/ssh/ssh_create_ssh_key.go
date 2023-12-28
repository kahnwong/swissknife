package ssh

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// helpers
func writeStringToFile(filePath, data string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}

func writePrivateKey(privateKey ed25519.PrivateKey) {
	privateKeyStr := fmt.Sprintf("-----BEGIN OPENSSH PRIVATE KEY-----\n%s\n-----END OPENSSH PRIVATE KEY-----\n", base64.StdEncoding.EncodeToString(privateKey))

	writeStringToFile("key.pem", privateKeyStr)
}

func writePublicKey(publicKey ed25519.PublicKey) {
	publicKeyStr := fmt.Sprintf("ssh-ed25519 %s", base64.StdEncoding.EncodeToString(publicKey))

	writeStringToFile("key.pub", publicKeyStr)
}

// main
func createSSHKeyEDSA() string {
	// Generate a new Ed25519 private key
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		os.Exit(1)
	}

	// Write key
	writePrivateKey(privateKey)
	writePublicKey(publicKey)

	return "foo"
}

var createSSHKey = &cobra.Command{
	Use:   "create-ssh-key",
	Short: "Create SSH key",
	Long:  `Create SSH key`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("SSH: create-ssh-key")

		fmt.Printf("\tSSH key created at: %s\n", createSSHKeyEDSA())
	},
}

func init() {
	Cmd.AddCommand(createSSHKey)
}
