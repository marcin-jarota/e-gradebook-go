include .env

start:
	go run ./cmd/web/. -dsn ${DSN} -port "8080"
build:
	go build -o ./tmp/main ./cmd/web/.
fe-dev:
	cd ./web && npm run dev