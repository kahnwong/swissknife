package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var siteinfoCmd = &cobra.Command{
	Use:   "siteinfo [url]",
	Short: "Get website technology information",
	Long:  `Analyze a website and display the technologies it uses, categorized by type`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return get.GetSiteInfo(args)
	},
}

func init() {
	Cmd.AddCommand(siteinfoCmd)
}
