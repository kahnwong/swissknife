package list

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"

	human "github.com/dustin/go-humanize"
	"github.com/kahnwong/swissknife/color"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/spf13/cobra"
)

func listVolumes() {
	// ref: <https://stackoverflow.com/a/64141403>
	// layout inspired by `duf`

	// init table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Mounted on", "Size", "Used", "Avail", "Use%", "Type", "Filesystem"})

	// get volumes info
	partitions, _ := disk.Partitions(false)
	for _, partition := range partitions {
		if partition.Fstype != "squashfs" && !strings.Contains(partition.Mountpoint, "snap") { // for example: snap partitions on ubuntu
			device := partition.Mountpoint
			stats, _ := disk.Usage(device)

			if stats.Total == 0 {
				continue
			}

			// disk available
			diskAvail := stats.Free
			diskAvailStr := ""
			if diskAvail < 50*1024*1024*1024 { // if free space less than 50 GB
				diskAvailStr = color.Red(human.Bytes(diskAvail))
			} else if diskAvail < 100*1024*1024*1024 { // if free space less than 100 GB
				diskAvailStr = color.Yellow(human.Bytes(diskAvail))
			} else {
				diskAvailStr = color.Green(human.Bytes(diskAvail))
			}

			// disk use percent
			percent := fmt.Sprintf("%2.f%%", stats.UsedPercent)
			percentRaw := stats.UsedPercent
			percentStr := ""
			if percentRaw > 80 {
				percentStr = color.Red(percent)
			} else if percentRaw > 70 {
				percentStr = color.Yellow(percent)
			} else {
				percentStr = color.Green(percent)
			}

			// append info to table
			t.AppendRows([]table.Row{
				{
					color.Blue(partition.Mountpoint),
					human.Bytes(stats.Total),
					human.Bytes(stats.Used),
					diskAvailStr,
					percentStr,
					color.Magenta(partition.Fstype),
					color.Magenta(stats.Fstype),
				},
			})
		}
	}

	// render
	t.Render()
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
