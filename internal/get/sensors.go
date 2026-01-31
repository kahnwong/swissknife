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
)

func Sensors() error {
	result := C.sensors()

	switch result.error {
	case C.SENSOR_SUCCESS:
		fmt.Printf("%s: %.2f\n", color.Green("Temperature"), float64(result.temperature))
		return nil
	case C.SENSOR_NO_COMPONENTS:
		return fmt.Errorf("no components found")
	case C.SENSOR_NO_TEMPERATURE:
		return fmt.Errorf("no temperature reading available")
	default:
		return fmt.Errorf("unknown error occurred")
	}
}
