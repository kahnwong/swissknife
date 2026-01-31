package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getSysInfoCmd = &cobra.Command{
	Use:   "sysinfo",
	Short: "Get system info",
	RunE: func(cmd *cobra.Command, args []string) error {
		return get.SysInfo()
	},
}

func init() {
	Cmd.AddCommand(getSysInfoCmd)
}
