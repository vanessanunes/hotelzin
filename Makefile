include .env
.PHONY: up

hello:
	echo "Hello"

docker-up:
	docker-compose up -d

docker-local:
	docker-compose down; docker-compose build; docker-compose up -d

docker-up-force:
	docker-compose up --force-recreate

app:
	go run main.go

