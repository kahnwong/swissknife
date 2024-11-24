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
				fmt.Printf("%s:   %s\n", color.Green("Server"), s.Name)

				err := s.PingTest(nil)
				if err != nil {
					log.Fatal().Err(err).Msg("Error pinging server")
				}

				err = s.DownloadTest()
				if err != nil {
					log.Fatal().Err(err).Msg("Error testing download speed")
				}

				err = s.UploadTest()
				if err != nil {
					log.Fatal().Err(err).Msg("Error testing upload speed")
				}
			}

			cancel()
		}()
		<-tests

		// print results
		fmt.Print(
			"\n" +
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
