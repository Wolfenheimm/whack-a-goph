package main

// Modules 1: Package Declarations and Imports 🐹
// Learn about Go's package system and how to organize code.

// I AM NOT DONE

import (
	"fmt"
	// TODO: Import the math package
	// TODO: Import the strings package
	// TODO: Import the time package
)

// TODO: Implement CalculateCircleArea function
// Use math.Pi constant to calculate area of circle with given radius
func CalculateCircleArea(radius float64) float64 {
	// Fix me! Use math.Pi
	return 0.0
}

// TODO: Implement FormatUserName function
// Use strings package to convert name to title case
func FormatUserName(name string) string {
	// Fix me! Use strings.Title or strings.ToTitle
	return ""
}

// TODO: Implement GetCurrentYear function
// Use time package to get current year
func GetCurrentYear() int {
	// Fix me! Use time.Now().Year()
	return 0
}

// TODO: Implement IsWorkingHour function
// Return true if hour is between 9 and 17 (inclusive)
func IsWorkingHour(hour int) bool {
	// Fix me!
	return false
}

// TODO: Implement ConvertToUpperCase function
// Use strings package to convert string to uppercase
func ConvertToUpperCase(s string) string {
	// Fix me! Use strings.ToUpper
	return ""
}

func main() {
	// Test circle area calculation
	radius := 5.0
	area := CalculateCircleArea(radius)
	fmt.Printf("Circle area with radius %.1f: %.2f\n", radius, area)

	// Test name formatting
	name := "john doe"
	formatted := FormatUserName(name)
	fmt.Printf("Formatted name: %s\n", formatted)

	// Test year function
	year := GetCurrentYear()
	fmt.Printf("Current year: %d\n", year)

	// Test working hours
	fmt.Printf("Is 14 a working hour? %v\n", IsWorkingHour(14))
	fmt.Printf("Is 20 a working hour? %v\n", IsWorkingHour(20))

	// Test uppercase conversion
	text := "hello world"
	upper := ConvertToUpperCase(text)
	fmt.Printf("Uppercase: %s\n", upper)
}
