package timers

import (
	"time"
)

type wrapped func()
type wrappedReturn[Result any] func() Result
type wrappedReturnErr[Result any] func() (Result, error)

//goland:noinspection GoUnusedConst
const (
	// TimedFormat format for writing time.Duration to Console or File
	TimedFormat string = "%s - Time %-20v %s\n"
	
	// TimedFormatTotal format for writing total time.Duration to Console or File
	TimedFormatTotal string = "%s - Time %-20v Total\n"
)

// TimedTotal takes X number of time.Durations adds them together and return the total
//
//goland:noinspection GoUnusedExportedFunction
func TimedTotal(ds ...time.Duration) time.Duration {
	var elapsed time.Duration
	for _, i := range ds {
		elapsed = elapsed + i
	}
	return elapsed
}

// Timed executes the 'fn' and times how long it takes to executes, returns the duration
//
//goland:noinspection GoUnusedExportedFunction
func Timed(fn wrapped) time.Duration {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	return elapsed
}

// TimedRes executes the 'fn' and times how long it takes to executes, returns the duration + 'fn' result
//
//goland:noinspection GoUnusedExportedFunction
func TimedRes[Result any](fn wrappedReturn[Result]) (time.Duration, Result) {
	start := time.Now()
	res := fn()
	elapsed := time.Since(start)
	return elapsed, res
}

// TimedResErr executes the 'fn' and times how long it takes to executes, returns the duration + 'fn' result + possible error
//
//goland:noinspection GoUnusedExportedFunction
func TimedResErr[Result any](fn wrappedReturnErr[Result]) (time.Duration, Result, error) {
	start := time.Now()
	res, err := fn()
	elapsed := time.Since(start)
	return elapsed, res, err
}
