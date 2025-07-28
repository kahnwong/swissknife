package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getSystemInfoCmd = &cobra.Command{
	Use:   "system-info",
	Short: "Get system info",
	Run: func(cmd *cobra.Command, args []string) {
		get.SystemInfo()
	},
}

func init() {
	Cmd.AddCommand(getSystemInfoCmd)
}
