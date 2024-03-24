IMAGE_NAME := migrate/migrate
MIGRATIONS_DIR := ./migrations
DATABASE_URL := mysql://root:root_password@tcp(auth-service-db:3306)/users_db
DB_NETWORK := db_auth_service_network 

.PHONY: run-migrations
run-migrations:
	docker run --rm -v $(MIGRATIONS_DIR):/migrations --network $(DB_NETWORK) $(IMAGE_NAME) -path /migrations -database "$(DATABASE_URL)" up

.PHONY: start-db
start-db:
	@docker-compose up -d

.PHONY: start-db-and-migrate
start-db-and-migrate: 
	@make start-db 
	@echo "DB started, waiting for being ready"
	@sleep 30
	@make run-migrations
