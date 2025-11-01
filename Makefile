# ⚠️ Auto-generated file. Review before use. It may contain suboptimal content.
.PHONY: help build up down logs clean
.PHONY: api-postgres api-sqlite cli-postgres cli-sqlite
.PHONY: postgres-up postgres-down sqlite-up sqlite-down
.PHONY: exec-cli-postgres exec-cli-sqlite

# Default target
help:
	@echo "Available commands:"
	@echo ""
	@echo "Build & General:"
	@echo "  make build              - Build Docker images"
	@echo "  make down               - Stop all services"
	@echo "  make clean              - Stop services and remove volumes"
	@echo "  make logs               - View logs from all services"
	@echo ""
	@echo "PostgreSQL API:"
	@echo "  make api-postgres       - Start PostgreSQL API (default)"
	@echo "  make postgres-up        - Start PostgreSQL API in background"
	@echo "  make postgres-down      - Stop PostgreSQL API"
	@echo ""
	@echo "SQLite API:"
	@echo "  make api-sqlite         - Start SQLite API"
	@echo "  make sqlite-up          - Start SQLite API in background"
	@echo "  make sqlite-down        - Stop SQLite API"
	@echo ""
	@echo "CLI (Interactive):"
	@echo "  make cli-postgres       - Start and exec into PostgreSQL CLI"
	@echo "  make cli-sqlite         - Start and exec into SQLite CLI"
	@echo "  make exec-cli-postgres  - Exec into running PostgreSQL CLI"
	@echo "  make exec-cli-sqlite    - Exec into running SQLite CLI"
	@echo ""
	@echo "Combined:"
	@echo "  make all-up             - Start both PostgreSQL and SQLite APIs"
	@echo "  make all-down           - Stop all APIs"

# Build images
build:
	docker-compose build

# PostgreSQL API
api-postgres:
	docker-compose up api

postgres-up:
	docker-compose up -d api

postgres-down:
	docker-compose stop api postgres

# SQLite API
api-sqlite:
	docker-compose --profile sqlite up api-sqlite

sqlite-up:
	docker-compose --profile sqlite up -d api-sqlite

sqlite-down:
	docker-compose --profile sqlite stop api-sqlite

# PostgreSQL CLI
cli-postgres:
	@echo "Starting PostgreSQL CLI..."
	@docker-compose --profile cli up -d cli
	@echo "Waiting for CLI container to be ready..."
	@sleep 2
	@echo "Entering CLI..."
	@docker exec -it contact-manager-cli-go /app/bin/cli

exec-cli-postgres:
	docker exec -it contact-manager-cli-go /app/bin/cli

# SQLite CLI
cli-sqlite:
	@echo "Starting SQLite CLI..."
	@docker-compose --profile sqlite-cli up -d cli-sqlite
	@echo "Waiting for CLI container to be ready..."
	@sleep 2
	@echo "Entering CLI..."
	@docker exec -it contact-manager-cli-sqlite-go /app/bin/cli

exec-cli-sqlite:
	docker exec -it contact-manager-cli-sqlite-go /app/bin/cli

# Combined operations
all-up: postgres-up sqlite-up
	@echo "Both PostgreSQL and SQLite APIs are running"
	@echo "PostgreSQL API: http://localhost:8080"
	@echo "SQLite API:     http://localhost:8081"

all-down:
	docker-compose --profile sqlite down

# Utility commands
up: postgres-up

down:
	docker-compose down

logs:
	docker-compose logs -f

clean:
	docker-compose --profile sqlite --profile cli --profile sqlite-cli down -v
	@echo "All services stopped and volumes removed"

# Status
status:
	@echo "Running containers:"
	@docker ps --filter "name=contact-manager" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
