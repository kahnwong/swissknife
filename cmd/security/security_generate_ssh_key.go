package security

import (
	"crypto"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// helpers
func writeStringToFile(filePath, data string, permission os.FileMode) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Chmod(permission)
	if err != nil {
		fmt.Println("Error setting file permissions:", err)
		return
	}
}

func returnKeyPath(fileName string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}

	keyPath := filepath.Join(currentDir, fileName)
	keyPath = keyPath + ".pem"

	return keyPath
}

// main
func generateSSHKeyEDSA(fileName string) {
	// Generate a new Ed25519 private key
	//// If rand is nil, crypto/rand.Reader will be used
	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	p, err := ssh.MarshalPrivateKey(crypto.PrivateKey(private), "")
	if err != nil {
		panic(err)
	}

	// private key
	privateKeyPem := pem.EncodeToMemory(p)
	privateKeyString := string(privateKeyPem)

	writeStringToFile(fmt.Sprintf("%s.pem", fileName), privateKeyString, 0600)

	// public key
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		panic(err)
	}
	publicKeyString := "ssh-ed25519" + " " + base64.StdEncoding.EncodeToString(publicKey.Marshal())
	writeStringToFile(fmt.Sprintf("%s.public", fileName), publicKeyString, 0644)
}

var generateSSHKeyCmd = &cobra.Command{
	Use:   "generate-ssh-key",
	Short: "Create SSH key",
	Long:  `Create SSH key`,
	Run: func(cmd *cobra.Command, args []string) {
		//init
		if len(args) == 0 {
			fmt.Println("Please specify key name")
			os.Exit(1)
		}

		// main
		color.Green("SSH: generate-ssh-key")

		generateSSHKeyEDSA(args[0])
		fmt.Printf("\tSSH key created at: %s\n", returnKeyPath(args[0]))
	},
}

func init() {
	Cmd.AddCommand(generateSSHKeyCmd)
}
