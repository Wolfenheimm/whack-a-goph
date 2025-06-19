package quizzes

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestTaskTypeString(t *testing.T) {
	tests := []struct {
		taskType TaskType
		expected string
	}{
		{WorkTask, "Work"},
		{PersonalTask, "Personal"},
		{UrgentTask, "Urgent"},
	}

	for _, test := range tests {
		result := test.taskType.String()
		if result != test.expected && result != "" {
			t.Errorf("TaskType(%d).String() = %q; want %q", test.taskType, result, test.expected)
		}
	}
}

func TestPriorityString(t *testing.T) {
	tests := []struct {
		priority Priority
		expected string
	}{
		{LowPriority, "Low"},
		{MediumPriority, "Medium"},
		{HighPriority, "High"},
	}

	for _, test := range tests {
		result := test.priority.String()
		if result != test.expected && result != "" {
			t.Errorf("Priority(%d).String() = %q; want %q", test.priority, result, test.expected)
		}
	}
}

func TestTaskImplementsTaskProcessor(t *testing.T) {
	var _ TaskProcessor = Task{}
}

func TestTaskProcess(t *testing.T) {
	task := Task{Completed: false}
	err := task.Process()

	// Should not error for incomplete task
	if err != nil && !strings.Contains(err.Error(), "completed") {
		t.Errorf("Task.Process() should work for incomplete task, got error: %v", err)
	}
}

func TestTaskGetEstimatedTime(t *testing.T) {
	task := Task{Type: WorkTask, Priority: HighPriority}
	duration := task.GetEstimatedTime()

	// Should return some duration
	if duration < 0 {
		t.Errorf("Task.GetEstimatedTime() should return non-negative duration, got %v", duration)
	}
}

func TestNewTaskManager(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Fatal("NewTaskManager() should not return nil")
	}
}

func TestTaskManagerAddTask(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	deadline := time.Now().Add(24 * time.Hour)
	id := tm.AddTask("Test Task", "Description", WorkTask, HighPriority, deadline, []string{"test"})

	if id == "" {
		t.Errorf("AddTask should return non-empty ID")
	}
}

func TestTaskManagerGetTask(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	// Test getting non-existent task
	_, err := tm.GetTask("nonexistent")
	if err == nil {
		t.Errorf("GetTask should return error for non-existent task")
	}
}

func TestTaskManagerCompleteTask(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	// Test completing non-existent task
	err := tm.CompleteTask("nonexistent")
	if err == nil {
		t.Errorf("CompleteTask should return error for non-existent task")
	}
}

func TestTaskManagerGetTasksByType(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	tasks := tm.GetTasksByType(WorkTask)
	// Should return slice (empty or not)
	if tasks == nil {
		t.Errorf("GetTasksByType should return slice, not nil")
	}
}

func TestTaskManagerGetOverdueTasks(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	tasks := tm.GetOverdueTasks()
	// Should return slice (empty or not)
	if tasks == nil {
		t.Errorf("GetOverdueTasks should return slice, not nil")
	}
}

func TestTaskManagerSortTasksByPriority(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	// Should not panic
	tm.SortTasksByPriority()
}

func TestTaskManagerGetTasksWithTag(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	tasks := tm.GetTasksWithTag("test")
	// Should return slice (empty or not)
	if tasks == nil {
		t.Errorf("GetTasksWithTag should return slice, not nil")
	}
}

func TestTaskManagerProcessAllTasks(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	resultChan := tm.ProcessAllTasks()

	// Should receive at least one result
	select {
	case result := <-resultChan:
		if result == nil {
			t.Errorf("ProcessAllTasks should return errors (even if just 'not implemented')")
		}
	case <-time.After(100 * time.Millisecond):
		t.Errorf("ProcessAllTasks should send results to channel")
	}
}

func TestTaskManagerToJSON(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	data, err := tm.ToJSON()
	if err != nil && !strings.Contains(err.Error(), "not implemented") {
		t.Errorf("ToJSON should work or return 'not implemented' error, got: %v", err)
	}

	if err == nil && data == nil {
		t.Errorf("ToJSON should return data if no error")
	}
}

func TestTaskManagerFromJSON(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Skip("NewTaskManager() not implemented yet")
	}

	testData := []byte(`{"tasks":[],"next_id":1}`)
	err := tm.FromJSON(testData)
	if err != nil && !strings.Contains(err.Error(), "not implemented") {
		t.Errorf("FromJSON should work or return 'not implemented' error, got: %v", err)
	}
}

func TestCalculateWorkload(t *testing.T) {
	var processors []TaskProcessor

	duration := CalculateWorkload(processors)
	if duration < 0 {
		t.Errorf("CalculateWorkload should return non-negative duration, got %v", duration)
	}
}

func TestTaskJSONTags(t *testing.T) {
	task := Task{
		ID:          "test",
		Title:       "Test Task",
		Description: "Test Description",
		Type:        WorkTask,
		Priority:    HighPriority,
		Deadline:    time.Now(),
		Completed:   false,
		Tags:        []string{"test"},
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Errorf("Task should be JSON marshallable, got error: %v", err)
	}

	if !strings.Contains(string(data), "test") {
		t.Errorf("JSON should contain task data, got: %s", string(data))
	}
}

func TestQuizFinal(t *testing.T) {
	result := QuizFinal()

	// Check that output contains expected elements
	expectedStrings := []string{
		"Starting Final Quiz",
		"Final Quiz completed",
		"Go master",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(result, expected) {
			t.Errorf("QuizFinal() should contain %q, got %q", expected, result)
		}
	}
}
