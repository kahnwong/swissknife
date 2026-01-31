package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getHwInfoCmd = &cobra.Command{
	Use:   "hwinfo",
	Short: "Get hardware info",
	RunE: func(cmd *cobra.Command, args []string) error {
		return get.HwInfo()
	},
}

func init() {
	Cmd.AddCommand(getHwInfoCmd)
}
