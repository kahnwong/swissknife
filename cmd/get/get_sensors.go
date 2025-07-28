package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var SensorsCmd = &cobra.Command{
	Use:   "sensors",
	Short: "Get sensors info",
	Run: func(cmd *cobra.Command, args []string) {
		get.Sensors()
	},
}

func init() {
	Cmd.AddCommand(SensorsCmd)
}
