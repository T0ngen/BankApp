postgres:
	docker run --name postgresbank -p 5433:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=banktest -d postgres

createdb:
	docker exec -it postgresbank createdb --username=root --owner=root bankdbtest2

dropdb:
	docker exec -it postgresbank dropdb bankdbtest2

migrateup:
	migrate -path database/migrations -database "postgresql://root:1234@localhost:5433/bankdbtest2?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://root:1234@localhost:5433/bankdbtest2?sslmode=disable" -verbose down
