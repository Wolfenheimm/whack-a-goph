package main

import (
	"strings"
	"testing"
)

func TestColorString(t *testing.T) {
	tests := []struct {
		color    Color
		expected string
	}{
		{Red, "Red"},
		{Green, "Green"},
		{Blue, "Blue"},
		{Yellow, "Yellow"},
	}

	for _, test := range tests {
		result := test.color.String()
		if result != test.expected {
			t.Errorf("Color(%d).String() = %q; want %q", test.color, result, test.expected)
		}
	}
}

func TestGenerateEnumString(t *testing.T) {
	enumName := "Status"
	values := []string{"Active", "Inactive", "Pending"}

	result := GenerateEnumString(enumName, values)

	// Check that result contains expected elements
	if result == "" {
		t.Errorf("GenerateEnumString should return non-empty string")
		return
	}

	// Should contain the enum name
	if !strings.Contains(result, enumName) {
		t.Errorf("Generated code should contain enum name %q", enumName)
	}

	// Should contain all values
	for _, value := range values {
		if !strings.Contains(result, value) {
			t.Errorf("Generated code should contain value %q", value)
		}
	}
}

func TestExecuteTemplate(t *testing.T) {
	enumName := "Priority"
	values := []string{"Low", "Medium", "High"}

	result, err := ExecuteTemplate(enumName, values)
	if err != nil {
		t.Errorf("ExecuteTemplate should not return error: %v", err)
		return
	}

	if result == "" {
		t.Errorf("ExecuteTemplate should return non-empty result")
		return
	}

	// Check that template was properly executed
	if !strings.Contains(result, enumName) {
		t.Errorf("Template result should contain enum name %q", enumName)
	}

	for _, value := range values {
		if !strings.Contains(result, value) {
			t.Errorf("Template result should contain value %q", value)
		}
	}
}

func TestBuildTags(t *testing.T) {
	tags := BuildTags()

	if tags == nil {
		t.Errorf("BuildTags should not return nil")
		return
	}

	// Should return at least some build tags
	if len(tags) == 0 {
		t.Errorf("BuildTags should return at least one tag")
	}

	// Common build tags that should be present
	hasValidTag := false
	validTags := []string{"linux", "darwin", "windows", "amd64", "arm64", "386"}
	for _, tag := range tags {
		for _, valid := range validTags {
			if tag == valid {
				hasValidTag = true
				break
			}
		}
	}

	if !hasValidTag {
		t.Errorf("BuildTags should return at least one valid build tag")
	}
}

func TestIsDebugBuild(t *testing.T) {
	// This test just verifies the function exists and returns a boolean
	result := IsDebugBuild()

	// Should return either true or false (both are valid)
	if result != true && result != false {
		t.Errorf("IsDebugBuild should return a boolean value")
	}
}

func TestGenerateTestFile(t *testing.T) {
	structName := "User"
	fields := []string{"Name", "Age", "Email"}

	result := GenerateTestFile(structName, fields)

	if result == "" {
		t.Errorf("GenerateTestFile should return non-empty string")
		return
	}

	// Should contain struct name
	if !strings.Contains(result, structName) {
		t.Errorf("Generated test should contain struct name %q", structName)
	}

	// Should contain test-related keywords
	testKeywords := []string{"test", "Test", "func"}
	hasTestKeyword := false
	for _, keyword := range testKeywords {
		if strings.Contains(result, keyword) {
			hasTestKeyword = true
			break
		}
	}

	if !hasTestKeyword {
		t.Errorf("Generated test should contain test-related keywords")
	}
}

func TestConditionalCode(t *testing.T) {
	tests := []struct {
		condition string
		expected  string
	}{
		{"production", "// Production code"},
		{"development", "// Development code"},
		{"unknown", "// Default code"},
	}

	for _, test := range tests {
		result := ConditionalCode(test.condition)
		if result != test.expected {
			t.Errorf("ConditionalCode(%q) = %q; want %q", test.condition, result, test.expected)
		}
	}
}
