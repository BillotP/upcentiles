###
### A Makefile to make things
###
### Note : use `make xxx DOCKER=podman SED=gsed` overrides
### 			style to match your available bins
###
SHELL           := /bin/bash
DOCKER			?= docker
STAGE        	?= dev
SED					?= sed
LOGS			?= logs
CMD				?= server
BUILD           := $(shell git rev-parse --short HEAD)
APP_NAME        := $(shell head -n 1 README.md | cut -d ' ' -f2 |  tr '[:upper:]' '[:lower:]')


REGISTRY_DEV	:= ghcr.io/billotp/upcentiles

help: ## Print this help message and exit
	@echo -e "\n\t\t$(APP_NAME)-$(BUILD) \033[1mmake\033[0m options:\n"
	@perl -nle 'print $& if m{^[a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "- \033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo -e "\n"

default: help

start: ## Start a cmd in background
	@mkdir -p logs
	@echo "[INFO] Will start $(CMD) listening on $(PORT)"
	@FOO=$(CMD) ; \
	FOO=$${FOO##*/} ; \
	{ PORT=$(PORT) HEALTH_PORT=$(HEALTH_PORT) CMD=$(CMD) go run cmd/$(CMD)/main.go &>"$(LOGS)/$$FOO.log" & }
	@echo "[INFO] Started $(CMD)"

test: ## Run all unit tests and get detailled coverage
	@rm cover.out coverage.html || true
	@go test -covermode=count -coverprofile=coverage.out -v ./... || true
	@go tool cover -func coverage.out
	@go tool cover -html=coverage.out -o coverage.html

benchmark: ## Run the stats package benchmarks
	@go test -bench=. ./...

lint: ## Run golangci-lint with the Upfluence config
	@$(DOCKER) run -t --rm -v $(shell pwd):/app -w /app "docker.io/golangci/golangci-lint:latest-alpine" golangci-lint run -v

container: ## Build a CMD container image
	@export version=$$(git rev-parse --short HEAD) && \
	export registry=$(REGISTRY_DEV) && \
	export imageName="$$registry/$(CMD)" && \
	export dockerfile=golang.dockerfile && \
	echo "$$dockerfile" && \
	$(DOCKER) build -t "$$imageName:$$version" \
	--build-arg version=$$version \
	--build-arg cmd=$(CMD) \
	-f $$dockerfile .

push: ## Push a previously built CMD container image
	@export version=$(git rev-parse --short HEAD) && \
	export registry=$(shell if [ $(STAGE) == "preprod" ]; then echo $(REGISTRY_PREPROD); else echo $(REGISTRY_DEV);fi) && \
	export imageName="$$registry/$(CMD)" && \
	echo "Will push $$imageName:$$version" && \
	$(DOCKER) push "$$imageName:$$version"

deploy: ## Deploy using flyctl if authenticated
	@fly deploy -c fly.toml --build-arg "version=v$(BUILD)" --build-arg "cmd=delegationz"

clean: ## Cleaning binary and temporary files 
	@rm -rf bin || true
	@rm -rf out || true
	@rm -rf logs || true
	@rm *.log **/*.log *.pid **/*.pid || true


.PHONY: help clean test all build push container
