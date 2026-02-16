SHELL := /bin/bash
ENV_FILE ?= .env

# If a .env file exists, include it so variables like POSTGRES_USER are available to Make
-include $(ENV_FILE)

.PHONY: all up down logs migrate run build test start

all: start

up:
	@if [ ! -f $(ENV_FILE) ]; then cp .env.example $(ENV_FILE); echo "Created $(ENV_FILE) from example"; fi
	@echo "Starting Postgres container..."
	@docker compose up -d

down:
	@echo "Stopping containers and removing volumes..."
	@docker compose down -v

logs:
	@echo "Tailing compose logs... (Ctrl-C to stop)"
	@docker compose logs -f

migrate:
	@echo "Applying migrations to Postgres container (idempotent)"
	@docker compose exec -T postgres psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -f /docker-entrypoint-initdb.d/001_create_amenity_requests.sql || true

run:
	@echo "Running Go server (reads DATABASE_URL from $(ENV_FILE) if present)"
	@DATABASE_URL='${DATABASE_URL}' go run ./cmd/server

build:
	@echo "Building binary"
	@go build -o bin/server ./cmd/server

test:
	@echo "Running tests"
	@go test ./...

start: up
	@echo "Waiting for Postgres to be ready..."
	@docker compose exec -T postgres bash -c 'until pg_isready -U ${POSTGRES_USER} -d ${POSTGRES_DB} >/dev/null 2>&1; do sleep 1; done'
	@$(MAKE) migrate
	@$(MAKE) run
