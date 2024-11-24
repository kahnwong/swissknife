package list

import (
	"fmt"

	human "github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/spf13/cobra"
)

func listVolumes() string {
	// ref: <https://stackoverflow.com/a/64141403>
	formatter := "%-14s %7s %7s %7s %4s %s\n"
	fmt.Printf(formatter, "Filesystem", "Size", "Used", "Avail", "Use%", "Mounted on")

	partitions, _ := disk.Partitions(false)
	for _, partition := range partitions {
		device := partition.Mountpoint
		stats, _ := disk.Usage(device)

		if stats.Total == 0 {
			continue
		}

		percent := fmt.Sprintf("%2.f%%", stats.UsedPercent)

		fmt.Printf(formatter,
			stats.Fstype,
			human.Bytes(stats.Total),
			human.Bytes(stats.Used),
			human.Bytes(stats.Free),
			percent,
			partition.Mountpoint,
		)
	}
	return "foo"
}

var listVolumesCmd = &cobra.Command{
	Use:   "volumes",
	Short: "List volumes",
	Run: func(cmd *cobra.Command, args []string) {
		listVolumes()
	},
}

func init() {
	Cmd.AddCommand(listVolumesCmd)
}
