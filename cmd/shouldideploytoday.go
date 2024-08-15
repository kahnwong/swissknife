package cmd

import (
	"fmt"
	"time"

	"context"

	"github.com/carlmjohnson/requests"
	"github.com/kahnwong/swissknife/color"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type ShouldIDeploy struct {
	Timezone      string    `json:"timezone"`
	Date          time.Time `json:"date"`
	ShouldIDeploy bool      `json:"shouldideploy"`
	Message       string    `json:"message"`
}

func ShouldIDeployToday() ShouldIDeploy {
	url := "https://shouldideploy.today"

	var response ShouldIDeploy
	err := requests.
		URL(url).
		Path("api").
		Param("tz", "UTC").
		Param("date", time.Now().UTC().Format("2006-01-02T15:04:05.000Z")).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		log.Fatal().Err(err).Msg("Error calling ShouldIDeploy API")
	}

	return response
}

var ShouldIDeployTodayCmd = &cobra.Command{
	Use:   "shouldideploytoday",
	Short: "Should I deploy today?",
	Run: func(cmd *cobra.Command, args []string) {
		response := ShouldIDeployToday()

		if response.ShouldIDeploy {
			fmt.Printf("%s\n", color.Green(response.Message))
		} else if !response.ShouldIDeploy {
			fmt.Printf("%s\n", color.Red(response.Message))
		}
	},
}

func init() {
	rootCmd.AddCommand(ShouldIDeployTodayCmd)
}
