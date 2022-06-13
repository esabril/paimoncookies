DOCKER_COMPOSE_FILE=docker-compose.yml

DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=paimon
DB_PASS=paimon
DB_NAME=paimoncookies
DB_SSL=disable

GO_BIN=$(shell go env GOPATH)/bin

.PHONY: start
start: run-database migrations run-app
silent-start: run-database migrations run-app-silent

run-database:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d postgres && sleep 1

run-app:
	BOT_TOKEN=$(BOT_TOKEN) docker-compose -f $(DOCKER_COMPOSE_FILE) up

run-app-silent:
	BOT_TOKEN=$(BOT_TOKEN) docker-compose -f $(DOCKER_COMPOSE_FILE) up -d app

.PHONY: migrations
migrations:
	 go install github.com/pressly/goose/v3/cmd/goose@latest && \
 		cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) up

run-local-tests:
	go test ./... -v

run-app-tests:
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec -it app sh && go test ./... -v

.PHONY: stop
stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

.PHONY: clean
clean:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --rmi local -v