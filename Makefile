APP_NAME := fs
BUILD_DIR := bin
MAIN_FILE := main.go

all: deps build

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME)

run:
	@go run $(MAIN_FILE)

test:
	@go test -v ./...

cover:
	@go test -cover ./...

clean:
	@rm -rf $(BUILD_DIR)

format:
	@go fmt ./...

deps:
	@go mod tidy

help:
	@echo "Usage: make [target]"
	@echo "Available targets:"
	@echo "  build      Build the application"
	@echo "  run        Run the application"
	@echo "  test       Run all tests"
	@echo "  cover      Run tests with coverage"
	@echo "  clean      Clean up build artifacts"
	@echo "  format     Format the code"
	@echo "  deps       Install dependencies"
