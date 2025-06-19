package intro

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Gopher!"
	actual := Hello()

	if actual != expected {
		t.Errorf("Hello() = %q; want %q", actual, expected)
	}
}
