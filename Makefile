postgres:
	docker run --name postgresbank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=banktest -d postgres

createdb:
	docker exec -it postgresbank createdb --username=root --owner=root bankdb

dropdb:
	docker exec -it postgresbank dropdb bankdb

migrateup:
	migrate -path database/migrations -database "postgresql://root:1234@localhost:5432/bankdb?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://root:1234@localhost:5432/bankdb?sslmode=disable" -verbose down
