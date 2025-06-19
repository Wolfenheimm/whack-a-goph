package main

// Advanced 1: Reflection and Runtime Type Information 🐹
// Learn about Go's reflection capabilities using the reflect package.

// I AM NOT DONE

import (
	"fmt"
)

// TODO: Define a User struct with Name, Age, and Email fields
type User struct {
	Name  string // Fix me!
	Age   int    // Fix me!
	Email string // Fix me!
}

// TODO: Implement GetTypeName function
// Use reflection to return the type name of any value
func GetTypeName(value interface{}) string {
	// Fix me! Use reflect.TypeOf() and .Name()
	return ""
}

// TODO: Implement GetFieldNames function
// Use reflection to return all field names of a struct
func GetFieldNames(value interface{}) []string {
	// Fix me! Use reflect.TypeOf() and iterate through fields
	return nil
}

// TODO: Implement GetFieldValue function
// Use reflection to get the value of a specific field by name
func GetFieldValue(value interface{}, fieldName string) interface{} {
	// Fix me! Use reflect.ValueOf() and .FieldByName()
	return nil
}

// TODO: Implement SetFieldValue function
// Use reflection to set the value of a specific field by name
func SetFieldValue(value interface{}, fieldName string, newValue interface{}) error {
	// Fix me! Use reflect.ValueOf() and .FieldByName().Set()
	return fmt.Errorf("not implemented")
}

// TODO: Implement IsPointer function
// Check if the given value is a pointer
func IsPointer(value interface{}) bool {
	// Fix me! Use reflect.TypeOf() and .Kind()
	return false
}

// TODO: Implement CallMethod function
// Use reflection to call a method by name on a value
func CallMethod(value interface{}, methodName string, args ...interface{}) []interface{} {
	// Fix me! Use reflect.ValueOf() and .MethodByName()
	return nil
}

// TODO: Implement method SayHello on User
func (u User) SayHello() string {
	// Fix me! Return "Hello, I'm " + u.Name
	return ""
}

// TODO: Implement CopyStruct function
// Create a deep copy of a struct using reflection
func CopyStruct(src interface{}) interface{} {
	// Fix me! Use reflection to create a copy
	return nil
}

func main() {
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Test type name
	fmt.Printf("Type name: %s\n", GetTypeName(user))

	// Test field names
	fields := GetFieldNames(user)
	fmt.Printf("Field names: %v\n", fields)

	// Test field value
	name := GetFieldValue(user, "Name")
	fmt.Printf("Name field: %v\n", name)

	// Test pointer check
	fmt.Printf("Is user a pointer? %v\n", IsPointer(user))
	fmt.Printf("Is &user a pointer? %v\n", IsPointer(&user))

	// Test method call
	result := CallMethod(user, "SayHello")
	if len(result) > 0 {
		fmt.Printf("Method result: %v\n", result[0])
	}
}
