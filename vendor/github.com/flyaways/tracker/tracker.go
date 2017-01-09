package tracker

import "fmt"

var Info = New(Bold, FgMagenta).SprintfFunc()

var Notice = New(Bold, FgGreen).SprintfFunc()

var Warning = New(Bold, FgYellow).SprintfFunc()

var Error = New(Bold, Underline, FgRed).SprintfFunc()

var Attributes = New(Bold, FgCyan).SprintfFunc()

func Tracker(color, format string, a ...interface{}) {
	var result string

	switch color {
	case "Black":
		result = Black(format, a...)
	case "Red":
		result = Red(format, a...)
	case "Green":
		result = Green(format, a...)
	case "Yellow":
		result = Yellow(format, a...)
	case "Blue":
		result = Blue(format, a...)
	case "Magenta":
		result = Magenta(format, a...)
	case "Cyan":
		result = Cyan(format, a...)
	case "White":
		result = White(format, a...)
	case "info":
		result = Info(format, a...)
	case "notice":
		result = Notice(format, a...)
	case "warning":
		result = Warning(format, a...)
	case "error":
		result = Error(format, a...)
	case "attribute":
		result = Attributes(format, a...)
	default:
		result = Error(format, a...)
	}

	fmt.Printf("%s\n", result)
}
