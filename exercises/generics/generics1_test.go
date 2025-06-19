package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinIntegers(t *testing.T) {
	result := Min(5, 3)
	expected := 3
	if result != expected {
		t.Errorf("Min(5, 3) = %d; want %d", result, expected)
	}
}

func TestMinStrings(t *testing.T) {
	result := Min("hello", "world")
	expected := "hello" // "hello" < "world" lexicographically
	if result != expected {
		t.Errorf("Min(\"hello\", \"world\") = %q; want %q", result, expected)
	}
}

func TestMaxIntegers(t *testing.T) {
	result := Max(5, 3)
	expected := 5
	if result != expected {
		t.Errorf("Max(5, 3) = %d; want %d", result, expected)
	}
}

func TestMaxStrings(t *testing.T) {
	result := Max("hello", "world")
	expected := "world" // "world" > "hello" lexicographically
	if result != expected {
		t.Errorf("Max(\"hello\", \"world\") = %q; want %q", result, expected)
	}
}

func TestSwap(t *testing.T) {
	a, b := Swap(10, 20)
	if a != 20 || b != 10 {
		t.Errorf("Swap(10, 20) = %d, %d; want 20, 10", a, b)
	}

	// Test with strings
	s1, s2 := Swap("hello", "world")
	if s1 != "world" || s2 != "hello" {
		t.Errorf("Swap(\"hello\", \"world\") = %q, %q; want \"world\", \"hello\"", s1, s2)
	}
}

func TestContains(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	if !Contains(numbers, 3) {
		t.Errorf("Contains([1,2,3,4,5], 3) should be true")
	}

	if Contains(numbers, 6) {
		t.Errorf("Contains([1,2,3,4,5], 6) should be false")
	}

	// Test with strings
	words := []string{"hello", "world", "go"}
	if !Contains(words, "go") {
		t.Errorf("Contains([\"hello\", \"world\", \"go\"], \"go\") should be true")
	}
}

func TestReverse(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := Reverse(numbers)
	expected := []int{5, 4, 3, 2, 1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Reverse([1,2,3,4,5]) = %v; want %v", result, expected)
	}

	// Test that original slice is unchanged
	if !reflect.DeepEqual(numbers, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Original slice should not be modified")
	}
}

func TestMap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	// Test doubling
	doubled := Map(numbers, func(x int) int { return x * 2 })
	expected := []int{2, 4, 6, 8, 10}

	if !reflect.DeepEqual(doubled, expected) {
		t.Errorf("Map(double) = %v; want %v", doubled, expected)
	}

	// Test converting to strings
	strings := Map(numbers, func(x int) string { return fmt.Sprintf("num_%d", x) })
	expectedStrings := []string{"num_1", "num_2", "num_3", "num_4", "num_5"}

	if !reflect.DeepEqual(strings, expectedStrings) {
		t.Errorf("Map(toString) = %v; want %v", strings, expectedStrings)
	}
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Test filtering even numbers
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	expected := []int{2, 4, 6, 8, 10}

	if !reflect.DeepEqual(evens, expected) {
		t.Errorf("Filter(even) = %v; want %v", evens, expected)
	}

	// Test filtering numbers > 5
	greater := Filter(numbers, func(x int) bool { return x > 5 })
	expectedGreater := []int{6, 7, 8, 9, 10}

	if !reflect.DeepEqual(greater, expectedGreater) {
		t.Errorf("Filter(>5) = %v; want %v", greater, expectedGreater)
	}
}
