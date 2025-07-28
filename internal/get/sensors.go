package get

/*
#cgo LDFLAGS: ./lib/libsystem.a -ldl
#include "../../lib/system.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"github.com/kahnwong/swissknife/configs/color"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/sensors"
)

func Sensors() {
	fmt.Println(float32(C.sensors()))

	//str1 := C.CString("world")
	//defer C.free(unsafe.Pointer(str1))
	//C.hello(str1)

	s, err := sensors.SensorsTemperatures()
	if err != nil {
		log.Fatal().Msg("Failed to get sensors info")
	}

	var temperature float64
	for _, sensor := range s {
		if temperature < sensor.Temperature {
			temperature = sensor.Temperature
		}
	}

	fmt.Printf("%s: %.2f\n", color.Green("Temperature"), temperature)
}
