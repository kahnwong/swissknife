package generate

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate stuff",
	Long:  `Generate stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify subcommand")
	},
}
