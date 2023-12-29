package security

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "security",
	Short: "Security tools",
	Long:  `Security tools`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify subcommand")
	},
}
