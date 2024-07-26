package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain information",
	Long:  `Obtain information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify subcommand")
	},
}

func init() {
	Cmd.AddCommand(getIPCmd)
	Cmd.AddCommand(getSystemInfoCmd)
}
