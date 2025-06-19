package quizzes

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestConcurrentSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := ConcurrentSum(numbers)
	expected := 55 // Sum of 1-10

	if result != expected {
		t.Errorf("ConcurrentSum(%v) = %d; want %d", numbers, result, expected)
	}
}

func TestConcurrentSumEmpty(t *testing.T) {
	numbers := []int{}
	result := ConcurrentSum(numbers)
	expected := 0

	if result != expected {
		t.Errorf("ConcurrentSum(%v) = %d; want %d", numbers, result, expected)
	}
}

func TestWorkerPool(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5}
	results := WorkerPool(jobs, 3)

	if len(results) != len(jobs) {
		t.Errorf("WorkerPool should return same number of results as jobs; got %d, want %d", len(results), len(jobs))
		return
	}

	// Results should be processed versions of input (exact processing depends on implementation)
	if results == nil {
		t.Errorf("WorkerPool should not return nil")
	}
}

func TestTimeoutOperationSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	result, err := TimeoutOperation(ctx, 50*time.Millisecond)
	if err != nil {
		t.Errorf("TimeoutOperation should succeed when duration < timeout: %v", err)
	}

	if result != "completed" {
		t.Errorf("TimeoutOperation should return 'completed'; got %q", result)
	}
}

func TestTimeoutOperationTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err := TimeoutOperation(ctx, 200*time.Millisecond)
	if err == nil {
		t.Errorf("TimeoutOperation should timeout when duration > timeout")
	}
}

func TestGenericCache(t *testing.T) {
	cache := NewGenericCache[string, int]()
	if cache == nil {
		t.Errorf("NewGenericCache should not return nil")
		return
	}

	// Test Set and Get
	cache.Set("key1", 42)
	value, found := cache.Get("key1")

	if !found {
		t.Errorf("Cache.Get should find existing key")
	}

	if value != 42 {
		t.Errorf("Cache.Get('key1') = %d; want 42", value)
	}

	// Test non-existent key
	_, found = cache.Get("nonexistent")
	if found {
		t.Errorf("Cache.Get should not find non-existent key")
	}
}

func TestGenericCacheThreadSafety(t *testing.T) {
	cache := NewGenericCache[int, string]()
	if cache == nil {
		t.Skip("NewGenericCache returned nil")
	}

	// Test concurrent access
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			cache.Set(id, "value")
			_, _ = cache.Get(id)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
	// If we get here without race conditions, the cache is thread-safe
}

func TestReflectStruct(t *testing.T) {
	type TestStruct struct {
		Name string
		Age  int
	}

	ts := TestStruct{Name: "Go", Age: 15}
	result := ReflectStruct(ts)

	if result == nil {
		t.Errorf("ReflectStruct should not return nil")
		return
	}

	if len(result) != 2 {
		t.Errorf("ReflectStruct should return 2 fields; got %d", len(result))
	}

	if name, ok := result["Name"]; !ok || name != "Go" {
		t.Errorf("ReflectStruct should include Name field with value 'Go'")
	}

	if age, ok := result["Age"]; !ok || age != 15 {
		t.Errorf("ReflectStruct should include Age field with value 15")
	}
}

func TestPipeline(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := Pipeline(input)

	if result == nil {
		t.Errorf("Pipeline should not return nil")
		return
	}

	// Expected: multiply by 2, add 10, filter even
	// 1->2->12 (even), 2->4->14 (even), 3->6->16 (even), 4->8->18 (even), 5->10->20 (even)
	// All should be even after processing
	for _, val := range result {
		if val%2 != 0 {
			t.Errorf("Pipeline result should contain only even numbers; got %d", val)
		}
	}
}

func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter(2) // 2 requests per second
	if rl == nil {
		t.Errorf("NewRateLimiter should not return nil")
		return
	}

	// First request should be allowed
	if !rl.Allow() {
		t.Errorf("First request should be allowed")
	}

	// Test that rate limiting works (implementation dependent)
	// This is a basic test - exact behavior depends on implementation
}

func TestQuiz4(t *testing.T) {
	result := Quiz4()

	if result == "" {
		t.Errorf("Quiz4 should return non-empty string")
		return
	}

	// Check that result contains expected sections
	expectedSections := []string{
		"Quiz 4 Results:",
		"Concurrent sum:",
		"Worker pool results:",
		"Timeout operation:",
		"Cache get:",
		"Reflection:",
		"Pipeline:",
		"Rate limiter:",
	}

	for _, section := range expectedSections {
		if !strings.Contains(result, section) {
			t.Errorf("Quiz4 result should contain '%s'", section)
		}
	}
}

func TestQuiz4Integration(t *testing.T) {
	// Test that Quiz4 runs without panicking
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Quiz4 should not panic: %v", r)
		}
	}()

	result := Quiz4()

	// Basic validation that it returns something meaningful
	if len(result) < 50 {
		t.Errorf("Quiz4 result seems too short; got %d characters", len(result))
	}
}
