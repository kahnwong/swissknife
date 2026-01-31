package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getIPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get IP information",
	Long:  `Get IP information`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return get.IP()
	},
}

func init() {
	Cmd.AddCommand(getIPCmd)
}
