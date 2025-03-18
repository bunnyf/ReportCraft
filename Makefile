.PHONY: all build clean test run

# Build settings
BINARY_NAME=genrep
BINARY_DIR=bin
SRC_MAIN=cmd/genrep/main.go
BUILD_FLAGS=-ldflags="-s -w"
PLUGINS_DIR=plugins

# Test settings
TEST_FLAGS=-v

# Run settings
CONFIG_FILE=examples/waveform-report.json

# Default target
all: clean build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BINARY_DIR)
	@go build $(BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME) $(SRC_MAIN)
	@echo "Build complete!"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BINARY_DIR)
	@echo "Clean complete!"

# Run tests
test:
	@echo "Running tests..."
	@go test $(TEST_FLAGS) ./...
	@echo "Tests complete!"

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BINARY_DIR)/$(BINARY_NAME) --config $(CONFIG_FILE)

# Create output directory
output:
	@mkdir -p output

# Create plugins directory
plugins:
	@mkdir -p $(PLUGINS_DIR)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@echo "Dependencies installed!"

# Build a plugin (usage: make plugin PLUGIN=myplugin)
plugin:
	@echo "Building plugin $(PLUGIN)..."
	@go build -buildmode=plugin -o $(PLUGINS_DIR)/$(PLUGIN).so plugins/$(PLUGIN)/main.go
	@echo "Plugin $(PLUGIN) built!"

# Help
help:
	@echo "Available targets:"
	@echo "  all      - Clean and build the application"
	@echo "  build    - Build the application"
	@echo "  clean    - Clean build artifacts"
	@echo "  test     - Run tests"
	@echo "  run      - Build and run the application with the default config"
	@echo "  output   - Create the output directory"
	@echo "  plugins  - Create the plugins directory"
	@echo "  deps     - Install dependencies"
	@echo "  plugin   - Build a plugin (usage: make plugin PLUGIN=myplugin)"
	@echo "  help     - Show this help message"
