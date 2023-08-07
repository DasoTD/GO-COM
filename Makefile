sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/Ecom?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/Ecom?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/Ecom?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/Ecom?sslmode=disable" -verbose down 1
installPG:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it  postgres12 createdb --username=root --owner=root Ecom

dropdb:
	sudo docker exec -it  postgres12 dropdb Ecom

# //to create the db/migration file
migration:
	migrate create -ext sql -dir  db/migration -seq {theName}

test:
	go test -v -cover ./...
server:
	go run main.go

mockdb:
	mockgen --build_flags=--mod=mod -package mockdb -destination db/mock/bank.go github.com/dasotd/Ecom/db/sqlc Bank

.PHONY: sqlc migrateup migratedown migrateup1 migratedown1 createdb dropdb test migration server mockdb installPG