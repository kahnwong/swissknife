package utils

import (
	"os"

	"github.com/atotto/clipboard"
	"github.com/rs/zerolog/log"
	clipboardImage "github.com/skanehira/clipboard-image"
)

func WriteToClipboard(text string) {
	err := clipboard.WriteAll(text)
	if err != nil {
		log.Error().Msg("Failed to write to clipboard")
	}
}

func WriteToClipboardImage(bytes []byte) {
	tempFilename := "/tmp/qr-image.png"
	err := os.WriteFile(tempFilename, bytes, 0644)
	if err != nil {
		log.Fatal().Msg("Failed to write temp image for clipboard")
	}

	f, err := os.Open(tempFilename)
	if err != nil {
		log.Fatal().Msg("Failed to open temp image for clipboard")
	}
	defer f.Close()

	if err = clipboardImage.CopyToClipboard(f); err != nil {
		log.Fatal().Msg("Failed to copy to clipboard")
	}
}

func ReadFromClipboard() string {
	text, _ := clipboard.ReadAll()
	return text
}
