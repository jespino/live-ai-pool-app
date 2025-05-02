.PHONY: build run test clean postgres-up postgres-down frontend-deps frontend-dev frontend-build docker-up docker-down dev dev-postgres hot-reload help

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

# Run with PostgreSQL database
run-postgres: postgres-up
	DB_TYPE=postgres go run ./cmd/server

# Run with hot reload
hot-reload:
	air

test:
	go test ./...

# Run PostgreSQL tests only
test-postgres:
	POSTGRES_TEST_DSN="host=localhost port=5432 user=postgres password=postgres dbname=poolapp_test sslmode=disable" go test -v ./internal/database -run TestPostgresDB

clean:
	rm -rf bin/
	rm -rf frontend/dist

deps:
	go mod tidy
	go mod download

# Frontend commands
frontend-deps:
	cd frontend && npm install --legacy-peer-deps

frontend-dev: frontend-deps
	cd frontend && npm run dev

frontend-build: frontend-deps
	cd frontend && npm run build

# Start PostgreSQL container for development
postgres-up:
	docker run --name pool-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=poolapp -p 5432:5432 -d postgres:15
	docker exec pool-postgres psql -U postgres -c "CREATE DATABASE poolapp_test;" || true

# Stop and remove PostgreSQL container
postgres-down:
	docker stop pool-postgres || true
	docker rm pool-postgres || true

# Start the application with Docker Compose
docker-up:
	docker-compose up --build -d

# Start with PostgreSQL enabled
docker-up-postgres:
	DB_TYPE=postgres docker-compose up --build -d

# Stop Docker Compose
docker-down:
	docker-compose down

# Run everything locally (backend and frontend)
run-all: postgres-up
	$(MAKE) -j2 run frontend-dev

# Full clean and rebuild
reset: clean docker-down postgres-down
	$(MAKE) docker-up-postgres

# Development targets
dev: deps frontend-deps
	$(MAKE) -j2 run frontend-dev

dev-postgres: deps frontend-deps postgres-up
	$(MAKE) -j2 run-postgres frontend-dev

# Help target
help:
	@echo "Pool App Makefile"
	@echo "================="
	@echo "Main targets:"
	@echo "  make build           - Build the server binary"
	@echo "  make run             - Run the server (in-memory database)"
	@echo "  make run-postgres    - Run the server with PostgreSQL"
	@echo "  make hot-reload      - Run the server with hot reload (using air)"
	@echo "  make test            - Run all tests"
	@echo "  make test-postgres   - Run PostgreSQL-specific tests"
	@echo ""
	@echo "Development targets:"
	@echo "  make dev             - Run backend and frontend for development"
	@echo "  make dev-postgres    - Run backend with PostgreSQL and frontend"
	@echo "  make frontend-dev    - Run the frontend dev server"
	@echo "  make frontend-build  - Build the frontend for production"
	@echo ""
	@echo "Container targets:"
	@echo "  make docker-up       - Start with Docker Compose (in-memory)"
	@echo "  make docker-up-postgres - Start with Docker Compose (PostgreSQL)"
	@echo "  make docker-down     - Stop Docker Compose containers"
	@echo ""
	@echo "Utility targets:"
	@echo "  make clean           - Clean build artifacts"
	@echo "  make postgres-up     - Start PostgreSQL container"
	@echo "  make postgres-down   - Stop PostgreSQL container"
	@echo "  make reset           - Clean and rebuild everything"
	@echo ""
	@echo "Use the devcontainer for the best development experience."