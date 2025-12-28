package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var getIPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get IP information",
	Long:  `Get IP information`,
	Run: func(cmd *cobra.Command, args []string) {
		get.IP()
	},
}

func init() {
	GetCmd.AddCommand(getIPCmd)
}
