package pointers

import "testing"

func TestSwapInts(t *testing.T) {
	a, b := 1, 2
	SwapInts(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("SwapInts(&1, &2) = %d, %d; want 2, 1", a, b)
	}
}

func TestDoubleValue(t *testing.T) {
	n := 5
	DoubleValue(&n)
	if n != 10 {
		t.Errorf("DoubleValue(&5) = %d; want 10", n)
	}
}

func TestCreateInt(t *testing.T) {
	n := CreateInt(42)
	if n == nil {
		t.Error("CreateInt(42) returned nil")
		return
	}
	if *n != 42 {
		t.Errorf("CreateInt(42) = %d; want 42", *n)
	}
}
