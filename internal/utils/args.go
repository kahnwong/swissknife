package utils

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func setValueFromArgsOrClipboard(args []string, validator func(string) bool, errorMsg string, allowEmpty bool) string {
	var value string
	if len(args) == 0 {
		clipboardValue, err := ReadFromClipboard()
		if err == nil && clipboardValue != "" && validator(clipboardValue) {
			value = strings.TrimSpace(clipboardValue)
		}
	}
	if value == "" {
		if len(args) == 0 {
			if allowEmpty {
				return ""
			}
			fmt.Println(errorMsg)
			os.Exit(1)
		} else if len(args) == 1 {
			value = args[0]
		}
	}
	return value
}

func SetURL(args []string) string {
	return setValueFromArgsOrClipboard(args,
		func(s string) bool { return strings.HasPrefix(s, "https://") },
		"Please specify URL", false)
}

func SetIP(args []string) string {
	return setValueFromArgsOrClipboard(args,
		func(s string) bool { return net.ParseIP(strings.TrimSpace(s)) != nil },
		"", true)
}
