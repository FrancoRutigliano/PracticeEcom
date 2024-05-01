include .env

build:
	@go build -o bin/ecom cmd/main.go

up:
	@echo "Starting containers"
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stoping containers"
	docker-compose down

start:
	./bin/ecom

restart: build start