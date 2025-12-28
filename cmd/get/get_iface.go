package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getIfaceCmd = &cobra.Command{
	Use:   "iface",
	Short: "Get iface",
	Long:  `Get iface used for public internet access`,
	Run: func(cmd *cobra.Command, args []string) {
		get.Iface()
	},
}

func init() {
	GetCmd.AddCommand(getIfaceCmd)
}
