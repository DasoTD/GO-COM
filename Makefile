DB_URL = postgresql://root:secret@localhost:5432/Ecom?sslmode=disable
sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

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

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

mockdb:
	mockgen --build_flags=--mod=mod -package mockdb -destination db/mock/bank.go github.com/dasotd/Ecom/db/sqlc Bank

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

.PHONY: proto db_docs db_schema sqlc migrateup migratedown migrateup1 migratedown1 createdb dropdb test migration server mockdb installPG