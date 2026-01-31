package shouldideploytoday

import (
	"context"
	"fmt"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/kahnwong/swissknife/configs/color"
)

type ShouldIDeploy struct {
	Timezone      string    `json:"timezone"`
	Date          time.Time `json:"date"`
	ShouldIDeploy bool      `json:"shouldideploy"`
	Message       string    `json:"message"`
}

func getResponse() (ShouldIDeploy, error) {
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
		return ShouldIDeploy{}, fmt.Errorf("error calling ShouldIDeploy API: %w", err)
	}

	return response, nil
}

func ShouldIDeployToday() error {
	response, err := getResponse()
	if err != nil {
		return err
	}

	if response.ShouldIDeploy {
		fmt.Printf("%s\n", color.Green(response.Message))
	} else if !response.ShouldIDeploy {
		fmt.Printf("%s\n", color.Red(response.Message))
	}
	return nil
}
