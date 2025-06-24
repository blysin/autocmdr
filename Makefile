# Makefile for LangChain Chat App

# Variables
BINARY_NAME=autocmdr
MAIN_PATH=./cmd/autocmdr
BUILD_DIR=./build
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "v1.0.0")
GIT_COMMIT=$(shell git rev-parse HEAD 2>/dev/null || echo "unknown")
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GO_VERSION=$(shell go version | cut -d' ' -f3)

# Build flags
LDFLAGS=-ldflags "-X github.com/blysin/autocmdr/pkg/version.Version=$(VERSION) \
                  -X github.com/blysin/autocmdr/pkg/version.GitCommit=$(GIT_COMMIT) \
                  -X github.com/blysin/autocmdr/pkg/version.BuildDate=$(BUILD_DATE)"

# Default target
.PHONY: all
all: clean test build

# Build the application
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# Build for multiple platforms
.PHONY: build-all
build-all: clean
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage: test
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmarks
.PHONY: bench
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Install the application
.PHONY: install
install:
	@echo "Installing $(BINARY_NAME)..."
	go install $(LDFLAGS) $(MAIN_PATH)

# Run the application
.PHONY: run
run:
	@echo "Running $(BINARY_NAME)..."
	go run $(MAIN_PATH)

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .

# Lint code
.PHONY: lint
lint:
	@echo "Linting code..."
	golangci-lint run

# Vet code
.PHONY: vet
vet:
	@echo "Vetting code..."
	go vet ./...

# Update dependencies
.PHONY: deps
deps:
	@echo "Updating dependencies..."
	go mod tidy
	go mod download

# Security check
.PHONY: security
security:
	@echo "Running security checks..."
	gosec ./...

# Generate documentation
.PHONY: docs
docs:
	@echo "Generating documentation..."
	godoc -http=:6060

# Docker build
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME):$(VERSION) .

# Docker run
.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run -it --rm $(BINARY_NAME):$(VERSION)

# Release preparation
.PHONY: release-prep
release-prep: clean test lint vet security build-all
	@echo "Release preparation complete"

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all          - Clean, test, and build"
	@echo "  build        - Build the application"
	@echo "  build-all    - Build for multiple platforms"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  bench        - Run benchmarks"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install the application"
	@echo "  run          - Run the application"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  vet          - Vet code"
	@echo "  deps         - Update dependencies"
	@echo "  security     - Run security checks"
	@echo "  docs         - Generate documentation"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  release-prep - Prepare for release"
	@echo "  help         - Show this help"
