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
		// Intro (Getting started)
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

		// Variables (6 exercises)
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
		{
			Name:        "variables3",
			Path:        "exercises/variables/variables3.go",
			Description: "Short variable declarations with :=",
			Hint:        "Use := operator for short variable declarations",
			Topic:       "variables",
		},
		{
			Name:        "variables4",
			Path:        "exercises/variables/variables4.go",
			Description: "Variable scope and shadowing",
			Hint:        "Understand how variable scope works in Go",
			Topic:       "variables",
		},
		{
			Name:        "variables5",
			Path:        "exercises/variables/variables5.go",
			Description: "Zero values and variable initialization",
			Hint:        "Learn about Go's zero values for different types",
			Topic:       "variables",
		},
		{
			Name:        "variables6",
			Path:        "exercises/variables/variables6.go",
			Description: "Constants and iota",
			Hint:        "Use const keyword and iota for constants",
			Topic:       "variables",
		},

		// Functions (5 exercises)
		{
			Name:        "functions1",
			Path:        "exercises/functions/functions1.go",
			Description: "Basic function definitions",
			Hint:        "Define functions that take parameters and return values",
			Topic:       "functions",
		},
		{
			Name:        "functions2",
			Path:        "exercises/functions/functions2.go",
			Description: "Multiple return values",
			Hint:        "Functions can return multiple values in Go",
			Topic:       "functions",
		},
		{
			Name:        "functions3",
			Path:        "exercises/functions/functions3.go",
			Description: "Named return values",
			Hint:        "Use named return values for cleaner code",
			Topic:       "functions",
		},
		{
			Name:        "functions4",
			Path:        "exercises/functions/functions4.go",
			Description: "Variadic functions",
			Hint:        "Functions can accept variable number of arguments",
			Topic:       "functions",
		},
		{
			Name:        "functions5",
			Path:        "exercises/functions/functions5.go",
			Description: "Function types and closures",
			Hint:        "Functions are first-class values in Go",
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

		// If (3 exercises)
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
			Description: "If with initialization",
			Hint:        "if statements can include initialization",
			Topic:       "if",
		},
		{
			Name:        "if3",
			Path:        "exercises/if/if3.go",
			Description: "Switch statements",
			Hint:        "Switch statements are cleaner than multiple if-else",
			Topic:       "if",
		},

		// Primitive Types (6 exercises)
		{
			Name:        "primitive_types1",
			Path:        "exercises/primitive_types/primitive_types1.go",
			Description: "Boolean and numeric types",
			Hint:        "Learn about Go's basic data types",
			Topic:       "primitive_types",
		},
		{
			Name:        "primitive_types2",
			Path:        "exercises/primitive_types/primitive_types2.go",
			Description: "String operations",
			Hint:        "Work with strings and string operations",
			Topic:       "primitive_types",
		},
		{
			Name:        "primitive_types3",
			Path:        "exercises/primitive_types/primitive_types3.go",
			Description: "Arrays and their properties",
			Hint:        "Arrays have fixed size and are value types",
			Topic:       "primitive_types",
		},
		{
			Name:        "primitive_types4",
			Path:        "exercises/primitive_types/primitive_types4.go",
			Description: "Slices and slice operations",
			Hint:        "Slices are dynamic arrays in Go",
			Topic:       "primitive_types",
		},
		{
			Name:        "primitive_types5",
			Path:        "exercises/primitive_types/primitive_types5.go",
			Description: "Working with tuples",
			Hint:        "Go doesn't have tuples, but you can return multiple values",
			Topic:       "primitive_types",
		},
		{
			Name:        "primitive_types6",
			Path:        "exercises/primitive_types/primitive_types6.go",
			Description: "Type conversions",
			Hint:        "Convert between different numeric types",
			Topic:       "primitive_types",
		},

		// Enums (3 exercises) - Go's closest equivalent
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

		// Vecs (2 exercises)
		{
			Name:        "vecs1",
			Path:        "exercises/vecs/vecs1.go",
			Description: "Creating and initializing slices",
			Hint:        "Use make() or slice literals to create slices",
			Topic:       "vecs",
		},
		{
			Name:        "vecs2",
			Path:        "exercises/vecs/vecs2.go",
			Description: "Slice operations and methods",
			Hint:        "Use append(), copy(), and other slice operations",
			Topic:       "vecs",
		},

		// Structs (3 exercises)
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
		{
			Name:        "structs3",
			Path:        "exercises/structs/structs3.go",
			Description: "Embedded structs and composition",
			Hint:        "Use struct embedding for composition",
			Topic:       "structs",
		},

		// Strings (4 exercises)
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
			Description: "String formatting with fmt",
			Hint:        "Use fmt.Sprintf and other formatting functions",
			Topic:       "strings",
		},
		{
			Name:        "strings3",
			Path:        "exercises/strings/strings3.go",
			Description: "String manipulation with strings package",
			Hint:        "Use strings package for manipulation functions",
			Topic:       "strings",
		},
		{
			Name:        "strings4",
			Path:        "exercises/strings/strings4.go",
			Description: "Runes and Unicode handling",
			Hint:        "Work with runes for proper Unicode support",
			Topic:       "strings",
		},

		// Modules (3 exercises)
		{
			Name:        "modules1",
			Path:        "exercises/modules/modules1.go",
			Description: "Package declarations and imports",
			Hint:        "Organize code into packages",
			Topic:       "modules",
		},
		{
			Name:        "modules2",
			Path:        "exercises/modules/modules2.go",
			Description: "Exported vs unexported identifiers",
			Hint:        "Use capitalization to control visibility",
			Topic:       "modules",
		},
		{
			Name:        "modules3",
			Path:        "exercises/modules/modules3.go",
			Description: "Package initialization and init functions",
			Hint:        "Use init() functions for package initialization",
			Topic:       "modules",
		},

		// Hashmaps (3 exercises) - Go's maps
		{
			Name:        "hashmaps1",
			Path:        "exercises/hashmaps/hashmaps1.go",
			Description: "Creating and using maps",
			Hint:        "Use make() or map literals to create maps",
			Topic:       "hashmaps",
		},
		{
			Name:        "hashmaps2",
			Path:        "exercises/hashmaps/hashmaps2.go",
			Description: "Map operations and iteration",
			Hint:        "Use range to iterate over maps",
			Topic:       "hashmaps",
		},
		{
			Name:        "hashmaps3",
			Path:        "exercises/hashmaps/hashmaps3.go",
			Description: "Advanced map patterns",
			Hint:        "Check for key existence and handle zero values",
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

		// Options (3 exercises) - Go's pointers and nil
		{
			Name:        "options1",
			Path:        "exercises/options/options1.go",
			Description: "Working with pointers and nil",
			Hint:        "Use pointers to represent optional values",
			Topic:       "options",
		},
		{
			Name:        "options2",
			Path:        "exercises/options/options2.go",
			Description: "Optional parameters patterns",
			Hint:        "Use functional options pattern",
			Topic:       "options",
		},
		{
			Name:        "options3",
			Path:        "exercises/options/options3.go",
			Description: "Interface{} for optional types",
			Hint:        "Use empty interface for generic optional values",
			Topic:       "options",
		},

		// Error Handling (6 exercises)
		{
			Name:        "errors1",
			Path:        "exercises/error_handling/errors1.go",
			Description: "Basic error handling",
			Hint:        "Use error return values for error handling",
			Topic:       "error_handling",
		},
		{
			Name:        "errors2",
			Path:        "exercises/error_handling/errors2.go",
			Description: "Creating custom errors",
			Hint:        "Implement the error interface",
			Topic:       "error_handling",
		},
		{
			Name:        "errors3",
			Path:        "exercises/error_handling/errors3.go",
			Description: "Error wrapping and unwrapping",
			Hint:        "Use fmt.Errorf with %w to wrap errors",
			Topic:       "error_handling",
		},
		{
			Name:        "errors4",
			Path:        "exercises/error_handling/errors4.go",
			Description: "Panic and recover",
			Hint:        "Use panic and recover for exceptional cases",
			Topic:       "error_handling",
		},
		{
			Name:        "errors5",
			Path:        "exercises/error_handling/errors5.go",
			Description: "Error checking patterns",
			Hint:        "Learn common error checking patterns",
			Topic:       "error_handling",
		},
		{
			Name:        "errors6",
			Path:        "exercises/error_handling/errors6.go",
			Description: "Errors package and stack traces",
			Hint:        "Use github.com/pkg/errors for enhanced error handling",
			Topic:       "error_handling",
		},

		// Generics (2 exercises) - Go 1.18+
		{
			Name:        "generics1",
			Path:        "exercises/generics/generics1.go",
			Description: "Basic generic functions",
			Hint:        "Use type parameters to write generic functions",
			Topic:       "generics",
		},
		{
			Name:        "generics2",
			Path:        "exercises/generics/generics2.go",
			Description: "Generic types and constraints",
			Hint:        "Define generic types with appropriate constraints",
			Topic:       "generics",
		},

		// Traits (5 exercises) - Go's interfaces
		{
			Name:        "traits1",
			Path:        "exercises/traits/traits1.go",
			Description: "Basic interface definitions",
			Hint:        "Define and implement simple interfaces",
			Topic:       "traits",
		},
		{
			Name:        "traits2",
			Path:        "exercises/traits/traits2.go",
			Description: "Interface composition",
			Hint:        "Compose interfaces from other interfaces",
			Topic:       "traits",
		},
		{
			Name:        "traits3",
			Path:        "exercises/traits/traits3.go",
			Description: "Empty interface and type assertions",
			Hint:        "Use interface{} and type assertions",
			Topic:       "traits",
		},
		{
			Name:        "traits4",
			Path:        "exercises/traits/traits4.go",
			Description: "Interface polymorphism",
			Hint:        "Use interfaces for polymorphic behavior",
			Topic:       "traits",
		},
		{
			Name:        "traits5",
			Path:        "exercises/traits/traits5.go",
			Description: "Interface best practices",
			Hint:        "Follow Go's interface design principles",
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

		// Tests (4 exercises)
		{
			Name:        "tests1",
			Path:        "exercises/tests/tests1.go",
			Description: "Basic unit testing",
			Hint:        "Write tests using the testing package",
			Topic:       "tests",
		},
		{
			Name:        "tests2",
			Path:        "exercises/tests/tests2.go",
			Description: "Table-driven tests",
			Hint:        "Use test tables for comprehensive testing",
			Topic:       "tests",
		},
		{
			Name:        "tests3",
			Path:        "exercises/tests/tests3.go",
			Description: "Benchmarks and examples",
			Hint:        "Write benchmarks and example functions",
			Topic:       "tests",
		},
		{
			Name:        "tests4",
			Path:        "exercises/tests/tests4.go",
			Description: "Test helpers and mocking",
			Hint:        "Create test helpers and mock interfaces",
			Topic:       "tests",
		},

		// Standard Library Types (5 exercises)
		{
			Name:        "iterators1",
			Path:        "exercises/iterators/iterators1.go",
			Description: "Range and iteration patterns",
			Hint:        "Use range for iteration over various types",
			Topic:       "iterators",
		},
		{
			Name:        "iterators2",
			Path:        "exercises/iterators/iterators2.go",
			Description: "Custom iteration with channels",
			Hint:        "Use channels to create custom iterators",
			Topic:       "iterators",
		},
		{
			Name:        "iterators3",
			Path:        "exercises/iterators/iterators3.go",
			Description: "Iterator pattern with interfaces",
			Hint:        "Implement iterator pattern using interfaces",
			Topic:       "iterators",
		},
		{
			Name:        "iterators4",
			Path:        "exercises/iterators/iterators4.go",
			Description: "Functional programming patterns",
			Hint:        "Implement map, filter, reduce patterns",
			Topic:       "iterators",
		},
		{
			Name:        "iterators5",
			Path:        "exercises/iterators/iterators5.go",
			Description: "Generator functions",
			Hint:        "Create generator functions using goroutines",
			Topic:       "iterators",
		},

		// Smart Pointers (4 exercises) - Go's reference types
		{
			Name:        "smart_pointers1",
			Path:        "exercises/smart_pointers/smart_pointers1.go",
			Description: "Pointers and memory management",
			Hint:        "Understand pointer basics and memory layout",
			Topic:       "smart_pointers",
		},
		{
			Name:        "smart_pointers2",
			Path:        "exercises/smart_pointers/smart_pointers2.go",
			Description: "Reference counting patterns",
			Hint:        "Implement reference counting manually",
			Topic:       "smart_pointers",
		},
		{
			Name:        "smart_pointers3",
			Path:        "exercises/smart_pointers/smart_pointers3.go",
			Description: "Weak references and cycles",
			Hint:        "Avoid circular references in data structures",
			Topic:       "smart_pointers",
		},
		{
			Name:        "smart_pointers4",
			Path:        "exercises/smart_pointers/smart_pointers4.go",
			Description: "Memory pools and object reuse",
			Hint:        "Use sync.Pool for object reuse",
			Topic:       "smart_pointers",
		},

		// Threads (3 exercises) - Go's concurrency
		{
			Name:        "threads1",
			Path:        "exercises/threads/threads1.go",
			Description: "Basic goroutines",
			Hint:        "Start goroutines with the go keyword",
			Topic:       "threads",
		},
		{
			Name:        "threads2",
			Path:        "exercises/threads/threads2.go",
			Description: "Channel communication",
			Hint:        "Use channels to communicate between goroutines",
			Topic:       "threads",
		},
		{
			Name:        "threads3",
			Path:        "exercises/threads/threads3.go",
			Description: "Synchronization with sync package",
			Hint:        "Use sync.Mutex, sync.WaitGroup for synchronization",
			Topic:       "threads",
		},

		// Macros (4 exercises) - Go's code generation
		{
			Name:        "macros1",
			Path:        "exercises/macros/macros1.go",
			Description: "go:generate directive",
			Hint:        "Use //go:generate for code generation",
			Topic:       "macros",
		},
		{
			Name:        "macros2",
			Path:        "exercises/macros/macros2.go",
			Description: "Build tags and conditional compilation",
			Hint:        "Use build tags for conditional compilation",
			Topic:       "macros",
		},
		{
			Name:        "macros3",
			Path:        "exercises/macros/macros3.go",
			Description: "Text/template for code generation",
			Hint:        "Use text/template to generate Go code",
			Topic:       "macros",
		},
		{
			Name:        "macros4",
			Path:        "exercises/macros/macros4.go",
			Description: "AST manipulation",
			Hint:        "Use go/ast package for code analysis and generation",
			Topic:       "macros",
		},

		// Advanced (3 exercises) - Advanced Go concepts
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
		{
			Name:        "advanced3",
			Path:        "exercises/advanced/advanced3.go",
			Description: "Performance optimization techniques",
			Hint:        "Apply memory optimization and profiling techniques",
			Topic:       "advanced",
		},

		// Conversions (5 exercises)
		{
			Name:        "conversions1",
			Path:        "exercises/conversions/conversions1.go",
			Description: "Type conversions and casting",
			Hint:        "Convert between compatible types",
			Topic:       "conversions",
		},
		{
			Name:        "conversions2",
			Path:        "exercises/conversions/conversions2.go",
			Description: "String conversions",
			Hint:        "Convert between strings and other types",
			Topic:       "conversions",
		},
		{
			Name:        "conversions3",
			Path:        "exercises/conversions/conversions3.go",
			Description: "JSON marshaling and unmarshaling",
			Hint:        "Use encoding/json for data conversion",
			Topic:       "conversions",
		},
		{
			Name:        "conversions4",
			Path:        "exercises/conversions/conversions4.go",
			Description: "Custom marshaling interfaces",
			Hint:        "Implement json.Marshaler and json.Unmarshaler",
			Topic:       "conversions",
		},
		{
			Name:        "conversions5",
			Path:        "exercises/conversions/conversions5.go",
			Description: "Binary encoding and decoding",
			Hint:        "Use encoding/binary for binary data",
			Topic:       "conversions",
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
