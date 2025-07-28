package get

import (
	"fmt"
	"net"

	"github.com/libp2p/go-netroute"
	"github.com/rs/zerolog/log"
)

// unix: `netstat -nr -f inet`
// get iface-owned app: `ifconfig -v utun3 | grep "agent domain"`

func getIface() string {
	r, err := netroute.New()
	if err != nil {
		log.Fatal().Msg("Error initializing netroute")
	}

	iface, _, _, err := r.Route(
		net.IPv4(104, 16, 133, 229), // cloudflare
	)
	if err != nil {
		log.Fatal().Msg("Error retrieving net route")
	}
	//fmt.Printf("%v, %v, %v, %v\n", iface, gw, src, err)

	return iface.Name
}

func Iface() {
	iface := getIface()
	fmt.Println(iface)
}
