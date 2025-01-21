package get

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/carlmjohnson/requests"
	"github.com/kahnwong/swissknife/color"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type PublicIPResponse struct {
	Ip string `json:"ip"`
}

type IPLocation struct {
	Ip         string `json:"ip"`
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
}

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal().Msg("Error on net.Dial")
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// sanitize
	parts := strings.Split(fmt.Sprintf("%v", localAddr), ":")

	localIP := ""
	if len(parts) > 0 {
		localIP = parts[0]
	} else {
		log.Error().Msg("Invalid address format")
	}

	return localIP
}

func getPublicIP() PublicIPResponse {
	var response PublicIPResponse
	err := requests.
		URL("https://api.ipify.org?format=json").
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		log.Fatal().Msg("Error getting public ip")
	}

	return response
}

func getIPLocation(ip string) IPLocation {
	var response IPLocation
	err := requests.
		URL(fmt.Sprintf("http://ip-api.com/json/%s", ip)).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		log.Fatal().Msg("Error getting ip location")
	}

	return response
}

var getIPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get IP information",
	Long:  `Get IP information`,
	Run: func(cmd *cobra.Command, args []string) {
		localIP := getLocalIP()
		fmt.Printf("%s: %s\n", color.Green("Local IP"), localIP)

		publicIP := getPublicIP()
		IPLocation := getIPLocation(publicIP.Ip)
		fmt.Printf("%s: %s (%s, %s)\n", color.Green("Public IP"), publicIP.Ip, color.Blue(IPLocation.RegionName), color.Blue(IPLocation.Country))
	},
}

func init() {
	Cmd.AddCommand(getIPCmd)
}
