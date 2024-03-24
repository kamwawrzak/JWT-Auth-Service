BINARY_NAME := jwt-auth-service 

.PHONY: build
build:
	cd cmd && go build -o ../bin/$(BINARY_NAME) .

.PHONY: run
run:
	./bin/$(BINARY_NAME)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run-db
run-db:
	make -C db start-db-and-migrate
