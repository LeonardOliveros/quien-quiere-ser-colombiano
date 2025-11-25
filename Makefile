# Makefile for Quiz App

# Variables
BINARY_NAME=quiz-app
GO=go
GOFLAGS=-v
PORT=8080

# Default target
.DEFAULT_GOAL := help

## help: Display this help message
.PHONY: help
help:
	@echo "Quiz de Naturalización Colombia - Makefile Commands"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@grep -E '^## ' Makefile | sed 's/## /  /'

## build: Build the application
.PHONY: build
build:
	@echo "Building application..."
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) .

## run: Run the application
.PHONY: run
run:
	@echo "Starting application on port $(PORT)..."
	$(GO) run .

## test: Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test ./... -v

## deps: Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GO) mod download
	$(GO) mod tidy

## clean: Clean build files and database
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -f quiz.db

## reset-db: Reset database (delete and recreate)
.PHONY: reset-db
reset-db:
	@echo "Resetting database..."
	rm -f quiz.db
	@echo "Database will be recreated on next run"

## install: Install dependencies and build
.PHONY: install
install: deps build
	@echo "Installation complete!"

## dev: Run in development mode with auto-reload
.PHONY: dev
dev:
	@echo "Starting in development mode..."
	@which air > /dev/null 2>&1 || (echo "Installing air..." && go install github.com/cosmtrek/air@latest)
	air

## docker-build: Build Docker image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t quiz-app:latest .

## docker-run: Run Docker container
.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run -p $(PORT):$(PORT) quiz-app:latest

## lint: Run linter
.PHONY: lint
lint:
	@echo "Running linter..."
	@which golangci-lint > /dev/null 2>&1 || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run

## format: Format code
.PHONY: format
format:
	@echo "Formatting code..."
	$(GO) fmt ./...
	gofmt -w .

## check: Run all checks (format, lint, test)
.PHONY: check
check: format lint test
	@echo "All checks passed!"

## frontend-install: Install frontend dependencies
.PHONY: frontend-install
frontend-install:
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

## frontend-build: Build frontend
.PHONY: frontend-build
frontend-build:
	@echo "Building frontend..."
	cd frontend && npm run build

## frontend-dev: Run frontend development server
.PHONY: frontend-dev
frontend-dev:
	@echo "Starting frontend development server..."
	cd frontend && npm run dev

## dev-full: Run both backend and frontend in development mode
.PHONY: dev-full
dev-full:
	@echo "Starting full development environment..."
	@trap 'kill 0' EXIT; \
	(cd frontend && npm run dev) & \
	$(GO) run .

## build-all: Build both frontend and backend
.PHONY: build-all
build-all: frontend-build build
	@echo "Full build complete!"

## clean-sessions: Delete all game sessions from database
.PHONY: clean-sessions
clean-sessions:
	@echo "Cleaning game sessions..."
	@sqlite3 quiz.db "DELETE FROM game_sessions; DELETE FROM game_answers;" 2>/dev/null || echo "No database found"
	@echo "Sessions cleaned!"

## clean-old-sessions: Delete old sessions with wrong question count
.PHONY: clean-old-sessions
clean-old-sessions:
	@echo "Cleaning old sessions..."
	@sqlite3 quiz.db "DELETE FROM game_sessions WHERE total_questions = 753;" 2>/dev/null || echo "No database found"
	@echo "Old sessions cleaned!"

## db-stats: Show database statistics
.PHONY: db-stats
db-stats:
	@echo "=== Database Statistics ==="
	@echo ""
	@sqlite3 quiz.db "SELECT 'Total Questions: ' || COUNT(*) FROM questions;" 2>/dev/null || echo "No database found"
	@echo ""
	@echo "Questions by Category:"
	@sqlite3 quiz.db "SELECT '  ' || category || ': ' || COUNT(*) FROM questions GROUP BY category ORDER BY category;" 2>/dev/null
	@echo ""
	@sqlite3 quiz.db "SELECT 'Active Sessions: ' || COUNT(*) FROM game_sessions WHERE status = 'ACTIVE';" 2>/dev/null
	@sqlite3 quiz.db "SELECT 'Total Sessions: ' || COUNT(*) FROM game_sessions;" 2>/dev/null

## setup: Complete setup (install dependencies and build)
.PHONY: setup
setup: deps frontend-install build-all
	@echo "Setup complete! Run 'make dev-full' to start development"
