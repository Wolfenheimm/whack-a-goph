package main

// Generics 1: Basic Generic Functions 🐹
// Learn how to write generic functions using type parameters.
// Note: Requires Go 1.18 or later.

// I AM NOT DONE

import "fmt"

// TODO: Implement a generic Min function
// It should work with any type that can be compared (use comparable constraint)
func Min[T comparable](a, b T) T {
	// Fix me! Return the smaller value
	var zero T
	return zero
}

// TODO: Implement a generic Max function
// It should work with any type that can be compared
func Max[T comparable](a, b T) T {
	// Fix me! Return the larger value
	var zero T
	return zero
}

// TODO: Implement a generic Swap function
// It should swap two values of any type
func Swap[T any](a, b T) (T, T) {
	// Fix me! Return b, a
	return a, b
}

// TODO: Implement a generic Contains function for slices
// Check if a slice contains a specific value
func Contains[T comparable](slice []T, value T) bool {
	// Fix me! Loop through slice and check for value
	return false
}

// TODO: Implement a generic Reverse function for slices
// Return a new slice with elements in reverse order
func Reverse[T any](slice []T) []T {
	// Fix me! Create reversed slice
	return nil
}

// TODO: Implement a generic Map function
// Apply a function to each element and return new slice
func Map[T, U any](slice []T, fn func(T) U) []U {
	// Fix me! Apply fn to each element
	return nil
}

// TODO: Implement a generic Filter function
// Return slice with elements that satisfy the predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// Fix me! Filter elements based on predicate
	return nil
}

func main() {
	// Test Min and Max with integers
	fmt.Printf("Min(5, 3): %d\n", Min(5, 3))
	fmt.Printf("Max(5, 3): %d\n", Max(5, 3))

	// Test Min and Max with strings
	fmt.Printf("Min(\"hello\", \"world\"): %s\n", Min("hello", "world"))
	fmt.Printf("Max(\"hello\", \"world\"): %s\n", Max("hello", "world"))

	// Test Swap
	a, b := Swap(10, 20)
	fmt.Printf("Swap(10, 20): %d, %d\n", a, b)

	// Test Contains
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Contains([1,2,3,4,5], 3): %v\n", Contains(numbers, 3))
	fmt.Printf("Contains([1,2,3,4,5], 6): %v\n", Contains(numbers, 6))

	// Test Reverse
	reversed := Reverse(numbers)
	fmt.Printf("Reverse([1,2,3,4,5]): %v\n", reversed)

	// Test Map - double each number
	doubled := Map(numbers, func(x int) int { return x * 2 })
	fmt.Printf("Map(double): %v\n", doubled)

	// Test Filter - only even numbers
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Filter(even): %v\n", evens)
}
