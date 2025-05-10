launchpostgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.5-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root readly
	docker exec -it postgres psql -U root -d readly -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

dropdb:
	docker exec -it postgres dropdb readly

# dependency migrate CLI(https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/readly?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/readly?sslmode=disable" -verbose down

# dependency sqlc CLI(https://github.com/sqlc-dev/sqlc#installation)
sqlc:
	sqlc generate

test:
	go test -v -cover ./... -tags test

launchserver:
	go run cmd/main.go

# dependencies
# Buf CLI(https://buf.build/docs/)
proto:
	buf generate

.PHONY: launchpostgres createdb dropdb migrateup migratedown sqlc test launchserver proto