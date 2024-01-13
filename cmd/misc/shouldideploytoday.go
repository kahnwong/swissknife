package misc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
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
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// decode
	var response ShouldIDeploy
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return response
}

var Cmd = &cobra.Command{
	Use:   "shouldideploytoday",
	Short: "Should I deploy today?",
	Long:  `Should I deploy today?`,
	Run: func(cmd *cobra.Command, args []string) {
		response := ShouldIDeployToday()

		// init output colors
		green := color.New(color.FgGreen).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()

		// set output colors
		if response.ShouldIDeploy {
			fmt.Printf("%s\n", green(response.Message))
		} else if !response.ShouldIDeploy {
			fmt.Printf("%s\n", red(response.Message))
		}
	},
}
