package structs

// Time to learn about structs! 🐹
// Structs group related data together.

// I AM NOT DONE

// TODO: Define a struct called Person with fields:
// - Name (string)
// - Age (int)
// Fix me! Make sure this struct has the right fields
type Person struct {
	Name string // Fix me! This field exists but might need adjustment
	Age  int    // Fix me! This field exists but might need adjustment
}

func CreatePerson(name string, age int) Person {
	// TODO: Create and return a Person with the given name and age
	// Fix me!
	return Person{}
}

func GetPersonInfo(p Person) string {
	// TODO: Return a string in format "Name: [name], Age: [age]"
	// Fix me!
	return ""
}

func IsAdult(p Person) bool {
	// TODO: Return true if person's age is 18 or older
	// Fix me!
	return false
}
