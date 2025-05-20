include .env

LOCAL_MIGRATION_DIR=.$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=${DB_PORT} dbname=${DB_DB} user=${DB_USER} password=${DB_PASSWORD} sslmode=disable"

swagger:
	swag init -g cmd/main.go -o pkg/swagger --outputTypes json
	mv pkg/swagger/swagger.json pkg/swagger/api.swagger.json
	go tool statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

install-deps:
	go get -tool  github.com/pressly/goose/v3/cmd/goose@v3.20.0
	go get -tool  github.com/rakyll/statik@v0.1.7

install-golangci-lint:
	go get -tool github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.6

lint:
	go mod tidy
	go tool golangci-lint run ./... --config .golangci.pipeline.yaml

docker-build:
	docker compose up -d --build

rebuild:
	docker compose up -d --build app
	
docker-run:
	docker compose up -d

run:
	go run cmd/main.go

local-migration-status:
	go tool goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	go tool goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	go tool goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

local-migration-reset:
	go tool goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} reset -v

migrate_down:
	go tool goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down 20250322151928_create_payments_table.sql

migrate_up:
	go tool goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up 20250322151928_create_payments_table.sql

#  go tool goose -dir ./migrations create NAME sql
