package quizzes

import (
	"strings"
	"testing"
)

func TestCalculateAge(t *testing.T) {
	age := CalculateAge(1990, 2024)
	expectedAge := 34
	if age != expectedAge {
		t.Errorf("CalculateAge(1990, 2024) = %d; want %d", age, expectedAge)
	}
}

func TestFormatName(t *testing.T) {
	fullName := FormatName("Jane", "Doe")
	expectedName := "Doe, Jane"
	if fullName != expectedName {
		t.Errorf("FormatName(\"Jane\", \"Doe\") = %q; want %q", fullName, expectedName)
	}
}

func TestQuiz1Demo(t *testing.T) {
	result := Quiz1()

	// Check that the result contains expected components
	expectedStrings := []string{
		"Name: Doe, Jane",
		"Age: 34",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(result, expected) {
			t.Errorf("Quiz1() result should contain %q, but got %q", expected, result)
		}
	}
}
