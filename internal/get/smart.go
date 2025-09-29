package get

import (
	"fmt"
	"runtime"

	"github.com/anatol/smart.go"
	"github.com/kahnwong/swissknife/configs/color"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/disk"
)

func nvmeSmart(device string) error {
	var err error
	dev, err := smart.OpenNVMe(device)
	if err != nil {
		return err
	} else {
		sm, err := dev.ReadSMART()
		if err != nil {
			return err
		}
		fmt.Printf("%s: %v\n", color.Green("Percent Used"), sm.PercentUsed)

		return nil
	}
}

func sataSmart(device string) error {
	var err error
	dev, err := smart.OpenSata(device)
	if err != nil {
		return err
	} else {
		sm, err := dev.ReadSMARTData()
		if err != nil {
			return err
		}
		fmt.Printf("%s: %v\n", color.Green("Media Wearout Indicator"), sm.Attrs[230].Current)

		return nil
	}
}

func Smart() {
	if runtime.GOOS == "linux" {
		device := getRootDiskVolume()

		err := nvmeSmart(device)
		if err != nil {
			err = sataSmart(device)
			if err != nil {
				log.Error().Err(err).Msg("Unrecognized device type. It's not NVME or SATA, or you didn't run as sudo")
			}
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
