package get

import (
	"fmt"
	"log"
	"math"
	"os/user"
	"strings"

	"github.com/kahnwong/swissknife/cmd/color"
	"github.com/shirou/gopsutil/v4/disk"

	"github.com/shirou/gopsutil/v4/mem"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/spf13/cobra"
)

func convertKBtoGB(v uint64) int {
	return int(math.Round(float64(v) / float64(1024) / float64(1024) / float64(1024)))
}

func convertToPercent(v float64) int {
	return int(math.Round(v * 100))
}

type SystemInfo struct {
	Username string
	Hostname string
	Platform string

	// cpu
	CPUModelName string
	CPUThreads   int

	// memory
	MemoryUsed        int
	MemoryTotal       int
	MemoryUsedPercent int

	// disk
	DiskUsed        int
	DiskTotal       int
	DiskUsedPercent int
}

func getSystemInfo() (SystemInfo, error) {
	// get info
	// ---- username ---- //
	username, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// ---- system ---- //
	// host
	hostStat, err := host.Info()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// cpu
	cpuStat, err := cpu.Info()
	if err != nil {
		log.Fatalf(err.Error())
	}

	cpuThreads, err := cpu.Counts(true)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// memory
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf(err.Error())
	}

	memoryUsed := convertKBtoGB(vmStat.Used)
	memoryTotal := convertKBtoGB(vmStat.Total)
	memoryUsedPercent := convertToPercent(float64(memoryUsed) / float64(memoryTotal))

	// disk
	diskStat, err := disk.Usage("/")
	if err != nil {
		log.Fatalf(err.Error())
	}

	diskUsed := convertKBtoGB(diskStat.Used)
	diskTotal := convertKBtoGB(diskStat.Total)
	diskUsedPercent := convertToPercent(float64(diskUsed) / float64(diskTotal))

	return SystemInfo{
		Username:          username.Username,
		Hostname:          hostStat.Hostname,
		Platform:          fmt.Sprintf("%s %s", hostStat.Platform, hostStat.PlatformVersion),
		CPUModelName:      cpuStat[0].ModelName,
		CPUThreads:        cpuThreads,
		MemoryUsed:        memoryUsed,
		MemoryTotal:       memoryTotal,
		MemoryUsedPercent: memoryUsedPercent,
		DiskUsed:          diskUsed,
		DiskTotal:         diskTotal,
		DiskUsedPercent:   diskUsedPercent,
	}, err
}

var getSystemInfoCmd = &cobra.Command{
	Use:   "system-info",
	Short: "Get system info",
	Run: func(cmd *cobra.Command, args []string) {
		systemInfo, err := getSystemInfo()
		if err != nil {
			fmt.Println(err)
		}

		// format message
		cpuInfo := fmt.Sprintf("%s (%v)", systemInfo.CPUModelName, systemInfo.CPUThreads)
		memoryInfo := fmt.Sprintf("%v/%v GB (%v%%)", systemInfo.MemoryUsed, systemInfo.MemoryTotal, color.Blue(systemInfo.MemoryUsedPercent))
		diskInfo := fmt.Sprintf("%v/%v GB (%v%%)", systemInfo.DiskUsed, systemInfo.DiskTotal, color.Blue(systemInfo.DiskUsedPercent))

		systemInfoStr := "" +
			fmt.Sprintf("%s@%s\n", color.Green(systemInfo.Username), color.Green(systemInfo.Hostname)) +
			strings.Repeat("-", len(systemInfo.Username)+len(systemInfo.Hostname)+1) + "\n" +
			fmt.Sprintf("%s:      %s\n", color.Green("OS"), systemInfo.Platform) +
			fmt.Sprintf("%s:     %s\n", color.Green("CPU"), cpuInfo) +
			fmt.Sprintf("%s:  %s\n", color.Green("Memory"), memoryInfo) +
			fmt.Sprintf("%s:    %s", color.Green("Disk"), diskInfo)

		fmt.Println(systemInfoStr)
	},
}

func init() {
	Cmd.AddCommand(getSystemInfoCmd)
}
