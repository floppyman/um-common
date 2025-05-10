package ulog

import (
	"path/filepath"
	"time"

	"github.com/phuslu/log"
	"github.com/robfig/cron/v3"

	"github.com/floppyman/um-common/ext"
)

type UmbrellaLoggerConfig struct {
	FileSize       int64  `default:"10"`
	FileBackups    int    `default:"1"`
	AllFileName    string `default:"all"`
	ErrorFileName  string `default:"error"`
	LogsFolderPath string `default:"./logs"`
	IncludeCaller  bool   `default:"false"`
	LogTimeFormat  string `default:"2006.01.02 15:04:05.000"`
}

type UmbrellaLogger struct {
	Console log.Logger
	File    log.Logger
}

var logger UmbrellaLogger
var isInitialized = false
var config UmbrellaLoggerConfig

// Console writes the logs directly to the console (not speed of writing is different on different OSs)
var Console log.Logger

// File writes the logs directly to a log file, either the 'all' file and/or the 'error' file.
var File log.Logger

// New creates an instance of the logger and sets the current instance config to the provided.
//
//goland:noinspection GoUnusedExportedFunction
func New(conf UmbrellaLoggerConfig) {
	if isInitialized {
		logger.Console.Trace().Msg("You have already initialized the logger, it cannot be done again.")
	}

	config = conf

	initLogger()

	Console = logger.Console
	File = logger.File

	initRotateTask()

	isInitialized = true
}

// NewSimple creates an instance of the logger and sets the current instance.
// Taking a 'fileSize' of the log files (all and error).
// The number `fileBackups` to keep for each log file.
// And a bool indicating if the caller+line number, should be included.
//
//goland:noinspection GoUnusedExportedFunction
func NewSimple(fileSize int64, fileBackups int, includeCaller bool) {
	if isInitialized {
		logger.Console.Trace().Msg("You have already initialized the logger, it cannot be done again.")
	}

	config = UmbrellaLoggerConfig{
		FileSize:      fileSize,
		FileBackups:   fileBackups,
		IncludeCaller: includeCaller,
	}

	initLogger()

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

	config = UmbrellaLoggerConfig{
		FileSize:       fileSize,
		FileBackups:    fileBackups,
		AllFileName:    allFileName,
		ErrorFileName:  errorFileName,
		LogsFolderPath: logsFolderPath,
		IncludeCaller:  true,
	}

	initLogger()

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

	config = UmbrellaLoggerConfig{
		FileSize:       fileSize,
		FileBackups:    fileBackups,
		AllFileName:    allFileName,
		ErrorFileName:  errorFileName,
		LogsFolderPath: logsFolderPath,
		IncludeCaller:  includeCaller,
	}

	initLogger()

	Console = logger.Console
	File = logger.File

	initRotateTask()

	isInitialized = true
}

func initLogger() {
	logger = UmbrellaLogger{
		Console: log.Logger{
			TimeFormat: config.LogTimeFormat,
			Caller:     ext.Iif(config.IncludeCaller, 1, 0),
			Writer: &log.ConsoleWriter{
				ColorOutput: true,
			},
		},
		File: log.Logger{
			TimeFormat: config.LogTimeFormat,
			Level:      log.TraceLevel,
			Writer: &log.MultiLevelWriter{
				ConsoleWriter: &log.FileWriter{
					TimeFormat:   config.LogTimeFormat,
					Filename:     filepath.Join(config.LogsFolderPath, config.AllFileName),
					MaxSize:      config.FileSize * 1024 * 1024, // 100mb
					MaxBackups:   config.FileBackups,
					LocalTime:    true,
					EnsureFolder: true,
				},
				ConsoleLevel: log.TraceLevel,
				ErrorWriter: &log.FileWriter{
					TimeFormat:   config.LogTimeFormat,
					Filename:     filepath.Join(config.LogsFolderPath, config.ErrorFileName),
					MaxSize:      config.FileSize * 1024 * 1024, // 100mb
					MaxBackups:   config.FileBackups,
					LocalTime:    true,
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
