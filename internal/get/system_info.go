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
	"math"
	"os/user"
	"strconv"
	"strings"

	"github.com/distatus/battery"
	"github.com/shirou/gopsutil/v4/cpu"

	"github.com/kahnwong/swissknife/configs/color"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type batteryStruct struct {
	BatteryCurrent        float64
	BatteryFull           float64
	BatteryDesignCapacity float64
	BatteryCycleCount     uint64
	BatteryTimeToEmpty    uint64
}

func SysInfo() {
	// host
	username := getUsername()
	hostInfo := getHostInfo()
	fmt.Printf("%s@%s\n", color.Green(username), color.Green(hostInfo.Hostname))
	fmt.Println(strings.Repeat("-", len(username)+len(hostInfo.Hostname)+1))

	// os
	fmt.Printf("%s: %s %s\n", color.Green("OS"), hostInfo.Platform, hostInfo.PlatformVersion)

	// cpu
	cpuModel, cpuThreads := getCpuInfo()
	fmt.Printf("%s: %s (%v)\n", color.Green("CPU"), cpuModel, cpuThreads)

	// memory
	memoryUsed, memoryTotal, memoryUsedPercent := getMemoryInfo()
	fmt.Printf("%s: %v GB / %v GB (%s)\n",
		color.Green("Memory"),
		convertKBtoGB(memoryUsed), convertKBtoGB(memoryTotal),
		color.Blue(strconv.Itoa(memoryUsedPercent)+"%"),
	)

	// disk
	diskUsed, diskTotal, diskUsedPercent := getDiskInfo()
	fmt.Printf("%s: %v GB / %v GB (%s)\n",
		color.Green("Disk"),
		convertKBtoGB(diskUsed), convertKBtoGB(diskTotal),
		color.Blue(strconv.Itoa(diskUsedPercent)+"%"),
	)

	// battery
	batteryInfo := getBatteryInfo()

	// only print battery info if is a laptop
	if batteryInfo.BatteryFull > 0 {
		batteryPercent := convertToPercent(batteryInfo.BatteryCurrent / batteryInfo.BatteryFull)
		batteryHealth := convertToPercent(batteryInfo.BatteryFull / batteryInfo.BatteryDesignCapacity)
		batteryFormat := fmt.Sprintf("%v%%", batteryPercent)
		var batteryPercentStr string
		if batteryPercent > 80 {
			batteryPercentStr = color.Green(batteryFormat)
		} else if batteryPercent > 70 {
			batteryPercentStr = color.Yellow(batteryFormat)
		} else {
			batteryPercentStr = color.Red(batteryFormat)
		}

		// convert BatteryTimeToEmpty from second to hour
		var batteryTimeToEmptyFormatted string
		if batteryInfo.BatteryTimeToEmpty > 0 {
			hours := batteryInfo.BatteryTimeToEmpty / 3600
			minutes := (batteryInfo.BatteryTimeToEmpty % 3600) / 60
			batteryTimeToEmptyFormatted = fmt.Sprintf("%02d:%02d", hours, minutes)
		} else {
			batteryTimeToEmptyFormatted = "--:--"
		}

		fmt.Printf(
			"%s: %s (Health: %s, Cycles: %s, Time Remaining: %s)\n",
			color.Green("Battery"), batteryPercentStr,
			color.Blue(strconv.Itoa(batteryHealth)+"%"),
			color.Blue(strconv.Itoa(int(batteryInfo.BatteryCycleCount))),
			color.Blue(batteryTimeToEmptyFormatted),
		)
	}
}

// ---- functions ----
func getUsername() string {
	username, err := user.Current()
	if err != nil {
		log.Fatal().Msg("Failed to get current user info")
	}
	return username.Username
}

func getHostInfo() *host.InfoStat {
	hostStat, err := host.Info()
	if err != nil {
		log.Fatal().Msg("Failed to get host info")
	}
	return hostStat
}

func getCpuInfo() (string, int) {
	cpuStat, err := cpu.Info()
	if err != nil {
		log.Fatal().Msg("Failed to get cpu info")
	}
	cpuThreads, err := cpu.Counts(true)
	if err != nil {
		log.Fatal().Msg("Failed to get cpu threads info")
	}

	return cpuStat[0].ModelName, cpuThreads
}

func getMemoryInfo() (uint64, uint64, int) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal().Msg("Failed to get memory info")
	}
	return vmStat.Used, vmStat.Total, convertToPercent(float64(vmStat.Used) / float64(vmStat.Total))
}

func getDiskInfo() (uint64, uint64, int) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		log.Fatal().Msg("Failed to get disk info")
	}
	return diskStat.Used, diskStat.Total, convertToPercent(float64(diskStat.Used) / float64(diskStat.Total))
}

func getBatteryInfo() batteryStruct {
	// battery
	batteries, err := battery.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			// ignore this happens on [linux on mac devices]
		} else {
			log.Fatal().Msg("Error getting battery info")
		}
	}

	//// charge stats
	var batteryCurrent float64
	var batteryFull float64
	var batteryDesignCapacity float64
	if len(batteries) > 0 {
		batteryCurrent = batteries[0].Current
		batteryFull = batteries[0].Full
		batteryDesignCapacity = batteries[0].Design
	}

	//// cycle count
	var batteryCycleCount uint64
	result := C.battery_cycle_count()

	switch result.error {
	case C.BATTERY_SUCCESS:
		batteryCycleCount = uint64(result.cycle_count)
	case C.BATTERY_NO_BATTERY:
		batteryCycleCount = 0
	case C.BATTERY_NO_CYCLE_COUNT:
		batteryCycleCount = 0
	case C.BATTERY_MANAGER_ERROR:
		log.Fatal().Msg("Battery manager error")
	default:
		log.Fatal().Msg("Unknown error occurred")
	}

	//// time to empty
	var batteryTimeToEmpty uint64
	resultTimeToEmpty := C.battery_time_to_empty()

	switch result.error {
	case C.BATTERY_TIME_TO_EMPTY_SUCCESS:
		batteryTimeToEmpty = uint64(resultTimeToEmpty.time_to_empty_seconds)
	case C.BATTERY_TIME_TO_EMPTY_NO_BATTERY:
		batteryTimeToEmpty = 0
	case C.BATTERY_TIME_TO_EMPTY_NO_TIME_TO_EMPTY:
		batteryTimeToEmpty = 0
	case C.BATTERY_TIME_TO_EMPTY_MANAGER_ERROR:
		log.Fatal().Msg("Battery manager error")
	default:
		log.Fatal().Msg("Unknown error occurred")
	}

	// return
	return batteryStruct{
		BatteryCurrent:        batteryCurrent,
		BatteryFull:           batteryFull,
		BatteryDesignCapacity: batteryDesignCapacity,
		BatteryCycleCount:     batteryCycleCount,
		BatteryTimeToEmpty:    batteryTimeToEmpty,
	}
}

// ---- utils ----
func convertKBtoGB(v uint64) float64 {
	return math.Round(float64(v)/float64(1024)/float64(1024)/float64(1024)*100) / 100
}

func convertToPercent(v float64) int {
	return int(math.Round(v * 100))
}
