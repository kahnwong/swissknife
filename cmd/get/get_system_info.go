package get

import (
	"fmt"
	"math"
	"os/user"
	"strconv"
	"strings"

	"github.com/distatus/battery"

	"github.com/kahnwong/swissknife/color"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/spf13/cobra"
)

type SystemInfo struct {
	Username string
	Hostname string
	Platform string

	// cpu
	CPUModelName string
	CPUThreads   int

	// memory
	MemoryUsed  uint64
	MemoryTotal uint64

	// disk
	DiskUsed  uint64
	DiskTotal uint64

	// battery
	BatteryCurrent        float64
	BatteryFull           float64
	BatteryDesignCapacity float64
}

func convertKBtoGB(v uint64) float64 {
	return math.Round(float64(v)/float64(1024)/float64(1024)/float64(1024)*100) / 100
}

func convertToPercent(v float64) int {
	return int(math.Round(v * 100))
}

func getSystemInfo() SystemInfo {
	// get info
	// ---- username ---- //
	username, err := user.Current()
	if err != nil {
		log.Fatal().Msg("Failed to get current user info")
	}

	// ---- system ---- //
	// host
	hostStat, err := host.Info()
	if err != nil {
		log.Fatal().Msg("Failed to get host info")
	}

	// cpu
	cpuStat, err := cpu.Info()
	if err != nil {
		log.Fatal().Msg("Failed to get cpu info")
	}
	cpuThreads, err := cpu.Counts(true)
	if err != nil {
		log.Fatal().Msg("Failed to get cpu threads info")
	}

	// memory
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal().Msg("Failed to get memory info")
	}
	memoryUsed := vmStat.Used
	memoryTotal := vmStat.Total

	// disk
	diskStat, err := disk.Usage("/")
	if err != nil {
		log.Fatal().Msg("Failed to get disk info")
	}
	diskUsed := diskStat.Used
	diskTotal := diskStat.Total

	// battery
	batteries, err := battery.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			// ignore this happens on [linux on mac devices]
		} else {
			log.Fatal().Msg("Error getting battery info")
		}
	}

	var batteryCurrent float64
	var batteryFull float64
	var batteryDesignCapacity float64
	if len(batteries) > 0 {
		batteryCurrent = batteries[0].Current
		batteryFull = batteries[0].Full
		batteryDesignCapacity = batteries[0].Design
	}

	// return
	return SystemInfo{
		Username:              username.Username,
		Hostname:              hostStat.Hostname,
		Platform:              fmt.Sprintf("%s %s", hostStat.Platform, hostStat.PlatformVersion),
		CPUModelName:          cpuStat[0].ModelName,
		CPUThreads:            cpuThreads,
		MemoryUsed:            memoryUsed,
		MemoryTotal:           memoryTotal,
		DiskUsed:              diskUsed,
		DiskTotal:             diskTotal,
		BatteryCurrent:        batteryCurrent,
		BatteryFull:           batteryFull,
		BatteryDesignCapacity: batteryDesignCapacity,
	}
}

var getSystemInfoCmd = &cobra.Command{
	Use:   "system-info",
	Short: "Get system info",
	Run: func(cmd *cobra.Command, args []string) {
		systemInfo := getSystemInfo()

		// format message
		hostStdout := fmt.Sprintf("%s@%s", color.Green(systemInfo.Username), color.Green(systemInfo.Hostname))
		linebreakStdout := strings.Repeat("-", len(systemInfo.Username)+len(systemInfo.Hostname)+1)
		osStdout := fmt.Sprintf("%s: %s", color.Green("OS"), systemInfo.Platform)

		cpuInfo := fmt.Sprintf("%s (%v)", systemInfo.CPUModelName, systemInfo.CPUThreads)
		cpuStdout := fmt.Sprintf("%s: %s", color.Green("CPU"), cpuInfo)

		memoryUsedPercent := convertToPercent(float64(systemInfo.MemoryUsed) / float64(systemInfo.MemoryTotal))
		memoryInfo := fmt.Sprintf("%v GB / %v GB (%s)", convertKBtoGB(systemInfo.MemoryUsed), convertKBtoGB(systemInfo.MemoryTotal), color.Blue(strconv.Itoa(memoryUsedPercent)+"%"))
		memoryStdout := fmt.Sprintf("%s: %s", color.Green("Memory"), memoryInfo)

		diskUsedPercent := convertToPercent(float64(systemInfo.DiskUsed) / float64(systemInfo.DiskTotal))
		diskInfo := fmt.Sprintf("%v GB / %v GB (%s)", convertKBtoGB(systemInfo.DiskUsed), convertKBtoGB(systemInfo.DiskTotal), color.Blue(strconv.Itoa(diskUsedPercent)+"%"))
		diskStdout := fmt.Sprintf("%s: %s", color.Green("Disk"), diskInfo)

		systemInfoStdout := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", hostStdout, linebreakStdout, osStdout, cpuStdout, memoryStdout, diskStdout)
		fmt.Print(systemInfoStdout)

		// only print battery info if is a laptop
		if systemInfo.BatteryFull > 0 {
			batteryPercent := convertToPercent(systemInfo.BatteryCurrent / systemInfo.BatteryFull)
			batteryHealth := convertToPercent(systemInfo.BatteryFull / systemInfo.BatteryDesignCapacity)
			batteryFormat := fmt.Sprintf("%v%%", batteryPercent)
			var batteryPercentStr string
			if batteryPercent > 80 {
				batteryPercentStr = color.Green(batteryFormat)
			} else if batteryPercent > 70 {
				batteryPercentStr = color.Yellow(batteryFormat)
			} else {
				batteryPercentStr = color.Red(batteryFormat)
			}

			batteryStdout := fmt.Sprintf("%s: %s (Health: %s)", color.Green("Battery"), batteryPercentStr, color.Blue(strconv.Itoa(batteryHealth)+"%"))
			fmt.Println(batteryStdout)
		}
	},
}

func init() {
	Cmd.AddCommand(getSystemInfoCmd)
}
