package methods

import "testing"

func TestRectangle(t *testing.T) {
	t.Run("Area", func(t *testing.T) {
		r := Rectangle{Width: 5, Height: 3}
		got := r.Area()
		want := 15.0
		if got != want {
			t.Errorf("Area() = %v; want %v", got, want)
		}
	})

	t.Run("Perimeter", func(t *testing.T) {
		r := Rectangle{Width: 5, Height: 3}
		got := r.Perimeter()
		want := 16.0
		if got != want {
			t.Errorf("Perimeter() = %v; want %v", got, want)
		}
	})

	t.Run("Scale", func(t *testing.T) {
		r := Rectangle{Width: 5, Height: 3}
		r.Scale(2)
		if r.Width != 10 || r.Height != 6 {
			t.Errorf("Scale(2) = %v; want width=10, height=6", r)
		}
	})
}
