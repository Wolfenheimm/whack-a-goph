package quizzes

// Quiz 3: Interfaces and Error Handling 🐹
// This quiz covers interfaces, error handling, methods, and advanced Go concepts.
//
// You need to implement interfaces and proper error handling.

// I AM NOT DONE

import (
	"errors"
	"fmt"
)

// TODO: Define an interface called Shape with methods:
// - Area() float64
// - Perimeter() float64
type Shape interface {
	Area() float64      // Fix me! Make sure this method signature is correct
	Perimeter() float64 // Fix me! Make sure this method signature is correct
}

// TODO: Define a Circle struct with field Radius (float64)
type Circle struct {
	Radius float64 // Fix me! Make sure this field exists
}

// TODO: Define a Rectangle struct with fields Width and Height (both float64)
type Rectangle struct {
	Width  float64 // Fix me! Make sure this field exists
	Height float64 // Fix me! Make sure this field exists
}

// TODO: Implement Area method for Circle (π * r²)
// Use 3.14159 for π
func (c Circle) Area() float64 {
	// Fix me!
	return 0.0
}

// TODO: Implement Perimeter method for Circle (2 * π * r)
func (c Circle) Perimeter() float64 {
	// Fix me!
	return 0.0
}

// TODO: Implement Area method for Rectangle (width * height)
func (r Rectangle) Area() float64 {
	// Fix me!
	return 0.0
}

// TODO: Implement Perimeter method for Rectangle (2 * (width + height))
func (r Rectangle) Perimeter() float64 {
	// Fix me!
	return 0.0
}

// TODO: Implement CalculateTotal function that takes a slice of Shape
// and returns the total area of all shapes
func CalculateTotal(shapes []Shape) float64 {
	// Fix me!
	return 0.0
}

// TODO: Implement CreateShape function that returns a Shape and error
// If shapeType is "circle" and radius > 0, return Circle
// If shapeType is "rectangle" and width > 0 and height > 0, return Rectangle
// Otherwise, return error "invalid shape parameters"
func CreateShape(shapeType string, params ...float64) (Shape, error) {
	// Fix me!
	return nil, errors.New("not implemented")
}

// TODO: Implement DescribeShape function that takes a Shape interface
// and returns a description string based on the concrete type
// Use type assertion or type switch
func DescribeShape(s Shape) string {
	// Fix me!
	return ""
}

// TODO: Implement the Quiz3 function
// This function should demonstrate the quiz functionality
func Quiz3() string {
	result := ""

	// TODO: Create shapes using CreateShape function
	circle, err1 := CreateShape("circle", 5.0)
	if err1 != nil {
		result += fmt.Sprintf("Error creating circle: %v\n", err1)
		return result
	}

	rectangle, err2 := CreateShape("rectangle", 4.0, 6.0)
	if err2 != nil {
		result += fmt.Sprintf("Error creating rectangle: %v\n", err2)
		return result
	}

	// TODO: Create a slice of shapes
	shapes := []Shape{circle, rectangle}

	// TODO: Print area and perimeter for each shape
	for _, shape := range shapes {
		result += fmt.Sprintf("Shape: %s\n", DescribeShape(shape))
		result += fmt.Sprintf("Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
	}

	// TODO: Print total area
	total := CalculateTotal(shapes)
	result += fmt.Sprintf("Total area: %.2f\n", total)

	// TODO: Test error case
	_, err3 := CreateShape("triangle", 3.0)
	if err3 != nil {
		result += fmt.Sprintf("Expected error: %v\n", err3)
	}

	result += "Quiz 3 completed!"
	return result
}
