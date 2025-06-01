package logr

import (
	"path/filepath"
	"time"

	"github.com/phuslu/log"
	"github.com/robfig/cron/v3"

	"github.com/floppyman/um-common/ext"
)

//goland:noinspection GoNameStartsWithPackageName
type LogrConfig struct {
	FileSize       int64
	FileBackups    int
	AllFileName    string
	ErrorFileName  string
	LogsFolderPath string
	IncludeCaller  bool
	LogTimeFormat  string
}

//goland:noinspection GoNameStartsWithPackageName
type Logr struct {
	Console log.Logger
	File    log.Logger
}

var logger Logr
var isInitialized = false
var config LogrConfig

// Console writes the logs directly to the console (not speed of writing is different on different OSs)
var Console log.Logger

// File writes the logs directly to a log file, either the 'all' file and/or the 'error' file.
var File log.Logger

// NewMinimal creates an instance of the logger with all defaults.
//
//goland:noinspection GoUnusedExportedFunction
func NewMinimal() {
	New(LogrConfig{
		FileSize:       10,
		FileBackups:    2,
		AllFileName:    "all",
		ErrorFileName:  "error",
		LogsFolderPath: "./logs",
		IncludeCaller:  false,
		LogTimeFormat:  "2006.01.02 15:04:05.000",
	})
}

// New creates an instance of the logger and sets the current instance config to the provided.
//
//goland:noinspection GoUnusedExportedFunction
func New(conf LogrConfig) {
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

func initLogger() {
	logger = Logr{
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
