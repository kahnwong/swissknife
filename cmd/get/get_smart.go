package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getSmartCmd = &cobra.Command{
	Use:   "smart",
	Short: "Get disk SMART info",
	Long:  "Get disk SMART info. Equivalent of `sudo smartctl -a /dev/nvme0n1p`",
	Run: func(cmd *cobra.Command, args []string) {
		get.Smart()
	},
}

func init() {
	Cmd.AddCommand(getSmartCmd)
}
