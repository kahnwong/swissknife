package networking

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

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

func getPublicIP() string {
	// make request
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// parse
	type Response struct {
		Origin string `json:"origin"`
	}

	var jsonResponse Response
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		log.Fatal(err)
	}

	return jsonResponse.Origin
}

var getIPCmd = &cobra.Command{
	Use:   "get-ip",
	Short: "Get IP information",
	Long:  `Get IP information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Local IP   : %s\n", getLocalIP())
		fmt.Printf("Public IP  : %s\n", getPublicIP())
	},
}

func init() {
	Cmd.AddCommand(getIPCmd)
}
