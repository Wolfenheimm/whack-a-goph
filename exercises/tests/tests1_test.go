package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	expected := 20
	if result != expected {
		t.Errorf("Multiply(4, 5) = %d; want %d", result, expected)
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-2, true},
		{-3, false},
	}

	for _, test := range tests {
		result := IsEven(test.input)
		if result != test.expected {
			t.Errorf("IsEven(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}
