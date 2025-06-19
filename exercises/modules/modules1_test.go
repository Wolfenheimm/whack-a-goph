package main

import (
	"math"
	"testing"
	"time"
)

func TestCalculateCircleArea(t *testing.T) {
	tests := []struct {
		radius   float64
		expected float64
	}{
		{1.0, math.Pi},
		{2.0, 4 * math.Pi},
		{5.0, 25 * math.Pi},
	}

	for _, test := range tests {
		result := CalculateCircleArea(test.radius)
		if result != test.expected {
			t.Errorf("CalculateCircleArea(%f) = %f; want %f", test.radius, result, test.expected)
		}
	}
}

func TestFormatUserName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"john doe", "John Doe"},
		{"jane smith", "Jane Smith"},
		{"bob wilson", "Bob Wilson"},
	}

	for _, test := range tests {
		result := FormatUserName(test.input)
		if result != test.expected {
			t.Errorf("FormatUserName(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestGetCurrentYear(t *testing.T) {
	result := GetCurrentYear()
	currentYear := time.Now().Year()

	if result != currentYear {
		t.Errorf("GetCurrentYear() = %d; want %d", result, currentYear)
	}
}

func TestIsWorkingHour(t *testing.T) {
	tests := []struct {
		hour     int
		expected bool
	}{
		{8, false},  // Before working hours
		{9, true},   // Start of working hours
		{12, true},  // Middle of day
		{17, true},  // End of working hours
		{18, false}, // After working hours
		{22, false}, // Evening
	}

	for _, test := range tests {
		result := IsWorkingHour(test.hour)
		if result != test.expected {
			t.Errorf("IsWorkingHour(%d) = %v; want %v", test.hour, result, test.expected)
		}
	}
}

func TestConvertToUpperCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"world", "WORLD"},
		{"Hello World", "HELLO WORLD"},
		{"", ""},
	}

	for _, test := range tests {
		result := ConvertToUpperCase(test.input)
		if result != test.expected {
			t.Errorf("ConvertToUpperCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
