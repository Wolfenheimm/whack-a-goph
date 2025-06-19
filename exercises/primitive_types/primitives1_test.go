package primitive_types

import "testing"

func TestGetAge(t *testing.T) {
	age := GetAge()
	if age <= 0 {
		t.Errorf("GetAge() = %d; want a positive number", age)
	}
}

func TestGetName(t *testing.T) {
	name := GetName()
	if name == "" {
		t.Errorf("GetName() = %q; want a non-empty string", name)
	}
}

func TestIsStudent(t *testing.T) {
	// This test just checks that a boolean is returned
	result := IsStudent()
	if result != true && result != false {
		t.Errorf("IsStudent() should return a boolean value")
	}
}

func TestGetPI(t *testing.T) {
	pi := GetPI()
	expected := 3.14159
	if pi < 3.0 || pi > 4.0 {
		t.Errorf("GetPI() = %f; want approximately %f", pi, expected)
	}
}
