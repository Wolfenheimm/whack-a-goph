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
		{
			Name:        "hello",
			Path:        "exercises/01_hello/hello.go",
			Description: "Get started with a simple hello function",
			Hint:        "Return the string 'Hello, Gopher!' from the Hello() function",
			Topic:       "basics",
		},
		{
			Name:        "basics",
			Path:        "exercises/02_basics/basics.go",
			Description: "Learn about variables and basic types",
			Hint:        "Declare variables with proper types and assign correct values",
			Topic:       "basics",
		},
		{
			Name:        "error_handling",
			Path:        "exercises/03_error_handling/error_handling.go",
			Description: "Learn to handle errors in Go",
			Hint:        "Use error handling patterns to safely handle potential failures",
			Topic:       "basics",
		},
		{
			Name:        "control_flow",
			Path:        "exercises/04_control_flow/control_flow.go",
			Description: "Use if statements and loops",
			Hint:        "Implement conditional logic and loops to solve the problems",
			Topic:       "control_flow",
		},
		{
			Name:        "functions",
			Path:        "exercises/05_functions/functions.go",
			Description: "Write functions with parameters and return values",
			Hint:        "Define functions that take parameters and return the expected values",
			Topic:       "functions",
		},
		{
			Name:        "structs",
			Path:        "exercises/06_structs/structs.go",
			Description: "Define and use structs",
			Hint:        "Create struct types and methods to organize your data",
			Topic:       "structs",
		},
		{
			Name:        "methods",
			Path:        "exercises/07_methods/methods.go",
			Description: "Define methods on types",
			Hint:        "Add methods to your types to define their behavior",
			Topic:       "structs",
		},
		{
			Name:        "testing",
			Path:        "exercises/08_testing/testing.go",
			Description: "Write tests for your Go code",
			Hint:        "Use the testing package to write unit tests for your functions",
			Topic:       "testing",
		},
		{
			Name:        "pointers",
			Path:        "exercises/09_pointers/pointers.go",
			Description: "Work with pointers in Go",
			Hint:        "Use pointers to modify values and work with memory efficiently",
			Topic:       "basics",
		},
		{
			Name:        "packages",
			Path:        "exercises/10_packages/packages.go",
			Description: "Organize code into packages",
			Hint:        "Structure your code into reusable packages with proper exports",
			Topic:       "packages",
		},
		{
			Name:        "strings",
			Path:        "exercises/11_strings/strings.go",
			Description: "Work with strings and text processing",
			Hint:        "Use the strings package to manipulate and process text data",
			Topic:       "standard_library",
		},
		{
			Name:        "file_io",
			Path:        "exercises/12_file_io/file_io.go",
			Description: "Read and write files",
			Hint:        "Use os and io packages to work with files and directories",
			Topic:       "standard_library",
		},
		{
			Name:        "json",
			Path:        "exercises/13_json/json.go",
			Description: "Encode and decode JSON data",
			Hint:        "Use the encoding/json package to work with JSON data",
			Topic:       "standard_library",
		},
		{
			Name:        "http_client",
			Path:        "exercises/14_http_client/http_client.go",
			Description: "Make HTTP requests and handle responses",
			Hint:        "Use the net/http package to make HTTP requests and parse responses",
			Topic:       "networking",
		},
		{
			Name:        "context",
			Path:        "exercises/15_context/context.go",
			Description: "Use context for cancellation and timeouts",
			Hint:        "Use the context package to manage request lifecycles and timeouts",
			Topic:       "advanced",
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
