package basic

import (
	"github.com/fatih/color"
)

var (
	normalColor = color.New(color.FgHiWhite)
	traceColor  = color.New(color.FgWhite)
	debugColor  = color.New(color.FgHiBlack)
	infoColor   = color.New(color.FgHiCyan)
	warnColor   = color.New(color.FgHiYellow)
	errorColor  = color.New(color.FgHiRed)
	fatalColor  = color.New(color.FgHiMagenta)
)

func Println(text string) {
	_, _ = normalColor.Println(text)
}

func Trace(text string) {
	_, _ = traceColor.Println(text)
}

func Debug(text string) {
	_, _ = debugColor.Println(text)
}

func Information(text string) {
	_, _ = infoColor.Println(text)
}

func Warning(text string) {
	_, _ = warnColor.Println(text)
}

func Error(text string) {
	_, _ = errorColor.Println(text)
}

func Fatal(text string) {
	_, _ = fatalColor.Println(text)
}
