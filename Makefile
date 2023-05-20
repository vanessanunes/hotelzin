include .env
.PHONY: up

local:
	docker-compose up -d

first:
	docker-compose down; docker-compose build; docker-compose up 

app:
	go run main.go

