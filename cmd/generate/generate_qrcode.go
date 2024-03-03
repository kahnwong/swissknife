package generate

import (
	"fmt"
	"os"

	"github.com/kahnwong/swissknife/cmd/utils"

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
	utils.CopyToClipboardImage(png)

	// for stdout
	//content := q.ToString(false)
	content := q.ToSmallString(false)

	return content, nil
}

var generateQRCodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "Generate QR code",
	Long:  `Generate QR code from URL and copy resulting image to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		//init
		if len(args) == 0 {
			fmt.Println("Please specify URL")
			os.Exit(1)
		}

		// main
		qrcodeStr, err := generateQRCode(args[0])
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(qrcodeStr)
	},
}

func init() {
	Cmd.AddCommand(generateQRCodeCmd)
}
