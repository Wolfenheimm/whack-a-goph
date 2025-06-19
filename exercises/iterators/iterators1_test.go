package main

import (
	"reflect"
	"testing"
)

func TestSumSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{10, 20, 30}, 60},
		{[]int{}, 0},
		{[]int{-1, 1, -2, 2}, 0},
	}

	for _, test := range tests {
		result := SumSlice(test.input)
		if result != test.expected {
			t.Errorf("SumSlice(%v) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello World", 3}, // e, o, o
		{"aeiou", 5},
		{"bcdfg", 0},
		{"Programming", 3}, // o, a, i
		{"", 0},
	}

	for _, test := range tests {
		result := CountVowels(test.input)
		if result != test.expected {
			t.Errorf("CountVowels(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestFindKeyInMap(t *testing.T) {
	m := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}

	// Test existing key
	value, found := FindKeyInMap(m, "Alice")
	if !found {
		t.Errorf("FindKeyInMap should find existing key 'Alice'")
	}
	if value != 95 {
		t.Errorf("FindKeyInMap('Alice') = %d; want 95", value)
	}

	// Test non-existing key
	value, found = FindKeyInMap(m, "David")
	if found {
		t.Errorf("FindKeyInMap should not find non-existing key 'David'")
	}
}

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"Go", "is", "awesome"}, []string{"awesome", "is", "Go"}},
		{[]string{}, []string{}},
		{[]string{"single"}, []string{"single"}},
	}

	for _, test := range tests {
		result := ReverseSlice(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ReverseSlice(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestNumberRangeForEach(t *testing.T) {
	nr := NumberRange{Start: 1, End: 5}
	var collected []int

	nr.ForEach(func(n int) {
		collected = append(collected, n)
	})

	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(collected, expected) {
		t.Errorf("NumberRange.ForEach collected %v; want %v", collected, expected)
	}
}

func TestNumberRangeToSlice(t *testing.T) {
	tests := []struct {
		nr       NumberRange
		expected []int
	}{
		{NumberRange{1, 5}, []int{1, 2, 3, 4, 5}},
		{NumberRange{0, 3}, []int{0, 1, 2, 3}},
		{NumberRange{10, 12}, []int{10, 11, 12}},
	}

	for _, test := range tests {
		result := test.nr.ToSlice()
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("NumberRange(%d, %d).ToSlice() = %v; want %v",
				test.nr.Start, test.nr.End, result, test.expected)
		}
	}
}

func TestNumberRangeFilter(t *testing.T) {
	nr := NumberRange{Start: 1, End: 10}

	// Test even numbers
	evens := nr.Filter(func(n int) bool { return n%2 == 0 })
	expectedEvens := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(evens, expectedEvens) {
		t.Errorf("Filter(even) = %v; want %v", evens, expectedEvens)
	}

	// Test numbers > 5
	greater := nr.Filter(func(n int) bool { return n > 5 })
	expectedGreater := []int{6, 7, 8, 9, 10}
	if !reflect.DeepEqual(greater, expectedGreater) {
		t.Errorf("Filter(>5) = %v; want %v", greater, expectedGreater)
	}
}

func TestEnumerateSlice(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := EnumerateSlice(input)

	if len(result) != 3 {
		t.Errorf("EnumerateSlice should return 3 items; got %d", len(result))
		return
	}

	expected := []struct {
		Index int
		Value string
	}{
		{0, "a"},
		{1, "b"},
		{2, "c"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("EnumerateSlice(%v) = %v; want %v", input, result, expected)
	}
}

func TestWordCount(t *testing.T) {
	words := []string{"go", "is", "great", "go", "go", "is"}
	result := WordCount(words)

	expected := map[string]int{
		"go":    3,
		"is":    2,
		"great": 1,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("WordCount(%v) = %v; want %v", words, result, expected)
	}
}
