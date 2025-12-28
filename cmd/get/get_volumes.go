package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var listVolumesCmd = &cobra.Command{
	Use:   "volumes",
	Short: "List volumes",
	Run: func(cmd *cobra.Command, args []string) {
		get.Volumes()
	},
}

func init() {
	GetCmd.AddCommand(listVolumesCmd)
}
