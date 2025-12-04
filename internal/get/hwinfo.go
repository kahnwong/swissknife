package get

import "C"
import (
	"fmt"

	"github.com/jaypipes/ghw"
	"github.com/kahnwong/swissknife/configs/color"
	"github.com/rs/zerolog/log"
	"github.com/yumaojun03/dmidecode"
)

func HwInfo() {
	// cpu
	cpuModel, cpuThreads := getCpuInfo() // shared with `sysinfo.go`
	fmt.Printf("%s: %s (%v)\n", color.Green("CPU"), cpuModel, cpuThreads)

	// gpu
	gpu, err := ghw.GPU()
	if err != nil {
		fmt.Printf("Error getting GPU info: %v", err)
	}

	if len(gpu.GraphicsCards) > 0 {
		fmt.Printf("%s:\n", color.Green("GPUs"))

		for _, card := range gpu.GraphicsCards {
			fmt.Printf("  - %s: %s\n", color.Blue("Vendor"), card.DeviceInfo.Vendor.Name)
			fmt.Printf("    %s: %s\n", color.Blue("Model"), card.DeviceInfo.Product.Name)
		}
	}

	// memory
	dmi, err := dmidecode.New()
	if err != nil {
		log.Fatal().Msg("Failed to get dmi info")
	}

	fmt.Printf("%s:\n", color.Green("Memory"))

	memoryDevices, err := dmi.MemoryDevice()
	if err != nil {
		log.Fatal().Msg("Failed to get memory info")
	}

	for _, i := range memoryDevices {
		fmt.Printf("  - %s: %s\n", color.Blue("Manufacturer"), i.Manufacturer)
		fmt.Printf("    %s: %s\n", color.Blue("Type"), i.Type)
		fmt.Printf("    %s: %v GB\n", color.Blue("Size"), i.Size/1024)
		fmt.Printf("    %s: %v MHz\n", color.Blue("Speed"), i.Speed)
		fmt.Printf("    %s: %s \n", color.Blue("Model"), i.PartNumber)
	}

	// disk
	block, err := ghw.Block()
	if err != nil {
		log.Fatal().Msg("Failed to get block storage info")
	}

	fmt.Printf("%s:\n", color.Green("Disks"))
	for _, disk := range block.Disks {
		if disk.DriveType.String() != "virtual" {
			fmt.Printf("  - %s: %s\n", color.Blue("Type"), disk.DriveType)
			fmt.Printf("    %s: %s\n", color.Blue("Model"), disk.Model)
			fmt.Printf("    %s: %v GB\n", color.Blue("Size"), disk.SizeBytes/1000/1000/1000)
		}
	}

	// mainboard
	//// mainboardInfo, err := dmi.BaseBoard()
	baseboard, err := ghw.Baseboard()
	if err != nil {
		log.Fatal().Msg("Failed to get baseboard info")
	}

	fmt.Printf("%s:\n", color.Green("Mainboard"))
	fmt.Printf("  - %s: %s\n", color.Blue("Manufacturer"), baseboard.Vendor)
	fmt.Printf("    %s: %s\n", color.Blue("Model"), baseboard.Product)
}
