package generate

import (
	"github.com/kahnwong/swissknife/internal/generate"
	"github.com/spf13/cobra"
)

var generatePasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate password",
	Long:  `Generate password. Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		generate.Password()
	},
}

func init() {
	Cmd.AddCommand(generatePasswordCmd)
}
