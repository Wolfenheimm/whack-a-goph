package main

import (
	"reflect"
	"testing"
)

func TestGetTypeName(t *testing.T) {
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	typeName := GetTypeName(user)
	expected := "User"
	if typeName != expected {
		t.Errorf("GetTypeName(user) = %q; want %q", typeName, expected)
	}

	// Test with built-in types
	if GetTypeName(42) != "int" {
		t.Errorf("GetTypeName(42) should return 'int'")
	}

	if GetTypeName("hello") != "string" {
		t.Errorf("GetTypeName(\"hello\") should return 'string'")
	}
}

func TestGetFieldNames(t *testing.T) {
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	fields := GetFieldNames(user)
	expected := []string{"Name", "Age", "Email"}

	if len(fields) != len(expected) {
		t.Errorf("GetFieldNames() returned %d fields; want %d", len(fields), len(expected))
		return
	}

	fieldSet := make(map[string]bool)
	for _, field := range fields {
		fieldSet[field] = true
	}

	for _, expectedField := range expected {
		if !fieldSet[expectedField] {
			t.Errorf("GetFieldNames() missing field %q", expectedField)
		}
	}
}

func TestGetFieldValue(t *testing.T) {
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	name := GetFieldValue(user, "Name")
	if name != "Alice" {
		t.Errorf("GetFieldValue(user, \"Name\") = %v; want \"Alice\"", name)
	}

	age := GetFieldValue(user, "Age")
	if age != 30 {
		t.Errorf("GetFieldValue(user, \"Age\") = %v; want 30", age)
	}

	email := GetFieldValue(user, "Email")
	if email != "alice@example.com" {
		t.Errorf("GetFieldValue(user, \"Email\") = %v; want \"alice@example.com\"", email)
	}
}

func TestSetFieldValue(t *testing.T) {
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Test setting field value (should work with pointer)
	err := SetFieldValue(&user, "Name", "Bob")
	if err != nil {
		t.Errorf("SetFieldValue should not return error for valid operation: %v", err)
	}

	// Verify the change
	if user.Name != "Bob" {
		t.Errorf("After SetFieldValue, user.Name = %q; want \"Bob\"", user.Name)
	}
}

func TestIsPointer(t *testing.T) {
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	if IsPointer(user) {
		t.Errorf("IsPointer(user) should be false")
	}

	if !IsPointer(&user) {
		t.Errorf("IsPointer(&user) should be true")
	}

	// Test with other types
	num := 42
	if IsPointer(num) {
		t.Errorf("IsPointer(num) should be false")
	}

	if !IsPointer(&num) {
		t.Errorf("IsPointer(&num) should be true")
	}
}

func TestCallMethod(t *testing.T) {
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	result := CallMethod(user, "SayHello")
	if len(result) != 1 {
		t.Errorf("CallMethod should return 1 result; got %d", len(result))
		return
	}

	expected := "Hello, I'm Alice"
	if result[0] != expected {
		t.Errorf("CallMethod(user, \"SayHello\") = %v; want %q", result[0], expected)
	}
}

func TestUserSayHello(t *testing.T) {
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	result := user.SayHello()
	expected := "Hello, I'm Alice"
	if result != expected {
		t.Errorf("user.SayHello() = %q; want %q", result, expected)
	}
}

func TestCopyStruct(t *testing.T) {
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	copied := CopyStruct(user)
	if copied == nil {
		t.Errorf("CopyStruct should not return nil")
		return
	}

	// Verify it's a User type
	copiedUser, ok := copied.(User)
	if !ok {
		t.Errorf("CopyStruct should return User type; got %T", copied)
		return
	}

	// Verify values are copied
	if !reflect.DeepEqual(user, copiedUser) {
		t.Errorf("CopyStruct did not create exact copy: original=%+v, copy=%+v", user, copiedUser)
	}
}
