package misc

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
	"github.com/spf13/cobra"
)

var SpeedTestCmd = &cobra.Command{
	Use:   "speedtest",
	Short: "Speedtest",
	Run: func(cmd *cobra.Command, args []string) {
		// https://github.com/showwin/speedtest-go#api-usage

		var speedtestClient = speedtest.New()

		serverList, _ := speedtestClient.FetchServers()
		targets, _ := serverList.FindServer([]int{})

		for _, s := range targets {
			fmt.Printf("Server: %s\n", s.Name)

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
				fmt.Sprintf("Latency: %s\n", s.Latency) +
				fmt.Sprintf("Download: %s\n", s.DLSpeed) +
				fmt.Sprintf("Upload: %s\n", s.ULSpeed),
			)

			s.Context.Reset()
		}
	},
}
