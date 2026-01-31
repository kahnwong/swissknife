package get

import (
	"github.com/kahnwong/swissknife/internal/get"
	"github.com/spf13/cobra"
)

var listVolumesCmd = &cobra.Command{
	Use:   "volumes",
	Short: "List volumes",
	RunE: func(cmd *cobra.Command, args []string) error {
		return get.Volumes()
	},
}

func init() {
	Cmd.AddCommand(listVolumesCmd)
}
