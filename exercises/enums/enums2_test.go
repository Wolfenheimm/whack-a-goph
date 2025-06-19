package main

import "testing"

func TestTemperatureConstants(t *testing.T) {
	if AbsoluteZeroC != -273.15 {
		t.Errorf("AbsoluteZeroC should be -273.15, got %f", AbsoluteZeroC)
	}
	if FreezingC != 0.0 {
		t.Errorf("FreezingC should be 0.0, got %f", FreezingC)
	}
	if BoilingC != 100.0 {
		t.Errorf("BoilingC should be 100.0, got %f", BoilingC)
	}
}

func TestColorConstants(t *testing.T) {
	if Red != "red" {
		t.Errorf("Red should be 'red', got %q", Red)
	}
	if Green != "green" {
		t.Errorf("Green should be 'green', got %q", Green)
	}
	if Blue != "blue" {
		t.Errorf("Blue should be 'blue', got %q", Blue)
	}
	if Yellow != "yellow" {
		t.Errorf("Yellow should be 'yellow', got %q", Yellow)
	}
}

func TestDirectionConstants(t *testing.T) {
	if North != 0 {
		t.Errorf("North should be 0, got %d", North)
	}
	if East != 1 {
		t.Errorf("East should be 1, got %d", East)
	}
	if South != 2 {
		t.Errorf("South should be 2, got %d", South)
	}
	if West != 3 {
		t.Errorf("West should be 3, got %d", West)
	}
}

func TestTemperatureToFahrenheit(t *testing.T) {
	tests := []struct {
		celsius    Temperature
		fahrenheit Temperature
	}{
		{0.0, 32.0},
		{25.0, 77.0},
		{100.0, 212.0},
		{-40.0, -40.0},
	}

	for _, test := range tests {
		result := test.celsius.ToFahrenheit()
		if result != test.fahrenheit {
			t.Errorf("ToFahrenheit(%f) = %f; want %f", test.celsius, result, test.fahrenheit)
		}
	}
}

func TestTemperatureIsFreezingOrBelow(t *testing.T) {
	tests := []struct {
		temp     Temperature
		expected bool
	}{
		{-10.0, true},
		{0.0, true},
		{10.0, false},
		{25.0, false},
	}

	for _, test := range tests {
		result := test.temp.IsFreezingOrBelow()
		if result != test.expected {
			t.Errorf("IsFreezingOrBelow(%f) = %v; want %v", test.temp, result, test.expected)
		}
	}
}

func TestColorIsPrimary(t *testing.T) {
	tests := []struct {
		color    Color
		expected bool
	}{
		{Red, true},
		{Green, true},
		{Blue, true},
		{Yellow, false},
	}

	for _, test := range tests {
		result := test.color.IsPrimary()
		if result != test.expected {
			t.Errorf("IsPrimary(%s) = %v; want %v", test.color, result, test.expected)
		}
	}
}

func TestDirectionOpposite(t *testing.T) {
	tests := []struct {
		direction Direction
		expected  Direction
	}{
		{North, South},
		{South, North},
		{East, West},
		{West, East},
	}

	for _, test := range tests {
		result := test.direction.Opposite()
		if result != test.expected {
			t.Errorf("Opposite(%d) = %d; want %d", test.direction, result, test.expected)
		}
	}
}

func TestDirectionTurnRight(t *testing.T) {
	tests := []struct {
		direction Direction
		expected  Direction
	}{
		{North, East},
		{East, South},
		{South, West},
		{West, North},
	}

	for _, test := range tests {
		result := test.direction.TurnRight()
		if result != test.expected {
			t.Errorf("TurnRight(%d) = %d; want %d", test.direction, result, test.expected)
		}
	}
}

func TestMixColors(t *testing.T) {
	// Test specific color mixing rules
	result := MixColors(Red, Green)
	if result != Yellow {
		t.Errorf("MixColors(Red, Green) should return Yellow, got %s", result)
	}
}
