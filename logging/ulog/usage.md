# ulog usage

## Initialization

In the beginning of your project, use either of these two methods to setup the logger ready for use.

**New**
```go
// First argument is the maximum size of the log, in megabytes
// Second argument is the number of log files to keep
ulog.New(100, 10)
```

**NewCustom**
```go
// First argument is the maximum size of the log, in megabytes
// Second argument is the number of log files to keep
// Third argument is the name of the 'all' log file, where every log made to file with be added
// Fourth argument is the name of the 'error' log file, where every error is added separate from the other log file
// Fifth argument is the folder and path to where to store the log files
ulog.NewCustom(100, 10, "all", "error", "./logs")
```

## Usage

Now you can everywhere you need to log something use methods like this.

**Basic logs**
```go
// To Console
// ----------
ulog.Console.Trace().Msg("test")
ulog.Console.Debug().Msg("test")
ulog.Console.Info().Msg("test")
ulog.Console.Warn().Msg("test")
ulog.Console.Error().Msg("test")

// Causes application to exit immediately with code 255, after writing message
ulog.Console.Fatal().Msg("test")

// Causes application to exit immediately with a panic, after writing message
ulog.Console.Panic().Msg("test")

// To File
// ----------
ulog.File.Trace().Msg("test")
ulog.File.Debug().Msg("test")
ulog.File.Info().Msg("test")
ulog.File.Warn().Msg("test")
ulog.File.Error().Msg("test")

// Causes application to exit immediately with code 255, after writing message
ulog.File.Fatal().Msg("test")

// Causes application to exit immediately with a panic, after writing message
ulog.File.Panic().Msg("test")
```

**Advanced logs**
```go
// This logs to console at level info, writing the messaage 'Total processing time' and has an extra 'argument" named 'sec' with the value of a statement
ulog.Console.Info().Float64("sec", time.Since(start).Seconds()).Msg("Total processing time")

// This logs to file at level error, writing the message 'Error happened' and adds the error and stacktrace from the 'err' variable
ulog.File.Error().Err(err).Msg("Error happened")
```
