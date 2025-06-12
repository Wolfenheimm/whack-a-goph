package control_flow

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		n, expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{-10, 10},
		{42, 42},
	}

	for _, test := range tests {
		actual := Abs(test.n)
		if actual != test.expected {
			t.Errorf("Abs(%d) = %d; want %d", test.n, actual, test.expected)
		}
	}
}

func TestClassify(t *testing.T) {
	tests := []struct {
		n        int
		expected string
	}{
		{5, "positive"},
		{-5, "negative"},
		{0, "zero"},
		{42, "positive"},
		{-1, "negative"},
	}

	for _, test := range tests {
		actual := Classify(test.n)
		if actual != test.expected {
			t.Errorf("Classify(%d) = %q; want %q", test.n, actual, test.expected)
		}
	}
}

func TestSumToN(t *testing.T) {
	tests := []struct {
		n, expected int
	}{
		{1, 1},   // 1
		{3, 6},   // 1+2+3
		{5, 15},  // 1+2+3+4+5
		{10, 55}, // 1+2+...+10
		{0, 0},   // edge case
	}

	for _, test := range tests {
		actual := SumToN(test.n)
		if actual != test.expected {
			t.Errorf("SumToN(%d) = %d; want %d", test.n, actual, test.expected)
		}
	}
}

func TestCountEvens(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3}, // 2, 4, 6 are even
		{[]int{1, 3, 5}, 0},          // no evens
		{[]int{2, 4, 6, 8}, 4},       // all evens
		{[]int{}, 0},                 // empty slice
		{[]int{0, -2, -1, 4}, 3},     // 0, -2, 4 are even
	}

	for _, test := range tests {
		actual := CountEvens(test.numbers)
		if actual != test.expected {
			t.Errorf("CountEvens(%v) = %d; want %d", test.numbers, actual, test.expected)
		}
	}
}
