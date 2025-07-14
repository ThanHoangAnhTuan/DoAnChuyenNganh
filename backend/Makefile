APP_NAME := server

GOOSE_DBSTRING := "root:root@tcp(localhost:3307)/ecommerce_go"
GOOSE_MIGRATION_DIR := sql/schemas
GOOSE_DRIVER := mysql

dev:
	go run ./cmd/${APP_NAME}/

docker-up:
	docker-compose -p ecommerce_go up -d

docker-up-with-observable:
	docker-compose -p ecommerce_go -f ./environment/docker-compose-dev.yaml up -d

docker-down:
	docker-compose -p ecommerce_go down

docker-build:
	docker-compose -p ecommerce_go up --build -d

wire:
	wire ./internal/wire/

# goose migration
create-migration:
	goose -dir $(GOOSE_MIGRATION_DIR) create $(name) sql

db-up:
	goose -dir $(GOOSE_MIGRATION_DIR) -s $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

db-down:
	goose -dir $(GOOSE_MIGRATION_DIR) -s $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

db-reset:
	goose -dir $(GOOSE_MIGRATION_DIR) -s $(GOOSE_DRIVER) $(GOOSE_DBSTRING) reset

sql-gen:
	sqlc generate

swagger:
	swag init -g cmd/server/main.go -o docs

.PHONY: dev docker-up docker-down docker-build wire db-up db-down db-reset

.PHONY: air