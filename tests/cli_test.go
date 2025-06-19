package tests

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/raithbheart/whack-a-goph/internal"
)

// TestCLIBuild tests that the CLI tool can be built successfully
func TestCLIBuild(t *testing.T) {
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

	cmd := exec.Command("go", "build", "-o", "whack-a-goph-test", ".")
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build CLI: %v", err)
	}
	defer os.Remove("whack-a-goph-test")

	// Make the binary executable
	err = os.Chmod("whack-a-goph-test", 0755)
	if err != nil {
		t.Fatalf("Failed to make CLI executable: %v", err)
	}

	// Verify the binary was created
	if _, err := os.Stat("whack-a-goph-test"); os.IsNotExist(err) {
		t.Error("CLI binary was not created")
	}
}

// TestCLIHelp tests the help command and root command
func TestCLIHelp(t *testing.T) {
	buildCLI(t)
	defer cleanupCLI()

	tests := []struct {
		name     string
		args     []string
		expected []string
	}{
		{
			name:     "root_command",
			args:     []string{},
			expected: []string{"Welcome to Whack-a-Goph", "whack-a-goph list", "whack-a-goph next"},
		},
		{
			name:     "help_command",
			args:     []string{"help"},
			expected: []string{"whack-a-goph", "Available Commands", "list", "run", "next"},
		},
		{
			name:     "help_list",
			args:     []string{"help", "list"},
			expected: []string{"List all exercises", "status"},
		},
		{
			name:     "help_run",
			args:     []string{"help", "run"},
			expected: []string{"Run a specific exercise"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := exec.Command(getBinaryPath(), test.args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Command failed: %v\nOutput: %s", err, output)
			}

			outputStr := string(output)
			for _, expected := range test.expected {
				if !strings.Contains(outputStr, expected) {
					t.Errorf("Output should contain %q, got: %s", expected, outputStr)
				}
			}
		})
	}
}

// TestCLIList tests the list command functionality
func TestCLIList(t *testing.T) {
	buildCLI(t)
	defer cleanupCLI()
	setupCleanProgress(t)
	defer cleanupProgress()

	cmd := exec.Command(getBinaryPath(), "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("List command failed: %v\nOutput: %s", err, output)
	}

	outputStr := string(output)

	// Check for expected content
	expectedContent := []string{
		"Exercise List",
		"intro1",
		"Progress:",
		"whack-a-goph next",
	}

	for _, expected := range expectedContent {
		if !strings.Contains(outputStr, expected) {
			t.Errorf("List output should contain %q, got: %s", expected, outputStr)
		}
	}

	// Check for exercise status indicators
	if !strings.Contains(outputStr, "🕳️  HOLE") && !strings.Contains(outputStr, "🎯 WHACKED") {
		t.Error("List should show exercise status indicators")
	}
}

// TestCLINext tests the next command functionality
func TestCLINext(t *testing.T) {
	buildCLI(t)
	defer cleanupCLI()
	setupCleanProgress(t)
	defer cleanupProgress()

	cmd := exec.Command(getBinaryPath(), "next")
	output, _ := cmd.CombinedOutput()

	// Note: This might fail because exercises are incomplete, but should not crash
	outputStr := string(output)

	// Check that it attempts to run the next exercise
	if !strings.Contains(outputStr, "Next exercise:") && !strings.Contains(outputStr, "Error") {
		t.Errorf("Next command should show next exercise or error, got: %s", outputStr)
	}
}

// TestCLIRun tests the run command functionality
func TestCLIRun(t *testing.T) {
	buildCLI(t)
	defer cleanupCLI()
	setupCleanProgress(t)
	defer cleanupProgress()

	// Test running a specific exercise
	cmd := exec.Command(getBinaryPath(), "run", "intro1")
	output, _ := cmd.CombinedOutput()

	outputStr := string(output)

	// Should attempt to run the exercise (might fail due to incomplete code)
	if !strings.Contains(outputStr, "Running exercise:") &&
		!strings.Contains(outputStr, "intro1") &&
		!strings.Contains(outputStr, "Error") {
		t.Errorf("Run command should show exercise info or error, got: %s", outputStr)
	}

	// Test running non-existent exercise
	cmd = exec.Command(getBinaryPath(), "run", "nonexistent")
	output, _ = cmd.CombinedOutput()

	// CLI doesn't return error codes, but should show error message
	outputStr = string(output)
	if !strings.Contains(outputStr, "not found") && !strings.Contains(outputStr, "error") {
		t.Errorf("Running non-existent exercise should show error message, got: %s", outputStr)
	}
}

// TestCLIReset tests the reset command functionality
func TestCLIReset(t *testing.T) {
	buildCLI(t)
	defer cleanupCLI()
	setupCleanProgress(t)
	defer cleanupProgress()

	// First mark an exercise as completed
	err := internal.MarkCompleted("intro1")
	if err != nil {
		t.Fatalf("Failed to mark exercise as completed: %v", err)
	}

	// Test resetting the exercise
	cmd := exec.Command(getBinaryPath(), "reset", "intro1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Reset command failed: %v\nOutput: %s", err, output)
	}

	outputStr := string(output)
	if !strings.Contains(outputStr, "reset") && !strings.Contains(outputStr, "intro1") {
		t.Errorf("Reset output should mention the exercise, got: %s", outputStr)
	}

	// Verify the exercise was actually reset
	progress, err := internal.LoadProgress()
	if err != nil {
		t.Fatalf("Failed to load progress: %v", err)
	}

	for _, exercise := range progress.Exercises {
		if exercise.Name == "intro1" && exercise.Completed {
			t.Error("Exercise should be reset (not completed)")
		}
	}

	// Test resetting non-existent exercise
	cmd = exec.Command(getBinaryPath(), "reset", "nonexistent")
	output, _ = cmd.CombinedOutput()

	// CLI doesn't return error codes, but should show error message
	outputStr = string(output)
	if !strings.Contains(outputStr, "not found") && !strings.Contains(outputStr, "error") {
		t.Errorf("Resetting non-existent exercise should show error message, got: %s", outputStr)
	}
}

// TestCLIHint tests the hint command functionality
func TestCLIHint(t *testing.T) {
	buildCLI(t)
	defer cleanupCLI()

	cmd := exec.Command(getBinaryPath(), "hint", "intro1")
	output, _ := cmd.CombinedOutput()

	outputStr := string(output)

	// Should show hint or "no hint available"
	if !strings.Contains(outputStr, "Hint") && !strings.Contains(outputStr, "hint") {
		t.Errorf("Hint command should mention hints, got: %s", outputStr)
	}
}

// TestCLIWatch tests the watch command (without actually watching)
func TestCLIWatch(t *testing.T) {
	buildCLI(t)
	defer cleanupCLI()
	setupCleanProgress(t)
	defer cleanupProgress()

	// Start watch command in background and kill it quickly
	cmd := exec.Command(getBinaryPath(), "watch", "intro1")

	// Start the command
	err := cmd.Start()
	if err != nil {
		t.Fatalf("Failed to start watch command: %v", err)
	}

	// Give it a moment to start
	time.Sleep(100 * time.Millisecond)

	// Kill the process
	err = cmd.Process.Kill()
	if err != nil {
		t.Fatalf("Failed to kill watch process: %v", err)
	}

	// Wait for it to finish
	cmd.Wait()

	// If we get here without hanging, the watch command started successfully
}

// TestCLIErrorHandling tests various error conditions
func TestCLIErrorHandling(t *testing.T) {
	buildCLI(t)
	defer cleanupCLI()

	errorTests := []struct {
		name        string
		args        []string
		expectError bool
	}{
		{"invalid_command", []string{"invalid"}, true},
		{"run_no_args", []string{"run"}, true},
		{"reset_no_args", []string{"reset"}, true},
		{"hint_no_args", []string{"hint"}, true},
		{"too_many_args", []string{"list", "extra", "args"}, false}, // list ignores extra args
	}

	for _, test := range errorTests {
		t.Run(test.name, func(t *testing.T) {
			cmd := exec.Command(getBinaryPath(), test.args...)
			output, err := cmd.CombinedOutput()
			outputStr := string(output)

			if test.expectError {
				// Should show error message (CLI doesn't use exit codes)
				if err == nil && !strings.Contains(outputStr, "error") && !strings.Contains(outputStr, "Error") {
					t.Errorf("Command %v should show error message, got output: %s", test.args, outputStr)
				}
			} else {
				// Should run successfully
				if err != nil {
					t.Errorf("Command %v should not return error, got: %v", test.args, err)
				}
			}
		})
	}
}

// Helper functions

func getBinaryPath() string {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "./whack-a-goph-test"
	}

	// If we're in the tests directory, the binary is in the parent directory
	if strings.HasSuffix(wd, "/tests") {
		return "../whack-a-goph-test"
	}

	// Otherwise, it's in the current directory
	return "./whack-a-goph-test"
}

func buildCLI(t *testing.T) {
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

	cmd := exec.Command("go", "build", "-o", "whack-a-goph-test", ".")
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build CLI: %v", err)
	}

	// Make the binary executable
	err = os.Chmod("whack-a-goph-test", 0755)
	if err != nil {
		t.Fatalf("Failed to make CLI executable: %v", err)
	}
}

func cleanupCLI() {
	os.Remove("whack-a-goph-test")
}

func setupCleanProgress(t *testing.T) {
	// Back up existing progress if it exists
	progressFile := ".whack-a-goph-progress.json"
	backupFile := progressFile + ".test-backup"

	if _, err := os.Stat(progressFile); err == nil {
		err = os.Rename(progressFile, backupFile)
		if err != nil {
			t.Fatalf("Failed to backup progress file: %v", err)
		}
	}
}

func cleanupProgress() {
	progressFile := ".whack-a-goph-progress.json"
	backupFile := progressFile + ".test-backup"

	// Remove test progress file
	os.Remove(progressFile)

	// Restore backup if it exists
	if _, err := os.Stat(backupFile); err == nil {
		os.Rename(backupFile, progressFile)
	}
}

// TestCLIProgressIntegration tests the integration between CLI commands and progress tracking
func TestCLIProgressIntegration(t *testing.T) {
	buildCLI(t)
	defer cleanupCLI()
	setupCleanProgress(t)
	defer cleanupProgress()

	// Check initial state
	cmd := exec.Command(getBinaryPath(), "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("List command failed: %v", err)
	}

	_ = string(output)

	// Try to complete an exercise (even if it fails, it should create progress file)
	cmd = exec.Command(getBinaryPath(), "run", "intro1")
	cmd.CombinedOutput() // Ignore error as exercise might fail

	// Check that progress file was created
	if _, err := os.Stat(".whack-a-goph-progress.json"); os.IsNotExist(err) {
		t.Error("Progress file should be created after running exercise")
	}

	// Verify progress file has valid JSON
	data, err := os.ReadFile(".whack-a-goph-progress.json")
	if err != nil {
		t.Fatalf("Failed to read progress file: %v", err)
	}

	var progress internal.Progress
	err = json.Unmarshal(data, &progress)
	if err != nil {
		t.Fatalf("Progress file should contain valid JSON: %v", err)
	}

	if len(progress.Exercises) == 0 {
		t.Error("Progress file should contain exercises")
	}
}
