package generate

import (
	"fmt"

	"github.com/kahnwong/swissknife/utils"
	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

func generatePassword() string {
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(32, 10, 0, false, false)
	if err != nil {
		log.Fatal().Msg("Failed to generate password")
	}

	return res
}

var generatePasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate password",
	Long:  `Generate password. Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		password := generatePassword()
		utils.WriteToClipboard(password)
		fmt.Printf("%s\n", password)
	},
}

func init() {
	Cmd.AddCommand(generatePasswordCmd)
}
