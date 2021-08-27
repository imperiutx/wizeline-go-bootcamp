db:
	docker run --name simple_service -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13.1-alpine

createdb:
	docker exec -it simple_service createdb --username=root --owner=root simple_service

mic:
	migrate create -ext sql -dir db/migration -seq init_schema

dropdb:
	docker exec -it simple_service dropdb simple_service

miup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_service?sslmode=disable" -verbose up

midown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_service?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: db createdb dropdb miup midown sqlc test server
