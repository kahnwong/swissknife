package cmd

import (
	"github.com/kahnwong/swissknife/internal/speedtest"
	"github.com/spf13/cobra"
)

var SpeedTestCmd = &cobra.Command{
	Use:   "speedtest",
	Short: "Speedtest",
	Run: func(cmd *cobra.Command, args []string) {
		speedtest.SpeedTest()
	},
}

func init() {
	rootCmd.AddCommand(SpeedTestCmd)
}
