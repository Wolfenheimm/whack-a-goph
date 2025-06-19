# Whack-a-Goph Test Suite 🐹

This directory contains comprehensive tests for the whack-a-goph project to ensure everything works correctly end-to-end.

## Test Categories

### 1. Integration Tests (`integration_test.go`)

- **End-to-End Workflow**: Tests the complete user journey from start to finish
- **CLI Commands**: Tests all CLI commands (list, run, next, reset, watch, hint)
- **Exercise Loading**: Validates that exercises can be loaded correctly
- **Progress Tracking**: Tests progress persistence and state management
- **Exercise Structure**: Ensures all exercises have proper file structure

### 2. Exercise Category Tests (`exercise_category_test.go`)

- **Category Validation**: Tests each exercise category (intro, functions, structs, etc.)
- **Exercise Structure**: Validates that each exercise has required files and markers
- **Compilation**: Ensures all exercises compile without errors
- **Test Execution**: Verifies that exercise tests can run (even if they fail)

### 3. CLI Tests (`cli_test.go`)

- **Command Functionality**: Tests all CLI commands in isolation
- **Error Handling**: Tests various error conditions and edge cases
- **Progress Integration**: Tests CLI commands with progress tracking
- **Help System**: Tests help commands and documentation

### 4. Performance Tests (`performance_test.go`)

- **Load Testing**: Tests system performance under various loads
- **Concurrency**: Tests concurrent access to progress system
- **Memory Usage**: Basic memory leak detection
- **Benchmarks**: Go benchmark tests for critical operations

## Running Tests

### Quick Start with Makefile

```bash
# From the project root:

# Run all comprehensive tests
make test-all

# Run quick validation tests
make test-quick

# Run specific test categories
make test-integration     # Integration tests
make test-cli            # CLI functionality tests
make test-exercise       # Exercise structure tests
make test-performance    # Performance tests

# Run with verbose output
make test-verbose

# Run benchmarks
make test-bench

# Full project validation (includes basic + comprehensive tests)
make validate
```

### Direct Go Commands

```bash
# Run all tests directly
go test ./tests/...

# Run specific test categories
go test ./tests/ -run TestEndToEndWorkflow     # Integration
go test ./tests/ -run TestExerciseStructure    # Structure
go test ./tests/ -run TestCLI                  # CLI tests
go test ./tests/ -run TestPerformance          # Performance

# Run benchmarks
go test ./tests/ -bench=.

# Verbose output
go test ./tests/... -v
```

## Test Requirements

### Prerequisites

- Go 1.18+ (for generics exercises)
- All project dependencies installed (`go mod download`)
- Write permissions in project directory (for progress file testing)

### Environment Setup

The tests automatically handle:

- Backing up existing progress files
- Creating temporary test files
- Cleaning up after tests complete

### CI/CD Integration

Tests are designed to work in CI environments:

- No interactive input required
- Proper exit codes for success/failure
- Comprehensive error reporting

## Test Scenarios Covered

### 🏗️ Project Structure

- ✅ All exercise files exist
- ✅ All test files exist
- ✅ Exercise files contain "I AM NOT DONE" marker
- ✅ Exercise files contain TODO/Fix me comments
- ✅ All categories have expected exercises

### 🖥️ CLI Functionality

- ✅ CLI tool builds successfully
- ✅ Help commands work correctly
- ✅ List command shows all exercises
- ✅ Run command executes exercises
- ✅ Next command finds next pending exercise
- ✅ Reset command clears exercise progress
- ✅ Watch command starts file monitoring
- ✅ Error handling for invalid commands

### 📊 Progress Management

- ✅ Progress file creation and persistence
- ✅ Exercise completion tracking
- ✅ Progress reset functionality
- ✅ Concurrent access handling
- ✅ Progress file corruption recovery

### 🏃‍♂️ Performance

- ✅ Exercise loading speed
- ✅ Progress operations performance
- ✅ Concurrent operation handling
- ✅ Memory usage validation
- ✅ Large dataset handling

### 🎯 Exercise Categories

- ✅ Intro exercises (intro1, intro2)
- ✅ Variables (variables1)
- ✅ Functions (functions1)
- ✅ Primitive types (primitives1, primitives2)
- ✅ If statements (if1, if2)
- ✅ Structs (structs1, structs2)
- ✅ Strings (strings1, strings2)
- ✅ Vectors/Slices (vecs1)
- ✅ Hashmaps (hashmaps1)
- ✅ Options/Pointers (options1)
- ✅ Error handling (errors1)
- ✅ Enums (enums1, enums2, enums3)
- ✅ Traits/Interfaces (traits1)
- ✅ Tests (tests1)
- ✅ Threads/Goroutines (threads1)
- ✅ Iterators (iterators1)
- ✅ Smart pointers (smart_pointers1)
- ✅ Macros/Code generation (macros1)
- ✅ Generics (generics1)
- ✅ Conversions (conversions1)
- ✅ Advanced (advanced1, advanced2)
- ✅ Modules (modules1)
- ✅ Quizzes (quiz1, quiz2, quiz3, quiz4, quiz_final)

## Expected Test Results

### When Tests Pass ✅

All tests should pass when:

- Project structure is correct
- All files are in place
- CLI builds successfully
- Exercise files compile
- Progress system works correctly

### When Tests Fail ❌

Tests may fail due to:

- Missing exercise files
- Compilation errors in exercises
- CLI build failures
- Progress system issues
- Performance regressions

## Troubleshooting

### Common Issues

1. **CLI Build Fails**: Check that `main.go` and dependencies are correct
2. **Exercise Not Found**: Verify exercise files exist in expected locations
3. **Permission Errors**: Ensure write permissions for progress file
4. **Timeout Errors**: Performance tests may fail on slower systems

### Debug Mode

For detailed debugging:

```bash
go test ./tests/... -v -timeout 30s
```

### Manual Validation

If automated tests fail, manually verify:

1. `go build .` works from project root
2. `./whack-a-goph list` shows exercises
3. Exercise files exist and contain expected markers
4. Progress file can be created/modified

## Contributing

When adding new exercises or features:

1. Add corresponding tests to appropriate test files
2. Update this README with new test scenarios
3. Ensure all tests pass before submitting changes
4. Add performance tests for new functionality if needed

## Test Data

Tests use:

- Temporary progress files (cleaned up automatically)
- Mock exercise data for large-scale testing
- Real exercise files for structure validation
- CLI binary built during test execution

---

🐹 **Happy Testing!** These tests ensure that learners have a smooth experience with whack-a-goph!
