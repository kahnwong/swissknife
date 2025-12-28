package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getSysInfoCmd = &cobra.Command{
	Use:   "sysinfo",
	Short: "Get system info",
	Run: func(cmd *cobra.Command, args []string) {
		get.SysInfo()
	},
}

func init() {
	GetCmd.AddCommand(getSysInfoCmd)
}
