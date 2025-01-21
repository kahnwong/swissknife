package get

import (
	"fmt"

	"github.com/kahnwong/swissknife/color"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/sensors"
	"github.com/spf13/cobra"
)

var SensorsCmd = &cobra.Command{
	Use:   "sensors",
	Short: "Get sensors info",
	Run: func(cmd *cobra.Command, args []string) {
		sensors, err := sensors.SensorsTemperatures()
		if err != nil {
			log.Fatal().Msg("Failed to get sensors info")
		}

		var temperature float64
		for _, sensor := range sensors {
			if temperature < sensor.Temperature {
				temperature = sensor.Temperature
			}
		}

		fmt.Printf("%s: %.2f\n", color.Green("Temperature"), temperature)
	},
}

func init() {
	Cmd.AddCommand(SensorsCmd)
}
