package context

import (
	"context"
	"time"
)

// TODO: Implement the following context-related functions:
// 1. DoWithTimeout - perform an operation with a timeout
// 2. DoWithCancel - perform an operation that can be cancelled
// 3. WorkerWithContext - simulate work that respects context cancellation
// 4. PassValue - pass and retrieve a value through context

func DoWithTimeout(duration time.Duration, work func() error) error {
	// Fix me!
	return nil
}

func DoWithCancel(work func(ctx context.Context) error) (func(), error) {
	// Fix me! Return cancel function and error
	return nil, nil
}

func WorkerWithContext(ctx context.Context, workDuration time.Duration) error {
	// Fix me! Simulate work that can be cancelled
	return nil
}

func PassValue(ctx context.Context, key, value string) context.Context {
	// Fix me!
	return nil
}

func GetValue(ctx context.Context, key string) string {
	// Fix me!
	return ""
}
