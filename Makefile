launchpostgres:
	docker run --name postgres17.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.2-alpine

createdb:
	docker exec -it postgres17.2 createdb --username=root --owner=root readly
	docker exec -it postgres17.2 psql -U root -d readly -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

dropdb:
	docker exec -it postgres17.2 dropdb readly

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/readly?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5432/readly?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./... -tags test

launchserver:
	go run cmd/main.go

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        proto/*.proto

.PHONY: launchpostgres createdb dropdb migrateup migratedown sqlc test launchserver proto