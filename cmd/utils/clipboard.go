package utils

import (
	"golang.design/x/clipboard"
)

func WriteToClipboard(text string) {
	err := clipboard.Init() // clipboard doesn't work from ssh session
	if err == nil {         // clipboard doesn't work from ssh session
		clipboard.Write(clipboard.FmtText, []byte(text))
	}
}

func WriteToClipboardImage(bytes []byte) {
	err := clipboard.Init()
	if err == nil {
		clipboard.Write(clipboard.FmtImage, bytes)
	}
}

func ReadFromClipboard() string {
	err := clipboard.Init()
	if err == nil {

		s := clipboard.Read(clipboard.FmtText)
		return string(s)
	}

	return ""
}
