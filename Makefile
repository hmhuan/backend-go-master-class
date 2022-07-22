postgres:
	docker run --name pg12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it pg12 createdb --username=root --owner=root simple_bank

drobdb:
	docker exec -it pg12 dropdb simple_bank

migrateup:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/hmhuan/simple-bank/db/sqlc Store

.PHONY: postgreq createdb drobdb migrateup migratedown sqlc test server mock