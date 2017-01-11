package tracker

func Black(format string, a ...interface{}) string { return printString(format, FgBlack, a...) }

func Red(format string, a ...interface{}) string { return printString(format, FgRed, a...) }

func Green(format string, a ...interface{}) string { return printString(format, FgGreen, a...) }

func Yellow(format string, a ...interface{}) string { return printString(format, FgYellow, a...) }

func Blue(format string, a ...interface{}) string { return printString(format, FgBlue, a...) }

func Magenta(format string, a ...interface{}) string { return printString(format, FgMagenta, a...) }

func Cyan(format string, a ...interface{}) string { return printString(format, FgCyan, a...) }

func White(format string, a ...interface{}) string { return printString(format, FgWhite, a...) }

func Info(format string, a ...interface{}) string { return printBoldString(format, FgMagenta, a...) }

func Notice(format string, a ...interface{}) string { return printBoldString(format, FgGreen, a...) }

func Warning(format string, a ...interface{}) string { return printBoldString(format, FgYellow, a...) }

func Error(format string, a ...interface{}) string { return printBoldString(format, FgRed, a...) }
