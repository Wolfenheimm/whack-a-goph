package options

import "testing"

func TestGetPointer(t *testing.T) {
	value := 42
	ptr := GetPointer(value)

	if ptr == nil {
		t.Fatal("GetPointer() returned nil")
	}

	if *ptr != value {
		t.Errorf("GetPointer(%d) points to %d; want %d", value, *ptr, value)
	}
}

func TestGetValue(t *testing.T) {
	value := 42
	ptr := &value
	result := GetValue(ptr)

	if result != value {
		t.Errorf("GetValue(&%d) = %d; want %d", value, result, value)
	}
}

func TestIsNil(t *testing.T) {
	var nilPtr *int
	value := 42
	validPtr := &value

	if !IsNil(nilPtr) {
		t.Errorf("IsNil(nil) = false; want true")
	}

	if IsNil(validPtr) {
		t.Errorf("IsNil(validPtr) = true; want false")
	}
}

func TestSafeGetValue(t *testing.T) {
	var nilPtr *int
	value := 42
	validPtr := &value
	defaultValue := 100

	// Test with nil pointer
	result := SafeGetValue(nilPtr, defaultValue)
	if result != defaultValue {
		t.Errorf("SafeGetValue(nil, %d) = %d; want %d", defaultValue, result, defaultValue)
	}

	// Test with valid pointer
	result = SafeGetValue(validPtr, defaultValue)
	if result != value {
		t.Errorf("SafeGetValue(validPtr, %d) = %d; want %d", defaultValue, result, value)
	}
}

func TestCreateOptionalString(t *testing.T) {
	// Test with value
	value := "hello"
	ptr := CreateOptionalString(true, value)
	if ptr == nil {
		t.Fatal("CreateOptionalString(true, value) returned nil")
	}
	if *ptr != value {
		t.Errorf("CreateOptionalString(true, %q) points to %q; want %q", value, *ptr, value)
	}

	// Test without value
	ptr = CreateOptionalString(false, value)
	if ptr != nil {
		t.Errorf("CreateOptionalString(false, value) = %v; want nil", ptr)
	}
}
