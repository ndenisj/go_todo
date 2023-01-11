postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root todo

dropdb:
	docker exec -it postgres15 dropdb --username=root --owner=root todo

create_migration:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/todo?sslmode=disable" -verbose up

migrateuplast:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/todo?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/todo?sslmode=disable" -verbose down

migratedownlast:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/todo?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/ndenisj/go_todo/db/sqlc Store


.PHONY: postgres createdb dropdb create_migration migrateup migratedown sqlc test server mock migratedownlast migrateuplast