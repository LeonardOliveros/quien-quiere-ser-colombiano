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

# ===== AWS serverless (Lambda + DynamoDB + API Gateway + S3 + CloudFront) =====

DDB_PORT=8000
DDB_ENDPOINT=http://localhost:$(DDB_PORT)

## dynamodb-local: Start DynamoDB Local on :8000 (Docker, or java + DDB_LOCAL_DIR=/path/to/dynamodb_local)
.PHONY: dynamodb-local
dynamodb-local:
	@docker run --rm -d --name quiz-ddb -p $(DDB_PORT):8000 amazon/dynamodb-local >/dev/null 2>&1 && \
		echo "DynamoDB Local (docker) on :$(DDB_PORT)" || \
	( [ -n "$(DDB_LOCAL_DIR)" ] && cd "$(DDB_LOCAL_DIR)" && \
		(java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -inMemory -port $(DDB_PORT) > /tmp/quiz-ddb.log 2>&1 &) && \
		echo "DynamoDB Local (java) on :$(DDB_PORT)" || \
		echo "Docker not available. Install Docker, or download DynamoDB Local and set DDB_LOCAL_DIR" )

## dynamodb-local-stop: Stop DynamoDB Local
.PHONY: dynamodb-local-stop
dynamodb-local-stop:
	@docker stop quiz-ddb >/dev/null 2>&1 || pkill -f DynamoDBLocal.jar || true
	@echo "DynamoDB Local stopped"

## test-integration: Run the storage conformance suite against DynamoDB Local
.PHONY: test-integration
test-integration:
	DYNAMODB_TEST_ENDPOINT=$(DDB_ENDPOINT) $(GO) test ./internal/storage/dynamodb/ -run TestConformance -v

## run-ddb: Run the app locally against DynamoDB Local (table quiz-local, auto-created + seeded)
.PHONY: run-ddb
run-ddb:
	DB_DRIVER=dynamodb DYNAMODB_TABLE=quiz-local DYNAMODB_ENDPOINT=$(DDB_ENDPOINT) $(GO) run .

## seed-local: One-off question-bank sync into DynamoDB Local
.PHONY: seed-local
seed-local:
	DB_DRIVER=dynamodb DYNAMODB_TABLE=quiz-local DYNAMODB_ENDPOINT=$(DDB_ENDPOINT) $(GO) run . -seed

## lambda-build: Verify the Lambda build (CGO off, linux/arm64)
.PHONY: lambda-build
lambda-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO) build -o /dev/null .
	@echo "Lambda build OK"

## infra-install: Install CDK dependencies
.PHONY: infra-install
infra-install:
	cd infra && npm install

## synth: Synthesize the CloudFormation template (also compiles the Go Lambda)
.PHONY: synth
synth:
	cd infra && npx cdk synth --quiet

## diff: Diff the CDK stack against what is deployed
.PHONY: diff
diff:
	cd infra && npx cdk diff

## deploy: Build the frontend and deploy the whole stack to AWS
.PHONY: deploy
deploy: frontend-build lambda-build
	cd infra && npx cdk deploy

## seed-remote: Re-run the question-bank sync on the deployed Lambda
.PHONY: seed-remote
seed-remote:
	aws lambda invoke --region us-east-1 \
		--function-name $$(aws cloudformation describe-stacks --region us-east-1 --stack-name QuizAppStack \
			--query "Stacks[0].Outputs[?OutputKey=='FunctionName'].OutputValue" --output text) \
		--cli-binary-format raw-in-base64-out \
		--payload '{"quizapp_action":"seed"}' /dev/stdout

## destroy: Tear down the AWS stack (the DynamoDB table is retained)
.PHONY: destroy
destroy:
	cd infra && npx cdk destroy
