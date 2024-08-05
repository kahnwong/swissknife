package generate

import (
	"fmt"
	"strings"

	"github.com/kahnwong/swissknife/utils"
	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-diceware/diceware"
	"github.com/spf13/cobra"
)

func generatePassphrase() string {
	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(6)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate passphrases")
	}

	res := strings.Join(list, "-")

	return res
}

var generatePassphraseCmd = &cobra.Command{
	Use:   "passphrase",
	Short: "Generate passphrase",
	Long:  `Generate passphrase. Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		passphrase := generatePassphrase()
		utils.WriteToClipboard(passphrase)
		fmt.Printf("%s\n", passphrase)
	},
}

func init() {
	Cmd.AddCommand(generatePassphraseCmd)
}
