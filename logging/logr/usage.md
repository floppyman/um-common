# Logr usage

## Initialization

In the beginning of your project, use either of these two methods to setup the logger ready for use.

**New**
```go
// miniaml logging instance, all default
logr.NewMinimal()

// logging instance, configurable
logr.New(LogrConfig{})
```

## Usage

Now you can everywhere you need to log something use methods like this.

**Basic logs**
```go
// To Console
// ----------
logr.Console.Trace().Msg("test")
logr.Console.Debug().Msg("test")
logr.Console.Info().Msg("test")
logr.Console.Warn().Msg("test")
logr.Console.Error().Msg("test")

// Causes application to exit immediately with code 255, after writing message
logr.Console.Fatal().Msg("test")

// Causes application to exit immediately with a panic, after writing message
logr.Console.Panic().Msg("test")

// To File
// ----------
logr.File.Trace().Msg("test")
logr.File.Debug().Msg("test")
logr.File.Info().Msg("test")
logr.File.Warn().Msg("test")
logr.File.Error().Msg("test")

// Causes application to exit immediately with code 255, after writing message
logr.File.Fatal().Msg("test")

// Causes application to exit immediately with a panic, after writing message
logr.File.Panic().Msg("test")
```

**Advanced logs**
```go
// This logs to console at level info, writing the messaage 'Total processing time' and has an extra 'argument" named 'sec' with the value of a statement
logr.Console.Info().Float64("sec", time.Since(start).Seconds()).Msg("Total processing time")

// This logs to file at level error, writing the message 'Error happened' and adds the error and stacktrace from the 'err' variable
logr.File.Error().Err(err).Msg("Error happened")
```
