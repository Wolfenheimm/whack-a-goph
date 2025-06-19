package intro

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrintGreeting(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	PrintGreeting("Gopher")

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := "Hello, Gopher! Welcome to Go!\n"
	if output != expected {
		t.Errorf("PrintGreeting(\"Gopher\") output = %q; want %q", output, expected)
	}
}
