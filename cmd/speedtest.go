package cmd

import (
	"fmt"

	"github.com/kahnwong/swissknife/color"
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
			fmt.Printf("%s:   %s\n", color.Green("Server"), s.Name)

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
				fmt.Sprintf("%s:  %s\n", color.Green("Latency"), s.Latency) +
				fmt.Sprintf("%s: %s\n", color.Green("Download"), s.DLSpeed) +
				fmt.Sprintf("%s:   %s\n", color.Green("Upload"), s.ULSpeed),
			)

			s.Context.Reset()
		}
	},
}

func init() {
	rootCmd.AddCommand(SpeedTestCmd)
}
