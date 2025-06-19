package main

import "testing"

func TestDogSpeak(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	result := dog.Speak()
	expected := "Woof! I'm Buddy"
	if result != expected {
		t.Errorf("Dog.Speak() = %q; want %q", result, expected)
	}
}

func TestCatSpeak(t *testing.T) {
	cat := Cat{Name: "Whiskers"}
	result := cat.Speak()
	expected := "Meow! I'm Whiskers"
	if result != expected {
		t.Errorf("Cat.Speak() = %q; want %q", result, expected)
	}
}

func TestDogMove(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	result := dog.Move()
	expected := "Buddy runs around"
	if result != expected {
		t.Errorf("Dog.Move() = %q; want %q", result, expected)
	}
}

func TestCatMove(t *testing.T) {
	cat := Cat{Name: "Whiskers"}
	result := cat.Move()
	expected := "Whiskers prowls silently"
	if result != expected {
		t.Errorf("Cat.Move() = %q; want %q", result, expected)
	}
}

func TestCarMove(t *testing.T) {
	car := Car{Brand: "Toyota"}
	result := car.Move()
	expected := "The Toyota drives smoothly"
	if result != expected {
		t.Errorf("Car.Move() = %q; want %q", result, expected)
	}
}

func TestMakeSound(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	result := MakeSound(dog)
	expected := "Woof! I'm Buddy"
	if result != expected {
		t.Errorf("MakeSound(dog) = %q; want %q", result, expected)
	}
}

func TestStartMoving(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	result := StartMoving(dog)
	expected := "Buddy runs around"
	if result != expected {
		t.Errorf("StartMoving(dog) = %q; want %q", result, expected)
	}
}

func TestDescribeAnimal(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	result := DescribeAnimal(dog)

	// Should contain both speaking and moving information
	if result == "" {
		t.Errorf("DescribeAnimal should return non-empty description")
	}
}
