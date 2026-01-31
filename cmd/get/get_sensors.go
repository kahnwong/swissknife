package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var SensorsCmd = &cobra.Command{
	Use:   "sensors",
	Short: "Get sensors info",
	RunE: func(cmd *cobra.Command, args []string) error {
		return get.Sensors()
	},
}

func init() {
	Cmd.AddCommand(SensorsCmd)
}
