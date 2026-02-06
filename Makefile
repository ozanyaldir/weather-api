.PHONY: help run setup test

help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup:
	@echo "Installing dependencies..."
	go mod download
	@if [ ! -f .env ]; then \
		echo "Creating .env file from .env.example..."; \
		cp .env.example .env; \
		echo "Please edit .env with your database credentials"; \
	fi
	@echo "Setup complete!"

run:
	@echo "Starting application..."
	go run cmd/server/main.go

test:
	@echo "Running tests..."
	go test -v ./...

fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "Format complete!"
