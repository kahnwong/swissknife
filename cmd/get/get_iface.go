package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getIfaceCmd = &cobra.Command{
	Use:   "iface",
	Short: "Get iface",
	Long:  `Get iface used for public internet access`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return get.Iface()
	},
}

func init() {
	Cmd.AddCommand(getIfaceCmd)
}
