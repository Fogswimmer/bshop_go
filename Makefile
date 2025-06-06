DB_USER=postgres
DB_NAME=books
DB_CONTAINER_NAME=postgres
GO_CONTAINER_NAME=goapi
MIGRATIONS_PATH=infra/db/create-tables.sql
UPLOADS_PATH=/public/uploads
USER_ID=$(shell id -u)
GROUP_ID=$(shell id -g)


.PHONY: init-prod init-prod-seed init-dev init-dev-seed build migrate seed fix-perms

init-prod: build migrate fix-perms run-dev

init-prod-seed: init-prod seed

init-dev: build run-dev

init-dev-seed: init-dev seed

run-dev:
	air
build:
	docker compose -f compose.prod.yaml up --build -d
migrate:
	docker exec -i $(DB_CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) < $(MIGRATIONS_PATH)
seed:
	docker exec -it $(GO_CONTAINER_NAME) ./docker-gs-ping seed
fix-perms:
	docker exec -it $(GO_CONTAINER_NAME) sh -c "chown -R 0:0 $(UPLOADS_PATH)"
	docker exec -it $(GO_CONTAINER_NAME) sh -c "chmod -R ug+rwX $(UPLOADS_PATH)"

