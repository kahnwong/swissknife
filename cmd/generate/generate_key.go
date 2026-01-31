package generate

import (
	"github.com/kahnwong/swissknife/internal/generate"
	"github.com/spf13/cobra"
)

var generateKeyCmd = &cobra.Command{
	Use:   "key",
	Short: "Generate key",
	Long:  `Generate key. Re-implementation of "openssl rand -base64 48"". Result is copied to clipboard.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return generate.Key()
	},
}

func init() {
	Cmd.AddCommand(generateKeyCmd)
}
