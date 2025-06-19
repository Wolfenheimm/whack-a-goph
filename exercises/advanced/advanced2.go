package main

// Advanced 2: Context and Cancellation 🐹
// Learn about Go's context package for managing cancellation,
// timeouts, and request-scoped values.

// I AM NOT DONE

import (
	"context"
	"fmt"
	"time"
)

// TODO: Implement ProcessWithTimeout function
// Process data with a timeout using context
func ProcessWithTimeout(ctx context.Context, data string, duration time.Duration) (string, error) {
	// Fix me! Use context.WithTimeout and select statement
	return "", fmt.Errorf("not implemented")
}

// TODO: Implement ProcessWithCancel function
// Process data that can be cancelled via context
func ProcessWithCancel(ctx context.Context, data string) (string, error) {
	// Fix me! Use context cancellation and select statement
	return "", fmt.Errorf("not implemented")
}

// TODO: Implement ProcessWithValue function
// Process data using context values
func ProcessWithValue(ctx context.Context, key, value interface{}) (interface{}, error) {
	// Fix me! Use context.WithValue and ctx.Value()
	return nil, fmt.Errorf("not implemented")
}

// TODO: Implement ProcessChain function
// Chain multiple context operations together
func ProcessChain(baseCtx context.Context) (string, error) {
	// Fix me! Create context chain with timeout and cancellation
	return "", fmt.Errorf("not implemented")
}

// TODO: Implement Demo function to test context functionality
func Demo() {
	// Fix me! Implement demo functionality here
	fmt.Println("Context demo not implemented")
}
