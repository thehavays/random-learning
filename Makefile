# Project binary name
BINARY_NAME=random-learning

# Build options
BUILD_DIR=build
OUTPUT_BINARY=$(BUILD_DIR)/$(BINARY_NAME)

# Go source files (exclude tests)
SRC=$(shell find . -name '*.go' -not -name '*_test.go')

# SQLite database file
DB_FILE=selected_sites.db

# Targets

# Build the Go application
build: $(OUTPUT_BINARY)

$(OUTPUT_BINARY): $(SRC)
	@mkdir -p $(BUILD_DIR)
	go build -o $(OUTPUT_BINARY) main.go

# Run the application
run: build
	./$(OUTPUT_BINARY)

# Run tests
test:
	go test ./...

# Clean build artifacts and SQLite database
clean:
	rm -rf $(BUILD_DIR) $(DB_FILE)

# Install dependencies
deps:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Update dependencies
update:
	go get -u ./...

.PHONY: build run test clean deps fmt lint update
