package get

/*
#cgo linux LDFLAGS: ./lib/libsystem.a
#cgo darwin LDFLAGS: ./lib/libsystem.a -framework IOKit
#cgo windows LDFLAGS: ./lib/libsystem.a
#include "../../lib/system.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"

	"github.com/kahnwong/swissknife/configs/color"
	"github.com/rs/zerolog/log"
)

func Sensors() {
	result := C.sensors()

	switch result.error {
	case C.SENSOR_SUCCESS:
		fmt.Printf("%s: %.2f\n", color.Green("Temperature"), float64(result.temperature))
	case C.SENSOR_NO_COMPONENTS:
		log.Fatal().Msg("No components found")
	case C.SENSOR_NO_TEMPERATURE:
		log.Fatal().Msg("No temperature reading available")
	default:
		log.Fatal().Msg("Unknown error occurred")
	}
}
