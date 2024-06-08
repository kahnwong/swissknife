package get

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestGetIp(t *testing.T) {
	ip, err := getLocalIP()
	if err != nil {
		t.Errorf("getLocalIP() error = %v", err)
	} else {
		counter := 0
		for _, v := range strings.Split(ip, ".") {
			vInt, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
			}
			if vInt >= 0 && vInt <= 256 {
				counter += 1
			}
		}

		if counter != 4 {
			t.Errorf("getLocalIP(): = %s, is not a valid IP", ip)
		}
	}
}
