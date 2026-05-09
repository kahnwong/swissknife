package get

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain information",
	Long:  `Obtain information`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				slog.Error("Failed to display help")
				os.Exit(1)
			}
			os.Exit(0)
		}
	},
}
