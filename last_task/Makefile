postgres:
	docker run --name=last_task_database_1  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root  postgres
createdb: 
	docker exec -it  last_task_database_1 createdb --username=root --owner=root requests
	

dropdb: 
	docker exec -it  last_task_database_1 dropdb requests
migrateup:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/requests?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/requests?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown