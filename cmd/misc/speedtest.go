package misc

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/showwin/speedtest-go/speedtest"
	"github.com/spf13/cobra"
)

var SpeedTestCmd = &cobra.Command{
	Use:   "speedtest",
	Short: "Speedtest",
	Run: func(cmd *cobra.Command, args []string) {
		// https://github.com/showwin/speedtest-go#api-usage

		green := color.New(color.FgHiGreen).SprintFunc()

		var speedtestClient = speedtest.New()

		serverList, _ := speedtestClient.FetchServers()
		targets, _ := serverList.FindServer([]int{})

		for _, s := range targets {
			fmt.Printf("%s:   %s\n", green("Server"), s.Name)

			err := s.PingTest(nil)
			if err != nil {
				return
			}

			err = s.DownloadTest()
			if err != nil {
				return
			}

			err = s.UploadTest()
			if err != nil {
				return
			}

			fmt.Printf("" +
				fmt.Sprintf("%s:  %s\n", green("Latency"), s.Latency) +
				fmt.Sprintf("%s: %s\n", green("Download"), s.DLSpeed) +
				fmt.Sprintf("%s:   %s\n", green("Upload"), s.ULSpeed),
			)

			s.Context.Reset()
		}
	},
}
