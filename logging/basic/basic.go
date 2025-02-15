package basic

import (
	"github.com/fatih/color"
)

var (
	printNormal      = color.New(color.FgHiWhite).PrintlnFunc()
	printTrace       = color.New(color.FgHiBlack).PrintlnFunc()
	printDebug       = color.New(color.FgWhite).PrintlnFunc()
	printInformation = color.New(color.FgHiCyan).PrintlnFunc()
	printWarning     = color.New(color.FgHiYellow).PrintlnFunc()
	printError       = color.New(color.FgHiRed).Add(color.Bold).PrintlnFunc()
	printFatal       = color.New(color.FgHiMagenta).Add(color.Bold).PrintlnFunc()
)

//goland:noinspection GoUnusedExportedFunction
func Println(text string) {
	printNormal(text)
}

//goland:noinspection GoUnusedExportedFunction
func Trace(text string) {
	printTrace(text)
}

//goland:noinspection GoUnusedExportedFunction
func Debug(text string) {
	printDebug(text)
}

//goland:noinspection GoUnusedExportedFunction
func Information(text string) {
	printInformation(text)
}

//goland:noinspection GoUnusedExportedFunction
func Warning(text string) {
	printWarning(text)
}

//goland:noinspection GoUnusedExportedFunction
func Error(text string) {
	printError(text)
}

//goland:noinspection GoUnusedExportedFunction
func Fatal(text string) {
	printFatal(text)
}
