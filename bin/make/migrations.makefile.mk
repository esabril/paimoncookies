# Migrations for Common DB Scheme
# -------------------------------
.PHONY: all-migrations
all-migrations:
	 go install github.com/pressly/goose/v3/cmd/goose@latest && \
 		cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) up

# make migration-common-create name=new_table_name type=sql
.PHONY: migration-common-create
migration-common-create:
	cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) create $(name) $(type)

.PHONY: migration-common-up
migration-common-up:
	cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) up

# make migration-common-up-to v=13
.PHONY: migration-common-up-to
migration-common-up-to:
	cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) up-to $(v)

.PHONY: migration-common-status
migration-common-status:
	cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) status

# make migration-common-down-to v=4
.PHONY: migration-common-down-to
migration-common-down-to:
	cd migrations/ && \
 		 $(GO_BIN)/goose $(DB_DRIVER) $(DB_DRIVER)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL) down-to $(v)