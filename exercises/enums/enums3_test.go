package main

import (
	"strings"
	"testing"
)

func TestLogLevelConstants(t *testing.T) {
	if Debug != 0 {
		t.Errorf("Debug should be 0, got %d", Debug)
	}
	if Info != 1 {
		t.Errorf("Info should be 1, got %d", Info)
	}
	if Fatal != 4 {
		t.Errorf("Fatal should be 4, got %d", Fatal)
	}
}

func TestLogLevelString(t *testing.T) {
	tests := []struct {
		level    LogLevel
		expected string
	}{
		{Debug, "DEBUG"},
		{Info, "INFO"},
		{Warn, "WARN"},
		{Error, "ERROR"},
		{Fatal, "FATAL"},
	}

	for _, test := range tests {
		result := test.level.String()
		if result != test.expected {
			t.Errorf("LogLevel(%d).String() = %q; want %q", test.level, result, test.expected)
		}
	}
}

func TestHttpStatusConstants(t *testing.T) {
	if StatusOK != 200 {
		t.Errorf("StatusOK should be 200, got %d", StatusOK)
	}
	if StatusNotFound != 404 {
		t.Errorf("StatusNotFound should be 404, got %d", StatusNotFound)
	}
	if StatusInternalServerError != 500 {
		t.Errorf("StatusInternalServerError should be 500, got %d", StatusInternalServerError)
	}
}

func TestHttpStatusString(t *testing.T) {
	tests := []struct {
		status   HttpStatus
		expected string
	}{
		{StatusOK, "200 OK"},
		{StatusNotFound, "404 Not Found"},
		{StatusInternalServerError, "500 Internal Server Error"},
	}

	for _, test := range tests {
		result := test.status.String()
		if result != test.expected {
			t.Errorf("HttpStatus(%d).String() = %q; want %q", test.status, result, test.expected)
		}
	}
}

func TestPlanetConstants(t *testing.T) {
	if Mercury != 0 {
		t.Errorf("Mercury should be 0, got %d", Mercury)
	}
	if Earth != 2 {
		t.Errorf("Earth should be 2, got %d", Earth)
	}
	if Neptune != 7 {
		t.Errorf("Neptune should be 7, got %d", Neptune)
	}
}

func TestPlanetString(t *testing.T) {
	tests := []struct {
		planet   Planet
		expected string
	}{
		{Mercury, "Mercury"},
		{Venus, "Venus"},
		{Earth, "Earth"},
		{Mars, "Mars"},
		{Jupiter, "Jupiter"},
		{Saturn, "Saturn"},
		{Uranus, "Uranus"},
		{Neptune, "Neptune"},
	}

	for _, test := range tests {
		result := test.planet.String()
		if result != test.expected {
			t.Errorf("Planet(%d).String() = %q; want %q", test.planet, result, test.expected)
		}
	}
}

func TestPlanetIsInnerPlanet(t *testing.T) {
	tests := []struct {
		planet   Planet
		expected bool
	}{
		{Mercury, true},
		{Venus, true},
		{Earth, true},
		{Mars, true},
		{Jupiter, false},
		{Saturn, false},
		{Uranus, false},
		{Neptune, false},
	}

	for _, test := range tests {
		result := test.planet.IsInnerPlanet()
		if result != test.expected {
			t.Errorf("Planet(%s).IsInnerPlanet() = %v; want %v", test.planet, result, test.expected)
		}
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected LogLevel
		success  bool
	}{
		{"DEBUG", Debug, true},
		{"INFO", Info, true},
		{"WARN", Warn, true},
		{"ERROR", Error, true},
		{"FATAL", Fatal, true},
		{"INVALID", Debug, false},
	}

	for _, test := range tests {
		result, ok := ParseLogLevel(test.input)
		if ok != test.success {
			t.Errorf("ParseLogLevel(%q) success = %v; want %v", test.input, ok, test.success)
		}
		if test.success && result != test.expected {
			t.Errorf("ParseLogLevel(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestLogMessage(t *testing.T) {
	result := LogMessage(Error, "Something went wrong")
	expected := "[ERROR] Something went wrong"
	if result != expected {
		t.Errorf("LogMessage() = %q; want %q", result, expected)
	}
}

func TestEnums3Demo(t *testing.T) {
	result := Enums3Demo()

	// Check that result contains expected strings
	expectedStrings := []string{
		"Log level:",
		"HTTP Status:",
		"Planet:",
		"inner planet",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(result, expected) {
			t.Errorf("Enums3Demo() should contain %q, but got %q", expected, result)
		}
	}
}
