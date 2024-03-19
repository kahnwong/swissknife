package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/fatih/color"

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
	// make request
	resp, err := http.Get("https://api.country.is")
	if err != nil {
		return PublicIPResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PublicIPResponse{}, err
	}

	// parse

	var jsonResponse PublicIPResponse
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return PublicIPResponse{}, err
	}

	return jsonResponse, nil
}

var getIPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get IP information",
	Long:  `Get IP information`,
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen).SprintFunc()

		localIP, err := getLocalIP()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Local IP   : %s\n", green(localIP))
		}

		publicIP, err := getPublicIP()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Public IP  : %s (%s)\n", green(publicIP.Ip), publicIP.Country)
		}
	},
}

func init() {
	Cmd.AddCommand(getIPCmd)
}
