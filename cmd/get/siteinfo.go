package get

import (
	"fmt"

	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

// siteinfoCmd represents the siteinfo command
var siteinfoCmd = &cobra.Command{
	Use:   "siteinfo [url]",
	Short: "display siteinfo",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		result, err := get.GetSiteInfo(url)
		if err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	},
}

func init() {
	GetCmd.AddCommand(siteinfoCmd)
}
