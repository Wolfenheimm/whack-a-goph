package if_statements

import "testing"

func TestCanVote(t *testing.T) {
	tests := []struct {
		age      int
		expected string
	}{
		{18, "Can vote"},
		{25, "Can vote"},
		{17, "Cannot vote"},
		{16, "Cannot vote"},
	}

	for _, test := range tests {
		result := CanVote(test.age)
		if result != test.expected {
			t.Errorf("CanVote(%d) = %q; want %q", test.age, result, test.expected)
		}
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		day      string
		expected bool
	}{
		{"Saturday", true},
		{"Sunday", true},
		{"Monday", false},
		{"Friday", false},
		{"Wednesday", false},
	}

	for _, test := range tests {
		result := IsWeekend(test.day)
		if result != test.expected {
			t.Errorf("IsWeekend(%q) = %v; want %v", test.day, result, test.expected)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"password123", true},   // >= 8 chars, has digit
		{"short1", false},       // < 8 chars
		{"longpassword", false}, // >= 8 chars, no digit
		{"pass1", false},        // < 8 chars, has digit
		{"mypassword9", true},   // >= 8 chars, has digit
	}

	for _, test := range tests {
		result := IsValidPassword(test.password)
		if result != test.expected {
			t.Errorf("IsValidPassword(%q) = %v; want %v", test.password, result, test.expected)
		}
	}
}

func TestGetTemperatureAdvice(t *testing.T) {
	tests := []struct {
		temp     int
		expected string
	}{
		{-5, "Stay inside!"},
		{5, "Wear a coat"},
		{20, "Perfect weather"},
		{30, "Stay hydrated"},
		{0, "Wear a coat"},
		{15, "Perfect weather"},
		{25, "Stay hydrated"},
	}

	for _, test := range tests {
		result := GetTemperatureAdvice(test.temp)
		if result != test.expected {
			t.Errorf("GetTemperatureAdvice(%d) = %q; want %q", test.temp, result, test.expected)
		}
	}
}
