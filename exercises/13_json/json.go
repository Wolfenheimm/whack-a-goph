package json

// TODO: Define a Person struct and implement JSON operations:
// 1. Define Person struct with Name (string), Age (int), Email (string)
// 2. PersonToJSON - convert Person to JSON string
// 3. PersonFromJSON - convert JSON string to Person
// 4. PeopleToJSON - convert slice of Person to JSON string
// 5. PeopleFromJSON - convert JSON string to slice of Person

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func PersonToJSON(p Person) (string, error) {
	// Fix me!
	return "", nil
}

func PersonFromJSON(jsonStr string) (Person, error) {
	// Fix me!
	return Person{}, nil
}

func PeopleToJSON(people []Person) (string, error) {
	// Fix me!
	return "", nil
}

func PeopleFromJSON(jsonStr string) ([]Person, error) {
	// Fix me!
	return nil, nil
}
