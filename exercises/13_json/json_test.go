package json

import (
	"strings"
	"testing"
)

func TestPersonJSON(t *testing.T) {
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}

	// Test PersonToJSON
	jsonStr, err := PersonToJSON(person)
	if err != nil {
		t.Errorf("PersonToJSON() returned error: %v", err)
		return
	}

	// Check if JSON contains expected fields
	if !strings.Contains(jsonStr, "John Doe") {
		t.Errorf("PersonToJSON() does not contain expected name")
	}

	// Test PersonFromJSON
	parsed, err := PersonFromJSON(jsonStr)
	if err != nil {
		t.Errorf("PersonFromJSON() returned error: %v", err)
		return
	}

	if parsed.Name != person.Name || parsed.Age != person.Age || parsed.Email != person.Email {
		t.Errorf("PersonFromJSON() = %+v; want %+v", parsed, person)
	}
}

func TestPeopleJSON(t *testing.T) {
	people := []Person{
		{Name: "Alice", Age: 25, Email: "alice@example.com"},
		{Name: "Bob", Age: 35, Email: "bob@example.com"},
	}

	// Test PeopleToJSON
	jsonStr, err := PeopleToJSON(people)
	if err != nil {
		t.Errorf("PeopleToJSON() returned error: %v", err)
		return
	}

	// Test PeopleFromJSON
	parsed, err := PeopleFromJSON(jsonStr)
	if err != nil {
		t.Errorf("PeopleFromJSON() returned error: %v", err)
		return
	}

	if len(parsed) != len(people) {
		t.Errorf("PeopleFromJSON() returned %d people; want %d", len(parsed), len(people))
		return
	}

	for i, p := range parsed {
		if p.Name != people[i].Name || p.Age != people[i].Age || p.Email != people[i].Email {
			t.Errorf("PeopleFromJSON()[%d] = %+v; want %+v", i, p, people[i])
		}
	}
}
