package get

import (
	"fmt"
	"log"
	"math"
	"os/user"

	"github.com/shirou/gopsutil/v4/disk"

	"github.com/shirou/gopsutil/v4/mem"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/spf13/cobra"
)

func convertKBtoGB(kb uint64) float64 {
	return math.Ceil(float64(kb)/float64(1024)/float64(1024)) / float64(1024)
}

type SystemInfo struct {
	Username string
	Hostname string
	Platform string
	CPU      string
	RAM      string
	Disk     string
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

	// disk
	diskStat, err := disk.Usage("/")
	if err != nil {
		log.Fatalf(err.Error())
	}

	return SystemInfo{
		Username: username.Username,
		Hostname: hostStat.Hostname,
		Platform: fmt.Sprintf("%s %s", hostStat.Platform, hostStat.PlatformVersion),
		CPU:      fmt.Sprintf("%s (%v)", cpuStat[0].ModelName, cpuThreads),
		RAM:      fmt.Sprintf("%.2f / %.2f GB", convertKBtoGB(vmStat.Used), convertKBtoGB(vmStat.Total)),
		Disk:     fmt.Sprintf("%.2f / %.2f GB", convertKBtoGB(diskStat.Used), convertKBtoGB(diskStat.Total)),
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
		green := color.New(color.FgHiGreen).SprintFunc()
		systemInfoStr := "" +
			fmt.Sprintf("%s@%s\n", green(systemInfo.Username), green(systemInfo.Hostname)) +
			fmt.Sprintf("%s:      %s\n", green("OS"), systemInfo.Platform) +
			fmt.Sprintf("%s:     %s\n", green("CPU"), systemInfo.CPU) +
			fmt.Sprintf("%s:  %s\n", green("Memory"), systemInfo.RAM) +
			fmt.Sprintf("%s:    %s", green("Disk"), systemInfo.Disk)

		fmt.Println(systemInfoStr)
	},
}

func init() {
	Cmd.AddCommand(getSystemInfoCmd)
}
