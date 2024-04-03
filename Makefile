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

.PHONY: test
test:
	go test ./...

.PHONY: test-coverage
test-coverage:
	@mkdir -p tests-results
	@go test -coverprofile=tests-results/coverage.out ./...
	@go tool cover -html=tests-results/coverage.out
