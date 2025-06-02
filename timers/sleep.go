package timers

import (
	"context"
	"time"
)

// SleepContext waits for the specified delay, but exits early if the context.Done() is called
// will return true, if it was cancelled with context, else false
func SleepContext(ctx context.Context, delay time.Duration) bool {
	select {
	case <-ctx.Done():
		return true
	case <-time.After(delay):
	}
	return false
}
