# SHELL := /bin/bash

# .PHONY: create_migration apply_migration sqlc_generate

migrate-create:
	./scripts/create_migrate.sh

migrate-up:
	docker run -v $(shell pwd)/scripts/migrations/:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://root:postgres@localhost:5432/whatsapp?sslmode=disable -verbose up

migrate-down:
	docker run -v $(shell pwd)/scripts/migrations/:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://root:postgres@localhost:5432/whatsapp?sslmode=disable -verbose down

sql-generate:
	sqlc generate

start:
	go run ./cmd/api/main.go

# go build -tags netgo -ldflags '-s -w' -o app
build:
	go build -o build/ci/api cmd/api/main.go
