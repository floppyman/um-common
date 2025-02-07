package policy

import (
	"context"
	"time"
)

type Effector[TOutput any] func(context.Context) (TOutput, error)
type LogCallback func(time.Duration, int, error)

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
