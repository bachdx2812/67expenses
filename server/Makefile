MIGRATIONS_FOLDER = $(PWD)/database/migrations

POSTGRES_DATABASE_NAME ?= home_expenses
POSTGRES_HOST ?= localhost
POSTGRES_PORT ?= 5432
POSTGRES_USERNAME ?= postgres
POSTGRES_PASSWORD ?= 
POSTGRES_SSL_MODE ?= disable

POSTGRESQL_URL=postgres://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE_NAME)?sslmode=$(POSTGRES_SSL_MODE)

db.migrate:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(POSTGRESQL_URL)" --verbose up

db.migrate.rollback:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(POSTGRESQL_URL)" --verbose down ${STEP}

db.migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(POSTGRESQL_URL)" --verbose force $(VERSION)