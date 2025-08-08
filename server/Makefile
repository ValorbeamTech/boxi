.PHONY: help build run test clean docker-build docker-run dev dev-down migrate-up migrate-down deps fmt lint security

# Variables
BINARY_NAME=main
DOCKER_IMAGE=myapp
DOCKER_TAG=latest
DATABASE_URL ?= postgresql://postgres:password@localhost:5432/myapp_dev?sslmode=disable

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

deps: ## Install dependencies
	go mod download
	go mod tidy

build: ## Build the application
	go build -o bin/$(BINARY_NAME) cmd/api/main.go

run: ## Run the application locally
	go run cmd/api/main.go

dev: ## Start development environment with Docker Compose
	docker-compose -f docker-compose.dev.yml up --build

dev-down: ## Stop development environment
	docker-compose -f docker-compose.dev.yml down

dev-logs: ## Show development logs
	docker-compose -f docker-compose.dev.yml logs -f

##@ Testing

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

benchmark: ## Run benchmarks
	go test -bench=. -benchmem ./...

##@ Database

migrate-up: ## Run database migrations
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down: ## Rollback database migrations
	migrate -down -path migrations -database "$(DATABASE_URL)" down

migrate-force: ## Force migration version
	@read -p "Enter migration version: " version; \
	migrate -path migrations -database "$(DATABASE_URL)" force $$version

migrate-version: ## Show current migration version
	migrate -path migrations -database "$(DATABASE_URL)" version

create-migration: ## Create a new migration file
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations $$name

##@ Docker

docker-build: ## Build Docker image
	docker build -f docker/Dockerfile -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-build-dev: ## Build development Docker image
	docker build -f docker/Dockerfile.dev -t $(DOCKER_IMAGE):dev .

docker-run: ## Run with Docker Compose (production)
	docker-compose up -d

docker-down: ## Stop Docker Compose
	docker-compose down

docker-logs: ## Show Docker logs
	docker-compose logs -f

docker-clean: ## Clean Docker images and containers
	docker-compose down -v --rmi all --remove-orphans

##@ Code Quality

fmt: ## Format Go code
	go fmt ./...
	goimports -w .

lint: ## Run linter
	golangci-lint run

security: ## Run security checks
	gosec ./...

##@ Database Operations

db-reset: ## Reset database (drop and recreate)
	docker-compose -f docker-compose.dev.yml exec db psql -U postgres -c "DROP DATABASE IF EXISTS myapp_dev;"
	docker-compose -f docker-compose.dev.yml exec db psql -U postgres -c "CREATE DATABASE myapp_dev;"
	$(MAKE) migrate-up

db-seed: ## Seed database with sample data
	go run scripts/seed.go

db-backup: ## Backup database
	docker-compose -f docker-compose.dev.yml exec db pg_dump -U postgres myapp_dev > backup.sql

db-restore: ## Restore database from backup
	docker-compose -f docker-compose.dev.yml exec -T db psql -U postgres myapp_dev < backup.sql

##@ Deployment

deploy-staging: ## Deploy to staging
	@echo "Deploying to staging..."
	docker build -f docker/Dockerfile -t $(DOCKER_IMAGE):staging .
	# Add your staging deployment commands here

deploy-prod: ## Deploy to production
	@echo "Deploying to production..."
	docker build -f docker/Dockerfile -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	# Add your production deployment commands here

##@ Utilities

clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf tmp/
	rm -f coverage.out coverage.html

install-tools: ## Install development tools
	go install github.com/cosmtrek/air@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	go install golang.org/x/tools/cmd/goimports@latest

check-env: ## Check environment variables
	@echo "Checking environment variables..."
	@echo "DATABASE_URL: $(DATABASE_URL)"
	@echo "Environment file exists: $$(test -f .env && echo 'Yes' || echo 'No')"

##@ Information

version: ## Show version info
	@echo "Go version: $$(go version)"
	@echo "Git commit: $$(git rev-parse --short HEAD 2>/dev/null || echo 'unknown')"
	@echo "Built: $$(date)"

ports: ## Show used ports
	@echo "Application ports:"
	@echo "  API: 8080"
	@echo "  PostgreSQL: 5432"
	@echo "  Redis: 6379"

status: ## Show service status
	@echo "Docker services status:"
	@docker-compose -f docker-compose.dev.yml ps

##@ Quick Commands

quick-start: deps dev ## Quick start for new developers

quick-test: fmt lint test ## Quick test suite