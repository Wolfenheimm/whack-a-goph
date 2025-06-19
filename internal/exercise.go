package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Exercise represents a single exercise
type Exercise struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Hint        string `json:"hint"`
	Completed   bool   `json:"completed"`
	Topic       string `json:"topic"`
}

// Progress tracks the user's progress through exercises
type Progress struct {
	Exercises []Exercise `json:"exercises"`
}

const progressFile = ".whack-a-goph-progress.json"

// GetExercises returns all exercises with their progress
func GetExercises() ([]Exercise, error) {
	exercises := []Exercise{
		// Intro (2 exercises)
		{
			Name:        "intro1",
			Path:        "exercises/intro/intro1.go",
			Description: "Get started with a simple hello function",
			Hint:        "Return the string 'Hello, Gopher!' from the Hello() function",
			Topic:       "intro",
		},
		{
			Name:        "intro2",
			Path:        "exercises/intro/intro2.go",
			Description: "Learn about Go package and imports",
			Hint:        "Add the correct package declaration and import statement",
			Topic:       "intro",
		},

		// Variables (2 exercises)
		{
			Name:        "variables1",
			Path:        "exercises/variables/variables1.go",
			Description: "Learn about variable declarations",
			Hint:        "Declare variables with proper types and assign correct values",
			Topic:       "variables",
		},
		{
			Name:        "variables2",
			Path:        "exercises/variables/variables2.go",
			Description: "Work with different variable types",
			Hint:        "Use var keyword to declare variables of different types",
			Topic:       "variables",
		},

		// Functions (1 exercise)
		{
			Name:        "functions1",
			Path:        "exercises/functions/functions1.go",
			Description: "Basic function definitions",
			Hint:        "Define functions that take parameters and return values",
			Topic:       "functions",
		},

		// Quiz 1
		{
			Name:        "quiz1",
			Path:        "exercises/quizzes/quiz1.go",
			Description: "Quiz: Variables and Functions",
			Hint:        "Combine knowledge from variables and functions sections",
			Topic:       "quiz",
		},

		// If (2 exercises)
		{
			Name:        "if1",
			Path:        "exercises/if/if1.go",
			Description: "Basic if statements",
			Hint:        "Use if statements for conditional logic",
			Topic:       "if",
		},
		{
			Name:        "if2",
			Path:        "exercises/if/if2.go",
			Description: "If with initialization and switch statements",
			Hint:        "if statements can include initialization, switch statements are cleaner than multiple if-else",
			Topic:       "if",
		},

		// Primitive Types (2 exercises)
		{
			Name:        "primitives1",
			Path:        "exercises/primitive_types/primitives1.go",
			Description: "Boolean and numeric types",
			Hint:        "Learn about Go's basic data types",
			Topic:       "primitive_types",
		},
		{
			Name:        "primitives2",
			Path:        "exercises/primitive_types/primitives2.go",
			Description: "String operations",
			Hint:        "Work with strings and string operations",
			Topic:       "primitive_types",
		},

		// Enums (3 exercises)
		{
			Name:        "enums1",
			Path:        "exercises/enums/enums1.go",
			Description: "Constants as enums",
			Hint:        "Use constants with iota to create enum-like types",
			Topic:       "enums",
		},
		{
			Name:        "enums2",
			Path:        "exercises/enums/enums2.go",
			Description: "Custom types for type safety",
			Hint:        "Create custom types for better type safety",
			Topic:       "enums",
		},
		{
			Name:        "enums3",
			Path:        "exercises/enums/enums3.go",
			Description: "String methods for enums",
			Hint:        "Implement String() method for your enum types",
			Topic:       "enums",
		},

		// Vecs (1 exercise)
		{
			Name:        "vecs1",
			Path:        "exercises/vecs/vecs1.go",
			Description: "Creating and working with slices",
			Hint:        "Use make() or slice literals to create slices, use append() for operations",
			Topic:       "vecs",
		},

		// Structs (2 exercises)
		{
			Name:        "structs1",
			Path:        "exercises/structs/structs1.go",
			Description: "Basic struct definitions and usage",
			Hint:        "Define structs and create instances",
			Topic:       "structs",
		},
		{
			Name:        "structs2",
			Path:        "exercises/structs/structs2.go",
			Description: "Struct methods and receivers",
			Hint:        "Add methods to structs using receivers",
			Topic:       "structs",
		},

		// Strings (2 exercises)
		{
			Name:        "strings1",
			Path:        "exercises/strings/strings1.go",
			Description: "String basics and operations",
			Hint:        "Work with string concatenation and basic operations",
			Topic:       "strings",
		},
		{
			Name:        "strings2",
			Path:        "exercises/strings/strings2.go",
			Description: "String formatting and manipulation",
			Hint:        "Use fmt.Sprintf and strings package functions",
			Topic:       "strings",
		},

		// Modules (1 exercise)
		{
			Name:        "modules1",
			Path:        "exercises/modules/modules1.go",
			Description: "Package declarations and imports",
			Hint:        "Organize code into packages and use standard library",
			Topic:       "modules",
		},

		// Hashmaps (1 exercise)
		{
			Name:        "hashmaps1",
			Path:        "exercises/hashmaps/hashmaps1.go",
			Description: "Creating and using maps",
			Hint:        "Use make() or map literals to create maps, iterate with range",
			Topic:       "hashmaps",
		},

		// Quiz 2
		{
			Name:        "quiz2",
			Path:        "exercises/quizzes/quiz2.go",
			Description: "Quiz: Data Structures",
			Hint:        "Combine knowledge from structs, slices, and maps",
			Topic:       "quiz",
		},

		// Options (1 exercise)
		{
			Name:        "options1",
			Path:        "exercises/options/options1.go",
			Description: "Working with pointers and nil",
			Hint:        "Use pointers to represent optional values",
			Topic:       "options",
		},

		// Error Handling (1 exercise)
		{
			Name:        "errors1",
			Path:        "exercises/error_handling/errors1.go",
			Description: "Basic error handling",
			Hint:        "Use error return values for error handling",
			Topic:       "error_handling",
		},

		// Generics (1 exercise)
		{
			Name:        "generics1",
			Path:        "exercises/generics/generics1.go",
			Description: "Basic generic functions",
			Hint:        "Use type parameters to write generic functions",
			Topic:       "generics",
		},

		// Traits/Interfaces (1 exercise)
		{
			Name:        "traits1",
			Path:        "exercises/traits/traits1.go",
			Description: "Basic interface definitions",
			Hint:        "Define and implement simple interfaces",
			Topic:       "traits",
		},

		// Quiz 3
		{
			Name:        "quiz3",
			Path:        "exercises/quizzes/quiz3.go",
			Description: "Quiz: Interfaces and Error Handling",
			Hint:        "Combine interfaces, error handling, and generics",
			Topic:       "quiz",
		},

		// Tests (1 exercise)
		{
			Name:        "tests1",
			Path:        "exercises/tests/tests1.go",
			Description: "Basic unit testing",
			Hint:        "Write tests using the testing package",
			Topic:       "tests",
		},

		// Iterators (1 exercise)
		{
			Name:        "iterators1",
			Path:        "exercises/iterators/iterators1.go",
			Description: "Range and iteration patterns",
			Hint:        "Use range for iteration over various types",
			Topic:       "iterators",
		},

		// Smart Pointers (1 exercise)
		{
			Name:        "smart_pointers1",
			Path:        "exercises/smart_pointers/smart_pointers1.go",
			Description: "Pointers and memory management",
			Hint:        "Understand pointer basics and memory layout",
			Topic:       "smart_pointers",
		},

		// Threads (1 exercise)
		{
			Name:        "threads1",
			Path:        "exercises/threads/threads1.go",
			Description: "Basic goroutines",
			Hint:        "Start goroutines with the go keyword",
			Topic:       "threads",
		},

		// Macros (1 exercise)
		{
			Name:        "macros1",
			Path:        "exercises/macros/macros1.go",
			Description: "go:generate directive and code generation",
			Hint:        "Use //go:generate for code generation",
			Topic:       "macros",
		},

		// Advanced (2 exercises)
		{
			Name:        "advanced1",
			Path:        "exercises/advanced/advanced1.go",
			Description: "Reflection and runtime type information",
			Hint:        "Use reflect package for runtime type inspection",
			Topic:       "advanced",
		},
		{
			Name:        "advanced2",
			Path:        "exercises/advanced/advanced2.go",
			Description: "Context and cancellation",
			Hint:        "Use context package for cancellation and timeouts",
			Topic:       "advanced",
		},

		// Conversions (1 exercise)
		{
			Name:        "conversions1",
			Path:        "exercises/conversions/conversions1.go",
			Description: "Type conversions and casting",
			Hint:        "Convert between compatible types",
			Topic:       "conversions",
		},

		// Quiz 4
		{
			Name:        "quiz4",
			Path:        "exercises/quizzes/quiz4.go",
			Description: "Quiz: Advanced Go Concepts",
			Hint:        "Combine concurrency, generics, reflection, and advanced patterns",
			Topic:       "quiz",
		},

		// Final Quiz
		{
			Name:        "quiz_final",
			Path:        "exercises/quizzes/quiz_final.go",
			Description: "🐹 FINAL QUIZ: Master of Go! Complete task management system",
			Hint:        "Build a comprehensive system using all Go concepts you've learned",
			Topic:       "quiz",
		},
	}

	// Load progress
	progress, err := LoadProgress()
	if err != nil {
		// If no progress file exists, create it
		progress = &Progress{Exercises: exercises}
		SaveProgress(progress)
	} else {
		// Update exercise list with saved progress
		progressMap := make(map[string]bool)
		for _, ex := range progress.Exercises {
			progressMap[ex.Name] = ex.Completed
		}

		for i := range exercises {
			if completed, exists := progressMap[exercises[i].Name]; exists {
				exercises[i].Completed = completed
			}
		}

		// Update progress with any new exercises
		progress.Exercises = exercises
		SaveProgress(progress)
	}

	return exercises, nil
}

// LoadProgress loads the user's progress from disk
func LoadProgress() (*Progress, error) {
	data, err := os.ReadFile(progressFile)
	if err != nil {
		return nil, err
	}

	var progress Progress
	err = json.Unmarshal(data, &progress)
	return &progress, err
}

// SaveProgress saves the user's progress to disk
func SaveProgress(progress *Progress) error {
	data, err := json.MarshalIndent(progress, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(progressFile, data, 0644)
}

// RunExercise runs the tests for a specific exercise
func RunExercise(exercise Exercise) (bool, string, error) {
	// Check if exercise file exists
	if _, err := os.Stat(exercise.Path); os.IsNotExist(err) {
		return false, "", fmt.Errorf("exercise file not found: %s", exercise.Path)
	}

	// Run the tests
	dir := filepath.Dir(exercise.Path)
	cmd := exec.Command("go", "test", "-v", ".")
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	outputStr := string(output)

	if err != nil {
		// Test failed
		return false, outputStr, nil
	}

	// Test passed
	return true, outputStr, nil
}

// MarkCompleted marks an exercise as completed
func MarkCompleted(exerciseName string) error {
	exercises, err := GetExercises()
	if err != nil {
		return err
	}

	for i := range exercises {
		if exercises[i].Name == exerciseName {
			exercises[i].Completed = true
			break
		}
	}

	progress := &Progress{Exercises: exercises}
	return SaveProgress(progress)
}

// ResetExercise resets an exercise's progress
func ResetExercise(exerciseName string) error {
	exercises, err := GetExercises()
	if err != nil {
		return err
	}

	for i := range exercises {
		if exercises[i].Name == exerciseName {
			exercises[i].Completed = false
			break
		}
	}

	progress := &Progress{Exercises: exercises}
	return SaveProgress(progress)
}

// GetNextExercise returns the next incomplete exercise
func GetNextExercise() (*Exercise, error) {
	exercises, err := GetExercises()
	if err != nil {
		return nil, err
	}

	for _, exercise := range exercises {
		if !exercise.Completed {
			return &exercise, nil
		}
	}

	return nil, fmt.Errorf("all exercises completed! 🎉")
}

// FindExercise finds an exercise by name
func FindExercise(name string) (*Exercise, error) {
	exercises, err := GetExercises()
	if err != nil {
		return nil, err
	}

	name = strings.ToLower(strings.TrimSpace(name))
	for _, exercise := range exercises {
		if strings.ToLower(exercise.Name) == name {
			return &exercise, nil
		}
	}

	return nil, fmt.Errorf("exercise not found: %s", name)
}
