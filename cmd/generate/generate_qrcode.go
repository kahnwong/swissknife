package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/kahnwong/swissknife/utils"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

// helpers
func generateQRCode(url string) (string, error) {
	// init
	var q *qrcode.QRCode
	q, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return "", err
	}

	// generate pdf
	png, err := q.PNG(1024)
	if err != nil {
		return "", err
	}

	// copy to clipboard
	utils.WriteToClipboardImage(png)

	// for stdout
	//content := q.ToString(false)
	content := q.ToSmallString(false)

	return content, nil
}

var generateQRCodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "Generate QR code",
	Long:  `Generate QR code from URL (either as an arg or from clipboard) and copy resulting image to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		// set URL
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
		fmt.Println(url)

		// main
		qrcodeStr, err := generateQRCode(url)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(qrcodeStr)
	},
}

func init() {
	Cmd.AddCommand(generateQRCodeCmd)
}
