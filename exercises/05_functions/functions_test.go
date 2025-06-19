package functions

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{10, 15, 25},
	}

	for _, test := range tests {
		actual := Add(test.a, test.b)
		if actual != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, actual, test.expected)
		}
	}
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name, expected string
	}{
		{"Alice", "Hello, Alice!"},
		{"Bob", "Hello, Bob!"},
		{"Gopher", "Hello, Gopher!"},
	}

	for _, test := range tests {
		actual := Greet(test.name)
		if actual != test.expected {
			t.Errorf("Greet(%q) = %q; want %q", test.name, actual, test.expected)
		}
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-2, true},
		{-3, false},
		{100, true},
		{101, false},
	}

	for _, test := range tests {
		actual := IsEven(test.n)
		if actual != test.expected {
			t.Errorf("IsEven(%d) = %v; want %v", test.n, actual, test.expected)
		}
	}
}
