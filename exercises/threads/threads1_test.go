package main

import (
	"testing"
	"time"
)

func TestPrintNumbers(t *testing.T) {
	// Test that PrintNumbers function exists and doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PrintNumbers() should not panic: %v", r)
		}
	}()

	PrintNumbers()
}

func TestPrintLetters(t *testing.T) {
	// Test that PrintLetters function exists and doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PrintLetters() should not panic: %v", r)
		}
	}()

	PrintLetters()
}

func TestCountDown(t *testing.T) {
	// Test that CountDown function exists and doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("CountDown() should not panic: %v", r)
		}
	}()

	start := time.Now()
	CountDown(2) // Should take at least 400ms (2 * 200ms)
	elapsed := time.Since(start)

	// Should take at least 300ms (allowing some tolerance)
	if elapsed < 300*time.Millisecond {
		t.Errorf("CountDown(2) should take at least 300ms, took %v", elapsed)
	}
}
