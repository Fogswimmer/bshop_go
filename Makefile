DB_USER=postgres
DB_NAME=books
DB_CONTAINER_NAME=db
GO_CONTAINER_NAME=goapi
MIGRATIONS_PATH=infra/db/create-tables.sql
UPLOADS_PATH=/public/uploads

.PHONY: init build migrate seed

init: build migrate

init-seed: build migrate seed

build:
	docker compose up --build -d
migrate:
	docker exec -i $(DB_CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) < $(MIGRATIONS_PATH)
seed:
	docker exec -it $(GO_CONTAINER_NAME) ./docker-gs-ping seed
