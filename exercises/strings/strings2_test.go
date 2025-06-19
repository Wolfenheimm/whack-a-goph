package strings

import (
	"reflect"
	"testing"
)

func TestToUpperCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"World", "WORLD"},
		{"gOpHeR", "GOPHER"},
		{"", ""},
	}

	for _, test := range tests {
		result := ToUpperCase(test.input)
		if result != test.expected {
			t.Errorf("ToUpperCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestSplitWords(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"hello world", []string{"hello", "world"}},
		{"Go is awesome", []string{"Go", "is", "awesome"}},
		{"", []string{}},
		{"single", []string{"single"}},
	}

	for _, test := range tests {
		result := SplitWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SplitWords(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestReplaceAll(t *testing.T) {
	tests := []struct {
		s, old, new, expected string
	}{
		{"hello world", "o", "0", "hell0 w0rld"},
		{"Go Go Go", "Go", "Rust", "Rust Rust Rust"},
		{"test", "x", "y", "test"},
		{"", "a", "b", ""},
	}

	for _, test := range tests {
		result := ReplaceAll(test.s, test.old, test.new)
		if result != test.expected {
			t.Errorf("ReplaceAll(%q, %q, %q) = %q; want %q", test.s, test.old, test.new, result, test.expected)
		}
	}
}

func TestTrimSpaces(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  hello  ", "hello"},
		{"\tworld\n", "world"},
		{"  ", ""},
		{"nospace", "nospace"},
	}

	for _, test := range tests {
		result := TrimSpaces(test.input)
		if result != test.expected {
			t.Errorf("TrimSpaces(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		s, prefix string
		expected  bool
	}{
		{"hello world", "hello", true},
		{"Go programming", "Go", true},
		{"test", "fail", false},
		{"", "", true},
	}

	for _, test := range tests {
		result := HasPrefix(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("HasPrefix(%q, %q) = %v; want %v", test.s, test.prefix, result, test.expected)
		}
	}
}
