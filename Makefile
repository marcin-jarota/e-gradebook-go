include .env

start:
	go run ./cmd/api/. -dsn ${DSN} -port "8080"
build:
	go build -o ./tmp/main ./cmd/api/.
fe-dev:
	cd ./web && npm run dev