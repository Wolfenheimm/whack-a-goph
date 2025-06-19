package structs

import "testing"

func TestPersonStruct(t *testing.T) {
	// Test creating a person
	person := NewPerson("Alice", 25, "alice@example.com")

	if person.Name != "Alice" {
		t.Errorf("person.Name = %q; want %q", person.Name, "Alice")
	}
	if person.Age != 25 {
		t.Errorf("person.Age = %d; want %d", person.Age, 25)
	}
	if person.Email != "alice@example.com" {
		t.Errorf("person.Email = %q; want %q", person.Email, "alice@example.com")
	}
}

func TestGreet(t *testing.T) {
	person := NewPerson("Bob", 30, "bob@example.com")
	expected := "Hi, I'm Bob and I'm 30 years old!"
	actual := person.Greet()

	if actual != expected {
		t.Errorf("person.Greet() = %q; want %q", actual, expected)
	}
}

func TestIsAdult(t *testing.T) {
	tests := []struct {
		person   Person
		expected bool
	}{
		{NewPerson("Adult", 18, "adult@example.com"), true},
		{NewPerson("Kid", 17, "kid@example.com"), false},
		{NewPerson("Elder", 65, "elder@example.com"), true},
		{NewPerson("Baby", 5, "baby@example.com"), false},
	}

	for _, test := range tests {
		actual := test.person.IsAdult()
		if actual != test.expected {
			t.Errorf("%s.IsAdult() = %v; want %v", test.person.Name, actual, test.expected)
		}
	}
}
