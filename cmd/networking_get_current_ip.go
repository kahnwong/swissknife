package cmd

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// sanitize
	parts := strings.Split(fmt.Sprintf("%v", localAddr), ":")

	localIP := ""
	if len(parts) > 0 {
		localIP = parts[0]
	} else {
		fmt.Println("Invalid address format")
	}

	return fmt.Sprintf("%v", localIP)
}

var getIPCmd = &cobra.Command{
	Use:   "get-ip",
	Short: "Get IP information",
	Long:  `Get IP information`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("Networking: get-ip")
		fmt.Printf("\tLocal IP: %s\n", getLocalIP())

		// public ip
	},
}

func init() {
	networkingCmd.AddCommand(getIPCmd)
}
