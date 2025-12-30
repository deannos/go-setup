.PHONY: help build run test clean docker-build docker-run migrate

help:
	@echo "Available commands:"
	@echo "  make build       - Build the application"
	@echo "  make run         - Run the application"
	@echo "  make test        - Run tests"
	@echo "  make test-coverage - Run tests with coverage"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make migrate     - Run database migrations"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run  - Run Docker container"

build:
	go build -o bin/go-setup cmd/main.go

run:
	go run cmd/main.go

test:
	GIN_MODE=test go test ./tests/unit/... -v

test-coverage:
	GIN_MODE=test go test ./tests/unit/... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out

clean:
	rm -f bin/go-setup
	rm -f coverage.out

migrate:
	psql -h localhost -U devuser -d go_setupdb -f migrations/001_create_users_table.sql

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

docker-build:
	docker build -t go-setup:latest .

docker-run:
	docker run -p 8080:8080 --network go-setup_default go-setup:latest