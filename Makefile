.PHONY: build run up down tidy swag

build:
	go build -o bin/agroapp ./cmd/agroapp

run:
	go run ./cmd/agroapp

up:
	docker compose up -d

down:
	docker compose down

tidy:
	go mod tidy

swag:
	swag init -g internal/interfaces/httpapi/router.go
