DOCKER_COMPOSE_FILE=docker-compose.yml

DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5433
DB_USER=paimon
DB_PASS=paimon
DB_NAME=paimoncookies
DB_SSL=disable

GO_BIN=$(shell go env GOPATH)/bin
GIT_TAG=$(shell git describe --tags)

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

.PHONY: run-app-silent
run-app-silent:
	PCOOKIES_BOT_TOKEN=$(BOT_TOKEN) docker-compose -f $(DOCKER_COMPOSE_FILE) up -d app

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

.PHONY: run-test-coverage
run-test-coverage:
	go test ./... -coverprofile test/cover.out && go tool cover -html=test/cover.out

# Pretty tests result
.PHONY: gts
gts:
	go install gotest.tools/gotestsum@latest && gotestsum

.PHONY: stop
stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

.PHONY: clean
clean:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --rmi local -v

.PHONY: dockerize
dockerize:
	docker build --build-arg tag_version=$(GIT_TAG) . --tag=$(DOCKER_REPO)paimoncookies:$(GIT_TAG)

.PHONY: go-export
go-export:
	export PATH=$PATH:/usr/local/go/bin && go version

include bin/make/migrations.makefile.mk