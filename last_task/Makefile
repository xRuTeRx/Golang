postgres:
	docker run --name=adress_book_db  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root  postgres
createdb: 
	docker exec -it  adress_book_db createdb --username=root --owner=root requests
dropdb: 
	docker exec -it  adress_book_db dropdb requests
migrateup:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/requests?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/requests?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown