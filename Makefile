init:
	go run github.com/99designs/gqlgen init

generate:
	go run github.com/99designs/gqlgen

run:
	go run server.go

psql:
	psql -h localhost -p 5432 -U postgres -W

migrateup:
	migrate -path app/database/migration/ -database "postgresql://postgres:root@localhost:5432/reports?sslmode=disable" -verbose up

migratedown:
	migrate -path app/database/migration/ -database "postgresql://postgres:root@localhost:5432/reports?sslmode=disable" -verbose down