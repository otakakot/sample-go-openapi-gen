SHELL := /bin/bash
include .env
export
export APP_NAME := $(basename $(notdir $(shell pwd)))

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PNONY: tool
tool: ## install the tools
	@go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: gen
gen: ## generate the code
	@go generate ./...
	@oapi-codegen -generate server -package oapi openapi.yaml > pkg/oapi/server.gen.go
	@oapi-codegen -generate spec -package oapi openapi.yaml > pkg/oapi/spec.gen.go
	@oapi-codegen -generate types -package oapi openapi.yaml > pkg/oapi/types.gen.go
	@kiota generate --language Go --class-name Kiota --namespace-name github.com/otakakot/sample-go-openapi-gen/pkg/kiota --openapi ./openapi.yaml --output ./pkg/kiota
	@go get -u -t ./...
	@go mod tidy

up: ## docker compose up with air hot reload
	@docker compose --project-name ${APP_NAME} --file ./.docker/compose.yaml up -d

.PHONY: down
down: ## docker compose down
	@docker compose --project-name ${APP_NAME} down --volumes
