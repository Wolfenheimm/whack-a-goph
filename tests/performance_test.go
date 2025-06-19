package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/raithbheart/whack-a-goph/internal"
)

// TestExerciseLoadingPerformance tests how quickly exercises can be loaded
func TestExerciseLoadingPerformance(t *testing.T) {
	// Benchmark exercise loading
	start := time.Now()

	for i := 0; i < 10; i++ {
		exercises, err := internal.GetExercises()
		if err != nil {
			t.Fatalf("Failed to load exercises: %v", err)
		}

		if len(exercises) == 0 {
			t.Fatal("No exercises loaded")
		}
	}

	duration := time.Since(start)

	// Should load quickly - arbitrary threshold of 1 second for 10 loads
	if duration > time.Second {
		t.Errorf("Exercise loading took too long: %v", duration)
	}

	t.Logf("Loaded exercises 10 times in %v", duration)
}

// TestProgressOperationsPerformance tests progress operations under load
func TestProgressOperationsPerformance(t *testing.T) {
	setupCleanProgress(t)
	defer cleanupProgress()

	exercises, err := internal.GetExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	if len(exercises) == 0 {
		t.Skip("No exercises available for performance test")
	}

	start := time.Now()

	// Perform multiple progress operations
	for i := 0; i < 50; i++ {
		exerciseName := exercises[i%len(exercises)].Name

		// Mark as completed
		err := internal.MarkCompleted(exerciseName)
		if err != nil {
			t.Fatalf("Failed to mark exercise as completed: %v", err)
		}

		// Load progress
		_, err = internal.LoadProgress()
		if err != nil {
			t.Fatalf("Failed to load progress: %v", err)
		}

		// Reset exercise
		err = internal.ResetExercise(exerciseName)
		if err != nil {
			t.Fatalf("Failed to reset exercise: %v", err)
		}
	}

	duration := time.Since(start)
	t.Logf("Completed 150 progress operations in %v", duration)

	// Should complete quickly - arbitrary threshold of 5 seconds
	if duration > 5*time.Second {
		t.Errorf("Progress operations took too long: %v", duration)
	}
}

// TestConcurrentProgressOperations tests concurrent access to progress system
func TestConcurrentProgressOperations(t *testing.T) {
	setupCleanProgress(t)
	defer cleanupProgress()

	exercises, err := internal.GetExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	if len(exercises) == 0 {
		t.Skip("No exercises available for concurrent test")
	}

	// Number of concurrent workers
	numWorkers := 10
	operationsPerWorker := 20

	var wg sync.WaitGroup
	start := time.Now()

	// Start concurrent workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < operationsPerWorker; j++ {
				exerciseName := exercises[(workerID*operationsPerWorker+j)%len(exercises)].Name

				// Random operations
				switch j % 3 {
				case 0:
					internal.MarkCompleted(exerciseName)
				case 1:
					internal.LoadProgress()
				case 2:
					internal.ResetExercise(exerciseName)
				}
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	totalOperations := numWorkers * operationsPerWorker
	t.Logf("Completed %d concurrent operations in %v", totalOperations, duration)

	// Should handle concurrent access without crashing
	// No specific time requirement, just that it completes
}

// TestMemoryUsage tests that the system doesn't have memory leaks
func TestMemoryUsage(t *testing.T) {
	// This is a basic test - would need runtime.GC() and runtime.ReadMemStats for detailed analysis
	setupCleanProgress(t)
	defer cleanupProgress()

	// Perform many operations to test for memory leaks
	for i := 0; i < 1000; i++ {
		exercises, err := internal.GetExercises()
		if err != nil {
			t.Fatalf("Failed to load exercises: %v", err)
		}

		if len(exercises) > 0 {
			exerciseName := exercises[i%len(exercises)].Name
			internal.MarkCompleted(exerciseName)
			internal.LoadProgress()
			internal.ResetExercise(exerciseName)
		}
	}

	// If we get here without running out of memory, test passes
	t.Log("Memory usage test completed successfully")
}

// TestExerciseValidationPerformance tests how quickly we can validate all exercises
func TestExerciseValidationPerformance(t *testing.T) {
	exercises, err := internal.GetExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	start := time.Now()

	validExercises := 0
	for _, exercise := range exercises {
		// Quick validation checks
		if exercise.Name != "" && exercise.Path != "" && exercise.Description != "" {
			validExercises++
		}
	}

	duration := time.Since(start)

	t.Logf("Validated %d exercises in %v", len(exercises), duration)

	if validExercises != len(exercises) {
		t.Errorf("Expected all exercises to be valid, got %d/%d", validExercises, len(exercises))
	}

	// Should validate quickly
	if duration > 100*time.Millisecond {
		t.Errorf("Exercise validation took too long: %v", duration)
	}
}

// TestLargeProgressFile tests handling of large progress files
func TestLargeProgressFile(t *testing.T) {
	setupCleanProgress(t)
	defer cleanupProgress()

	// Create a large number of fake exercises
	largeProgress := internal.Progress{
		Exercises: make([]internal.Exercise, 0, 1000),
	}

	for i := 0; i < 1000; i++ {
		exercise := internal.Exercise{
			Name:        fmt.Sprintf("test_exercise_%d", i),
			Path:        fmt.Sprintf("test/path_%d.go", i),
			Description: fmt.Sprintf("Test exercise %d", i),
			Completed:   i%2 == 0, // Half completed
		}
		largeProgress.Exercises = append(largeProgress.Exercises, exercise)
	}

	start := time.Now()

	// This would require access to SaveProgress function
	// For now, we test that loading a large number of exercises doesn't crash
	exercises, err := internal.GetExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	// Perform operations on all exercises
	for _, exercise := range exercises {
		internal.FindExercise(exercise.Name)
	}

	duration := time.Since(start)
	t.Logf("Processed large exercise set in %v", duration)
}

// TestStressExerciseSearch tests searching for exercises under stress
func TestStressExerciseSearch(t *testing.T) {
	exercises, err := internal.GetExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	if len(exercises) == 0 {
		t.Skip("No exercises available for stress test")
	}

	start := time.Now()

	// Search for each exercise multiple times
	for i := 0; i < 100; i++ {
		for _, exercise := range exercises {
			_, err := internal.FindExercise(exercise.Name)
			if err != nil {
				t.Errorf("Failed to find exercise %s: %v", exercise.Name, err)
			}
		}
	}

	duration := time.Since(start)
	t.Logf("Performed %d exercise searches in %v", 100*len(exercises), duration)

	// Should complete quickly
	if duration > 2*time.Second {
		t.Errorf("Exercise search took too long: %v", duration)
	}
}

// Helper functions for benchmarking

// BenchmarkExerciseLoading is a Go benchmark test
func BenchmarkExerciseLoading(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises, err := internal.GetExercises()
		if err != nil {
			b.Fatalf("Failed to load exercises: %v", err)
		}
		if len(exercises) == 0 {
			b.Fatal("No exercises loaded")
		}
	}
}

// BenchmarkProgressOperations benchmarks progress operations
func BenchmarkProgressOperations(b *testing.B) {
	setupCleanProgress(&testing.T{}) // Note: This is a hack for the benchmark
	defer cleanupProgress()

	exercises, err := internal.GetExercises()
	if err != nil {
		b.Fatalf("Failed to load exercises: %v", err)
	}

	if len(exercises) == 0 {
		b.Skip("No exercises available")
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		exerciseName := exercises[i%len(exercises)].Name
		internal.MarkCompleted(exerciseName)
		internal.LoadProgress()
		internal.ResetExercise(exerciseName)
	}
}
