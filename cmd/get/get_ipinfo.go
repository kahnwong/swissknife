package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getIPInfoCmd = &cobra.Command{
	Use:   "ipinfo [ip]",
	Short: "Get detailed IP information (location, ISP, etc.)",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		get.IPInfo(args)
	},
}

func init() {
	Cmd.AddCommand(getIPInfoCmd)
}
