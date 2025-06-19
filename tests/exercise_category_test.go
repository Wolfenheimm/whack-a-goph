package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/raithbheart/whack-a-goph/internal"
)

// TestIntroExercises tests the intro exercises
func TestIntroExercises(t *testing.T) {
	testExerciseCategory(t, "intro", []string{"intro1", "intro2"})
}

// TestVariableExercises tests the variable exercises
func TestVariableExercises(t *testing.T) {
	testExerciseCategory(t, "variables", []string{"variables1"})
}

// TestFunctionExercises tests the function exercises
func TestFunctionExercises(t *testing.T) {
	testExerciseCategory(t, "functions", []string{"functions1"})
}

// TestPrimitiveTypeExercises tests the primitive type exercises
func TestPrimitiveTypeExercises(t *testing.T) {
	testExerciseCategory(t, "primitive_types", []string{"primitives1", "primitives2"})
}

// TestIfExercises tests the if statement exercises
func TestIfExercises(t *testing.T) {
	testExerciseCategory(t, "if", []string{"if1", "if2"})
}

// TestStructExercises tests the struct exercises
func TestStructExercises(t *testing.T) {
	testExerciseCategory(t, "structs", []string{"structs1", "structs2"})
}

// TestStringExercises tests the string exercises
func TestStringExercises(t *testing.T) {
	testExerciseCategory(t, "strings", []string{"strings1", "strings2"})
}

// TestVecExercises tests the vector/slice exercises
func TestVecExercises(t *testing.T) {
	testExerciseCategory(t, "vecs", []string{"vecs1"})
}

// TestHashmapExercises tests the hashmap exercises
func TestHashmapExercises(t *testing.T) {
	testExerciseCategory(t, "hashmaps", []string{"hashmaps1"})
}

// TestOptionExercises tests the option exercises
func TestOptionExercises(t *testing.T) {
	testExerciseCategory(t, "options", []string{"options1"})
}

// TestErrorHandlingExercises tests the error handling exercises
func TestErrorHandlingExercises(t *testing.T) {
	testExerciseCategory(t, "error_handling", []string{"errors1"})
}

// TestEnumExercises tests the enum exercises
func TestEnumExercises(t *testing.T) {
	testExerciseCategory(t, "enums", []string{"enums1", "enums2", "enums3"})
}

// TestTraitExercises tests the trait exercises
func TestTraitExercises(t *testing.T) {
	testExerciseCategory(t, "traits", []string{"traits1"})
}

// TestTestExercises tests the test exercises
func TestTestExercises(t *testing.T) {
	testExerciseCategory(t, "tests", []string{"tests1"})
}

// TestThreadExercises tests the thread exercises
func TestThreadExercises(t *testing.T) {
	testExerciseCategory(t, "threads", []string{"threads1"})
}

// TestIteratorExercises tests the iterator exercises
func TestIteratorExercises(t *testing.T) {
	testExerciseCategory(t, "iterators", []string{"iterators1"})
}

// TestSmartPointerExercises tests the smart pointer exercises
func TestSmartPointerExercises(t *testing.T) {
	testExerciseCategory(t, "smart_pointers", []string{"smart_pointers1"})
}

// TestMacroExercises tests the macro exercises
func TestMacroExercises(t *testing.T) {
	testExerciseCategory(t, "macros", []string{"macros1"})
}

// TestGenericsExercises tests the generics exercises
func TestGenericsExercises(t *testing.T) {
	testExerciseCategory(t, "generics", []string{"generics1"})
}

// TestConversionExercises tests the conversion exercises
func TestConversionExercises(t *testing.T) {
	testExerciseCategory(t, "conversions", []string{"conversions1"})
}

// TestAdvancedExercises tests the advanced exercises
func TestAdvancedExercises(t *testing.T) {
	testExerciseCategory(t, "advanced", []string{"advanced1", "advanced2"})
}

// TestModuleExercises tests the module exercises
func TestModuleExercises(t *testing.T) {
	testExerciseCategory(t, "modules", []string{"modules1"})
}

// testExerciseCategory is a helper function that tests a category of exercises
func testExerciseCategory(t *testing.T, category string, expectedExercises []string) {
	t.Run(category, func(t *testing.T) {
		exercises, err := internal.GetExercises()
		if err != nil {
			t.Fatalf("Failed to load exercises: %v", err)
		}

		// Find exercises in this category
		categoryExercises := make(map[string]internal.Exercise)
		for _, exercise := range exercises {
			if strings.Contains(exercise.Path, category) {
				categoryExercises[exercise.Name] = exercise
			}
		}

		// Check that all expected exercises exist
		for _, expectedName := range expectedExercises {
			if _, found := categoryExercises[expectedName]; !found {
				t.Errorf("Expected exercise %s not found in category %s", expectedName, category)
			}
		}

		// Test that each exercise in this category has valid structure
		for name, exercise := range categoryExercises {
			t.Run(name, func(t *testing.T) {
				testExerciseStructure(t, exercise)
				testExerciseCompiles(t, exercise)
			})
		}
	})
}

// testExerciseStructure checks that an exercise has the proper structure
func testExerciseStructure(t *testing.T, exercise internal.Exercise) {
	// Adjust path to be relative to tests directory
	exercisePath := filepath.Join("..", exercise.Path)

	// Check that the exercise file exists
	if _, err := os.Stat(exercisePath); os.IsNotExist(err) {
		t.Errorf("Exercise file does not exist: %s", exercisePath)
		return
	}

	// Check that the corresponding test file exists
	dir := filepath.Dir(exercisePath)
	base := filepath.Base(exercisePath)
	testFile := filepath.Join(dir, strings.Replace(base, ".go", "_test.go", 1))

	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Errorf("Test file does not exist: %s", testFile)
		return
	}

	// Check that the exercise contains the "I AM NOT DONE" marker
	content, err := os.ReadFile(exercisePath)
	if err != nil {
		t.Errorf("Failed to read exercise file: %v", err)
		return
	}

	if !strings.Contains(string(content), "I AM NOT DONE") {
		t.Errorf("Exercise file should contain 'I AM NOT DONE' marker: %s", exercisePath)
	}

	// Check that the exercise has TODO comments
	if !strings.Contains(string(content), "TODO") && !strings.Contains(string(content), "Fix me") {
		t.Errorf("Exercise file should contain TODO or 'Fix me' comments: %s", exercisePath)
	}
}

// testExerciseCompiles checks that an exercise compiles (even if tests fail)
func testExerciseCompiles(t *testing.T, exercise internal.Exercise) {
	// Adjust directory path to be relative to tests directory
	dir := filepath.Join("..", filepath.Dir(exercise.Path))

	// Run go build to check if it compiles
	cmd := exec.Command("go", "build", ".")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("Exercise %s does not compile: %v\nOutput: %s", exercise.Name, err, output)
	}
}

// TestExerciseTestsRun verifies that all exercise tests can run (even if they fail)
func TestExerciseTestsRun(t *testing.T) {
	exercises, err := internal.GetExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	for _, exercise := range exercises {
		t.Run(exercise.Name+"_tests_run", func(t *testing.T) {
			// Adjust directory path to be relative to tests directory
			dir := filepath.Join("..", filepath.Dir(exercise.Path))

			// Run go test to see if tests execute
			cmd := exec.Command("go", "test", "-v", ".")
			cmd.Dir = dir
			output, err := cmd.CombinedOutput()

			// We expect tests to run, but they might fail (that's the point of exercises)
			// We're just checking that they execute without compilation errors
			if err != nil {
				outputStr := string(output)
				// Check if it's a compilation error vs test failure
				if strings.Contains(outputStr, "build failed") ||
					strings.Contains(outputStr, "syntax error") ||
					strings.Contains(outputStr, "undefined:") {
					t.Errorf("Exercise %s has compilation errors: %v\nOutput: %s", exercise.Name, err, outputStr)
				}
				// Test failures are expected for incomplete exercises
			}
		})
	}
}

// TestQuizStructure tests that all quizzes are properly structured
func TestQuizStructure(t *testing.T) {
	// Adjust path to be relative to tests directory
	quizDir := "../exercises/quizzes"
	quizFiles := []string{"quiz1.go", "quiz2.go", "quiz3.go", "quiz4.go", "quiz_final.go"}

	for _, quizFile := range quizFiles {
		t.Run(quizFile, func(t *testing.T) {
			quizPath := filepath.Join(quizDir, quizFile)

			// Check that quiz file exists
			if _, err := os.Stat(quizPath); os.IsNotExist(err) {
				t.Errorf("Quiz file does not exist: %s", quizPath)
				return
			}

			// Check that corresponding test file exists
			testFile := filepath.Join(quizDir, strings.Replace(quizFile, ".go", "_test.go", 1))
			if _, err := os.Stat(testFile); os.IsNotExist(err) {
				t.Errorf("Quiz test file does not exist: %s", testFile)
				return
			}

			// Check that quiz compiles
			cmd := exec.Command("go", "build", quizFile)
			cmd.Dir = quizDir
			if err := cmd.Run(); err != nil {
				t.Errorf("Quiz %s does not compile: %v", quizFile, err)
			}
		})
	}
}
