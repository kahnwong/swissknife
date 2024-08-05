package get

import (
	"fmt"
	"math"
	"os/user"
	"strconv"
	"strings"

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
	MemoryUsed  int
	MemoryTotal int

	// disk
	DiskUsed  int
	DiskTotal int
}

func convertKBtoGB(v uint64) int {
	return int(math.Round(float64(v) / float64(1024) / float64(1024) / float64(1024)))
}

func convertToPercent(v float64) int {
	return int(math.Round(v * 100))
}

func getSystemInfo() SystemInfo {
	// get info
	// ---- username ---- //
	username, err := user.Current()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get current user info")
	}

	// ---- system ---- //
	// host
	hostStat, err := host.Info()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get host info")
	}

	// cpu
	cpuStat, err := cpu.Info()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get cpu info")
	}
	cpuThreads, err := cpu.Counts(true)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get cpu threads info")
	}

	// memory
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get memory info")
	}
	memoryUsed := convertKBtoGB(vmStat.Used)
	memoryTotal := convertKBtoGB(vmStat.Total)

	// disk
	diskStat, err := disk.Usage("/")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get disk info")
	}
	diskUsed := convertKBtoGB(diskStat.Used)
	diskTotal := convertKBtoGB(diskStat.Total)

	return SystemInfo{
		Username:     username.Username,
		Hostname:     hostStat.Hostname,
		Platform:     fmt.Sprintf("%s %s", hostStat.Platform, hostStat.PlatformVersion),
		CPUModelName: cpuStat[0].ModelName,
		CPUThreads:   cpuThreads,
		MemoryUsed:   memoryUsed,
		MemoryTotal:  memoryTotal,
		DiskUsed:     diskUsed,
		DiskTotal:    diskTotal,
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
		memoryInfo := fmt.Sprintf("%v/%v GB (%s)", systemInfo.MemoryUsed, systemInfo.MemoryTotal, color.Blue(strconv.Itoa(memoryUsedPercent)+"%"))
		memoryStdout := fmt.Sprintf("%s: %s", color.Green("Memory"), memoryInfo)

		diskUsedPercent := convertToPercent(float64(systemInfo.DiskUsed) / float64(systemInfo.DiskTotal))
		diskInfo := fmt.Sprintf("%v/%v GB (%s)", systemInfo.DiskUsed, systemInfo.DiskTotal, color.Blue(strconv.Itoa(diskUsedPercent)+"%"))
		diskStdout := fmt.Sprintf("%s: %s", color.Green("Disk"), diskInfo)

		systemInfoStdout := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", hostStdout, linebreakStdout, osStdout, cpuStdout, memoryStdout, diskStdout)

		fmt.Println(systemInfoStdout)
	},
}

func init() {
	Cmd.AddCommand(getSystemInfoCmd)
}
