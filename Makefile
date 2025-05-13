DB_USER=postgres
DB_NAME=books
DB_CONTAINER_NAME=db
MIGRATIONS_PATH=db/migrations/create-tables.sql

.PHONY: prepare run run_air init_db init_prod init_dev prod dev

init_prod: prepare init_db prod
init_dev: prepare init_db dev

prepare:
	mkdir -p $(UPLOADS_PATH)
	chmod -R 755 $(UPLOADS_PATH)

init_db:
	docker exec -i $(DB_CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) < $(MIGRATIONS_PATH)

prod:
	go run .

dev:
	air
