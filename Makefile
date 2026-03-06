.DEFAULT_GOAL := help

.PHONY: help docker-up

## help: Display available commands
help:
	@echo "Available commands:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' |  sed -e 's/^/ /'

## start: Start docker-compose containers (PostgreSQL) and run the application
start:
	docker compose up -d
	go run ./cmd/api/main.go
