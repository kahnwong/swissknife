package cmd

import (
	"fmt"
	"os"

	"github.com/anatol/smart.go"

	"github.com/kahnwong/swissknife/cmd/generate"
	"github.com/kahnwong/swissknife/cmd/get"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

var rootCmd = &cobra.Command{
	Use:     "swissknife",
	Version: version,
	Short:   "Various utils",
	Long:    `Various utils`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// skip the error handling for more compact API example
	dev, _ := smart.OpenNVMe("/dev/disk3s1s1")
	c, nss, _ := dev.Identify()
	fmt.Println("Model number: ", c.ModelNumber())
	fmt.Println("Serial number: ", c.SerialNumber())
	fmt.Println("Size: ", c.Tnvmcap.Val[0])

	// namespace #1
	ns := nss[0]
	fmt.Println("Namespace 1 utilization: ", ns.Nuse*ns.LbaSize())

	sm, _ := dev.ReadSMART()
	fmt.Println("Temperature: ", sm.Temperature, "K")
	// PowerOnHours is reported as 128-bit value and represented by this library as an array of uint64
	fmt.Println("Power-on hours: ", sm.PowerOnHours.Val[0])
	fmt.Println("Power cycles: ", sm.PowerCycles.Val[0])

	fmt.Println(sm.PercentUsed)

	rootCmd.AddCommand(get.Cmd)
	rootCmd.AddCommand(generate.Cmd)
}
