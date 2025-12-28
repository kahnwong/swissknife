package generate

import (
	"fmt"

	"github.com/kahnwong/swissknife/internal/utils"
	"github.com/rs/zerolog/log"
	qrcode "github.com/skip2/go-qrcode"
)

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

func QRCode(args []string) {
	// set URL
	url := utils.SetURL(args)
	fmt.Println(url)

	// main
	png, stdout := generateQRCode(url)
	utils.WriteToClipboardImage(png)
	fmt.Println(stdout)
}
