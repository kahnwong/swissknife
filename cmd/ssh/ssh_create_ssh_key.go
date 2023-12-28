package ssh

import (
	"crypto"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"

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

// main
func createSSHKeyEDSA(fileName string) {
	// Generate a new Ed25519 private key
	//// If rand is nil, crypto/rand.Reader will be used
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	p, err := ssh.MarshalPrivateKey(crypto.PrivateKey(priv), "")
	if err != nil {
		panic(err)
	}

	// private key
	privateKeyPem := pem.EncodeToMemory(p)
	privateKeyString := string(privateKeyPem)

	writeStringToFile(fmt.Sprintf("%s.pem", fileName), privateKeyString, 0600)

	// public key
	publicKey, err := ssh.NewPublicKey(pub)
	if err != nil {
		panic(err)
	}
	publicKeyString := "ssh-ed25519" + " " + base64.StdEncoding.EncodeToString(publicKey.Marshal())
	writeStringToFile(fmt.Sprintf("%s.pub", fileName), publicKeyString, 0644)
}

var createSSHKey = &cobra.Command{
	Use:   "create-ssh-key",
	Short: "Create SSH key",
	Long:  `Create SSH key`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("SSH: create-ssh-key")

		fileName := "foo"
		createSSHKeyEDSA(fileName)
		fmt.Printf("\tSSH key created at: %s\n", fileName)
	},
}

func init() {
	Cmd.AddCommand(createSSHKey)
}
