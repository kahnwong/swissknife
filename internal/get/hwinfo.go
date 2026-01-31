package get

import "C"
import (
	"fmt"
	"os"

	"github.com/jaypipes/ghw"
	"github.com/kahnwong/swissknife/configs/color"
	"github.com/yumaojun03/dmidecode"
)

func HwInfo() error {
	// need to run as sudo
	if os.Geteuid() != 0 {
		return fmt.Errorf("need to run as sudo")
	}

	// cpu
	cpuModel, cpuThreads, err := getCpuInfo() // shared with `sysinfo.go`
	if err != nil {
		return err
	}
	fmt.Printf("%s: %s (%v)\n", color.Green("CPU"), cpuModel, cpuThreads)

	// gpu
	gpu, err := ghw.GPU()
	if err != nil {
		fmt.Printf("Error getting GPU info: %v\n", err)
	} else {
		fmt.Printf("%s:\n", color.Green("GPUs"))

		for _, card := range gpu.GraphicsCards {
			fmt.Printf("  - %s: %s\n", color.Blue("Vendor"), card.DeviceInfo.Vendor.Name)
			fmt.Printf("    %s: %s\n", color.Blue("Model"), card.DeviceInfo.Product.Name)
		}
	}

	// memory
	dmi, err := dmidecode.New()
	if err != nil {
		return fmt.Errorf("failed to get dmi info: %w", err)
	}

	fmt.Printf("%s:\n", color.Green("Memory"))

	memoryDevices, err := dmi.MemoryDevice()
	if err != nil {
		return fmt.Errorf("failed to get memory info: %w", err)
	}

	for _, i := range memoryDevices {
		if i.Type.String() != "Unknown" {
			fmt.Printf("  - %s: %s\n", color.Blue("Manufacturer"), i.Manufacturer)
			fmt.Printf("    %s: %s\n", color.Blue("Type"), i.Type)
			fmt.Printf("    %s: %v GB\n", color.Blue("Size"), i.Size/1024)
			fmt.Printf("    %s: %v MHz\n", color.Blue("Speed"), i.Speed)
			fmt.Printf("    %s: %s \n", color.Blue("Model"), i.PartNumber)
		}
	}

	// disk
	block, err := ghw.Block()
	if err != nil {
		return fmt.Errorf("failed to get block storage info: %w", err)
	}

	fmt.Printf("%s:\n", color.Green("Disks"))
	for _, disk := range block.Disks {
		if (disk.DriveType.String() != "virtual") && (disk.DriveType.String() != "Unknown") && (disk.DriveType.String() != "ODD") {
			fmt.Printf("  - %s: %s\n", color.Blue("Type"), disk.DriveType)
			fmt.Printf("    %s: %s\n", color.Blue("Model"), disk.Model)
			fmt.Printf("    %s: %v GB\n", color.Blue("Size"), disk.SizeBytes/1000/1000/1000)
		}
	}

	// mainboard
	//// mainboardInfo, err := dmi.BaseBoard()
	baseboard, err := ghw.Baseboard()
	if err != nil {
		return fmt.Errorf("failed to get baseboard info: %w", err)
	}

	fmt.Printf("%s:\n", color.Green("Mainboard"))
	fmt.Printf("  - %s: %s\n", color.Blue("Manufacturer"), baseboard.Vendor)
	fmt.Printf("    %s: %s\n", color.Blue("Model"), baseboard.Product)
	return nil
}
