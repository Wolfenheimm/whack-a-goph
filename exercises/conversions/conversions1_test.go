package conversions

import "testing"

func TestIntToFloat(t *testing.T) {
	tests := []struct {
		input    int
		expected float64
	}{
		{42, 42.0},
		{0, 0.0},
		{-5, -5.0},
	}

	for _, test := range tests {
		result := IntToFloat(test.input)
		if result != test.expected {
			t.Errorf("IntToFloat(%d) = %f; want %f", test.input, result, test.expected)
		}
	}
}

func TestFloatToInt(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{42.7, 42},
		{0.9, 0},
		{-5.3, -5},
	}

	for _, test := range tests {
		result := FloatToInt(test.input)
		if result != test.expected {
			t.Errorf("FloatToInt(%f) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestStringToInt(t *testing.T) {
	// Simple test - for now just check that it returns some integer
	result := StringToInt("42")
	if result == 0 {
		t.Errorf("StringToInt should convert string to integer, got %d", result)
	}
}

func TestIntToString(t *testing.T) {
	// Simple test - for now just check that it returns some string
	result := IntToString(42)
	if result == "" {
		t.Errorf("IntToString should convert int to string, got empty string")
	}
}

func TestByteToRune(t *testing.T) {
	tests := []struct {
		input    byte
		expected rune
	}{
		{65, 'A'},
		{97, 'a'},
		{48, '0'},
	}

	for _, test := range tests {
		result := ByteToRune(test.input)
		if result != test.expected {
			t.Errorf("ByteToRune(%d) = %c; want %c", test.input, result, test.expected)
		}
	}
}
