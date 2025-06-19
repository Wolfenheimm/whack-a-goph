package main

// Macros 1: go:generate Directive and Code Generation 🐹
// Learn about Go's code generation using //go:generate comments
// and build-time code generation patterns.

// I AM NOT DONE

//go:generate echo "Code generation example"

import (
	"fmt"
)

// TODO: Define a Color type
type Color int

const (
	Red Color = iota
	Green
	Blue
	Yellow
)

// TODO: Implement String method for Color that will be "generated"
// For now, implement manually - in real world this would be generated
func (c Color) String() string {
	// Fix me! Return color names: "Red", "Green", "Blue", "Yellow"
	return ""
}

// TODO: Implement GenerateEnumString function
// This simulates what a code generator would do
func GenerateEnumString(enumName string, values []string) string {
	// Fix me! Create a string method implementation for enum
	// Return Go code as string that implements String() method
	return ""
}

// TODO: Implement template for generating code
var enumTemplate = `
// Generated code for {{.EnumName}}
func ({{.Receiver}} {{.EnumName}}) String() string {
	switch {{.Receiver}} {
	{{range $index, $value := .Values}}case {{$index}}: return "{{$value}}"
	{{end}}default: return "Unknown"
	}
}
`

// TODO: Implement ExecuteTemplate function
// Use text/template to generate code from template
func ExecuteTemplate(enumName string, values []string) (string, error) {
	// Fix me! Parse template and execute with data
	return "", fmt.Errorf("not implemented")
}

// TODO: Implement BuildTags function
// Simulate reading build tags (this would be environment-dependent)
func BuildTags() []string {
	// Fix me! Return slice of build tags like ["linux", "amd64"]
	return nil
}

// TODO: Implement IsDebugBuild function
// Check if this is a debug build (simulate build tag checking)
func IsDebugBuild() bool {
	// Fix me! Return true if "debug" build tag is present
	return false
}

// TODO: Implement GenerateTestFile function
// Generate test file content for a given struct
func GenerateTestFile(structName string, fields []string) string {
	// Fix me! Generate basic test file content
	return ""
}

// TODO: Implement ConditionalCode function
// Return different code based on build conditions
func ConditionalCode(condition string) string {
	// Fix me! Return different code based on condition
	switch condition {
	case "production":
		return "// Production code"
	case "development":
		return "// Development code"
	default:
		return "// Default code"
	}
}

func main() {
	// Test color enum
	color := Red
	fmt.Printf("Color: %s\n", color.String())

	// Test code generation
	values := []string{"Red", "Green", "Blue", "Yellow"}
	generated := GenerateEnumString("Color", values)
	fmt.Printf("Generated code length: %d\n", len(generated))

	// Test template execution
	result, err := ExecuteTemplate("Status", []string{"Active", "Inactive"})
	if err == nil {
		fmt.Printf("Template result length: %d\n", len(result))
	}

	// Test build tags
	tags := BuildTags()
	fmt.Printf("Build tags: %v\n", tags)

	// Test debug check
	fmt.Printf("Is debug build: %v\n", IsDebugBuild())

	// Test conditional code
	fmt.Printf("Production code: %s\n", ConditionalCode("production"))
}
