start-app:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	# Install godotenv with 'go install github.com/joho/godotenv/cmd/godotenv@latest'
	reflex -s -r '\.go$$' -- godotenv -f .env go run archmark.go

start-view:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -r '\.qtpl$$' -- qtc -dir=view

start-worker:
	reflex -s -r '\.go$$' -- godotenv -f .env go run worker/cmd/archmark_worker.go

start-faktory:
	docker run --rm -it -p 127.0.0.1:7419:7419 -p 127.0.0.1:7420:7420 contribsys/faktory:latest


db-migrate:
	migrate -path ./migrations -database "postgres://localhost:5432/archmark?sslmode=disable" up

db-schema-dump:
	pg_dump --schema-only -O archmark > database/schema.sql

sqlc-gen:
	sqlc --experimental generate

.PHONY: start-app start-view db-migrate db-schema-dump sqlc-gen
