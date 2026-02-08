.PHONY: help setup build run test lint fmt

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup: ## Install dependencies and set up environment
	@echo "Installing dependencies..."
	go mod download
	@if [ ! -f .env ]; then \
		echo "Creating .env file from .env.example..."; \
		cp .env.example .env; \
		echo "Please edit .env with your database credentials"; \
	fi
	@echo "Setup complete!"

build: ## Build the application
	@echo "Building application..."
	go build -o bin/weather-api cmd/server/main.go
	@echo "Build complete! Binary: bin/weather-api"

run: ## Run the application
	@echo "Starting application..."
	go run cmd/server/main.go

test: ## Run tests
	@echo "Running tests..."
	go test ./test/...

lint: ## Run linter
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install it from https://golangci-lint.run/usage/install/"; \
	fi

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	@echo "Format complete!"

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t weather-api:latest .
	@echo "Docker image built!"

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker-compose up --build

dev: ## Run in development mode with auto-reload (requires air)
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "air not installed. Install it: go install github.com/cosmtrek/air@latest"; \
		echo "Running without auto-reload..."; \
		make run; \
	fi

db-create: ## Create database (requires mysql client)
	@echo "Creating database..."
	@mysql -h $${DB_HOST:-localhost} -u root -p -e "CREATE DATABASE IF NOT EXISTS $${DB_NAME:-weather-db} CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
	@echo "Database created!"

db-drop: ## Drop database (requires mysql client)
	@echo "Dropping database..."
	@mysql -h $${DB_HOST:-localhost} -u root -p -e "DROP DATABASE IF EXISTS $${DB_NAME:-weather-db};"
	@echo "Database dropped!"