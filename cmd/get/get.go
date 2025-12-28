package get

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain information",
	Long:  `Obtain information`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				log.Fatal().Msg("Failed to display help")
			}
			os.Exit(0)
		}
	},
}
