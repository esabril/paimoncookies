DOCKER_COMPOSE_FILE=docker-compose.yml

DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=paimon
DB_PASS=paimon
DB_NAME=paimoncookies
DB_SSL=disable

GO_BIN=$(shell go env GOPATH)/bin

.PHONY: all
all: start

.PHONY: start
start: run-database all-migrations
silent-start: run-database all-migrations

.PHONY: run-database
run-database:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d postgres && sleep 1

.PHONY: run-database
run-app:
	PCOOKIES_BOT_TOKEN=$(BOT_TOKEN) docker-compose -f $(DOCKER_COMPOSE_FILE) up

.PHONY: run-database
run-app-silent:
	PCOOKIES_BOT_TOKEN=$(BOT_TOKEN) docker-compose -f $(DOCKER_COMPOSE_FILE) up -d app

.PHONY: all-migrations
all-migrations:
	 go install github.com/pressly/goose/v3/cmd/goose@latest && \
 		cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) up

# make migration-create name=new_table_name type=sql
.PHONY: migration-create
migration-create:
	cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) create $(name) $(type)

.PHONY: migration-up
migration-up:
	cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) up

.PHONY: migration-status
migration-status:
	cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) status

# make migration-down-to v=4
.PHONY: migration-down-to
migration-down-to:
	cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) down-to $(v)

.PHONY: mock-generate
mock-generate:
	go install github.com/golang/mock/mockgen@v1.6.0 && \
	$(GO_BIN)/mockgen -source=internal/service/world/repository/repository.go -destination=test/world/repository/repository.go -package=world_repo && \
	$(GO_BIN)/mockgen -source=internal/service/characters/repository/repository.go -destination=test/characters/repository/repository.go -package=characters_repo

.PHONY: run-local-tests
run-local-tests: mock-generate
	go test ./... -v

.PHONY: run-app-tests
run-app-tests: mock-generate
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec -it app sh && go test ./... -v

.PHONY: stop
stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

.PHONY: clean
clean:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --rmi local -v

# todo: docker build target (make dockerize)