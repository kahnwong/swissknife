package utils

import (
	"net"
	"strings"
)

func SetIP(args []string) string {
	var ip string
	if len(args) == 0 {
		ipFromClipboard := ReadFromClipboard()
		if ipFromClipboard != "" {
			// Check if it's a valid IP address
			if net.ParseIP(strings.TrimSpace(ipFromClipboard)) != nil {
				ip = strings.TrimSpace(ipFromClipboard)
			}
		}
	}
	if ip == "" {
		if len(args) == 0 {
			// Return empty string to use public IP
			return ""
		} else if len(args) == 1 {
			ip = args[0]
		}
	}

	return ip
}
