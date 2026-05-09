package cmd

import (
	"log/slog"
	"os"

	"github.com/kahnwong/swissknife/cmd/generate"
	"github.com/kahnwong/swissknife/cmd/get"
	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

var rootCmd = &cobra.Command{
	Use:           "swissknife",
	Version:       version,
	Short:         "Various utils",
	Long:          `Various utils`,
	SilenceErrors: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	logger := zerolog.New(output).With().Timestamp().Logger()
	slog.SetDefault(slog.New(slogzerolog.Option{Logger: &logger}.NewZerologHandler()))

	rootCmd.AddCommand(get.Cmd)
	rootCmd.AddCommand(generate.Cmd)
}
