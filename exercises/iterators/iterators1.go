package main

// Iterators 1: Range and Iteration Patterns 🐹
// Learn about Go's iteration patterns using range and custom iterators.

// I AM NOT DONE

import "fmt"

// TODO: Define a NumberRange struct with Start and End fields
type NumberRange struct {
	Start int // Fix me!
	End   int // Fix me!
}

// TODO: Implement SumSlice function
// Use range to sum all numbers in a slice
func SumSlice(numbers []int) int {
	// Fix me! Use range to iterate and sum
	return 0
}

// TODO: Implement CountVowels function
// Use range to count vowels in a string
func CountVowels(s string) int {
	// Fix me! Use range over string, check for vowels (a,e,i,o,u)
	return 0
}

// TODO: Implement FindKeyInMap function
// Use range to find if a key exists in map, return value and bool
func FindKeyInMap(m map[string]int, key string) (int, bool) {
	// Fix me! Use range over map
	return 0, false
}

// TODO: Implement ReverseSlice function
// Create reversed copy without using built-in reverse
func ReverseSlice(slice []string) []string {
	// Fix me! Iterate backwards using range
	return nil
}

// TODO: Implement method ForEach on NumberRange
// Call the given function for each number in the range
func (nr NumberRange) ForEach(fn func(int)) {
	// Fix me! Iterate from Start to End, call fn for each
}

// TODO: Implement method ToSlice on NumberRange
// Return slice containing all numbers in the range
func (nr NumberRange) ToSlice() []int {
	// Fix me! Create slice with numbers from Start to End
	return nil
}

// TODO: Implement method Filter on NumberRange
// Return new slice with numbers that satisfy the predicate
func (nr NumberRange) Filter(predicate func(int) bool) []int {
	// Fix me! Use range and predicate function
	return nil
}

// TODO: Implement EnumerateSlice function
// Return slice of index-value pairs
func EnumerateSlice(slice []string) []struct {
	Index int
	Value string
} {
	// Fix me! Use range with index and value
	return nil
}

// TODO: Implement WordCount function
// Count occurrences of each word in a slice
func WordCount(words []string) map[string]int {
	// Fix me! Use range and map
	return nil
}

func main() {
	// Test sum slice
	numbers := []int{1, 2, 3, 4, 5}
	sum := SumSlice(numbers)
	fmt.Printf("Sum of %v: %d\n", numbers, sum)

	// Test count vowels
	text := "Hello World"
	vowels := CountVowels(text)
	fmt.Printf("Vowels in '%s': %d\n", text, vowels)

	// Test map lookup
	scores := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}
	value, found := FindKeyInMap(scores, "Alice")
	fmt.Printf("Alice's score: %d (found: %v)\n", value, found)

	// Test reverse
	words := []string{"Go", "is", "awesome"}
	reversed := ReverseSlice(words)
	fmt.Printf("Reversed %v: %v\n", words, reversed)

	// Test number range
	nr := NumberRange{Start: 1, End: 5}
	fmt.Print("Numbers in range: ")
	nr.ForEach(func(n int) {
		fmt.Printf("%d ", n)
	})
	fmt.Println()

	// Test range to slice
	slice := nr.ToSlice()
	fmt.Printf("Range as slice: %v\n", slice)

	// Test filter
	evens := nr.Filter(func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers in range: %v\n", evens)
}
