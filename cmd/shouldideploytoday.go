package cmd

import (
	"github.com/kahnwong/swissknife/internal/shouldideploytoday"
	"github.com/spf13/cobra"
)

var ShouldIDeployTodayCmd = &cobra.Command{
	Use:   "shouldideploytoday",
	Short: "Should I deploy today?",
	RunE: func(cmd *cobra.Command, args []string) error {
		return shouldideploytoday.ShouldIDeployToday()
	},
}

func init() {
	rootCmd.AddCommand(ShouldIDeployTodayCmd)
}
