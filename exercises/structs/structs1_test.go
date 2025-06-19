package structs

import "testing"

func TestCreatePerson(t *testing.T) {
	person := CreatePerson("Alice", 25)

	if person.Name != "Alice" {
		t.Errorf("Expected Name to be 'Alice', got %q", person.Name)
	}

	if person.Age != 25 {
		t.Errorf("Expected Age to be 25, got %d", person.Age)
	}
}

func TestGetPersonInfo(t *testing.T) {
	person := Person{Name: "Bob", Age: 30}
	expected := "Name: Bob, Age: 30"
	result := GetPersonInfo(person)

	if result != expected {
		t.Errorf("GetPersonInfo() = %q; want %q", result, expected)
	}
}

func TestIsAdult(t *testing.T) {
	tests := []struct {
		person   Person
		expected bool
	}{
		{Person{Name: "Alice", Age: 25}, true},
		{Person{Name: "Bob", Age: 17}, false},
		{Person{Name: "Charlie", Age: 18}, true},
		{Person{Name: "Dave", Age: 16}, false},
	}

	for _, test := range tests {
		result := IsAdult(test.person)
		if result != test.expected {
			t.Errorf("IsAdult(%v) = %v; want %v", test.person, result, test.expected)
		}
	}
}
