package utils

import (
	"fmt"
	"os"
	"strings"
)

func SetURL(args []string) string {
	var url string
	if len(args) == 0 {
		urlFromClipboard := ReadFromClipboard()
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
