postgres:
	docker run --name postgresdb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:14.3-alpine

createdb:
	docker exec -it postgresdb createdb --username=root --owner=root buss

dropdb:
	docker exec -it postgresdb dropdb buss

migrateup:
	 migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/buss?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/buss?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc