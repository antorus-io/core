## help: show this help tool
.PHONY: help
help:
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## audit: tidy dependencies, vet
.PHONY: audit
audit:
	@go mod tidy
	@go mod verify
	@go vet ./...

## build: build project
.PHONY: build
build:
	@TARGET=build docker compose --file=./docker/production/docker-compose.yml --project-name=antorus up --build --exit-code-from=build --remove-orphans


## clean: clean up build artifacts and temporary data
.PHONY: clean
clean:
	@rm -rf core coverage/ tmp/ vendor/

## lint/all: run all lint targets
.PHONY: lint/all
lint/all:
	@$(MAKE) lint/actionlint
	@$(MAKE) lint/golangci-lint
	@$(MAKE) lint/hadolint
	@$(MAKE) lint/yamllint

## lint/actionlint: run actionlint target
.PHONY: lint/actionlint
lint/actionlint:
	@TARGET=actionlint docker compose --file=./docker/lint/actionlint/docker-compose.yml pull && \
	TARGET=actionlint docker compose --file=./docker/lint/actionlint/docker-compose.yml --project-name=antorus up --exit-code-from=actionlint --remove-orphans

## lint/golangci-lint: run golangci-lint target
.PHONY: lint/golangci-lint
lint/golangci-lint:
	@TARGET=golangci-lint docker compose --file=./docker/lint/golangci-lint/docker-compose.yml pull && \
	TARGET=golangci-lint docker compose --file=./docker/lint/golangci-lint/docker-compose.yml --project-name=antorus up --exit-code-from=golangci-lint --remove-orphans

## lint/hadolint: run hadolint target
.PHONY: lint/hadolint
lint/hadolint:
	@TARGET=hadolint docker compose --file=./docker/lint/hadolint/docker-compose.yml pull && \
	TARGET=hadolint docker compose --file=./docker/lint/hadolint/docker-compose.yml --project-name=antorus up --exit-code-from=hadolint --remove-orphans

## lint/yamllint: run yamllint target
.PHONY: lint/yamllint
lint/yamllint:
	@TARGET=yamllint docker compose --file=./docker/lint/yamllint/docker-compose.yml pull && \
	TARGET=yamllint docker compose --file=./docker/lint/yamllint/docker-compose.yml --project-name=antorus up --exit-code-from=yamllint --remove-orphans

## test/unit: run unit test suites
.PHONY: test/unit
test/unit:
	@TARGET=test docker compose --file=./docker/test/docker-compose.yml --project-name=antorus up --build --exit-code-from=test --remove-orphans

## test/coverage: show coverage statistics
.PHONY: test/coverage
test/coverage:
	@TARGET=coverage docker compose --file=./docker/test/docker-compose.yml --project-name=antorus up --build --exit-code-from=test --remove-orphans

## upgrade: run dependency upgrade command along with tidy and vendoring
.PHONY: upgrade
upgrade:
	@go get -u ./... && \
	go mod tidy && \
	go mod vendor && \
	go mod verify
