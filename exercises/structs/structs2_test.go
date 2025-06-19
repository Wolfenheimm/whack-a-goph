package structs

import "testing"

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	expected := 15.0
	result := rect.Area()

	if result != expected {
		t.Errorf("Rectangle.Area() = %f; want %f", result, expected)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	expected := 16.0 // 2 * (5 + 3)
	result := rect.Perimeter()

	if result != expected {
		t.Errorf("Rectangle.Perimeter() = %f; want %f", result, expected)
	}
}

func TestRectangleString(t *testing.T) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	expected := "Rectangle: Width=5.000000, Height=3.000000"
	result := rect.String()

	if result != expected {
		t.Errorf("Rectangle.String() = %q; want %q", result, expected)
	}
}

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 2.0}
	expected := 12.56636 // 3.14159 * 2 * 2
	result := circle.Area()

	if result != expected {
		t.Errorf("Circle.Area() = %f; want %f", result, expected)
	}
}

func TestCreateRectangle(t *testing.T) {
	rect := CreateRectangle(4.0, 6.0)

	if rect.Width != 4.0 {
		t.Errorf("Expected Width to be 4.0, got %f", rect.Width)
	}

	if rect.Height != 6.0 {
		t.Errorf("Expected Height to be 6.0, got %f", rect.Height)
	}
}
