.PHONY: build run test up down

build:
	go build -o bin/agroapp ./cmd/agroapp

run:
	go run ./cmd/agroapp

up:
	docker compose up -d

down:
	docker compose down
