postgres:
	docker run --name pg12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it pg12 createdb --username=root --owner=root simple_bank

drobdb:
	docker exec -it pg12 dropdb simple_bank

migrateup:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgreq createdb drobdb migrateup migratedown sqlc