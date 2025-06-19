package main

// Traits 1: Basic Interface Definitions 🐹
// Learn how to define and implement interfaces in Go.

// I AM NOT DONE

import "fmt"

// TODO: Define a Speaker interface with a Speak method
type Speaker interface {
	Speak() string // Fix me! Add the Speak method
}

// TODO: Define a Mover interface with a Move method
type Mover interface {
	Move() string // Fix me! Add the Move method
}

// TODO: Define a Dog struct with Name field
type Dog struct {
	Name string // Fix me! Add the Name field
}

// TODO: Define a Cat struct with Name field
type Cat struct {
	Name string // Fix me! Add the Name field
}

// TODO: Define a Car struct with Brand field
type Car struct {
	Brand string // Fix me! Add the Brand field
}

// TODO: Implement Speak method for Dog
func (d Dog) Speak() string {
	// Fix me! Return "Woof! I'm " + d.Name
	return ""
}

// TODO: Implement Move method for Dog
func (d Dog) Move() string {
	// Fix me! Return d.Name + " runs around"
	return ""
}

// TODO: Implement Speak method for Cat
func (c Cat) Speak() string {
	// Fix me! Return "Meow! I'm " + c.Name
	return ""
}

// TODO: Implement Move method for Cat
func (c Cat) Move() string {
	// Fix me! Return c.Name + " prowls silently"
	return ""
}

// TODO: Implement Move method for Car
func (c Car) Move() string {
	// Fix me! Return "The " + c.Brand + " drives smoothly"
	return ""
}

// TODO: Implement function MakeSound that takes a Speaker
func MakeSound(s Speaker) string {
	// Fix me! Return s.Speak()
	return ""
}

// TODO: Implement function StartMoving that takes a Mover
func StartMoving(m Mover) string {
	// Fix me! Return m.Move()
	return ""
}

// TODO: Implement function DescribeAnimal that takes both Speaker and Mover
func DescribeAnimal(animal interface{}) string {
	// Fix me! Use type assertion to check if animal implements both interfaces
	// Return description combining Speak() and Move() if possible
	return ""
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	car := Car{Brand: "Toyota"}

	// Test speaking
	fmt.Println("Speaking:")
	fmt.Println(MakeSound(dog))
	fmt.Println(MakeSound(cat))

	// Test moving
	fmt.Println("\nMoving:")
	fmt.Println(StartMoving(dog))
	fmt.Println(StartMoving(cat))
	fmt.Println(StartMoving(car))

	// Test describing animals
	fmt.Println("\nDescribing:")
	fmt.Println(DescribeAnimal(dog))
	fmt.Println(DescribeAnimal(cat))
	fmt.Println(DescribeAnimal(car))
}
