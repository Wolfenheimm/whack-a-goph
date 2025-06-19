package context

import (
	"context"
	"testing"
	"time"
)

func TestDoWithTimeout(t *testing.T) {
	t.Run("completes within timeout", func(t *testing.T) {
		work := func() error {
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		err := DoWithTimeout(50*time.Millisecond, work)
		if err != nil {
			t.Errorf("DoWithTimeout() returned error: %v", err)
		}
	})

	t.Run("times out", func(t *testing.T) {
		work := func() error {
			time.Sleep(100 * time.Millisecond)
			return nil
		}

		err := DoWithTimeout(10*time.Millisecond, work)
		if err == nil {
			t.Error("DoWithTimeout() should have timed out")
		}
	})
}

func TestDoWithCancel(t *testing.T) {
	workStarted := make(chan bool)
	workDone := make(chan bool)

	work := func(ctx context.Context) error {
		workStarted <- true
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(100 * time.Millisecond):
			workDone <- true
			return nil
		}
	}

	cancel, err := DoWithCancel(work)
	if err != nil {
		t.Errorf("DoWithCancel() returned error: %v", err)
		return
	}

	// Wait for work to start
	<-workStarted

	// Cancel the work
	cancel()

	// Work should not complete
	select {
	case <-workDone:
		t.Error("Work completed despite being cancelled")
	case <-time.After(50 * time.Millisecond):
		// Expected - work was cancelled
	}
}

func TestWorkerWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := WorkerWithContext(ctx, 100*time.Millisecond)
	if err == nil {
		t.Error("WorkerWithContext() should have been cancelled")
	}
}

func TestPassAndGetValue(t *testing.T) {
	ctx := context.Background()
	key := "user"
	value := "john"

	ctx = PassValue(ctx, key, value)
	got := GetValue(ctx, key)

	if got != value {
		t.Errorf("GetValue() = %q; want %q", got, value)
	}
}
