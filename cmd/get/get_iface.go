package get

import (
	"fmt"
	"net"

	netroute "github.com/libp2p/go-netroute"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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

var getIfaceCmd = &cobra.Command{
	Use:   "iface",
	Short: "Get iface",
	Long:  `Get iface used for public internet access`,
	Run: func(cmd *cobra.Command, args []string) {
		iface := getIface()
		fmt.Println(iface)
	},
}

func init() {
	Cmd.AddCommand(getIfaceCmd)
}
