package cmd

import (
	"os"

	"github.com/kahnwong/swissknife/cmd/networking"
	"github.com/kahnwong/swissknife/cmd/security"
	"github.com/kahnwong/swissknife/cmd/ssh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "swissknife",
	Version: "0.1.0",
	Short:   "Various utils",
	Long:    `Various utils`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(networking.Cmd)
	rootCmd.AddCommand(security.Cmd)
	rootCmd.AddCommand(ssh.Cmd)
}
