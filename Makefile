start-app:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	# Install godotenv with 'go install github.com/joho/godotenv/cmd/godotenv@latest'
	reflex -s -r '\.go$$' -- godotenv -f .env go run archmark.go

start-view:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -r '\.qtpl$$' -- qtc -dir=view

db-migrate:
	migrate -path ./migrations -database "postgres://localhost:5432/archmark?sslmode=disable" up

db-schema-dump:
	pg_dump --schema-only -O archmark > database/schema.sql

sqlc-gen:
	sqlc --experimental generate

.PHONY: start-app start-view db-migrate db-schema-dump sqlc-gen
