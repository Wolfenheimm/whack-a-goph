package main

// Enums 2: Custom Types for Type Safety 🐹
// Learn how to create strongly-typed enums that prevent
// mixing up different types of constants.

// I AM NOT DONE

import "fmt"

// TODO: Define a Temperature type based on float64
type Temperature float64 // Fix me!

// TODO: Define temperature constants for absolute zero and boiling points
const (
	AbsoluteZeroC Temperature = 0.0 // Fix me! Set to -273.15
	FreezingC     Temperature = 0.0 // Fix me! Set to 0.0
	BoilingC      Temperature = 0.0 // Fix me! Set to 100.0
)

// TODO: Define a Color type based on string
type Color string // Fix me!

// TODO: Define color constants
const (
	Red    Color = "" // Fix me! Set to "red"
	Green  Color = "" // Fix me! Set to "green"
	Blue   Color = "" // Fix me! Set to "blue"
	Yellow Color = "" // Fix me! Set to "yellow"
)

// TODO: Define a Direction type based on int
type Direction int // Fix me!

// TODO: Define direction constants using iota
const (
	North Direction = 0 // Fix me! Use iota
	East  Direction = 0 // Fix me!
	South Direction = 0 // Fix me!
	West  Direction = 0 // Fix me!
)

// TODO: Implement method ToFahrenheit on Temperature type
// Formula: F = C * 9/5 + 32
func (t Temperature) ToFahrenheit() Temperature {
	// Fix me!
	return 0.0
}

// TODO: Implement method IsFreezingOrBelow on Temperature type
func (t Temperature) IsFreezingOrBelow() bool {
	// Fix me!
	return false
}

// TODO: Implement method IsPrimary on Color type
// Return true for red, green, blue
func (c Color) IsPrimary() bool {
	// Fix me!
	return false
}

// TODO: Implement method Opposite on Direction type
// North->South, East->West, etc.
func (d Direction) Opposite() Direction {
	// Fix me!
	return North
}

// TODO: Implement method TurnRight on Direction type
// North->East->South->West->North
func (d Direction) TurnRight() Direction {
	// Fix me!
	return North
}

// TODO: Implement function MixColors that takes two colors and returns a new color
// Define your own mixing rules (e.g., Red + Green = Yellow)
func MixColors(c1, c2 Color) Color {
	// Fix me!
	return ""
}

// Enums2Demo demonstrates the type safety features
func Enums2Demo() string {
	result := ""

	// Test temperature operations
	temp := Temperature(25.0)
	result += fmt.Sprintf("25°C = %.1f°F\n", temp.ToFahrenheit())
	result += fmt.Sprintf("Is 25°C freezing or below? %v\n", temp.IsFreezingOrBelow())

	// Test color operations
	result += fmt.Sprintf("Is red a primary color? %v\n", Red.IsPrimary())
	result += fmt.Sprintf("Red + Green = %s\n", MixColors(Red, Green))

	// Test direction operations
	result += fmt.Sprintf("Opposite of North: %v\n", North.Opposite())
	result += fmt.Sprintf("Turn right from North: %v\n", North.TurnRight())

	return result
}
