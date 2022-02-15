#!/bin/sh

# --silent drops the need to prepend `@` to suppress command output.
MAKEFLAGS += --silent

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Run unit tests
	go test ./...

.PHONY: docker-start
docker-start: ## Start the go runner container
	docker compose up -d

.PHONY: docker-stop
docker-stop: ## Stop the go runner container
	docker compose down

.PHONY: docker-test
docker-test:
	docker exec go-runner bash -c "cd /workdir && go test ./..."
