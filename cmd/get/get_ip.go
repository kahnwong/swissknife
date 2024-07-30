package get

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/carlmjohnson/requests"
	"github.com/kahnwong/swissknife/color"
	"github.com/spf13/cobra"
)

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
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

	return fmt.Sprintf("%v", localIP), nil
}

type PublicIPResponse struct {
	Ip      string `json:"ip"`
	Country string `json:"country"`
}

func getPublicIP() (PublicIPResponse, error) {
	var response PublicIPResponse
	err := requests.
		URL("https://api.country.is").
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		fmt.Println("Error getting public ip:", err)
	}
	return response, nil
}

var getIPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get IP information",
	Long:  `Get IP information`,
	Run: func(cmd *cobra.Command, args []string) {
		localIP, err := getLocalIP()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s: %s\n", color.Green("Local IP"), localIP)
		}

		publicIP, err := getPublicIP()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s: %s (%s)\n", color.Green("Public IP"), publicIP.Ip, color.Blue(publicIP.Country))
		}
	},
}
