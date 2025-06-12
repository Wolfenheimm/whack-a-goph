.PHONY: help build test clean install reset-all reset list next run watch hint
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
	@echo "  make test          Test all exercises compile"
	@echo ""
	@echo "Exercise Commands:"
	@echo "  make list          List all exercises and progress"
	@echo "  make next          Run next pending exercise"
	@echo "  make run EXERCISE=<name>     Run specific exercise"
	@echo "  make watch EXERCISE=<name>   Watch mode for exercise"  
	@echo "  make hint EXERCISE=<name>    Get hint for exercise"
	@echo ""
	@echo "Reset Commands:"
	@echo "  make reset-all               Reset ALL progress"
	@echo "  make reset EXERCISE=<name>   Reset specific exercise"
	@echo ""
	@echo "Maintenance:"
	@echo "  make clean         Clean build artifacts"
	@echo "  make add-path      Show commands to add Go bin to PATH"
	@echo "  make help          Show this help"
	@echo ""
	@echo "Examples:"
	@echo "  make run EXERCISE=hello"
	@echo "  make watch EXERCISE=functions"
	@echo "  make reset EXERCISE=structs" 