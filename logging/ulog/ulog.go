package ulog

import (
	"path/filepath"
	"time"

	"github.com/phuslu/log"
	"github.com/robfig/cron/v3"
	"github.com/umbrella-sh/um-common/ext"
)

type UmbrellaLogger struct {
	Console log.Logger
	File    log.Logger
}

var logger UmbrellaLogger
var isInitialized = false

// Console writes the logs directly to the console (not speed of writing is different on different OSs)
var Console log.Logger

// File writes the logs directly to a log file, either the 'all' file and/or the 'error' file.
var File log.Logger

// New creates an instance of the logger and sets the current instance.
// Taking a 'fileSize' of the log files (all and error).
// The number `fileBackups` to keep for each log file.
// And a bool indicating if the caller+line number, should be included.
//
//goland:noinspection GoUnusedExportedFunction
func New(fileSize int64, fileBackups int, includeCaller bool) {
	if isInitialized {
		logger.Console.Trace().Msg("You have already initialized the logger, it cannot be done again.")
	}

	initLogger(fileSize, fileBackups, "all", "error", "./logs", includeCaller)

	Console = logger.Console
	File = logger.File

	initRotateTask()

	isInitialized = true
}

// NewCustom creates an instance of the logger and sets the current instance.
// Taking a 'fileSize' of the log files (all and error).
// The number `fileBackups` to keep for each log file.
// A name for the 'all' file.
// A name for the 'error' file.
// And the path to the folder to store the logs.
//
//goland:noinspection GoUnusedExportedFunction
func NewCustom(fileSize int64, fileBackups int, allFileName string, errorFileName string, logsFolderPath string) {
	if isInitialized {
		logger.Console.Trace().Msg("You have already initialized the logger, it cannot be done again.")
	}

	initLogger(fileSize, fileBackups, allFileName, errorFileName, logsFolderPath, true)

	Console = logger.Console
	File = logger.File

	initRotateTask()

	isInitialized = true
}

// NewCustomFull creates an instance of the logger and sets the current instance.
// Taking a 'fileSize' of the log files (all and error).
// The number `fileBackups` to keep for each log file.
// A name for the 'all' file.
// A name for the 'error' file.
// The path to the folder to store the logs.
// And a bool indicating if the caller+line number, should be included.
//
//goland:noinspection GoUnusedExportedFunction
func NewCustomFull(fileSize int64, fileBackups int, allFileName string, errorFileName string, logsFolderPath string, includeCaller bool) {
	if isInitialized {
		logger.Console.Trace().Msg("You have already initialized the logger, it cannot be done again.")
	}

	initLogger(fileSize, fileBackups, allFileName, errorFileName, logsFolderPath, includeCaller)

	Console = logger.Console
	File = logger.File

	initRotateTask()

	isInitialized = true
}

func initLogger(fileSize int64, fileBackups int, allFileName string, errorFileName string, logsFolderPath string, includeCaller bool) {
	logger = UmbrellaLogger{
		Console: log.Logger{
			TimeFormat: "2006.01.02 15:04:05.000",
			Caller:     ext.Iif(includeCaller, 1, 0),
			Writer: &log.ConsoleWriter{
				ColorOutput: true,
			},
		},
		File: log.Logger{
			Level: log.TraceLevel,
			Writer: &log.MultiLevelWriter{
				ConsoleWriter: &log.FileWriter{
					Filename:     filepath.Join(logsFolderPath, allFileName),
					MaxSize:      fileSize * 1024 * 1024, //100mb
					MaxBackups:   fileBackups,
					LocalTime:    false,
					EnsureFolder: true,
				},
				ConsoleLevel: log.TraceLevel,
				ErrorWriter: &log.FileWriter{
					Filename:     filepath.Join(logsFolderPath, errorFileName),
					MaxSize:      fileSize * 1024 * 1024, //100mb
					MaxBackups:   fileBackups,
					LocalTime:    false,
					EnsureFolder: true,
				},
			},
		},
	}
}

func initRotateTask() {
	runner := cron.New(cron.WithLocation(time.Local))
	_, err := runner.AddFunc("0 0 * * *", func() {
		_ = logger.File.Writer.(*log.MultiLevelWriter).ConsoleWriter.(*log.FileWriter).Rotate()
		_ = logger.File.Writer.(*log.MultiLevelWriter).ErrorWriter.(*log.FileWriter).Rotate()
	})
	if err != nil {
		panic(err)
	}
	go runner.Run()
}
