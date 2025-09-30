package get

import (
	"os"
	"reflect"
	"runtime"

	"github.com/anatol/smart.go"
	"github.com/jedib0t/go-pretty/v6/table"
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

		// render table
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Name", "Value"})

		v := reflect.ValueOf(sm)
		typeOfS := v.Type()

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
			typeOfS = typeOfS.Elem()
		}

		for i := 0; i < v.NumField(); i++ {
			fieldName := typeOfS.Field(i).Name
			field := v.Field(i)

			if !field.CanInterface() {
				continue
			} else {
				fieldValue := field.Interface()

				t.AppendRows([]table.Row{
					{
						color.Green(fieldName),
						fieldValue,
					},
				})
			}
		}

		t.Render()

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

		// render table
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Name", "Current", "Worst", "Raw Value"})

		for _, i := range sm.Attrs {
			t.AppendRows([]table.Row{
				{
					color.Green(i.Name),
					i.Current,
					i.Worst,
					i.ValueRaw,
				},
			})
		}

		t.Render()

		return nil
	}
}

func Smart(args []string) {
	if runtime.GOOS == "linux" {
		var device string
		if len(args) == 0 {
			device = getRootDiskVolume()
		} else if len(args) == 1 {
			device = args[0]
		} else {
			log.Error().Msg("Too many arguments")
		}
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
