package cmd

import (
	"github.com/kahnwong/swissknife/internal/stopwatch"
	"github.com/spf13/cobra"
)

var StopwatchCmd = &cobra.Command{
	Use:   "stopwatch",
	Short: "Create a stopwatch",
	Run: func(cmd *cobra.Command, args []string) {
		stopwatch.Stopwatch()
	},
}

func init() {
	rootCmd.AddCommand(StopwatchCmd)
}
