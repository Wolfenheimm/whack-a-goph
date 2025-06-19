package quizzes

// Quiz 1: Variables and Functions 🐹
// This quiz combines knowledge from the variables and functions sections.
//
// You need to implement a function that processes user information
// and declare variables to store the results.

// I AM NOT DONE

import "fmt"

// TODO: Implement the CalculateAge function
// It should take a birth year (int) and current year (int)
// and return the calculated age (int)
func CalculateAge(birthYear, currentYear int) int {
	// Fix me!
	return 0
}

// TODO: Implement the FormatName function
// It should take a first name and last name (both strings)
// and return the full name in the format "Last, First"
func FormatName(firstName, lastName string) string {
	// Fix me!
	return ""
}

// TODO: Implement the Quiz1Demo function
// This function should demonstrate the quiz functionality
// It should declare variables and use the functions above
func Quiz1() string {
	// TODO: Declare variables for the following:
	// - currentYear with value 2024
	// - birthYear with value 1990
	// - firstName with value "Jane"
	// - lastName with value "Doe"
	currentYear := 0 // Fix me!
	birthYear := 0   // Fix me!
	firstName := ""  // Fix me!
	lastName := ""   // Fix me!

	// TODO: Format the name using your FormatName function
	// Store the result in a variable called 'fullName'
	fullName := "FIX ME!" // Fix me! Use FormatName(firstName, lastName)

	// TODO: Calculate the age using your CalculateAge function
	// Store the result in a variable called 'age'
	age := 0 // Fix me! Use CalculateAge(birthYear, currentYear)

	// Remove these lines when you implement the TODOs above
	_ = currentYear
	_ = birthYear
	_ = firstName
	_ = lastName

	// Return formatted result for testing
	return fmt.Sprintf("Name: %s\nAge: %d", fullName, age)
}
