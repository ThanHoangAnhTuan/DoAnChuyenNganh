APP_NAME := server

GOOSE_DBSTRING := "root:root@tcp(localhost:3306)/ecommerce_go"
GOOSE_MIGRATION_DIR := sql/schemas
GOOSE_DRIVER := mysql

dev:
	go run ./cmd/${APP_NAME}/

docker_up:
	docker-compose -p ecommerce_go up -d

docker_down:
	docker-compose -p ecommerce_go down

docker_build:
	docker-compose -p ecommerce_go up --build -d

wire:
	wire ./internal/wire/

# goose migration
create_migration:
	@goose -dir $(GOOSE_MIGRATION_DIR) create $(name) sql

db_up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) up

db_down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) down

db_reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) reset

sql_gen:
	sqlc generate

swagger:
	swag init -g cmd/server/main.go -o docs

.PHONY: dev docker_up docker_down docker_build wire db_up db_down db_reset
  
.PHONY: air