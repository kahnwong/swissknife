package get

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/kahnwong/swissknife/configs/color"
	"github.com/kahnwong/swissknife/internal/utils"
	"github.com/rs/zerolog/log"
)

type IPInfoResponse struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

func getIPInfo(ip string) IPInfoResponse {
	var response IPInfoResponse
	err := requests.
		URL(fmt.Sprintf("http://ip-api.com/json/%s", ip)).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		log.Fatal().Msg("Error getting detailed ip info")
	}

	return response
}

func IPInfo(args []string) {
	ip := utils.SetIP(args)

	var targetIP string
	if ip == "" {
		publicIP := getPublicIP()
		targetIP = publicIP.Ip
	} else {
		targetIP = ip
	}

	ipInfo := getIPInfo(targetIP)

	fmt.Printf("%s: %s\n", color.Green("IP Address"), ipInfo.Query)
	fmt.Printf("%s: %s, %s, %s\n", color.Green("Location"), ipInfo.City, ipInfo.RegionName, color.Blue(ipInfo.Country))
	fmt.Printf("%s: %s\n", color.Green("ISP"), ipInfo.Isp)
	fmt.Printf("%s: %s\n", color.Green("Organization"), ipInfo.Org)
}
