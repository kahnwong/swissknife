package utils

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	clipboardImage "github.com/skanehira/clipboard-image/v2"
)

func WriteToClipboard(text string) error {
	if err := clipboard.WriteAll(text); err != nil {
		return fmt.Errorf("failed to write to clipboard: %w", err)
	}
	return nil
}

func WriteToClipboardImage(bytes []byte) error {
	tempFilename := "/tmp/qr-image.png"
	if err := os.WriteFile(tempFilename, bytes, 0644); err != nil {
		return fmt.Errorf("failed to write temp image for clipboard: %w", err)
	}

	f, err := os.Open(tempFilename)
	if err != nil {
		return fmt.Errorf("failed to open temp image for clipboard: %w", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			// Log this error but don't override the main return error
			fmt.Fprintf(os.Stderr, "warning: error closing temp image file: %v\n", err)
		}
	}(f)

	if err = clipboardImage.Write(f); err != nil {
		return fmt.Errorf("failed to copy to clipboard: %w", err)
	}
	return nil
}

func ReadFromClipboard() (string, error) {
	text, err := clipboard.ReadAll()
	if err != nil {
		return "", fmt.Errorf("failed to read from clipboard: %w", err)
	}
	return text, nil
}
