#!/bin/sh

# --silent drops the need to prepend `@` to suppress command output.
MAKEFLAGS += --silent

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Run all unit tests
	go test ./...

.PHONY: test-roman
test-roman: ## Run roman unit tests
	go test -v -run ^TestRomanNumerals$ ./...

.PHONY: test-bowling
test-bowling: ## Run bowling unit tests
	go test -run ^TestCalculateBowlingScore$ ./...

.PHONY: test-password
test-password: ## Run password unit tests
	go test -run ^TestGeneratePassword$ ./...

.PHONY: test-generate-data
test-generate-data: ## Generate data for pipeline
	go test -run ^TestGenerateData$ ./...

.PHONY: test-pipeline
test-pipeline: ## Run pipeline
	go test -run ^TestProcessPipeline$ ./...

.PHONY: docker-start
docker-start: ## Start the go runner container
	docker compose up -d

.PHONY: docker-stop
docker-stop: ## Stop the go runner container
	docker compose down

.PHONY: docker-test
docker-test:
	docker exec go-runner bash -c "cd /workdir && go test ./..."
