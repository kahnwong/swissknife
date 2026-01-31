package cmd

import (
	"github.com/kahnwong/swissknife/internal/stopwatch"
	"github.com/spf13/cobra"
)

var StopwatchCmd = &cobra.Command{
	Use:   "stopwatch",
	Short: "Create a stopwatch",
	RunE: func(cmd *cobra.Command, args []string) error {
		return stopwatch.Stopwatch()
	},
}

func init() {
	rootCmd.AddCommand(StopwatchCmd)
}
