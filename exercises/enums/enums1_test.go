package main

import "testing"

func TestDayOfWeekConstants(t *testing.T) {
	// Test that constants are properly defined with iota
	if Sunday != 0 {
		t.Errorf("Sunday should be 0, got %d", Sunday)
	}
	if Monday != 1 {
		t.Errorf("Monday should be 1, got %d", Monday)
	}
	if Saturday != 6 {
		t.Errorf("Saturday should be 6, got %d", Saturday)
	}
}

func TestStatusConstants(t *testing.T) {
	// Test that status constants start from 1
	if Pending != 1 {
		t.Errorf("Pending should be 1, got %d", Pending)
	}
	if InProgress != 2 {
		t.Errorf("InProgress should be 2, got %d", InProgress)
	}
	if Completed != 3 {
		t.Errorf("Completed should be 3, got %d", Completed)
	}
	if Cancelled != 4 {
		t.Errorf("Cancelled should be 4, got %d", Cancelled)
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		day      DayOfWeek
		expected bool
	}{
		{Sunday, true},
		{Monday, false},
		{Tuesday, false},
		{Wednesday, false},
		{Thursday, false},
		{Friday, false},
		{Saturday, true},
	}

	for _, test := range tests {
		result := IsWeekend(test.day)
		if result != test.expected {
			t.Errorf("IsWeekend(%d) = %v; want %v", test.day, result, test.expected)
		}
	}
}

func TestGetNextDay(t *testing.T) {
	tests := []struct {
		day      DayOfWeek
		expected DayOfWeek
	}{
		{Sunday, Monday},
		{Monday, Tuesday},
		{Tuesday, Wednesday},
		{Wednesday, Thursday},
		{Thursday, Friday},
		{Friday, Saturday},
		{Saturday, Sunday}, // Wrap around
	}

	for _, test := range tests {
		result := GetNextDay(test.day)
		if result != test.expected {
			t.Errorf("GetNextDay(%d) = %d; want %d", test.day, result, test.expected)
		}
	}
}

func TestIsValidStatus(t *testing.T) {
	tests := []struct {
		status   Status
		expected bool
	}{
		{Status(0), false}, // Before range
		{Pending, true},
		{InProgress, true},
		{Completed, true},
		{Cancelled, true},
		{Status(5), false}, // After range
	}

	for _, test := range tests {
		result := IsValidStatus(test.status)
		if result != test.expected {
			t.Errorf("IsValidStatus(%d) = %v; want %v", test.status, result, test.expected)
		}
	}
}

func TestGetStatusMessage(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{Pending, "Task is pending"},
		{InProgress, "Task is in progress"},
		{Completed, "Task is completed"},
		{Cancelled, "Task is cancelled"},
	}

	for _, test := range tests {
		result := GetStatusMessage(test.status)
		if result != test.expected {
			t.Errorf("GetStatusMessage(%d) = %q; want %q", test.status, result, test.expected)
		}
	}
}
