# ==============================================================================
# Makefile — golangfromzero
# ==============================================================================
# Infrastructure, build and deploy management
# ==============================================================================

.PHONY: help up down down/volumes restart logs logs/app logs/db ps shell \
        build/dev build/dev/clean build/prod \
        migrate/up migrate/down migrate/down/all migrate/create migrate/force migrate/version \
        test test/local test/coverage test/coverage/html \
        deploy/build deploy/push deploy \
        tidy fmt vet lint clean

# ==============================================================================
# Variables
# ==============================================================================

APP_NAME   := golangfromzero
IMAGE_NAME := fabianoflorentino/golangfromzero
IMAGE_TAG  := v0.0.1
MIGRATIONS := database/migrations

# ==============================================================================
# Colors
# ==============================================================================

RED    := \033[0;31m
GREEN  := \033[0;32m
YELLOW := \033[0;33m
BLUE   := \033[0;34m
NC     := \033[0m

# ==============================================================================
# Default target
# ==============================================================================

.DEFAULT_GOAL := help

##@ Help

help: ## Show this help message
	@echo ""
	@echo -e "$(BLUE)╔══════════════════════════════════════════════════════════╗$(NC)"
	@echo -e "$(BLUE)║        golangfromzero - Available Commands               ║$(NC)"
	@echo -e "$(BLUE)╚══════════════════════════════════════════════════════════╝$(NC)"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"; printf ""} /^[a-zA-Z_/-]+:.*?##/ { printf "  $(GREEN)%-22s$(NC) %s\n", $$1, $$2 } /^##@/ { printf "\n$(YELLOW)%s$(NC)\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
	@echo ""

##@ Local Development

up: ## Start all containers in development mode (with hot-reload)
	@echo -e "$(BLUE)🚀 Starting development environment...$(NC)"
	@docker compose up --build -d
	@echo -e "$(GREEN)✓ Containers started!$(NC)"

down: ## Stop and remove containers
	@echo -e "$(BLUE)🛑 Stopping containers...$(NC)"
	@docker compose down
	@echo -e "$(GREEN)✓ Containers stopped!$(NC)"

down/volumes: ## Stop and remove containers and persistent volumes (deletes database data)
	@echo -e "$(YELLOW)⚠️  Stopping containers and removing volumes...$(NC)"
	@docker compose down -v
	@echo -e "$(GREEN)✓ Done!$(NC)"

restart: down up ## Restart all containers

logs: ## Stream logs from all containers
	@docker compose logs -f

logs/app: ## Stream logs from the application container only
	@docker compose logs -f app

logs/db: ## Stream logs from the database container only
	@docker compose logs -f db

ps: ## List container status
	@echo -e "$(BLUE)📊 Container status:$(NC)"
	@docker compose ps

shell: ## Open an interactive shell inside the application container
	@echo -e "$(BLUE)🐚 Opening shell in container...$(NC)"
	@docker compose exec app sh

##@ Docker Image Builds

build/dev: ## Build the Docker image in development mode (uses cache)
	@echo -e "$(BLUE)🔨 Building development image...$(NC)"
	@docker compose build
	@echo -e "$(GREEN)✓ Build complete!$(NC)"

build/dev/clean: ## Build the Docker image in development mode without cache
	@echo -e "$(BLUE)🔨 Building development image (no cache)...$(NC)"
	@docker compose build --no-cache
	@echo -e "$(GREEN)✓ Build complete!$(NC)"

build/prod: ## Build the Docker image in production mode (distroless)
	@echo -e "$(BLUE)🔨 Building production image...$(NC)"
	@docker build \
		--target production \
		--tag $(IMAGE_NAME):$(IMAGE_TAG) \
		--tag $(IMAGE_NAME):latest \
		.
	@echo -e "$(GREEN)✓ Production build complete!$(NC)"

##@ Database Migrations

migrate/up: ## Apply all pending migrations
	@echo -e "$(BLUE)⬆️  Applying migrations...$(NC)"
	@docker compose exec app sh -c \
		'migrate -path $(MIGRATIONS) -database "$$DATABASE_URL" up'
	@echo -e "$(GREEN)✓ Migrations applied!$(NC)"

migrate/down: ## Revert the last applied migration
	@echo -e "$(YELLOW)⬇️  Reverting last migration...$(NC)"
	@docker compose exec app sh -c \
		'migrate -path $(MIGRATIONS) -database "$$DATABASE_URL" down 1'
	@echo -e "$(GREEN)✓ Migration reverted!$(NC)"

migrate/down/all: ## Revert ALL migrations (drops the entire schema)
	@echo -e "$(RED)⚠️  Reverting ALL migrations...$(NC)"
	@docker compose exec app sh -c \
		'echo y | migrate -path $(MIGRATIONS) -database "$$DATABASE_URL" down'
	@echo -e "$(GREEN)✓ All migrations reverted!$(NC)"

migrate/create: ## Create a new migration file pair. Usage: make migrate/create NAME=migration_name
	@test -n "$(NAME)" || (echo -e "$(RED)❌ Provide a name: make migrate/create NAME=<name>$(NC)"; exit 1)
	@echo -e "$(BLUE)📝 Creating migration: $(NAME)...$(NC)"
	@docker compose exec app migrate create \
		-ext sql \
		-dir $(MIGRATIONS) \
		-seq $(NAME)
	@echo -e "$(GREEN)✓ Migration files created!$(NC)"

migrate/force: ## Force a migration version (fix dirty state). Usage: make migrate/force VERSION=1
	@test -n "$(VERSION)" || (echo -e "$(RED)❌ Provide a version: make migrate/force VERSION=<number>$(NC)"; exit 1)
	@echo -e "$(YELLOW)⚠️  Forcing migration version $(VERSION)...$(NC)"
	@docker compose exec app sh -c \
		'migrate -path $(MIGRATIONS) -database "$$DATABASE_URL" force $(VERSION)'
	@echo -e "$(GREEN)✓ Version forced!$(NC)"

migrate/version: ## Show the current database schema version
	@echo -e "$(BLUE)🔍 Current migration version:$(NC)"
	@docker compose exec app sh -c \
		'migrate -path $(MIGRATIONS) -database "$$DATABASE_URL" version'

##@ Tests

test: ## Run all tests inside the application container
	@echo -e "$(BLUE)🧪 Running tests...$(NC)"
	@docker compose exec app go test ./... -v

test/local: ## Run all tests locally (requires Go installed)
	@echo -e "$(BLUE)🧪 Running tests locally...$(NC)"
	@go test ./... -v

test/coverage: ## Run tests with coverage report
	@echo -e "$(BLUE)🧪 Running tests with coverage...$(NC)"
	@docker compose exec app sh -c \
		'go test ./... -coverprofile=tmp/coverage.out && go tool cover -func=tmp/coverage.out'

test/coverage/html: ## Run tests and open HTML coverage report in the browser
	@echo -e "$(BLUE)🧪 Running tests with HTML coverage report...$(NC)"
	@docker compose exec app sh -c \
		'go test ./... -coverprofile=tmp/coverage.out && go tool cover -html=tmp/coverage.out -o tmp/coverage.html'
	@echo -e "$(GREEN)✓ Coverage report generated: tmp/coverage.html$(NC)"
	@echo -e "$(YELLOW)🌐 Open in browser (WSL path):$(NC)"
	@echo -e "$(GREEN)  \\\\\\\\wsl.localhost\\FedoraLinux-43$(PWD)/tmp/coverage.html$(NC)"

##@ Production Deploy

deploy/build: ## Build the production image
	$(MAKE) build/prod

deploy/push: ## Push the image to the registry (Docker Hub)
	@echo -e "$(BLUE)📤 Pushing image to registry...$(NC)"
	@docker push $(IMAGE_NAME):$(IMAGE_TAG)
	@docker push $(IMAGE_NAME):latest
	@echo -e "$(GREEN)✓ Image pushed!$(NC)"

deploy: ## Build + push the production image
	$(MAKE) deploy/build
	$(MAKE) deploy/push

##@ Utilities

tidy: ## Update Go dependencies (go mod tidy)
	@echo -e "$(BLUE)📦 Updating dependencies...$(NC)"
	@go mod tidy
	@echo -e "$(GREEN)✓ Dependencies updated!$(NC)"

fmt: ## Format Go source code
	@echo -e "$(BLUE)🎨 Formatting code...$(NC)"
	@go fmt ./...
	@echo -e "$(GREEN)✓ Code formatted!$(NC)"

vet: ## Run Go static analysis
	@echo -e "$(BLUE)🔍 Running static analysis...$(NC)"
	@go vet ./...
	@echo -e "$(GREEN)✓ Static analysis complete!$(NC)"

lint: ## Run linter inside the application container (requires golangci-lint)
	@echo -e "$(BLUE)🔍 Running linter...$(NC)"
	@docker compose exec app golangci-lint run ./...

clean: ## Remove local binaries and temporary artifacts
	@echo -e "$(YELLOW)🧹 Cleaning temporary artifacts...$(NC)"
	@rm -f tmp/*.out tmp/*.log tmp/*.html
	@echo -e "$(GREEN)✓ Temporary artifacts removed.$(NC)"
