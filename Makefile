.PHONY: up down build

up:
	docker-compose up -d

down:
	docker-compose down

build:
	go build -o platformctl