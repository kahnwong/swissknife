package cmd

import (
	"github.com/kahnwong/swissknife/internal/speedtest"
	"github.com/spf13/cobra"
)

var SpeedTestCmd = &cobra.Command{
	Use:   "speedtest",
	Short: "Speedtest",
	RunE: func(cmd *cobra.Command, args []string) error {
		return speedtest.SpeedTest()
	},
}

func init() {
	rootCmd.AddCommand(SpeedTestCmd)
}
