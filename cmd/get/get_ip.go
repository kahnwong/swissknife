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

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal().Err(err).Msg("Error on net.Dial")
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

type PublicIPResponse struct {
	Ip      string `json:"ip"`
	Country string `json:"country"`
}

func getPublicIP() PublicIPResponse {
	var response PublicIPResponse
	err := requests.
		URL("https://api.country.is").
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		log.Fatal().Err(err).Msg("Error getting public ip")
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
		fmt.Printf("%s: %s (%s)\n", color.Green("Public IP"), publicIP.Ip, color.Blue(publicIP.Country))
	},
}

func init() {
	Cmd.AddCommand(getIPCmd)
}
