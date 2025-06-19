package tests

import (
	"os"
	"testing"
)

// TestConfiguration holds test settings
var TestConfig = struct {
	Verbose      bool
	Timeout      int
	SkipSlow     bool
	TestDataDir  string
	BackupSuffix string
}{
	Verbose:      false,
	Timeout:      30,
	SkipSlow:     false,
	TestDataDir:  "test_data",
	BackupSuffix: ".test-backup",
}

// SetupTestEnvironment prepares the test environment
func SetupTestEnvironment() {
	// Set test-friendly environment variables
	os.Setenv("WHACK_A_GOPH_TEST_MODE", "true")

	// Create test data directory if it doesn't exist
	os.MkdirAll(TestConfig.TestDataDir, 0755)
}

// CleanupTestEnvironment cleans up after tests
func CleanupTestEnvironment() {
	// Remove test data directory
	os.RemoveAll(TestConfig.TestDataDir)

	// Unset test environment variables
	os.Unsetenv("WHACK_A_GOPH_TEST_MODE")
}

// IsVerbose returns whether verbose mode is enabled
func IsVerbose() bool {
	return TestConfig.Verbose || testing.Verbose()
}

// ShouldSkipSlow returns whether slow tests should be skipped
func ShouldSkipSlow() bool {
	return TestConfig.SkipSlow || os.Getenv("SKIP_SLOW_TESTS") == "true"
}

// GetProgressBackupFile returns the backup filename for progress file
func GetProgressBackupFile() string {
	return ".whack-a-goph-progress.json" + TestConfig.BackupSuffix
}
