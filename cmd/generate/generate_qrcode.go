package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/kahnwong/swissknife/utils"
	"github.com/rs/zerolog/log"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

func setURL(args []string) string {
	var url string
	if len(args) == 0 {
		urlFromClipboard := utils.ReadFromClipboard()
		if urlFromClipboard != "" {
			if strings.HasPrefix(urlFromClipboard, "https://") {
				url = urlFromClipboard
			}
		}
	}
	if url == "" {
		if len(args) == 0 {
			fmt.Println("Please specify URL")
			os.Exit(1)
		} else if len(args) == 1 {
			url = args[0]
		}
	}

	return url
}

func generateQRCode(url string) ([]byte, string) {
	// init
	var q *qrcode.QRCode
	q, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		log.Fatal().Msg("Failed to initialize QRCode object")
	}

	// generate png
	png, err := q.PNG(1024)
	if err != nil {
		log.Fatal().Msg("Failed to generate QRCode PNG")
	}

	// for stdout
	//stdout := q.ToString(false)
	stdout := q.ToSmallString(false)

	return png, stdout
}

var generateQRCodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "Generate QR code",
	Long:  `Generate QR code from URL (either as an arg or from clipboard) and copy resulting image to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		// set URL
		url := setURL(args)
		fmt.Println(url)

		// main
		png, stdout := generateQRCode(url)
		utils.WriteToClipboardImage(png)
		fmt.Println(stdout)
	},
}

func init() {
	Cmd.AddCommand(generateQRCodeCmd)
}
