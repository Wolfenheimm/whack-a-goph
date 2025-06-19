package variables

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestVariables1(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call Variables1 function
	Variables1()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check if output contains expected values
	expectedLines := []string{
		"x = 5",
		"name = Gopher",
		"done = true",
	}

	for _, expected := range expectedLines {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain %q, but got %q", expected, output)
		}
	}
}
