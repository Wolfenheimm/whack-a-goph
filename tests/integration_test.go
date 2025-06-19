package tests

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/raithbheart/whack-a-goph/internal"
)

// TestEndToEndWorkflow tests the complete user workflow
func TestEndToEndWorkflow(t *testing.T) {
	// Save original progress file if it exists
	progressFile := ".whack-a-goph-progress.json"
	backupFile := progressFile + ".backup"

	if _, err := os.Stat(progressFile); err == nil {
		err = os.Rename(progressFile, backupFile)
		if err != nil {
			t.Fatalf("Failed to backup progress file: %v", err)
		}
		defer func() {
			os.Remove(progressFile)
			os.Rename(backupFile, progressFile)
		}()
	} else {
		defer os.Remove(progressFile)
	}

	// Test CLI commands
	t.Run("CLI_Commands", func(t *testing.T) {
		testCLICommands(t)
	})

	// Test exercise loading
	t.Run("Exercise_Loading", func(t *testing.T) {
		testExerciseLoading(t)
	})

	// Test progress tracking
	t.Run("Progress_Tracking", func(t *testing.T) {
		testProgressTracking(t)
	})
}

func testCLICommands(t *testing.T) {
	// Change to the parent directory to build from the root
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	// Go to the project root (parent of tests directory)
	if strings.HasSuffix(originalDir, "/tests") {
		err = os.Chdir("..")
		if err != nil {
			t.Fatalf("Failed to change to parent directory: %v", err)
		}
		defer os.Chdir(originalDir) // Return to original directory
	}

	// Build the CLI tool first
	cmd := exec.Command("go", "build", "-o", "whack-a-goph-test", ".")
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build CLI: %v", err)
	}
	defer os.Remove("whack-a-goph-test")

	// Test list command
	t.Run("List_Command", func(t *testing.T) {
		binaryPath := "../whack-a-goph-test"
		if wd, err := os.Getwd(); err == nil && !strings.HasSuffix(wd, "/tests") {
			binaryPath = "./whack-a-goph-test"
		}
		cmd := exec.Command(binaryPath, "list")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("List command failed: %v\nOutput: %s", err, output)
		}

		outputStr := string(output)
		if !strings.Contains(outputStr, "Exercise List") {
			t.Errorf("List command should contain 'Exercise List', got: %s", outputStr)
		}
		if !strings.Contains(outputStr, "intro1") {
			t.Errorf("List command should contain 'intro1' exercise, got: %s", outputStr)
		}
	})

	// Test help command
	t.Run("Help_Command", func(t *testing.T) {
		binaryPath := "../whack-a-goph-test"
		if wd, err := os.Getwd(); err == nil && !strings.HasSuffix(wd, "/tests") {
			binaryPath = "./whack-a-goph-test"
		}
		cmd := exec.Command(binaryPath, "help")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Help command failed: %v", err)
		}

		outputStr := string(output)
		if !strings.Contains(outputStr, "whack-a-goph") {
			t.Errorf("Help command should contain 'whack-a-goph', got: %s", outputStr)
		}
	})
}

func testExerciseLoading(t *testing.T) {
	exercises, err := internal.GetExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	if len(exercises) == 0 {
		t.Fatal("No exercises loaded")
	}

	// Check that we have exercises from different categories
	categories := make(map[string]bool)
	for _, exercise := range exercises {
		// Extract category from exercise name (e.g., "intro1" -> "intro")
		if len(exercise.Name) > 0 {
			category := strings.TrimRight(exercise.Name, "0123456789")
			categories[category] = true
		}
	}

	expectedCategories := []string{"intro", "variables", "functions", "primitives", "if", "structs"}
	for _, expected := range expectedCategories {
		if !categories[expected] {
			t.Errorf("Expected category %s not found in exercises", expected)
		}
	}
}

func testProgressTracking(t *testing.T) {
	// Test initial progress
	progress, err := internal.LoadProgress()
	if err != nil {
		t.Fatalf("Failed to load progress: %v", err)
	}

	// Should start with no completed exercises
	completedCount := 0
	for _, exercise := range progress.Exercises {
		if exercise.Completed {
			completedCount++
		}
	}

	// Test marking an exercise as completed
	if len(progress.Exercises) > 0 {
		firstExercise := progress.Exercises[0].Name
		err = internal.MarkCompleted(firstExercise)
		if err != nil {
			t.Fatalf("Failed to mark exercise as completed: %v", err)
		}

		// Verify it was marked as completed
		updatedProgress, err := internal.LoadProgress()
		if err != nil {
			t.Fatalf("Failed to reload progress: %v", err)
		}

		found := false
		for _, exercise := range updatedProgress.Exercises {
			if exercise.Name == firstExercise && exercise.Completed {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Exercise %s should be marked as completed", firstExercise)
		}

		// Test resetting the exercise
		err = internal.ResetExercise(firstExercise)
		if err != nil {
			t.Fatalf("Failed to reset exercise: %v", err)
		}

		// Verify it was reset
		resetProgress, err := internal.LoadProgress()
		if err != nil {
			t.Fatalf("Failed to reload progress after reset: %v", err)
		}

		for _, exercise := range resetProgress.Exercises {
			if exercise.Name == firstExercise && exercise.Completed {
				t.Errorf("Exercise %s should not be completed after reset", firstExercise)
			}
		}
	}
}

// TestExerciseStructure tests that all exercises have the expected structure
func TestExerciseStructure(t *testing.T) {
	exercises, err := internal.GetExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	for _, exercise := range exercises {
		t.Run(exercise.Name, func(t *testing.T) {
			// Adjust path to be relative to tests directory
			exercisePath := filepath.Join("..", exercise.Path)

			// Check that exercise file exists
			if _, err := os.Stat(exercisePath); os.IsNotExist(err) {
				t.Errorf("Exercise file does not exist: %s", exercisePath)
			}

			// Check that exercise has required fields
			if exercise.Name == "" {
				t.Error("Exercise must have a name")
			}
			if exercise.Description == "" {
				t.Error("Exercise must have a description")
			}
			if exercise.Path == "" {
				t.Error("Exercise must have a path")
			}

			// Check that corresponding test file exists
			dir := filepath.Dir(exercisePath)
			base := filepath.Base(exercisePath)
			testFile := filepath.Join(dir, strings.Replace(base, ".go", "_test.go", 1))

			if _, err := os.Stat(testFile); os.IsNotExist(err) {
				t.Errorf("Test file does not exist: %s", testFile)
			}
		})
	}
}

// TestProgressFilePersistence tests that progress is properly saved and loaded
func TestProgressFilePersistence(t *testing.T) {
	progressFile := ".whack-a-goph-progress-test.json"
	defer os.Remove(progressFile)

	// Create test progress data
	testProgress := internal.Progress{
		Exercises: []internal.Exercise{
			{
				Name:        "test1",
				Path:        "test/path1.go",
				Description: "Test exercise 1",
				Completed:   true,
			},
			{
				Name:        "test2",
				Path:        "test/path2.go",
				Description: "Test exercise 2",
				Completed:   false,
			},
		},
	}

	// Save progress
	data, err := json.MarshalIndent(testProgress, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal progress: %v", err)
	}

	err = os.WriteFile(progressFile, data, 0644)
	if err != nil {
		t.Fatalf("Failed to write progress file: %v", err)
	}

	// Load progress
	loadedData, err := os.ReadFile(progressFile)
	if err != nil {
		t.Fatalf("Failed to read progress file: %v", err)
	}

	var loadedProgress internal.Progress
	err = json.Unmarshal(loadedData, &loadedProgress)
	if err != nil {
		t.Fatalf("Failed to unmarshal progress: %v", err)
	}

	// Verify loaded data matches original
	if len(loadedProgress.Exercises) != len(testProgress.Exercises) {
		t.Errorf("Expected %d exercises, got %d", len(testProgress.Exercises), len(loadedProgress.Exercises))
	}

	for i, exercise := range loadedProgress.Exercises {
		if i < len(testProgress.Exercises) {
			expected := testProgress.Exercises[i]
			if exercise.Name != expected.Name {
				t.Errorf("Exercise %d: expected name %s, got %s", i, expected.Name, exercise.Name)
			}
			if exercise.Completed != expected.Completed {
				t.Errorf("Exercise %d: expected completed %v, got %v", i, expected.Completed, exercise.Completed)
			}
		}
	}
}

// TestConcurrentAccess tests that the progress system handles concurrent access
func TestConcurrentAccess(t *testing.T) {
	// This test ensures that multiple operations on progress don't cause race conditions
	exercises, err := internal.GetExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	if len(exercises) == 0 {
		t.Skip("No exercises available for concurrent test")
	}

	// Run multiple operations concurrently
	done := make(chan bool, 10)
	exerciseName := exercises[0].Name

	for i := 0; i < 5; i++ {
		go func() {
			defer func() { done <- true }()

			// Mark as completed
			internal.MarkCompleted(exerciseName)

			// Load progress
			internal.LoadProgress()

			// Reset
			internal.ResetExercise(exerciseName)
		}()
	}

	// Wait for all goroutines to complete
	timeout := time.After(10 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case <-done:
			// Success
		case <-timeout:
			t.Fatal("Concurrent access test timed out")
		}
	}
}
