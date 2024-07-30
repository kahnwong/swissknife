package cmd

import (
	"fmt"
	"time"

	"context"

	"github.com/carlmjohnson/requests"
	"github.com/kahnwong/swissknife/color"
	"github.com/spf13/cobra"
)

type ShouldIDeploy struct {
	Timezone      string    `json:"timezone"`
	Date          time.Time `json:"date"`
	ShouldIDeploy bool      `json:"shouldideploy"`
	Message       string    `json:"message"`
}

func ShouldIDeployToday() ShouldIDeploy {
	url := "https://shouldideploy.today/api?tz=Asia%2FBangkok"

	var response ShouldIDeploy
	err := requests.
		URL(url).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		fmt.Println("Error calling ShouldIDeploy API:", err)
	}

	return response
}

var ShouldIDeployTodayCmd = &cobra.Command{
	Use:   "shouldideploytoday",
	Short: "Should I deploy today?",
	Long:  `Should I deploy today?`,
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
