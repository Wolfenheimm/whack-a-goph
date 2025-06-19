package main

import (
	"context"
	"testing"
	"time"
)

func TestProcessWithTimeout(t *testing.T) {
	// Test successful completion within timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	result, err := ProcessWithTimeout(ctx, "test", 50*time.Millisecond)
	if err != nil {
		t.Errorf("ProcessWithTimeout should complete successfully: %v", err)
	}
	if result == "" {
		t.Errorf("ProcessWithTimeout should return non-empty result")
	}
}

func TestProcessWithTimeoutExpired(t *testing.T) {
	// Test timeout expiration
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err := ProcessWithTimeout(ctx, "test", 100*time.Millisecond)
	if err == nil {
		t.Errorf("ProcessWithTimeout should timeout and return error")
	}
}

func TestProcessWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// Start processing in goroutine
	done := make(chan struct{})
	var result string
	var err error

	go func() {
		result, err = ProcessWithCancel(ctx, "test data")
		close(done)
	}()

	// Cancel after short delay
	time.Sleep(10 * time.Millisecond)
	cancel()

	// Wait for completion
	select {
	case <-done:
		if err == nil {
			t.Errorf("ProcessWithCancel should return error when cancelled")
		}
		// Use result to avoid unused variable warning
		_ = result
	case <-time.After(100 * time.Millisecond):
		t.Errorf("ProcessWithCancel should complete quickly after cancellation")
	}
}

func TestProcessWithValue(t *testing.T) {
	// Test retrieving context value
	ctx := context.WithValue(context.Background(), "userID", "12345")

	result, err := ProcessWithValue(ctx, "userID", nil)
	if err != nil {
		t.Errorf("ProcessWithValue should not return error: %v", err)
	}

	if result != "12345" {
		t.Errorf("ProcessWithValue should return '12345', got %v", result)
	}
}

func TestProcessWithValueNotFound(t *testing.T) {
	// Test missing context value
	ctx := context.Background()

	result, err := ProcessWithValue(ctx, "missingKey", "default")
	if err != nil {
		t.Errorf("ProcessWithValue should handle missing keys gracefully: %v", err)
	}

	// Should return default value or nil
	if result != nil && result != "default" {
		t.Errorf("ProcessWithValue should return nil or default for missing key")
	}
}

func TestProcessChain(t *testing.T) {
	baseCtx := context.Background()

	result, err := ProcessChain(baseCtx)
	if err != nil {
		t.Errorf("ProcessChain should not return error: %v", err)
	}

	if result == "" {
		t.Errorf("ProcessChain should return non-empty result")
	}
}

func TestDemo(t *testing.T) {
	// Test that demo function exists and can be called
	Demo()
	// This test just verifies the function exists
}
