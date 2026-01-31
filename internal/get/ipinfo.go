package get

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/kahnwong/swissknife/configs/color"
	"github.com/kahnwong/swissknife/internal/utils"
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

func getIPInfo(ip string) (IPInfoResponse, error) {
	var response IPInfoResponse
	err := requests.
		URL(fmt.Sprintf("http://ip-api.com/json/%s", ip)).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		return IPInfoResponse{}, fmt.Errorf("error getting detailed ip info: %w", err)
	}

	return response, nil
}

func IPInfo(args []string) error {
	ip := utils.SetIP(args)

	var targetIP string
	if ip == "" {
		publicIP, err := getPublicIP()
		if err != nil {
			return err
		}
		targetIP = publicIP.Ip
	} else {
		targetIP = ip
	}

	ipInfo, err := getIPInfo(targetIP)
	if err != nil {
		return err
	}

	fmt.Printf("%s: %s\n", color.Green("IP Address"), ipInfo.Query)
	fmt.Printf("%s: %s, %s, %s\n", color.Green("Location"), ipInfo.City, ipInfo.RegionName, color.Blue(ipInfo.Country))
	fmt.Printf("%s: %s\n", color.Green("ISP"), ipInfo.Isp)
	fmt.Printf("%s: %s\n", color.Green("Organization"), ipInfo.Org)
	return nil
}
