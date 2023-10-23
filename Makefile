# app commands

up:
	docker-compose -f ./deploy/docker-compose.yml up --build -d

down:
	docker-compose -f ./deploy/docker-compose.yml down

lint:
	golangci-lint run

go-mod:
	go mod download && go mod tidy

test:
	go test ./...

pre-push: go-mod test lint

# tools aliases

gen-all: swag-v1 sqlboiler-generate

install-tools: goose-install sqlboiler-install swag-install

# goose commands

goose-install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

DB="postgresql://user:password@localhost:5432/db?sslmode=disable"

goose-up:
	goose --dir ./migrations postgres $(DB) up

goose-down:
	goose --dir ./migrations postgres $(DB) down

# sqlboiler commands

sqlboiler-install:
	go install github.com/volatiletech/sqlboiler/v4@latest && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

sqlboiler-generate:
	sqlboiler psql --config ./tools/sqlboiler/cfg.toml

# swag commands

swag-install:
	go install github.com/swaggo/swag/cmd/swag@latest

swag-v1:
	swag init -g ./internal/controller/server.go -o ./api