package generate

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate stuff",
	Long:  `Generate stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to display help")
			}
			os.Exit(0)
		}
	},
}
