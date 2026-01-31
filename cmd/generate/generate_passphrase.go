package generate

import (
	"github.com/kahnwong/swissknife/internal/generate"
	"github.com/spf13/cobra"
)

var generatePassphraseCmd = &cobra.Command{
	Use:   "passphrase",
	Short: "Generate passphrase",
	Long:  `Generate passphrase. Result is copied to clipboard.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return generate.Passphrase()
	},
}

func init() {
	Cmd.AddCommand(generatePassphraseCmd)
}
