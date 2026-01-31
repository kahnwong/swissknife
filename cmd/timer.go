package cmd

import (
	"github.com/kahnwong/swissknife/internal/timer"
	"github.com/spf13/cobra"
)

var TimerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Create a timer",
	RunE: func(cmd *cobra.Command, args []string) error {
		return timer.Timer(args)
	},
}

func init() {
	rootCmd.AddCommand(TimerCmd)
}
