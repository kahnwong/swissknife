package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var networkingCmd = &cobra.Command{
	Use:   "networking",
	Short: "Networking tools",
	Long:  `Networking tools`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify subcommand")
	},
}

func init() {
	rootCmd.AddCommand(networkingCmd)
}
