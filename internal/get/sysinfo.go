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

func SysInfo() error {
	// host
	username, err := getUsername()
	if err != nil {
		return err
	}
	hostInfo, err := getHostInfo()
	if err != nil {
		return err
	}
	fmt.Printf("%s@%s\n", color.Green(username), color.Green(hostInfo.Hostname))
	fmt.Println(strings.Repeat("-", len(username)+len(hostInfo.Hostname)+1))

	// os
	fmt.Printf("%s: %s %s\n", color.Green("OS"), hostInfo.Platform, hostInfo.PlatformVersion)

	// cpu
	cpuModel, cpuThreads, err := getCpuInfo()
	if err != nil {
		return err
	}
	fmt.Printf("%s: %s (%v)\n", color.Green("CPU"), cpuModel, cpuThreads)

	// memory
	memoryUsed, memoryTotal, memoryUsedPercent, err := getMemoryInfo()
	if err != nil {
		return err
	}
	fmt.Printf("%s: %.2f GB / %v GB (%s)\n",
		color.Green("Memory"),
		convertKBtoGB(memoryUsed, true), convertKBtoGB(memoryTotal, false),
		color.Blue(strconv.Itoa(memoryUsedPercent)+"%"),
	)

	// disk
	diskUsed, diskTotal, diskUsedPercent, err := getDiskInfo()
	if err != nil {
		return err
	}
	fmt.Printf("%s: %v GB / %v GB (%s)\n",
		color.Green("Disk"),
		convertKBtoGiB(diskUsed), convertKBtoGiB(diskTotal),
		color.Blue(strconv.Itoa(diskUsedPercent)+"%"),
	)

	// battery
	batteryInfo, err := getBatteryInfo()
	if err != nil {
		return err
	}

	// only print battery info if is a laptop
	if batteryInfo.BatteryFull > 0 {
		batteryPercent := convertToPercent(batteryInfo.BatteryCurrent / batteryInfo.BatteryFull)
		batteryHealth := convertToPercent(batteryInfo.BatteryFull / batteryInfo.BatteryDesignCapacity)
		batteryFormat := fmt.Sprintf("%v%%", batteryPercent)
		var batteryPercentStr string
		if batteryPercent > 80 {
			batteryPercentStr = color.Green(batteryFormat)
		} else if batteryPercent > 60 {
			batteryPercentStr = color.Yellow(batteryFormat)
		} else if batteryPercent > 30 {
			batteryPercentStr = color.Red(batteryFormat)
		} else {
			batteryPercentStr = batteryFormat
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
	return nil
}

// ---- functions ----
func getUsername() (string, error) {
	username, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get current user info: %w", err)
	}
	return username.Username, nil
}

func getHostInfo() (*host.InfoStat, error) {
	hostStat, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get host info: %w", err)
	}
	return hostStat, nil
}

func getCpuInfo() (string, int, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return "", 0, fmt.Errorf("failed to get cpu info: %w", err)
	}
	cpuThreads, err := cpu.Counts(true)
	if err != nil {
		return "", 0, fmt.Errorf("failed to get cpu threads info: %w", err)
	}

	return cpuStat[0].ModelName, cpuThreads, nil
}

func getMemoryInfo() (uint64, uint64, int, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to get memory info: %w", err)
	}
	return vmStat.Used, vmStat.Total, convertToPercent(float64(vmStat.Used) / float64(vmStat.Total)), nil
}

func getDiskInfo() (uint64, uint64, int, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to get disk info: %w", err)
	}
	return diskStat.Used, diskStat.Total, convertToPercent(float64(diskStat.Used) / float64(diskStat.Total)), nil
}

func getBatteryInfo() (batteryStruct, error) {
	// battery
	batteries, err := battery.GetAll()
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			return batteryStruct{}, fmt.Errorf("error getting battery info: %w", err)
		}
		// ignore this happens on [linux on mac devices]
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
		return batteryStruct{}, fmt.Errorf("battery manager error")
	default:
		return batteryStruct{}, fmt.Errorf("unknown error occurred")
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
		return batteryStruct{}, fmt.Errorf("battery manager error")
	default:
		return batteryStruct{}, fmt.Errorf("unknown error occurred")
	}

	// return
	return batteryStruct{
		BatteryCurrent:        batteryCurrent,
		BatteryFull:           batteryFull,
		BatteryDesignCapacity: batteryDesignCapacity,
		BatteryCycleCount:     batteryCycleCount,
		BatteryTimeToEmpty:    batteryTimeToEmpty,
	}, nil
}

// ---- utils ----
func convertKBtoGiB(v uint64) float64 {
	return math.Round(float64(v)/float64(1024)/float64(1024)/float64(1024)*100) / 100
}

func convertKBtoGB(v uint64, isFloat bool) float64 {
	gb := (float64(v) / float64(1024) / float64(1024) / float64(1000) * 100) / 100
	if isFloat {
		return gb
	} else {
		return math.Round(gb)
	}
}

func convertToPercent(v float64) int {
	return int(math.Round(v * 100))
}
