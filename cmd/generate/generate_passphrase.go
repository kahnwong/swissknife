package generate

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/kahnwong/swissknife/cmd/utils"

	"github.com/sethvargo/go-diceware/diceware"
	"github.com/spf13/cobra"
)

func generatePassphrase() (string, error) {
	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(6)
	if err != nil {
		return "", err
	}

	res := strings.Join(list, "-")

	return res, nil
}

var generatePassphraseCmd = &cobra.Command{
	Use:   "passphrase",
	Short: "Generate passphrase",
	Long:  `Generate passphrase. Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		// main
		passphrase, err := generatePassphrase()
		if err != nil {
			slog.Error("Error generating passphrase")
		}
		utils.WriteToClipboard(passphrase)
		fmt.Printf("%s\n", passphrase)
	},
}

func init() {
	Cmd.AddCommand(generatePassphraseCmd)
}
