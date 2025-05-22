DB_USER=postgres
DB_NAME=books
DB_CONTAINER_NAME=db
MIGRATIONS_PATH=db/create-tables.sql
UPLOADS_PATH=/public/uploads

.PHONY: init

init:
	docker exec -i $(DB_CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) < $(MIGRATIONS_PATH)

