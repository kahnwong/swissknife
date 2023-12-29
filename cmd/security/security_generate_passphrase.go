package security

import (
	"fmt"
	"log"
	"strings"

	"github.com/kahnwong/swissknife/cmd/utils"

	"github.com/sethvargo/go-diceware/diceware"
	"github.com/spf13/cobra"
)

func generatePassphrase() string {
	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(6)
	if err != nil {
		log.Fatal(err)
	}

	res := strings.Join(list, "-")

	return res
}

var generatePassphraseCmd = &cobra.Command{
	Use:   "generate-passphrase",
	Short: "Generate passphrase",
	Long:  `Generate passphrase. Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		// main
		passphrase := generatePassphrase()
		utils.CopyToClipboard(passphrase)
		fmt.Printf("%s\n", passphrase)
	},
}

func init() {
	Cmd.AddCommand(generatePassphraseCmd)
}
