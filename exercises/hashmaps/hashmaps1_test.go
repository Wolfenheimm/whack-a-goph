package hashmaps

import (
	"reflect"
	"sort"
	"testing"
)

func TestCreateMap(t *testing.T) {
	result := CreateMap()

	if result == nil {
		t.Fatal("CreateMap() returned nil")
	}

	expected := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CreateMap() = %v; want %v", result, expected)
	}
}

func TestGetValue(t *testing.T) {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
	}

	tests := []struct {
		key      string
		expected int
	}{
		{"apple", 5},
		{"banana", 3},
		{"nonexistent", 0},
	}

	for _, test := range tests {
		result := GetValue(m, test.key)
		if result != test.expected {
			t.Errorf("GetValue(m, %q) = %d; want %d", test.key, result, test.expected)
		}
	}
}

func TestSetValue(t *testing.T) {
	m := make(map[string]int)
	SetValue(m, "test", 42)

	if m["test"] != 42 {
		t.Errorf("After SetValue(m, \"test\", 42), m[\"test\"] = %d; want 42", m["test"])
	}
}

func TestDeleteKey(t *testing.T) {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
	}

	DeleteKey(m, "apple")

	if _, exists := m["apple"]; exists {
		t.Errorf("After DeleteKey(m, \"apple\"), key \"apple\" still exists")
	}

	if m["banana"] != 3 {
		t.Errorf("After DeleteKey(m, \"apple\"), m[\"banana\"] = %d; want 3", m["banana"])
	}
}

func TestGetKeys(t *testing.T) {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}

	result := GetKeys(m)

	if result == nil {
		t.Fatal("GetKeys() returned nil")
	}

	sort.Strings(result)
	expected := []string{"apple", "banana", "orange"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetKeys(m) = %v; want %v", result, expected)
	}
}
