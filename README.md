# 🐹 Whack-a-Goph

Small exercises to get you used to reading and writing Go code!

## About

Whack-a-Goph is a collection of small Go exercises designed to get you used to reading and writing Go code. Each broken exercise is like a gopher popping up from a hole - fix the code to make the tests pass and "whack" the gopher!

This project is perfect for:

- Go beginners who want hands-on practice
- Developers coming from other languages
- Anyone wanting to brush up on Go fundamentals
- Learning through interactive, immediate feedback

## Installation

### Prerequisites

- Go 1.21 or later installed on your system
- Basic familiarity with the command line

### Install Whack-a-Goph

**Option 1: Quick Setup with Make**

```bash
# Clone the repository
git clone https://github.com/yourusername/whack-a-goph.git
cd whack-a-goph

# Setup everything (install deps + build)
make setup

# See all available commands
make help
```

**Option 2: Manual Setup**

```bash
# Clone the repository
git clone https://github.com/yourusername/whack-a-goph.git
cd whack-a-goph

# Install dependencies
go mod tidy

# Build the binary
go build -o whack-a-goph .

# Test that it works
./whack-a-goph

```

## Quick Start

**Option 1: Using Make Commands (Handles everything)**

```bash
# See all exercises and their status
make list

# Start with the first exercise
make next

# Run a specific exercise
make run EXERCISE=hello

# Get a hint for an exercise
make hint EXERCISE=hello

# Watch mode - automatically runs tests when you save changes
make watch EXERCISE=hello

# Reset a specific exercise
make reset EXERCISE=hello

# Reset ALL progress (start completely fresh)
make reset-all
```

**Option 2: Using Local Binary (Simple)**

```bash
# Use the local binary with ./ prefix
./whack-a-goph list
./whack-a-goph next
./whack-a-goph run hello
./whack-a-goph hint hello
./whack-a-goph watch hello
```

**Option 3: Global Installation (No ./ needed)**

```bash
# Install globally (optional - modifies your PATH)
make install
# or: go install .

# Then use without ./ prefix
whack-a-goph list
whack-a-goph next
whack-a-goph run hello
```

> **Choose your preference:** Use `make` commands for convenience, `./whack-a-goph` for simplicity, or install globally if you don't mind adding to your PATH.

## How It Works

1. **List exercises**: See all available exercises and your progress
2. **Pick an exercise**: Each exercise is a broken Go program in the `exercises/` directory
3. **Fix the code**: Edit the `.go` files to make the tests pass
4. **Run tests**: Use `./whack-a-goph run <name>` or `./whack-a-goph next` to test your solution
5. **Get feedback**: See immediate results with helpful error messages and hints
6. **Progress tracking**: Your progress is automatically saved

## Exercise Topics

- **Basics** - Hello world, variables, and basic types
- **Functions** - Function definitions, parameters, and return values
- **Control Flow** - If statements, loops, and conditional logic
- **Structs** - Defining structs and methods
- **Collections** - Slices, arrays, and maps
- **Interfaces** - Interface definitions and implementations
- **Concurrency** - Goroutines and channels

## Commands

- `./whack-a-goph` - Show welcome message and basic usage
- `./whack-a-goph list` - List all exercises and their completion status
- `./whack-a-goph run <exercise>` - Run a specific exercise
- `./whack-a-goph next` - Run the next pending exercise
- `./whack-a-goph watch [exercise]` - Watch mode for automatic testing
- `./whack-a-goph hint <exercise>` - Get a hint for an exercise
- `./whack-a-goph reset <exercise>` - Reset progress for an exercise
- `./whack-a-goph help` - Show detailed help information

_Note: Commands shown use local binary. Remove `./` if you've installed globally._

## Example Session

```bash
$ ./whack-a-goph list
🐹  WHACK-A-GOPH!  🐹

🐹 Exercise List
────────────────────────────────────────────────────────────
hello           🕳️  HOLE Get started with a simple hello function
basics          🕳️  HOLE Learn about variables and basic types
functions       🕳️  HOLE Write functions with parameters and return values
...

$ ./whack-a-goph next
🐹 Next exercise: hello
Description: Get started with a simple hello function
File: exercises/01_hello/hello.go

         🐹 *WHOOPS!*

✗ Exercise failed!

Test output:
--- FAIL: TestHello (0.00s)
    hello_test.go:8: Hello() = ""; want "Hello, Gopher!"

💡 Hint: Return the string 'Hello, Gopher!' from the Hello() function

# Edit exercises/01_hello/hello.go to fix the function...

$ ./whack-a-goph next
🐹 Next exercise: hello
Description: Get started with a simple hello function
File: exercises/01_hello/hello.go

         🐹 *WHACKED!*

✓ Exercise passed! 🎉
```

## Progress Tracking

Your progress is automatically saved in `.whack-a-goph-progress.json`. This file tracks which exercises you've completed, so you can continue where you left off.

## Watch Mode

Watch mode is perfect for iterative development:

```bash
./whack-a-goph watch hello
```

This will:

- Run the exercise immediately
- Watch the file for changes
- Automatically re-run tests when you save
- Show live feedback as you work

Press `Ctrl+C` to exit watch mode.
