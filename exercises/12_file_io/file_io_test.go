package file_io

import (
	"os"
	"testing"
)

func TestWriteAndReadFile(t *testing.T) {
	filename := "test.txt"
	content := "Hello, Go!"

	// Clean up
	defer os.Remove(filename)

	// Test WriteToFile
	err := WriteToFile(filename, content)
	if err != nil {
		t.Errorf("WriteToFile(%q, %q) returned error: %v", filename, content, err)
	}

	// Test ReadFromFile
	got, err := ReadFromFile(filename)
	if err != nil {
		t.Errorf("ReadFromFile(%q) returned error: %v", filename, err)
	}
	if got != content {
		t.Errorf("ReadFromFile(%q) = %q; want %q", filename, got, content)
	}
}

func TestFileExists(t *testing.T) {
	filename := "test_exists.txt"

	// File should not exist initially
	if FileExists(filename) {
		t.Errorf("FileExists(%q) = true; want false", filename)
	}

	// Create file
	WriteToFile(filename, "test")
	defer os.Remove(filename)

	// File should exist now
	if !FileExists(filename) {
		t.Errorf("FileExists(%q) = false; want true", filename)
	}
}

func TestCountLines(t *testing.T) {
	filename := "test_lines.txt"
	content := "line 1\nline 2\nline 3"

	// Clean up
	defer os.Remove(filename)

	WriteToFile(filename, content)

	got, err := CountLines(filename)
	if err != nil {
		t.Errorf("CountLines(%q) returned error: %v", filename, err)
	}
	want := 3
	if got != want {
		t.Errorf("CountLines(%q) = %d; want %d", filename, got, want)
	}
}
