package basic

import (
	"github.com/fatih/color"
)

var (
	printNormal   = color.RGB(255, 255, 255).PrintFunc()
	printLnNormal = color.RGB(255, 255, 255).PrintlnFunc()
	printFNormal  = color.RGB(255, 255, 255).PrintfFunc()

	printTrace   = color.RGB(92, 92, 92).PrintFunc()
	printLnTrace = color.RGB(92, 92, 92).PrintlnFunc()
	printfTrace  = color.RGB(92, 92, 92).PrintfFunc()

	printGui = color.RGB(116, 116, 116).PrintFunc()

	printDebug   = color.RGB(160, 160, 160).PrintFunc()
	printLnDebug = color.RGB(160, 160, 160).PrintlnFunc()
	printfDebug  = color.RGB(160, 160, 160).PrintfFunc()

	printInfo   = color.RGB(0, 255, 255).PrintFunc()
	printLnInfo = color.RGB(0, 255, 255).PrintlnFunc()
	printfInfo  = color.RGB(0, 255, 255).PrintfFunc()

	printWarn   = color.RGB(255, 255, 0).PrintFunc()
	printLnWarn = color.RGB(255, 255, 0).PrintlnFunc()
	printfWarn  = color.RGB(255, 255, 0).PrintfFunc()

	printError   = color.RGB(255, 0, 0).PrintFunc()
	printLnError = color.RGB(255, 0, 0).PrintlnFunc()
	printfError  = color.RGB(255, 0, 0).PrintfFunc()

	printFatal   = color.BgRGB(160, 0, 160).AddRGB(255, 255, 255).PrintFunc()
	printLnFatal = color.BgRGB(160, 0, 160).AddRGB(255, 255, 255).PrintlnFunc()
	printfFatal  = color.BgRGB(160, 0, 160).AddRGB(255, 255, 255).PrintfFunc()

	printSuccess   = color.RGB(0, 255, 0).PrintFunc()
	printLnSuccess = color.RGB(0, 255, 0).PrintlnFunc()
	printfSuccess  = color.RGB(0, 255, 0).PrintfFunc()
)

func Print(a ...interface{})                 { printNormal(a...) }
func Println(a ...interface{})               { printLnNormal(a...) }
func Printf(format string, a ...interface{}) { printFNormal(format, a...) }

func Trace(a ...interface{})                 { printTrace(a...) }
func Traceln(a ...interface{})               { printLnTrace(a...) }
func Tracef(format string, a ...interface{}) { printfTrace(format, a...) }

func Debug(a ...interface{})                 { printDebug(a...) }
func Debugln(a ...interface{})               { printLnDebug(a...) }
func Debugf(format string, a ...interface{}) { printfDebug(format, a...) }

func Info(a ...interface{})                 { printInfo(a...) }
func Infoln(a ...interface{})               { printLnInfo(a...) }
func Infof(format string, a ...interface{}) { printfInfo(format, a...) }
func infoPrintHeader() {
	printGui("[")
	printInfo(" info ")
	printGui("] ")
}
func InfoPrint(a ...interface{}) {
	infoPrintHeader()
	printLnNormal(a...)
}
func InfoPrintf(format string, a ...interface{}) {
	infoPrintHeader()
	printFNormal(format, a...)
}

func Warn(a ...interface{})                 { printWarn(a...) }
func Warnln(a ...interface{})               { printLnWarn(a...) }
func Warnf(format string, a ...interface{}) { printfWarn(format, a...) }
func warnPrintHeader() {
	printGui("[")
	printWarn(" warn ")
	printGui("] ")
}
func WarnPrint(a ...interface{}) {
	warnPrintHeader()
	printLnNormal(a...)
}
func WarnPrintf(format string, a ...interface{}) {
	warnPrintHeader()
	printFNormal(format, a...)
}

func Error(a ...interface{})                 { printError(a...) }
func Errorln(a ...interface{})               { printLnError(a...) }
func Errorf(format string, a ...interface{}) { printfError(format, a...) }

func Fatal(a ...interface{})                 { printFatal(a...) }
func Fatalln(a ...interface{})               { printLnFatal(a...) }
func Fatalf(format string, a ...interface{}) { printfFatal(format, a...) }

func Success(a ...interface{})                 { printSuccess(a...) }
func Successln(a ...interface{})               { printLnSuccess(a...) }
func Successf(format string, a ...interface{}) { printfSuccess(format, a...) }
func successPrintHeader() {
	printGui("[")
	printSuccess("  ok  ")
	printGui("] ")
}
func SuccessPrint(a ...interface{}) {
	successPrintHeader()
	printLnNormal(a...)
}
func SuccessPrintf(format string, a ...interface{}) {
	successPrintHeader()
	printFNormal(format, a...)
}

func blankPrintHeader() {
	printGui("[")
	printInfo("      ")
	printGui("] ")
}
func BlankPrint(a ...interface{}) {
	blankPrintHeader()
	printLnNormal(a...)
}
func BlankPrintf(format string, a ...interface{}) {
	blankPrintHeader()
	printFNormal(format, a...)
}

func failPrintHeader() {
	printGui("[")
	printError("failed")
	printGui("] ")
}
func FailPrint(a ...interface{}) {
	failPrintHeader()
	printLnNormal(a...)
}
func FailPrintf(format string, a ...interface{}) {
	failPrintHeader()
	printFNormal(format, a...)
}

func waitPrintHeader() {
	printGui("[")
	printInfo(" wait ")
	printGui("] ")
}
func WaitPrint(a ...interface{}) {
	waitPrintHeader()
	printLnNormal(a...)
}
func WaitPrintf(format string, a ...interface{}) {
	waitPrintHeader()
	printFNormal(format, a...)
}

func TestColors() {
	Traceln("Trace")
	Debugln("Debug")
	Infoln("Information")
	Warnln("Warning")
	Errorln("Error")
	Fatalln("Fatal")
	Println("Normal printing")
}
