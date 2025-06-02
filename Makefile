# Makefile for NOBI Mutual Fund Engine

include .env

MYSQL_CMD=mysql -h $(DB_HOST) -P $(DB_PORT) -u $(DB_USER) -p$(DB_PASS) $(DB_NAME)

.PHONY: migrate
migrate:
	$(MYSQL_CMD) < migrations/001_initial_schema.sql

.PHONY: run
run:
	go run ./cmd/server/main.go

.PHONY: proto
proto:
	protoc --go_out=pkg/pb/nobi --go-grpc_out=pkg/pb/nobi --proto_path=proto/nobi proto/nobi/nobi.proto

.PHONY: tidy
tidy:
	go mod tidy
