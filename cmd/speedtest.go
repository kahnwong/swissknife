package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/kahnwong/swissknife/color"
	"github.com/rs/zerolog/log"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/spf13/cobra"
)

var SpeedTestCmd = &cobra.Command{
	Use:   "speedtest",
	Short: "Speedtest",
	Run: func(cmd *cobra.Command, args []string) {
		// ref: <https://github.com/showwin/speedtest-go#api-usage>
		var speedtestClient = speedtest.New()
		serverList, _ := speedtestClient.FetchServers()
		targets, _ := serverList.FindServer([]int{})

		// print test server location
		fmt.Printf("%s:   %s\n", color.Green("Server"), targets[0].Name)

		// start tests
		var s *speedtest.Server
		tests := make(chan struct{})
		ctx, cancel := context.WithCancel(context.Background())

		// -- background progress report -- //
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					tests <- struct{}{}
					return
				default:
					fmt.Print(".") // for progress report
				}

				time.Sleep(500 * time.Millisecond)
			}
		}(ctx)

		// -- actual tests -- //
		go func() {
			for _, s = range targets {
				fmt.Print("Pinging")
				err := s.PingTest(nil)
				if err != nil {
					log.Fatal().Err(err).Msg("Error pinging server")
				}
				fmt.Printf("\033[2K\r") // clear line

				fmt.Print("Downloading")
				err = s.DownloadTest()
				if err != nil {
					log.Fatal().Err(err).Msg("Error testing download speed")
				}
				fmt.Printf("\033[2K\r") // clear line

				fmt.Print("Uploading")
				err = s.UploadTest()
				if err != nil {
					log.Fatal().Err(err).Msg("Error testing upload speed")
				}
				fmt.Printf("\033[2K\r") // clear line
			}

			cancel()
		}()
		<-tests

		// print results
		fmt.Print(
			fmt.Sprintf("%s:  %s\n", color.Green("Latency"), s.Latency.Truncate(time.Millisecond)) +
				fmt.Sprintf("%s: %s\n", color.Green("Download"), s.DLSpeed) +
				fmt.Sprintf("%s:   %s\n", color.Green("Upload"), s.ULSpeed),
		)

		s.Context.Reset()
	},
}

func init() {
	rootCmd.AddCommand(SpeedTestCmd)
}
