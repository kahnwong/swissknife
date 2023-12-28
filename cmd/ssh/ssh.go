package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH tools",
	Long:  `SSH tools`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify subcommand")
	},
}
