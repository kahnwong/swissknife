package get

import (
	"fmt"
	"runtime"

	"github.com/anatol/smart.go"
	"github.com/kahnwong/swissknife/configs/color"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/disk"
)

func Smart() {
	if runtime.GOOS == "linux" {
		device := getRootDiskVolume()

		isNvme := true
		dev, err := smart.OpenNVMe(device)
		if err != nil {
			log.Fatal().Err(err).Msgf("Failed to open %s device. You probaby have to run as sudo", device)
		}
		defer func(dev smart.Device) {
			err = dev.Close()
			if err != nil {
				log.Fatal().Err(err).Msgf("Failed to close %s device", device)
			}
		}(dev)

		if isNvme {
			sm, err := dev.ReadSMART()
			if err != nil {
				log.Fatal().Msg("Failed to read SMART info")
			}
			fmt.Printf("%s: %v\n", color.Green("Percent Used"), sm.PercentUsed)
		}
	} else {
		log.Error().Msgf("%s is not supported\n", runtime.GOOS)
	}
}

func getRootDiskVolume() string {
	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Fatal().Msg("Failed to get disk partitions")
	}

	var volume string
	for _, partition := range partitions {
		if partition.Mountpoint == "/" {
			volume = partition.Device
		}
	}

	return volume
}
