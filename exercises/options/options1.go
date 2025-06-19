package options

// Time to learn about pointers and nil! 🐹
// Go uses pointers and nil to represent optional values, similar to Rust's Option.

// I AM NOT DONE

func GetPointer(value int) *int {
	// TODO: Return a pointer to the value
	// Hint: Use &value
	// Fix me!
	return nil
}

func GetValue(ptr *int) int {
	// TODO: Return the value that the pointer points to
	// Assume pointer is not nil
	// Hint: Use *ptr
	// Fix me!
	return 0
}

func IsNil(ptr *int) bool {
	// TODO: Return true if pointer is nil, false otherwise
	// Fix me!
	return false
}

func SafeGetValue(ptr *int, defaultValue int) int {
	// TODO: Return the value that pointer points to if not nil,
	// otherwise return defaultValue
	// Fix me!
	return 0
}

func CreateOptionalString(hasValue bool, value string) *string {
	// TODO: Return a pointer to the string if hasValue is true,
	// otherwise return nil
	// Fix me!
	return nil
}
