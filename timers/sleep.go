package timers

import (
	"context"
	"time"
)

func SleepContext(ctx context.Context, delay time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(delay):
	}
}
