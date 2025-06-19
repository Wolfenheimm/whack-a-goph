package quizzes

// Quiz 2: Data Structures and Control Flow 🐹
// This quiz covers structs, strings, if statements, and primitive types.
//
// You need to implement functions that work with structs and control flow.

// I AM NOT DONE

// TODO: Define a struct called Student with fields:
// - Name (string)
// - Age (int)
// - Grade (float64)
// - IsActive (bool)
type Student struct {
	Name     string  // Fix me! Make sure this has the right type
	Age      int     // Fix me! Make sure this has the right type
	Grade    float64 // Fix me! Make sure this has the right type
	IsActive bool    // Fix me! Make sure this has the right type
}

// TODO: Implement the CreateStudent function
// It should create and return a Student with the given parameters
func CreateStudent(name string, age int, grade float64, isActive bool) Student {
	// Fix me!
	return Student{}
}

// TODO: Implement the GetGradeLetter function
// Convert numeric grade to letter grade:
// 90-100: "A", 80-89: "B", 70-79: "C", 60-69: "D", below 60: "F"
func GetGradeLetter(grade float64) string {
	// Fix me!
	return ""
}

// TODO: Implement the FormatStudentInfo function
// Return a formatted string: "Name: [name], Age: [age], Grade: [letter] ([numeric])"
// Example: "Name: Alice, Age: 20, Grade: A (95.5)"
func FormatStudentInfo(student Student) string {
	// Fix me!
	return ""
}

// TODO: Implement the IsEligibleForHonors function
// Return true if student is active AND grade >= 85
func IsEligibleForHonors(student Student) bool {
	// Fix me!
	return false
}

// TODO: Implement the CountActiveStudents function
// Count how many students in the slice are active
func CountActiveStudents(students []Student) int {
	// Fix me!
	return 0
}

// TODO: Implement the Quiz2 function
// This function should demonstrate the quiz functionality
func Quiz2() string {
	// TODO: Create students using CreateStudent function
	// Create at least 3 students with different grades and active status
	student1 := Student{} // Fix me! Use CreateStudent
	student2 := Student{} // Fix me! Use CreateStudent
	student3 := Student{} // Fix me! Use CreateStudent

	// TODO: Create a slice of students
	students := []Student{student1, student2, student3} // Fix me!

	result := ""

	// TODO: Print information for each student using FormatStudentInfo
	// Also print if they're eligible for honors
	for _, student := range students {
		result += "FIXME: Print student info and honors eligibility\n"
		_ = student // Remove this line when you implement the loop
	}

	// TODO: Print the total count of active students
	result += "FIXME: Print active student count\n"

	// Final completion message
	result += "Quiz 2 completed!"
	return result
}
