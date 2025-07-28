package cmd

import (
	"github.com/kahnwong/swissknife/internal/timer"
	"github.com/spf13/cobra"
)

var TimerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Create a timer",
	Run: func(cmd *cobra.Command, args []string) {
		timer.Timer(args)
	},
}

func init() {
	rootCmd.AddCommand(TimerCmd)
}
