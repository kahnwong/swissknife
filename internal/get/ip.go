package get

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/carlmjohnson/requests"
	"github.com/kahnwong/swissknife/configs/color"
	"github.com/rs/zerolog/log"
)

type PublicIPResponse struct {
	Ip string `json:"ip"`
}

type IPLocationResponse struct {
	Ip         string `json:"ip"`
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
}

func getInternalIP() (string, error) {
	var internalIP string

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("failed to get network interfaces: %w", err)
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // skip loopback interface
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", fmt.Errorf("failed to get interface addresses: %w", err)
		}

		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			if ip.To4() != nil {
				if internalIP == "" {
					internalIP = ip.String()
				}
			}
		}
	}

	return internalIP, nil
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", fmt.Errorf("error on net.Dial: %w", err)
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("error closing connection")
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// sanitize
	parts := strings.Split(fmt.Sprintf("%v", localAddr), ":")

	localIP := ""
	if len(parts) > 0 {
		localIP = parts[0]
	} else {
		return "", fmt.Errorf("invalid address format")
	}

	return localIP, nil
}

func getPublicIP() (PublicIPResponse, error) {
	var response PublicIPResponse
	err := requests.
		URL("https://api.ipify.org?format=json").
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		return PublicIPResponse{}, fmt.Errorf("error getting public ip: %w", err)
	}

	return response, nil
}

func getIPLocation(ip string) (IPLocationResponse, error) {
	var response IPLocationResponse
	err := requests.
		URL(fmt.Sprintf("http://ip-api.com/json/%s?fields=query,country,regionName", ip)).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		return IPLocationResponse{}, fmt.Errorf("error getting ip location: %w", err)
	}

	return response, nil
}

func IP() error {
	internalIP, err := getInternalIP()
	if err != nil {
		return err
	}
	fmt.Printf("%s: %s\n", color.Green("Internal IP"), internalIP)

	localIP, err := getLocalIP()
	if err != nil {
		return err
	}
	fmt.Printf("%s: %s\n", color.Green("Local IP"), localIP)

	publicIP, err := getPublicIP()
	if err != nil {
		return err
	}

	IPLocation, err := getIPLocation(publicIP.Ip)
	if err != nil {
		return err
	}
	fmt.Printf("%s: %s (%s, %s)\n", color.Green("Public IP"), publicIP.Ip, color.Blue(IPLocation.RegionName), color.Blue(IPLocation.Country))
	return nil
}
