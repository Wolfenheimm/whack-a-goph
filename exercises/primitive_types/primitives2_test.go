package primitive_types

import "testing"

func TestMaxInt(t *testing.T) {
	if MaxInt != 2147483647 {
		t.Errorf("MaxInt = %d; want 2147483647", MaxInt)
	}
}

func TestAppName(t *testing.T) {
	if AppName != "WhackAGoph" {
		t.Errorf("AppName = %q; want %q", AppName, "WhackAGoph")
	}
}

func TestGetByte(t *testing.T) {
	result := GetByte()
	if result != 65 {
		t.Errorf("GetByte() = %d; want 65", result)
	}
}

func TestGetRune(t *testing.T) {
	result := GetRune()
	if result != '🐹' {
		t.Errorf("GetRune() = %v; want '🐹'", result)
	}
}

func TestGetFloat32(t *testing.T) {
	result := GetFloat32()
	if result != 1.5 {
		t.Errorf("GetFloat32() = %f; want 1.5", result)
	}
}
