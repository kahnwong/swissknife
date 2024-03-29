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
func generateSSHKeyEDSA(fileName string) error {
	// Generate a new Ed25519 private key
	//// If rand is nil, crypto/rand.Reader will be used
	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}
	p, err := ssh.MarshalPrivateKey(crypto.PrivateKey(private), "")
	if err != nil {
		return err
	}

	// private key
	privateKeyPem := pem.EncodeToMemory(p)
	privateKeyString := string(privateKeyPem)

	writeStringToFile(fmt.Sprintf("%s.pem", fileName), privateKeyString, 0600)

	// public key
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return err
	}
	publicKeyString := "ssh-ed25519" + " " + base64.StdEncoding.EncodeToString(publicKey.Marshal())
	writeStringToFile(fmt.Sprintf("%s.pub", fileName), publicKeyString, 0644)

	return nil
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
		err := generateSSHKeyEDSA(args[0])
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("SSH key created at: %s\n", returnKeyPath(args[0]))
	},
}

func init() {
	Cmd.AddCommand(generateSSHKeyCmd)
}
