package vecs

import (
	"reflect"
	"testing"
)

func TestCreateSlice(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	result := CreateSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CreateSlice() = %v; want %v", result, expected)
	}
}

func TestGetLength(t *testing.T) {
	tests := []struct {
		slice    []int
		expected int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{}, 0},
		{[]int{42}, 1},
	}

	for _, test := range tests {
		result := GetLength(test.slice)
		if result != test.expected {
			t.Errorf("GetLength(%v) = %d; want %d", test.slice, result, test.expected)
		}
	}
}

func TestGetFirst(t *testing.T) {
	tests := []struct {
		slice    []int
		expected int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{42}, 42},
		{[]int{5, 4, 3, 2, 1}, 5},
	}

	for _, test := range tests {
		result := GetFirst(test.slice)
		if result != test.expected {
			t.Errorf("GetFirst(%v) = %d; want %d", test.slice, result, test.expected)
		}
	}
}

func TestGetLast(t *testing.T) {
	tests := []struct {
		slice    []int
		expected int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{42}, 42},
		{[]int{5, 4, 3, 2, 1}, 1},
	}

	for _, test := range tests {
		result := GetLast(test.slice)
		if result != test.expected {
			t.Errorf("GetLast(%v) = %d; want %d", test.slice, result, test.expected)
		}
	}
}

func TestAppendElement(t *testing.T) {
	slice := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}
	result := AppendElement(slice, 4)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AppendElement(%v, 4) = %v; want %v", slice, result, expected)
	}
}
