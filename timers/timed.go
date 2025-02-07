package timers

import (
	"time"

	"github.com/umbrella-sh/um-common/logging/ulog"
)

type wrapped func()
type wrappedReturn[Result any] func() Result
type wrappedReturnErr[Result any] func() (Result, error)

const timedFormat string = "%s - Time %-20v %s\n"
const timedFormatTotal string = "%s - Time %-20v Total\n"

var DoTimedLogging = false

//goland:noinspection GoUnusedExportedFunction
func TimedTotal(method string, ds ...time.Duration) {
	var elapsed time.Duration
	for _, i := range ds {
		elapsed = elapsed + i
	}
	//goland:noinspection GoBoolExpressions
	if DoTimedLogging {
		ulog.Console.Trace().Msgf(timedFormatTotal, method, elapsed)
	}
}

//goland:noinspection GoUnusedExportedFunction
func Timed(fn wrapped, method string, key string) time.Duration {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	//goland:noinspection GoBoolExpressions
	if DoTimedLogging {
		ulog.Console.Trace().Msgf(timedFormat, method, elapsed, key)
	}
	return elapsed
}

//goland:noinspection GoUnusedExportedFunction
func TimedRes[Result any](fn wrappedReturn[Result], method string, key string) (time.Duration, Result) {
	start := time.Now()
	res := fn()
	elapsed := time.Since(start)
	//goland:noinspection GoBoolExpressions
	if DoTimedLogging {
		ulog.Console.Trace().Msgf(timedFormat, method, elapsed, key)
	}
	return elapsed, res
}

//goland:noinspection GoUnusedExportedFunction
func TimedResErr[Result any](fn wrappedReturnErr[Result], method string, key string) (time.Duration, Result, error) {
	start := time.Now()
	res, err := fn()
	elapsed := time.Since(start)
	//goland:noinspection GoBoolExpressions
	if DoTimedLogging {
		ulog.Console.Trace().Msgf(timedFormat, method, elapsed, key)
	}
	return elapsed, res, err
}
