package configs

import "github.com/fatih/color"

type ColorsStruct struct {
	Green   func(a ...interface{}) string
	Magenta func(a ...interface{}) string
	Red     func(a ...interface{}) string
	Yellow  func(a ...interface{}) string
	Blue    func(a ...interface{}) string
}

var Colors = ColorsStruct{
	Green:   color.New(color.FgHiGreen).SprintFunc(),
	Magenta: color.New(color.FgHiMagenta).SprintFunc(),
	Red:     color.New(color.FgRed).SprintFunc(),
	Yellow:  color.New(color.FgYellow).SprintFunc(),
	Blue:    color.New(color.FgBlue).SprintFunc(),
}
