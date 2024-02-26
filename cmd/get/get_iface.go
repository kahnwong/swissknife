package get

import (
	"fmt"
	"net"

	netroute "github.com/libp2p/go-netroute"
	"github.com/spf13/cobra"
)

// unix: `netstat -nr -f inet`
// get iface-owned app: `ifconfig -v utun3 | grep "agent domain"`

func getIface() (string, error) {
	r, err := netroute.New()
	if err != nil {
		return "", err
	}

	iface, _, _, err := r.Route(
		net.IPv4(104, 16, 133, 229), // cloudflare
	)
	if err != nil {
		return "", err
	}

	//fmt.Printf("%v, %v, %v, %v\n", iface, gw, src, err)

	return iface.Name, nil
}

var getIfaceCmd = &cobra.Command{
	Use:   "iface",
	Short: "Get iface",
	Long:  `Get iface used for public internet access`,
	Run: func(cmd *cobra.Command, args []string) {
		iface, err := getIface()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(iface)
	},
}

func init() {
	Cmd.AddCommand(getIfaceCmd)
}
