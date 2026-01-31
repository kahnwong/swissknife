package get

import (
	"fmt"
	"net"

	"github.com/libp2p/go-netroute"
)

// unix: `netstat -nr -f inet`
// get iface-owned app: `ifconfig -v utun3 | grep "agent domain"`

func getIface() (string, error) {
	r, err := netroute.New()
	if err != nil {
		return "", fmt.Errorf("error initializing netroute: %w", err)
	}

	iface, _, _, err := r.Route(
		net.IPv4(104, 16, 133, 229), // cloudflare
	)
	if err != nil {
		return "", fmt.Errorf("error retrieving net route: %w", err)
	}
	//fmt.Printf("%v, %v, %v, %v\n", iface, gw, src, err)

	return iface.Name, nil
}

func Iface() error {
	iface, err := getIface()
	if err != nil {
		return err
	}
	fmt.Println(iface)
	return nil
}
