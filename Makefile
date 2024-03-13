all: archmark

archmark:
	CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o archmark archmark.go

archmark-worker:
	CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o archmark-worker worker/main.go

start-app:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	# Install godotenv with 'go install github.com/joho/godotenv/cmd/godotenv@latest'
	reflex -s -r '\.go$$' -- godotenv -f .env go run archmark.go

start-worker:
	reflex -s -r '\.go$$' -- godotenv -f .env go run worker/main.go

start-view:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -r '\.qtpl$$' -- qtc -dir=internal/view

db-migrate:
	migrate -path migrations -database "postgres://127.0.0.1/archmark?sslmode=disable" up

db-schema-dump:
	pg_dump --schema-only -O archmark > internal/database/schema.sql

sqlc-gen:
	sqlc --experimental generate

.PHONY: archmark archmark-worker start-app start-worker start-view db-migrate db-schema-dump sqlc-gen
