package main

// Enums 1: Constants as Enums 🐹
// Go doesn't have traditional enums, but we can use constants with iota
// to create enum-like behavior for type safety and readability.

// I AM NOT DONE

import "fmt"

// TODO: Define a custom type called DayOfWeek based on int
type DayOfWeek int // Fix me!

// TODO: Define constants for days of the week using iota
// Start with Sunday = 0, Monday = 1, etc.
const (
	Sunday    DayOfWeek = 0 // Fix me! Use iota
	Monday    DayOfWeek = 0 // Fix me!
	Tuesday   DayOfWeek = 0 // Fix me!
	Wednesday DayOfWeek = 0 // Fix me!
	Thursday  DayOfWeek = 0 // Fix me!
	Friday    DayOfWeek = 0 // Fix me!
	Saturday  DayOfWeek = 0 // Fix me!
)

// TODO: Define a Status type for task status
type Status int // Fix me!

// TODO: Define status constants using iota, starting from 1
const (
	Pending    Status = 0 // Fix me! Start from 1 using iota
	InProgress Status = 0 // Fix me!
	Completed  Status = 0 // Fix me!
	Cancelled  Status = 0 // Fix me!
)

// TODO: Implement IsWeekend function
// Return true if the day is Saturday or Sunday
func IsWeekend(day DayOfWeek) bool {
	// Fix me!
	return false
}

// TODO: Implement GetNextDay function
// Return the next day of the week (Sunday comes after Saturday)
func GetNextDay(day DayOfWeek) DayOfWeek {
	// Fix me!
	return Sunday
}

// TODO: Implement IsValidStatus function
// Return true if status is between Pending and Cancelled (inclusive)
func IsValidStatus(status Status) bool {
	// Fix me!
	return false
}

// TODO: Implement GetStatusMessage function
// Return appropriate message for each status
func GetStatusMessage(status Status) string {
	// Fix me! Use a switch statement
	return ""
}

func main() {
	// Test day operations
	fmt.Printf("Is Friday a weekend? %v\n", IsWeekend(Friday))
	fmt.Printf("Day after Saturday: %v\n", GetNextDay(Saturday))

	// Test status operations
	fmt.Printf("Is status 2 valid? %v\n", IsValidStatus(InProgress))
	fmt.Printf("Status message for Completed: %s\n", GetStatusMessage(Completed))
}
