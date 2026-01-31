package generate

import (
	"fmt"

	"github.com/kahnwong/swissknife/internal/utils"
	qrcode "github.com/skip2/go-qrcode"
)

func generateQRCode(url string) ([]byte, string, error) {
	// init
	var q *qrcode.QRCode
	q, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return nil, "", fmt.Errorf("failed to initialize QRCode object: %w", err)
	}

	// generate png
	png, err := q.PNG(1024)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate QRCode PNG: %w", err)
	}

	// for stdout
	//stdout := q.ToString(false)
	stdout := q.ToSmallString(false)

	return png, stdout, nil
}

func QRCode(args []string) error {
	// set URL
	url := utils.SetURL(args)
	fmt.Println(url)

	// main
	png, stdout, err := generateQRCode(url)
	if err != nil {
		return err
	}

	if err = utils.WriteToClipboardImage(png); err != nil {
		return fmt.Errorf("failed to write to clipboard: %w", err)
	}

	fmt.Println(stdout)
	return nil
}
