
APP_NAME := user_auth
SRC_DIR := ./cmd/$(APP_NAME)

GO := go
GOTEST := $(GO) test
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean
GOVET := $(GO) vet
GOMOD := $(GO) mod
GOLINT := golint

BUILD_DIR := ./bin

all: build

build:
	@echo "Building $(APP_NAME)..."
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) .

run: build
	@echo "Running $(APP_NAME)..."
	$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "Running tests..."
	$(GOTEST) ./...

lint:
	@echo "Linting..."
	$(GOLINT) ./...

vet:
	@echo "Running go vet..."
	$(GOVET) ./...

clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/$(APP_NAME)

tidy:
	@echo "Tidying Go modules..."
	$(GOMOD) tidy

help:
	@echo "Available targets:"
	@echo "  make build      - Build the project"
	@echo "  make run        - Build and run the project"
	@echo "  make test       - Run tests"
	@echo "  make lint       - Run linter"
	@echo "  make vet        - Run go vet for static analysis"
	@echo "  make clean      - Remove generated binaries and cache"
	@echo "  make tidy       - Clean up unused Go modules"

.PHONY: all build run test lint vet clean tidy help
