package generate

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

	"github.com/spf13/cobra"
)

// helpers
func writeStringToFile(filePath string, data string, permission os.FileMode) {
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
func generateSSHKeyEDSA() (string, string, error) {
	// Generate a new Ed25519 private key
	//// If rand is nil, crypto/rand.Reader will be used
	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}
	p, err := ssh.MarshalPrivateKey(crypto.PrivateKey(private), "")
	if err != nil {
		return "", "", err
	}

	// public key
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return "", "", err
	}
	publicKeyString := "ssh-ed25519" + " " + base64.StdEncoding.EncodeToString(publicKey.Marshal())

	// private key
	privateKeyPem := pem.EncodeToMemory(p)
	privateKeyString := string(privateKeyPem)

	return publicKeyString, privateKeyString, nil
}

var generateSSHKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Create SSH key",
	Long:  `Create SSH key`,
	Run: func(cmd *cobra.Command, args []string) {
		//init
		if len(args) == 0 {
			fmt.Println("Please specify key name")
			os.Exit(1)
		}

		// main
		publicKeyString, privateKeyString, err := generateSSHKeyEDSA()
		if err != nil {
			fmt.Print(err)
		}

		// write key to file
		writeStringToFile(fmt.Sprintf("%s.pub", args[0]), publicKeyString, 0644)
		writeStringToFile(fmt.Sprintf("%s.pem", args[0]), privateKeyString, 0600)

		fmt.Printf("SSH key created at: %s\n", returnKeyPath(args[0]))
	},
}

func init() {
	Cmd.AddCommand(generateSSHKeyCmd)
}
