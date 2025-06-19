package quizzes

import (
	"strings"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	student := CreateStudent("Alice", 20, 95.5, true)

	if student.Name != "Alice" {
		t.Errorf("Expected Name to be 'Alice', got %q", student.Name)
	}
	if student.Age != 20 {
		t.Errorf("Expected Age to be 20, got %d", student.Age)
	}
	if student.Grade != 95.5 {
		t.Errorf("Expected Grade to be 95.5, got %f", student.Grade)
	}
	if student.IsActive != true {
		t.Errorf("Expected IsActive to be true, got %v", student.IsActive)
	}
}

func TestGetGradeLetter(t *testing.T) {
	tests := []struct {
		grade    float64
		expected string
	}{
		{95.0, "A"},
		{85.0, "B"},
		{75.0, "C"},
		{65.0, "D"},
		{55.0, "F"},
	}

	for _, test := range tests {
		result := GetGradeLetter(test.grade)
		if result != test.expected {
			t.Errorf("GetGradeLetter(%f) = %q; want %q", test.grade, result, test.expected)
		}
	}
}

func TestFormatStudentInfo(t *testing.T) {
	student := Student{Name: "Alice", Age: 20, Grade: 95.5, IsActive: true}
	result := FormatStudentInfo(student)

	// Check that the result contains the expected components
	if !strings.Contains(result, "Alice") {
		t.Errorf("FormatStudentInfo result should contain 'Alice', got %q", result)
	}
	if !strings.Contains(result, "20") {
		t.Errorf("FormatStudentInfo result should contain '20', got %q", result)
	}
	if !strings.Contains(result, "A") {
		t.Errorf("FormatStudentInfo result should contain 'A', got %q", result)
	}
}

func TestIsEligibleForHonors(t *testing.T) {
	tests := []struct {
		student  Student
		expected bool
	}{
		{Student{Name: "Alice", Age: 20, Grade: 90.0, IsActive: true}, true},
		{Student{Name: "Bob", Age: 19, Grade: 90.0, IsActive: false}, false},
		{Student{Name: "Charlie", Age: 21, Grade: 80.0, IsActive: true}, false},
		{Student{Name: "Dave", Age: 22, Grade: 80.0, IsActive: false}, false},
	}

	for _, test := range tests {
		result := IsEligibleForHonors(test.student)
		if result != test.expected {
			t.Errorf("IsEligibleForHonors(%v) = %v; want %v", test.student, result, test.expected)
		}
	}
}

func TestCountActiveStudents(t *testing.T) {
	students := []Student{
		{Name: "Alice", Age: 20, Grade: 90.0, IsActive: true},
		{Name: "Bob", Age: 19, Grade: 85.0, IsActive: false},
		{Name: "Charlie", Age: 21, Grade: 88.0, IsActive: true},
	}

	expected := 2
	result := CountActiveStudents(students)

	if result != expected {
		t.Errorf("CountActiveStudents() = %d; want %d", result, expected)
	}
}

func TestQuiz2(t *testing.T) {
	result := Quiz2()

	// Check that output contains expected elements
	if !strings.Contains(result, "Quiz 2 completed!") {
		t.Errorf("Quiz2() should contain 'Quiz 2 completed!', got %q", result)
	}
}
