package security

import (
	"fmt"
	"log"

	"github.com/sethvargo/go-password/password"

	"github.com/fatih/color"
	"github.com/kahnwong/swissknife/cmd/utils"
	"github.com/spf13/cobra"
)

func generatePassword() string {
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(32, 10, 0, false, false)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

var generatePasswordCmd = &cobra.Command{
	Use:   "generate-password",
	Short: "Generate password",
	Long:  `Generate password. Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		// main
		color.Green("Security: generate password")

		password := generatePassword()
		utils.CopyToClipboard(password)
		fmt.Printf("%s\n", password)
	},
}

func init() {
	Cmd.AddCommand(generatePasswordCmd)
}
