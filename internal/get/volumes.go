package get

import (
	"fmt"
	"math"
	"os"
	"strings"

	human "github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kahnwong/swissknife/configs/color"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/disk"
)

func listVolumes() {
	// ref: <https://stackoverflow.com/a/64141403>
	// layout inspired by `duf`

	// init table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Mounted on", "Size", "Used", "Avail", "Use%", "Type", "Filesystem"})

	// get volumes info
	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Fatal().Msg("Error getting partitions info")
	}

	for _, partition := range partitions {
		// linux
		isSquashFs := partition.Fstype == "squashfs"
		isSnap := strings.Contains(partition.Mountpoint, "snap")
		isKubernetes := strings.Contains(partition.Mountpoint, "kubelet")
		// osx
		isOsx := strings.Contains(partition.Mountpoint, "System/Volumes")
		isDevFs := partition.Fstype == "devfs"
		isOsxNix := strings.Contains(partition.Mountpoint, "/nix")
		if !isSquashFs && !isSnap && !isKubernetes && !isOsx && !isDevFs && !isOsxNix {
			device := partition.Mountpoint
			stats, err := disk.Usage(device)
			if err != nil {
				log.Fatal().Msg("Error getting disk info")
			}

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
			percentRaw := stats.UsedPercent
			diskUseTicks := int(math.Round(stats.UsedPercent)) / 10
			diskUseBar := fmt.Sprintf(
				"[%s%s] %2.f%%",
				strings.Repeat("#", diskUseTicks),
				strings.Repeat(".", 10-diskUseTicks),
				stats.UsedPercent,
			)
			diskUseStr := ""
			if percentRaw > 80 {
				diskUseStr = color.Red(diskUseBar)
			} else if percentRaw > 70 {
				diskUseStr = color.Yellow(diskUseBar)
			} else {
				diskUseStr = color.Green(diskUseBar)
			}

			// append info to table
			t.AppendRows([]table.Row{
				{
					color.Blue(partition.Mountpoint),
					human.Bytes(stats.Total),
					human.Bytes(stats.Used),
					diskAvailStr,
					diskUseStr,
					color.Magenta(partition.Fstype),
					color.Magenta(partition.Device),
				},
			})
		}
	}

	// render
	t.Render()
}

func Volumes() {
	listVolumes()
}
