.PHONY: help build test clean install reset-all reset list next run watch hint test-all test-integration test-cli test-exercise test-performance test-quick test-verbose test-bench setup-tests clean-tests validate test-ci coverage cleanup test-solutions
.DEFAULT_GOAL := help

# Build the binary
build:
	@echo "🔨 Building whack-a-goph..."
	@go build -o whack-a-goph .
	@echo "✅ Build complete! Run './whack-a-goph' to get started."

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	@go mod tidy
	@echo "✅ Dependencies installed!"

# Install globally
install: build
	@echo "🌍 Installing globally..."
	@go install .
	@echo "✅ whack-a-goph installed globally!"
	@echo ""
	@echo "Testing global installation..."
	@if command -v whack-a-goph >/dev/null 2>&1; then \
		echo "🎉 Global installation successful! You can use 'whack-a-goph' anywhere."; \
	else \
		echo "⚠️  Command 'whack-a-goph' not found in PATH."; \
		echo "Add Go bin to your PATH with:"; \
		echo "  echo 'export PATH=\$$PATH:\$$(go env GOPATH)/bin' >> ~/.zshrc"; \
		echo "  source ~/.zshrc"; \
		echo "Or use: export PATH=\$$PATH:\$$(go env GOPATH)/bin"; \
	fi

# Test all exercises (build tests to verify they compile)
test:
	@echo "🧪 Testing all exercises..."
	@for dir in exercises/*/; do \
		if [ -f "$$dir"*.go ]; then \
			echo "Testing $$dir..."; \
			cd "$$dir" && go test -c > /dev/null 2>&1 && echo "  ✅ $$dir compiles" || echo "  ❌ $$dir has issues"; \
			cd ../..; \
		fi \
	done

# ===============================================
# 🧪 COMPREHENSIVE TEST SUITE
# ===============================================

# Run all comprehensive tests
test-all: setup-tests
	@echo "🧪 Running comprehensive test suite..."
	go test ./tests/... -timeout 60s

# Run quick validation tests
test-quick: setup-tests
	@echo "⚡ Running quick validation tests..."
	go test ./tests/ -run "TestCLIBuild|TestCLIHelp|TestEndToEndWorkflow" -timeout 30s

# Run integration tests
test-integration: setup-tests
	@echo "🔗 Running integration tests..."
	go test ./tests/ -run "TestEndToEnd|TestExerciseStructure|TestProgressPersistence" -timeout 60s

# Run CLI tests
test-cli: setup-tests
	@echo "💻 Running CLI tests..."
	go test ./tests/ -run "TestCLI" -timeout 45s

# Run exercise category tests
test-exercise: setup-tests
	@echo "📚 Running exercise category tests..."
	go test ./tests/ -run "Exercise" -timeout 60s

# Run performance tests
test-performance: setup-tests
	@echo "🏃‍♂️ Running performance tests..."
	go test ./tests/ -run "Performance|TestMemory|TestConcurrent" -timeout 120s

# Run solution validation tests
test-solutions: setup-tests
	@echo "🎯 Running comprehensive solution validation tests..."
	go test ./tests/ -run "TestComprehensiveExerciseSolutions|TestSolutionWorkflow|TestExerciseCanBeCompleted|TestRandomExerciseSampling" -timeout 180s

# Run tests with verbose output
test-verbose: setup-tests
	@echo "🔍 Running tests with verbose output..."
	go test ./tests/... -v -timeout 60s

# Run benchmark tests
test-bench: setup-tests
	@echo "📊 Running benchmark tests..."
	go test ./tests/ -bench=. -benchmem

# Setup test environment
setup-tests:
	@echo "🛠️  Setting up test environment..."
	@mkdir -p tests/test_data
	@go mod download

# Clean test artifacts
clean-tests:
	@echo "🧹 Cleaning test artifacts..."
	@rm -rf tests/test_data/
	@rm -f whack-a-goph-test
	@rm -f .whack-a-goph-progress.json.test-backup
	@go clean -testcache

# Full project validation
validate: clean clean-tests setup-tests test test-all test-solutions
	@echo ""
	@echo "✅ Project validation complete!"
	@echo "🐹 Ready for gopher whacking!"

# Run tests in CI mode
test-ci: setup-tests
	@echo "🤖 Running tests in CI mode..."
	go test ./tests/... -timeout 120s -race -coverprofile=coverage.out

# Generate test coverage report
coverage: test-ci
	@echo "📈 Generating coverage report..."
	go tool cover -html=coverage.out -o tests/coverage.html
	@echo "Coverage report generated: tests/coverage.html"

# Reset ALL progress (removes progress file completely)
reset-all:
	@echo "🗑️  Resetting ALL progress..."
	@if [ -f .whack-a-goph-progress.json ]; then \
		rm .whack-a-goph-progress.json; \
		echo "✅ All progress reset! All gophers are back in their holes! 🕳️"; \
	else \
		echo "ℹ️  No progress file found - you're already starting fresh!"; \
	fi

# Reset a specific exercise
# Usage: make reset EXERCISE=hello
reset:
	@if [ -z "$(EXERCISE)" ]; then \
		echo "❌ Please specify an exercise: make reset EXERCISE=hello"; \
		exit 1; \
	fi
	@echo "🔄 Resetting exercise: $(EXERCISE)..."
	@if [ ! -f whack-a-goph ]; then \
		echo "⚠️  Binary not found, building first..."; \
		$(MAKE) build; \
	fi
	@./whack-a-goph reset $(EXERCISE)

# Quick commands that require the binary
list: build
	@./whack-a-goph list

next: build  
	@./whack-a-goph next

# Run a specific exercise
# Usage: make run EXERCISE=hello
run: build
	@if [ -z "$(EXERCISE)" ]; then \
		echo "❌ Please specify an exercise: make run EXERCISE=hello"; \
		exit 1; \
	fi
	@./whack-a-goph run $(EXERCISE)

# Watch a specific exercise  
# Usage: make watch EXERCISE=hello
watch: build
	@if [ -z "$(EXERCISE)" ]; then \
		echo "❌ Please specify an exercise: make watch EXERCISE=hello"; \
		exit 1; \
	fi
	@./whack-a-goph watch $(EXERCISE)

# Get hint for a specific exercise
# Usage: make hint EXERCISE=hello  
hint: build
	@if [ -z "$(EXERCISE)" ]; then \
		echo "❌ Please specify an exercise: make hint EXERCISE=hello"; \
		exit 1; \
	fi
	@./whack-a-goph hint $(EXERCISE)

# Clean build artifacts
clean:
	@echo "🧹 Cleaning up..."
	@rm -f whack-a-goph
	@echo "✅ Clean complete!"

# Comprehensive cleanup - remove all junk files and build artifacts
cleanup:
	@echo "🗑️  Performing comprehensive cleanup..."
	@echo "Removing main binaries..."
	@rm -f whack-a-goph whack-a-goph.exe whack-a-goph-test
	@echo "Removing exercise binaries..."
	@rm -f exercises/advanced/advanced exercises/enums/enums exercises/generics/generics
	@rm -f exercises/iterators/iterators exercises/macros/macros exercises/modules/modules
	@rm -f exercises/smart_pointers/smart_pointers exercises/tests/tests exercises/threads/threads
	@rm -f exercises/traits/traits
	@echo "Removing test artifacts..."
	@rm -rf tests/test_data/ tests/coverage.html
	@rm -f .whack-a-goph-progress.json.test-backup
	@rm -f coverage.out coverage.txt
	@echo "Removing temp and backup files..."
	@find . -name "*.tmp" -delete 2>/dev/null || true
	@find . -name "*.temp" -delete 2>/dev/null || true
	@find . -name "*.bak" -delete 2>/dev/null || true
	@find . -name "*.backup" -delete 2>/dev/null || true
	@find . -name "*.log" -delete 2>/dev/null || true
	@find . -name ".DS_Store" -delete 2>/dev/null || true
	@find . -name "._*" -delete 2>/dev/null || true
	@find . -name "Thumbs.db" -delete 2>/dev/null || true
	@find . -name "ehthumbs.db" -delete 2>/dev/null || true
	@echo "Cleaning Go test cache..."
	@go clean -testcache -cache -modcache -fuzzcache 2>/dev/null || true
	@echo ""
	@echo "✨ Deep cleanup complete! Project is squeaky clean!"
	@echo "🐹 All gophers have been thoroughly whacked out of the system!"

# Add Go bin to PATH (for current session)
add-path:
	@echo "Adding Go bin to PATH for current session..."
	@echo "Run: export PATH=\$$PATH:\$$(go env GOPATH)/bin"
	@echo "To make permanent, add to your shell config:"
	@echo "  echo 'export PATH=\$$PATH:\$$(go env GOPATH)/bin' >> ~/.zshrc"
	@echo "  source ~/.zshrc"

# Setup everything for first time users
setup: deps build
	@echo ""
	@echo "🎉 Setup complete! You're ready to start whacking gophers!"
	@echo ""
	@echo "Quick start:"
	@echo "  make list          # See all exercises"
	@echo "  make next          # Start with first exercise"  
	@echo "  make run EXERCISE=hello    # Run specific exercise"
	@echo "  make watch EXERCISE=hello  # Watch mode"
	@echo "  make reset-all     # Reset all progress"
	@echo ""

# Show help
help:
	@echo "🐹 Whack-a-Goph Makefile Commands"
	@echo ""
	@echo "Setup & Build:"
	@echo "  make setup         Setup everything for first time"
	@echo "  make deps          Install Go dependencies"
	@echo "  make build         Build the binary" 
	@echo "  make install       Install binary globally"
	@echo ""
	@echo "Testing:"
	@echo "  make test          Test exercise compilation (basic)"
	@echo "  make test-all      Run comprehensive test suite"
	@echo "  make test-quick    Run quick validation tests"
	@echo "  make test-integration  Run integration tests"
	@echo "  make test-cli      Run CLI functionality tests"
	@echo "  make test-exercise Run exercise structure tests"
	@echo "  make test-solutions Validate exercise solutions work"
	@echo "  make test-performance  Run performance tests"
	@echo "  make test-verbose  Run tests with verbose output"
	@echo "  make test-bench    Run benchmark tests"
	@echo "  make validate      Full project validation"
	@echo ""
	@echo "Exercise Commands:"
	@echo "  make list          List all exercises and progress"
	@echo "  make next          Run next pending exercise"
	@echo "  make run EXERCISE=<n>     Run specific exercise"
	@echo "  make watch EXERCISE=<n>   Watch mode for exercise"  
	@echo "  make hint EXERCISE=<n>    Get hint for exercise"
	@echo ""
	@echo "Reset Commands:"
	@echo "  make reset-all               Reset ALL progress"
	@echo "  make reset EXERCISE=<n>   Reset specific exercise"
	@echo ""
	@echo "Maintenance:"
	@echo "  make clean         Clean build artifacts"
	@echo "  make clean-tests   Clean test artifacts"
	@echo "  make cleanup       Deep cleanup - remove ALL junk files"
	@echo "  make add-path      Show commands to add Go bin to PATH"
	@echo "  make help          Show this help"
	@echo ""
	@echo "Examples:"
	@echo "  make run EXERCISE=hello"
	@echo "  make watch EXERCISE=functions"
	@echo "  make reset EXERCISE=structs" 