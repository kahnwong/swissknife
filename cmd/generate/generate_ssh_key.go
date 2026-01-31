package generate

import (
	"github.com/kahnwong/swissknife/internal/generate"
	"github.com/spf13/cobra"
)

var generateSSHKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Create SSH key",
	Long:  `Create SSH key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return generate.SSHKey(args)
	},
}

func init() {
	Cmd.AddCommand(generateSSHKeyCmd)
}
