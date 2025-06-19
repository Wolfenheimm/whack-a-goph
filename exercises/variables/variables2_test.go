package variables

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestVariables2(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call Variables2 function
	Variables2()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check if output contains expected values with correct types
	expectedLines := []string{
		"number is 42 and has type int",
		"pi is 3.14 and has type float64",
		"message is Hello, Go! and has type string",
	}

	for _, expected := range expectedLines {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain %q, but got %q", expected, output)
		}
	}
}
