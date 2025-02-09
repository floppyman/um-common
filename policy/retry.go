package policy

import (
	"context"
	"time"
)

type Effector[TOutput any] func(context.Context) (TOutput, error)
type LogCallback func(time.Duration, int, error)

// Retry calls 'effector' method the number of 'retries' waiting the 'delay' in between, will return the response of 'effector' if it was success, to when the 'retries' has been reached
//
//goland:noinspection GoUnusedExportedFunction
func Retry[TOutput any](effector Effector[TOutput], failVal TOutput, retries int, delay time.Duration) Effector[TOutput] {
	return func(ctx context.Context) (TOutput, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				// Return when there is no error or the maximum number of retries is reached.
				return response, err
			}

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return failVal, ctx.Err()
			}
		}
	}
}

// RetryWithLog calls 'effector' function the number of 'retries' waiting the 'delay' in between, will return the response of 'effector' if it was success, to when the 'retries' has been reached.
// For each try, the 'callbackLog' function is called with the delay, current retryCount and the error
//
//goland:noinspection GoUnusedExportedFunction
func RetryWithLog[TOutput any](effector Effector[TOutput], failVal TOutput, retries int, delay time.Duration, callbackLog LogCallback) Effector[TOutput] {
	return func(ctx context.Context) (TOutput, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				// Return when there is no error or the maximum number of retries is reached.
				return response, err
			}

			callbackLog(delay, r, err)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return failVal, ctx.Err()
			}
		}
	}
}
