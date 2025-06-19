package structs

// Structs help organize related data! 🐹
// Help these gophers learn about struct definitions and methods!

// TODO: Define a struct called Person with the following fields:
// - Name (string)
// - Age (int)
// - Email (string)
type Person struct {
	Name  string
	Age   int
	Email string
}

// TODO: Write a method for Person called Greet that returns a greeting string
// It should return: "Hi, I'm [Name] and I'm [Age] years old!"
func (p Person) Greet() string {
	// Fix me!
	return ""
}

// TODO: Write a function called NewPerson that takes name, age, and email
// as parameters and returns a new Person struct
func NewPerson(name string, age int, email string) Person {
	// Fix me!
	return Person{}
}

// TODO: Write a method for Person called IsAdult that returns true if age >= 18
func (p Person) IsAdult() bool {
	// Fix me!
	return false
}
