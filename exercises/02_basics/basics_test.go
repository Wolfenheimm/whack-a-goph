package basics

import "testing"

func TestGetName(t *testing.T) {
	expected := "Gopher"
	actual := GetName()

	if actual != expected {
		t.Errorf("GetName() = %q; want %q", actual, expected)
	}
}

func TestGetAge(t *testing.T) {
	expected := 5
	actual := GetAge()

	if actual != expected {
		t.Errorf("GetAge() = %d; want %d", actual, expected)
	}
}

func TestIsAwesome(t *testing.T) {
	expected := true
	actual := IsAwesome()

	if actual != expected {
		t.Errorf("IsAwesome() = %v; want %v", actual, expected)
	}
}
