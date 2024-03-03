package utils

import (
	"golang.design/x/clipboard"
)

func CopyToClipboard(text string) {
	err := clipboard.Init() // clipboard doesn't work from ssh session
	if err == nil {         // clipboard doesn't work from ssh session
		clipboard.Write(clipboard.FmtText, []byte(text))
	}
}

func CopyToClipboardImage(bytes []byte) {
	err := clipboard.Init()
	if err == nil {
		clipboard.Write(clipboard.FmtImage, bytes)
	}
}
