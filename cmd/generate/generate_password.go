package generate

import (
	"fmt"
	"log/slog"

	"github.com/sethvargo/go-password/password"

	"github.com/kahnwong/swissknife/cmd/utils"
	"github.com/spf13/cobra"
)

func generatePassword() (string, error) {
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(32, 10, 0, false, false)
	if err != nil {
		return "", err
	}

	return res, nil
}

var generatePasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate password",
	Long:  `Generate password. Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		// main
		password, err := generatePassword()
		if err != nil {
			slog.Error("Error generating password")
		}
		utils.WriteToClipboard(password)
		fmt.Printf("%s\n", password)
	},
}

func init() {
	Cmd.AddCommand(generatePasswordCmd)
}
