package main

// Enums 3: String Methods for Enums 🐹
// Learn how to implement the String() method to make
// your enums more readable and debuggable.

// I AM NOT DONE

import "fmt"

// TODO: Define a LogLevel type based on int
type LogLevel int // Fix me!

// TODO: Define log level constants using iota
const (
	Debug LogLevel = 0 // Fix me! Use iota
	Info  LogLevel = 0 // Fix me!
	Warn  LogLevel = 0 // Fix me!
	Error LogLevel = 0 // Fix me!
	Fatal LogLevel = 0 // Fix me!
)

// TODO: Implement String method for LogLevel
// Return the string representation: "DEBUG", "INFO", "WARN", "ERROR", "FATAL"
func (l LogLevel) String() string {
	// Fix me!
	return ""
}

// TODO: Define a HttpStatus type based on int
type HttpStatus int // Fix me!

// TODO: Define HTTP status constants
const (
	StatusOK                  HttpStatus = 0 // Fix me! Set to 200
	StatusNotFound            HttpStatus = 0 // Fix me! Set to 404
	StatusInternalServerError HttpStatus = 0 // Fix me! Set to 500
)

// TODO: Implement String method for HttpStatus
// Return descriptive strings: "200 OK", "404 Not Found", "500 Internal Server Error"
func (h HttpStatus) String() string {
	// Fix me!
	return ""
}

// TODO: Define a Planet type based on int
type Planet int // Fix me!

// TODO: Define planet constants using iota
const (
	Mercury Planet = 0 // Fix me! Use iota
	Venus   Planet = 0 // Fix me!
	Earth   Planet = 0 // Fix me!
	Mars    Planet = 0 // Fix me!
	Jupiter Planet = 0 // Fix me!
	Saturn  Planet = 0 // Fix me!
	Uranus  Planet = 0 // Fix me!
	Neptune Planet = 0 // Fix me!
)

// TODO: Implement String method for Planet
// Return the planet name: "Mercury", "Venus", etc.
func (p Planet) String() string {
	// Fix me!
	return ""
}

// TODO: Implement IsInnerPlanet method for Planet
// Return true for Mercury, Venus, Earth, Mars
func (p Planet) IsInnerPlanet() bool {
	// Fix me!
	return false
}

// TODO: Implement function ParseLogLevel that converts string to LogLevel
// Return the LogLevel and a bool indicating success
func ParseLogLevel(s string) (LogLevel, bool) {
	// Fix me!
	return Debug, false
}

// TODO: Implement function LogMessage that takes level and message
// Return formatted string: "[LEVEL] message"
func LogMessage(level LogLevel, message string) string {
	// Fix me!
	return ""
}

// Enums3Demo demonstrates string methods for enums
func Enums3Demo() string {
	result := ""

	// Test log levels
	result += fmt.Sprintf("Log level: %s\n", Info)
	result += fmt.Sprintf("Error message: %s\n", LogMessage(Error, "Something went wrong"))

	// Test HTTP status
	result += fmt.Sprintf("HTTP Status: %s\n", StatusNotFound)

	// Test planets
	result += fmt.Sprintf("Planet: %s\n", Earth)
	result += fmt.Sprintf("Is Earth an inner planet? %v\n", Earth.IsInnerPlanet())

	// Test parsing
	if level, ok := ParseLogLevel("WARN"); ok {
		result += fmt.Sprintf("Parsed level: %s\n", level)
	}

	return result
}
