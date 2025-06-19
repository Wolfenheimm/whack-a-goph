package main

// Threads 1: Basic Goroutines 🐹
// Learn about goroutines and concurrent programming in Go.

// I AM NOT DONE

import (
	"fmt"
	"time"
)

// TODO: Implement PrintNumbers function
// Print numbers 1-5 with 100ms delay between each
func PrintNumbers() {
	// Fix me! Loop from 1 to 5, print each number with time.Sleep(100 * time.Millisecond)
}

// TODO: Implement PrintLetters function
// Print letters A-E with 150ms delay between each
func PrintLetters() {
	// Fix me! Loop through letters A-E, print each with time.Sleep(150 * time.Millisecond)
}

// TODO: Implement CountDown function
// Count down from given number to 1
func CountDown(from int) {
	// Fix me! Count down and print each number
	for i := from; i >= 1; i-- {
		fmt.Printf("Countdown: %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Starting concurrent execution...")

	// TODO: Run PrintNumbers as a goroutine
	// Fix me! Use 'go' keyword
	PrintNumbers()

	// TODO: Run PrintLetters as a goroutine
	// Fix me! Use 'go' keyword
	PrintLetters()

	// TODO: Run CountDown as a goroutine
	// Fix me! Use 'go' keyword
	CountDown(3)

	// Wait for goroutines to complete
	time.Sleep(2 * time.Second)
	fmt.Println("All done!")
}
