package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getHwInfoCmd = &cobra.Command{
	Use:   "hwinfo",
	Short: "Get hardware info",
	Run: func(cmd *cobra.Command, args []string) {
		get.HwInfo()
	},
}

func init() {
	GetCmd.AddCommand(getHwInfoCmd)
}
