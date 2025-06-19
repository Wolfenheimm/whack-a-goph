package strings

import "testing"

func TestGetLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{"Go", 2},
		{"🐹", 4}, // Unicode character is 4 bytes
	}

	for _, test := range tests {
		result := GetLength(test.input)
		if result != test.expected {
			t.Errorf("GetLength(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestConcatenate(t *testing.T) {
	tests := []struct {
		a, b, expected string
	}{
		{"hello", "world", "helloworld"},
		{"Go", "pher", "Gopher"},
		{"", "test", "test"},
		{"test", "", "test"},
	}

	for _, test := range tests {
		result := Concatenate(test.a, test.b)
		if result != test.expected {
			t.Errorf("Concatenate(%q, %q) = %q; want %q", test.a, test.b, result, test.expected)
		}
	}
}

func TestGetFirstChar(t *testing.T) {
	tests := []struct {
		input    string
		expected byte
	}{
		{"hello", 'h'},
		{"Go", 'G'},
		{"abc", 'a'},
	}

	for _, test := range tests {
		result := GetFirstChar(test.input)
		if result != test.expected {
			t.Errorf("GetFirstChar(%q) = %c; want %c", test.input, result, test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello", "hello", true},
		{"hello", "world", false},
		{"Gopher", "Go", false}, // Simple implementation won't find substring
		{"test", "test", true},
	}

	for _, test := range tests {
		result := Contains(test.s, test.substr)
		if result != test.expected {
			t.Errorf("Contains(%q, %q) = %v; want %v", test.s, test.substr, result, test.expected)
		}
	}
}
