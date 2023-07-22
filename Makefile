GO := go
DB_FILE := test.db
SQL_SCRIPT := db/init.sql

run:
	$(GO) run ./cmd/server.go
