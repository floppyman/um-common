package basic

import (
	"github.com/fatih/color"
)

var (
	printNormal   = color.New(color.FgHiWhite).PrintFunc()
	printLnNormal = color.New(color.FgHiWhite).PrintlnFunc()
	printFNormal  = color.New(color.FgHiWhite).PrintfFunc()

	printTrace   = color.New(color.FgHiBlack).PrintFunc()
	printLnTrace = color.New(color.FgHiBlack).PrintlnFunc()
	printfTrace  = color.New(color.FgHiBlack).PrintfFunc()

	printDebug   = color.New(color.FgWhite).PrintFunc()
	printLnDebug = color.New(color.FgWhite).PrintlnFunc()
	printfDebug  = color.New(color.FgWhite).PrintfFunc()

	printInfo   = color.New(color.FgHiCyan).PrintFunc()
	printLnInfo = color.New(color.FgHiCyan).PrintlnFunc()
	printfInfo  = color.New(color.FgHiCyan).PrintfFunc()

	printWarn   = color.New(color.FgHiYellow).PrintFunc()
	printLnWarn = color.New(color.FgHiYellow).PrintlnFunc()
	printfWarn  = color.New(color.FgHiYellow).PrintfFunc()

	printError   = color.New(color.FgHiRed).PrintFunc()
	printLnError = color.New(color.FgHiRed).PrintlnFunc()
	printfError  = color.New(color.FgHiRed).PrintfFunc()

	printFatal   = color.New(color.BgMagenta, color.FgHiWhite).PrintFunc()
	printLnFatal = color.New(color.BgMagenta, color.FgHiWhite).PrintlnFunc()
	printfFatal  = color.New(color.BgMagenta, color.FgHiWhite).PrintfFunc()
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

func Warn(a ...interface{})                 { printWarn(a...) }
func Warnln(a ...interface{})               { printLnWarn(a...) }
func Warnf(format string, a ...interface{}) { printfWarn(format, a...) }

func Error(a ...interface{})                 { printError(a...) }
func Errorln(a ...interface{})               { printLnError(a...) }
func Errorf(format string, a ...interface{}) { printfError(format, a...) }

func Fatal(a ...interface{})                 { printFatal(a...) }
func Fatalln(a ...interface{})               { printLnFatal(a...) }
func Fatalf(format string, a ...interface{}) { printfFatal(format, a...) }

func TestColors() {
	Println("Normal printing")
	Traceln("Trace")
	Debugln("Debug")
	Infoln("Information")
	Warnln("Warning")
	Errorln("Error")
	Fatalln("Fatal")
}
