package quizzes

import (
	"strings"
	"testing"
)

func TestCircleImplementsShape(t *testing.T) {
	var _ Shape = Circle{}
}

func TestRectangleImplementsShape(t *testing.T) {
	var _ Shape = Rectangle{}
}

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 2.0}
	expected := 12.56636 // 3.14159 * 2 * 2
	result := circle.Area()

	if result != expected {
		t.Errorf("Circle.Area() = %f; want %f", result, expected)
	}
}

func TestCirclePerimeter(t *testing.T) {
	circle := Circle{Radius: 2.0}
	expected := 12.56636 // 2 * 3.14159 * 2
	result := circle.Perimeter()

	if result != expected {
		t.Errorf("Circle.Perimeter() = %f; want %f", result, expected)
	}
}

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 4.0, Height: 6.0}
	expected := 24.0
	result := rect.Area()

	if result != expected {
		t.Errorf("Rectangle.Area() = %f; want %f", result, expected)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	rect := Rectangle{Width: 4.0, Height: 6.0}
	expected := 20.0 // 2 * (4 + 6)
	result := rect.Perimeter()

	if result != expected {
		t.Errorf("Rectangle.Perimeter() = %f; want %f", result, expected)
	}
}

func TestCalculateTotal(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 2.0},
		Rectangle{Width: 4.0, Height: 6.0},
	}

	expected := 36.56636 // 12.56636 + 24.0
	result := CalculateTotal(shapes)

	if result != expected {
		t.Errorf("CalculateTotal() = %f; want %f", result, expected)
	}
}

func TestCreateShape(t *testing.T) {
	// Test valid circle
	shape, err := CreateShape("circle", 5.0)
	if err != nil {
		t.Errorf("CreateShape(\"circle\", 5.0) returned error: %v", err)
	}
	if _, ok := shape.(Circle); !ok {
		t.Errorf("CreateShape(\"circle\", 5.0) should return Circle")
	}

	// Test valid rectangle
	shape, err = CreateShape("rectangle", 4.0, 6.0)
	if err != nil {
		t.Errorf("CreateShape(\"rectangle\", 4.0, 6.0) returned error: %v", err)
	}
	if _, ok := shape.(Rectangle); !ok {
		t.Errorf("CreateShape(\"rectangle\", 4.0, 6.0) should return Rectangle")
	}

	// Test invalid shape
	_, err = CreateShape("triangle", 3.0)
	if err == nil {
		t.Errorf("CreateShape(\"triangle\", 3.0) should return error")
	}
}

func TestDescribeShape(t *testing.T) {
	circle := Circle{Radius: 5.0}
	result := DescribeShape(circle)

	if result == "" {
		t.Errorf("DescribeShape should return non-empty string for Circle")
	}

	rect := Rectangle{Width: 4.0, Height: 6.0}
	result = DescribeShape(rect)

	if result == "" {
		t.Errorf("DescribeShape should return non-empty string for Rectangle")
	}
}

func TestQuiz3(t *testing.T) {
	result := Quiz3()

	// Check that output contains expected elements
	if !strings.Contains(result, "Quiz 3 completed!") {
		t.Errorf("Quiz3() should contain 'Quiz 3 completed!', got %q", result)
	}
}
