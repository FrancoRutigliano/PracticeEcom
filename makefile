include .env

up:
	@echo "Starting containers"
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stoping containers"
	docker-compose down

migration:
	@migrate create -ext sql -dir internal/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run internal/migrate/main.go up

migrate-down:
	@go run internal/migrate/main.go down

build:
	@go build -o bin/ecom cmd/main.go

start:
	./bin/ecom

restart: build start