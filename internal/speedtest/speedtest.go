package speedtest

import (
	"context"
	"fmt"
	"time"

	"github.com/kahnwong/swissknife/configs/color"
	"github.com/showwin/speedtest-go/speedtest"
)

func SpeedTest() error {
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

	// Store error from goroutine
	errChan := make(chan error, 1)

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
				errChan <- fmt.Errorf("error pinging server: %w", err)
				cancel()
				return
			}
			fmt.Printf("\033[2K\r") // clear line

			fmt.Print("Downloading")
			err = s.DownloadTest()
			if err != nil {
				errChan <- fmt.Errorf("error testing download speed: %w", err)
				cancel()
				return
			}
			fmt.Printf("\033[2K\r") // clear line

			fmt.Print("Uploading")
			err = s.UploadTest()
			if err != nil {
				errChan <- fmt.Errorf("error testing upload speed: %w", err)
				cancel()
				return
			}
			fmt.Printf("\033[2K\r") // clear line
		}

		errChan <- nil
		cancel()
	}()
	<-tests

	// Check for errors
	if err := <-errChan; err != nil {
		return err
	}

	// print results
	fmt.Print(
		fmt.Sprintf("%s:  %s\n", color.Green("Latency"), s.Latency.Truncate(time.Millisecond)) +
			fmt.Sprintf("%s: %s\n", color.Green("Download"), s.DLSpeed) +
			fmt.Sprintf("%s:   %s\n", color.Green("Upload"), s.ULSpeed),
	)

	s.Context.Reset()
	return nil
}
