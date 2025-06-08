# Variables
BINARY_NAME=last_week
BINARY_DIR=./bin
SOURCE_DIR=./cmd

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	@mkdir -p $(BINARY_DIR)
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(SOURCE_DIR)/main.go

# Run the binary
.PHONY: run
run: build
	@$(BINARY_DIR)/$(BINARY_NAME)

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BINARY_DIR)

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build  - Build the binary to ./bin/"
	@echo "  run    - Build and run the binary"
	@echo "  clean  - Remove build artifacts"
	@echo "  help   - Show this help message" 