# Make dependencies, clean and build proto.
deps:
	dep version || (curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh)
	dep ensure -v

docker_login:
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin https://$(REGISTRY)

docker_build: deps build_linux
	docker build -f build/Dockerfile -t $(DOCKER_IMAGE):latest .

PWD = $(shell pwd)

MODULE = dobi-oms
IMAGE_TAG ?= $(MODULE)
GITHUB_SHA ?= $(MODULE)

SRC = `go list -f {{.Dir}} ./... | grep -v /vendor/`

install:
	@echo "==> Installing tools..."
	@go install golang.org/x/tools/...
	@go install golang.org/x/lint/golint
	@GO111MODULE=off go get github.com/golang/mock/mockgen
	@GO111MODULE=off go get mvdan.cc/gofumpt/gofumports
	@GO111MODULE=off go get github.com/daixiang0/gci
	@brew install golangci/tap/golangci-lint
	@brew upgrade golangci/tap/golangci-lint

test-coverage:
	@echo "==> Running unit tests with HTML coverage report..."
	@go clean -testcache ./...
	@go test ./... -p 1 -coverprofile cover.out
	@go tool cover -html cover.out -o cover.out.html

generate:
	@echo "==> Generating code..."
	@go generate ./...

build:
	@docker build \
		--build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
		--build-arg COMMIT_HASH=$(COMMIT_HASH) \
		--secret id=passwordaeskey,src=$(PASSWORD_AES_KEY) \
		--target release \
		-f build/Dockerfile \
		-t $(IMAGE_TAG) .

test-up:
	@COMPOSE_HTTP_TIMEOUT=180 docker-compose \
		-f build/docker-compose.test.yml \
		-p $(GITHUB_SHA) up \
		--force-recreate \
		--abort-on-container-exit \
		--exit-code-from app \
		--build

test-down:
	@COMPOSE_HTTP_TIMEOUT=180 docker-compose \
		-f build/docker-compose.test.yml \
		-p $(GITHUB_SHA) down \
 		-v --rmi local

dev-up:
	@docker compose \
		-f build/docker-compose.dev.yml \
		-p $(GITHUB_SHA) up -d

dev-down:
	@docker compose \
		-f build/docker-compose.dev.yml \
		-p $(GITHUB_SHA) down -v --rmi local

.PHONY: all fmt lint test test-unit test-integration install dev-up dev-down dev-ps test-up test-down build
