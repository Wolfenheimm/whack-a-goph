package strings

import "testing"

func TestCountWords(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"hello world", 2},
		{"", 0},
		{"   ", 0},
		{"one", 1},
		{"  hello   world  go  ", 3},
	}

	for _, tt := range tests {
		got := CountWords(tt.input)
		if got != tt.want {
			t.Errorf("CountWords(%q) = %d; want %d", tt.input, got, tt.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"Go", "oG"},
	}

	for _, tt := range tests {
		got := ReverseString(tt.input)
		if got != tt.want {
			t.Errorf("ReverseString(%q) = %q; want %q", tt.input, got, tt.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A man a plan a canal Panama", true},
		{"race a car", false},
		{"", true},
	}

	for _, tt := range tests {
		got := IsPalindrome(tt.input)
		if got != tt.want {
			t.Errorf("IsPalindrome(%q) = %t; want %t", tt.input, got, tt.want)
		}
	}
}

func TestTitleCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello world", "Hello World"},
		{"go programming", "Go Programming"},
		{"", ""},
		{"HELLO", "Hello"},
	}

	for _, tt := range tests {
		got := TitleCase(tt.input)
		if got != tt.want {
			t.Errorf("TitleCase(%q) = %q; want %q", tt.input, got, tt.want)
		}
	}
}
