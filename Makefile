# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

init: deps setup-hooks ### Initialize project

mock: ### Gen mock
	rm -rf internal/core/ports/mocks && \
	mockery --dir=internal/core/ports --all --output=internal/core/ports/mocks && \
	go mod tidy

dev:
	set -a && \
		source configs/local.env && \
	set +a && \
	nodemon --exec go run --tags dynamic $(shell pwd)/cmd/server/main.go --signal SIGTERM

run: ### Run server
	set -a && \
		source .env && \
	set +a && \
	go run --tags dynamic $(shell pwd)/cmd/server/main.go

deps: ### Install dependencies
	go mod tidy -v
	go mod vendor -v

swag-init: ### swag init
	go run cmd/docs/main.go init --parseDependency --parseInternal --parseVendor -g cmd/server/main.go

setup-hooks: ### Setup git hooks
	cp .githooks/pre-push .git/hooks/pre-push
	cp .githooks/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-push
	chmod +x .git/hooks/pre-commit
	@echo Hooks installed successfully.

lint: ### Run Go linter
	test -s ./golangci-lint || \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b "$(PWD)" v1.61.0
	./golangci-lint run --config .golangci.yml --timeout 10m -v --modules-download-mode vendor

.PHONY: unit-test
unit-test: ### Run unit tests
	go test ./...

.PHONY: integration-test
integration-test: ### Run integration tests
	go test -v -failfast -tags=integration ./...

.PHONY: coverage
coverage: ### Run coverage tests
	go test -coverprofile=coverage.out -tags=all ./...
	go tool cover -html=coverage.out

.PHONY: gen-mocks
gen-mocks:
	mockgen -destination internal/mocks/repository/sale_hierarchy_repository.go -package repository github.com/centraldigital/cfw-bw-sale-api/internal/core/port SaleHierarchyRepository
	mockgen -destination internal/mocks/adapter/staff_adapter.go -package adapter github.com/centraldigital/cfw-bw-sale-api/internal/core/port StaffAdapter
	mockgen -destination internal/mocks/service/sale_hierarchy_service.go -package adapter github.com/centraldigital/cfw-bw-sale-api/internal/core/port SaleHierarchyService
