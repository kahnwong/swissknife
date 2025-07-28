package generate

import (
	"github.com/kahnwong/swissknife/internal/generate"
	"github.com/spf13/cobra"
)

var generatePassphraseCmd = &cobra.Command{
	Use:   "passphrase",
	Short: "Generate passphrase",
	Long:  `Generate passphrase. Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		generate.Passphrase()
	},
}

func init() {
	Cmd.AddCommand(generatePassphraseCmd)
}
