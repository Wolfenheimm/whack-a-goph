package if_statements

import "testing"

func TestIsPositive(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{5, true},
		{-3, false},
		{0, false},
		{100, true},
	}

	for _, test := range tests {
		result := IsPositive(test.input)
		if result != test.expected {
			t.Errorf("IsPositive(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestGetGrade(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{95, "A"},
		{85, "B"},
		{75, "C"},
		{65, "D"},
		{55, "F"},
		{100, "A"},
		{80, "B"},
	}

	for _, test := range tests {
		result := GetGrade(test.score)
		if result != test.expected {
			t.Errorf("GetGrade(%d) = %q; want %q", test.score, result, test.expected)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 5},
		{1, 10, 10},
		{7, 7, 7},
		{-5, -3, -3},
	}

	for _, test := range tests {
		result := Max(test.a, test.b)
		if result != test.expected {
			t.Errorf("Max(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}
