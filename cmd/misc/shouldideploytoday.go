package misc

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/kahnwong/swissknife/color"
	"github.com/spf13/cobra"
)

type ShouldIDeploy struct {
	Timezone      string    `json:"timezone"`
	Date          time.Time `json:"date"`
	ShouldIDeploy bool      `json:"shouldideploy"`
	Message       string    `json:"message"`
}

func ShouldIDeployToday() (ShouldIDeploy, error) {
	url := "https://shouldideploy.today/api?tz=Asia%2FBangkok"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return ShouldIDeploy{}, err
	}

	if err != nil {
		return ShouldIDeploy{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return ShouldIDeploy{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ShouldIDeploy{}, err
	}

	// decode
	var response ShouldIDeploy
	if err := json.Unmarshal(body, &response); err != nil {
		slog.Error("Can not unmarshal JSON")
	}

	return response, nil
}

var ShouldIDeployTodayCmd = &cobra.Command{
	Use:   "shouldideploytoday",
	Short: "Should I deploy today?",
	Long:  `Should I deploy today?`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := ShouldIDeployToday()
		if err != nil {
			fmt.Println(err)
		}

		if response.ShouldIDeploy {
			fmt.Printf("%s\n", color.Green(response.Message))
		} else if !response.ShouldIDeploy {
			fmt.Printf("%s\n", color.Red(response.Message))
		}
	},
}
