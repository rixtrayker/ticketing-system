.PHONY: all build run test clean generate migrate-up migrate-down

# Default target
all: build

# Build the application
build:
	@echo "Building application..."
	go build -o bin/server ./cmd/server

# Run the application
run:
	@echo "Running application..."
	go run ./cmd/server

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

# Generate GraphQL code
generate:
	@echo "Generating GraphQL code..."
	go run github.com/99designs/gqlgen generate

# Database migrations
migrate-up:
	@echo "Running database migrations up..."
	cd migrations && make up

migrate-down:
	@echo "Running database migrations down..."
	cd migrations && make down

# Development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/99designs/gqlgen@latest
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Help
help:
	@echo "Available targets:"
	@echo "  all            - Build the application (default)"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  test           - Run tests"
	@echo "  clean          - Clean build artifacts"
	@echo "  generate       - Generate GraphQL code"
	@echo "  migrate-up     - Run database migrations up"
	@echo "  migrate-down   - Run database migrations down"
	@echo "  install-tools  - Install development tools"
	@echo "  help           - Show this help message" 