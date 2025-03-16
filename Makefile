GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= "admin:admin@tcp(127.0.0.1:3306)/ecomerce"
GOOSE_MIGRATION_DIR ?= sql/schema

# app name
APP_NAME = server

# docker
docker_build:
	docker-compose up -d --build
	docker-compose ps
docker_up:
	docker compose up -d
docker_stop:
	docker-compose down

up_by_one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one

dev:
	go run ./cmd/$(APP_NAME)

create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlcgen:
	sqlc generate

.PHONY: upse downse resetse docker_build docker_stop docker_up
.PHONY: air
