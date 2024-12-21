launchpostgres:
	docker run --name postgres17.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.2-alpine

createdb:
	docker exec -it postgres17.2 createdb --username=root --owner=root readly

dropdb:
	docker exec -it postgres17.2 dropdb readly

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/readly?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5432/readly?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: launchpostgres createdb dropdb migrateup migratedown sqlc