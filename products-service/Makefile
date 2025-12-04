.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY: lint run-build run clean docker-build docker-run \
docker-clean help gen build rebuild vendor mod format vet update \
aligo-install aligo-check aligo-view

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

LOCAL_BIN:=$(CURDIR)/bin

APP_NAME := products-service
DOCKER_NETWORK := test-services
PORT := 10000
VERSION := v1.0.0
COMMIT_LAST := $(shell git rev-parse HEAD)
COMMIT_DATE := $(shell git log -1 --date=format:"%Y-%m-%dT%TZ" --format="%ad")

LDFLAGS="-X main.name=$(APP_NAME) \
       -X 'main.version=$(VERSION)' \
       -X 'main.commit=$(COMMIT_LAST)' \
       -X 'main.date=$(COMMIT_DATE)'"

help:
	@echo "\033[1;33m🚀 Go commands:\033[0m"
	@echo "\033[1;34m  build                \t Build the application\033[0m"
	@echo "\033[1;34m  clean                \t Clean build artifacts\033[0m"
	@echo "\033[1;34m  format               \t Format code with 'go fmt'\033[0m"
	@echo "\033[1;34m  gen                  \t Run go generate\033[0m"
	@echo "\033[1;34m  lint                 \t Run golangci-lint\033[0m"
	@echo "\033[1;34m  mod                  \t Tidy go modules\033[0m"
	@echo "\033[1;34m  rebuild              \t Clean and rebuild the application\033[0m"
	@echo "\033[1;34m  run                  \t Run the application\033[0m"
	@echo "\033[1;34m  run-build            \t Build and run the application\033[0m"
	@echo "\033[1;34m  update               \t Update dependencies\033[0m"
	@echo "\033[1;34m  vendor               \t Create vendor directory\033[0m"
	@echo "\033[1;34m  vet                  \t Run 'go vet'\033[0m"
	@echo ""
	@echo "\033[1;33m🐳 Docker commands:\033[0m"
	@echo "\033[1;35m  docker-build         \t Build Docker image\033[0m"
	@echo "\033[1;35m  docker-clean         \t Stop and remove Docker container and image\033[0m"
	@echo "\033[1;35m  docker-run           \t Build and run Docker container\033[0m"
	@echo ""
	@echo "\033[1;33m🛠 Other commands:\033[0m"
	@echo "\033[1;32m  aligo-install        \t Install aligo\033[0m"
	@echo "\033[1;32m  aligo-check          \t Run aligo check\033[0m"
	@echo "\033[1;32m  aligo-view           \t Run aligo view\033[0m"
	@echo ""

build:
	go build -ldflags=$(LDFLAGS) -o ./build/$(APP_NAME) ./cmd/.

clean:
	go clean ./...
	rm -f ./cmd/$(APP_NAME)

format:
	go fmt ./...

gen:
	go generate ./...

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... -v --fix --config golangci.pipeline.yaml

mod:
	go mod tidy

rebuild: clean build

run:
	go run -ldflags=$(LDFLAGS) ./cmd/main.go -c config.yaml

run-build: build
	./build/$(APP_NAME)

update:
	go get -u -v ./...
	go mod tidy

vendor:
	go mod vendor

vet:
	go vet ./...

docker-build:
	docker build --build-arg SERVICE_NAME=$(APP_NAME) \
              --build-arg BUILD_DATE=$(COMMIT_DATE) \
              --build-arg VERSION=$(VERSION) \
              --build-arg COMMIT=$(COMMIT_LAST) \
              -t $(APP_NAME):$(VERSION) .

docker-clean:
	-docker ps -a --filter "name=$(APP_NAME)" -q | xargs -r docker stop
	-docker ps -a --filter "name=$(APP_NAME)" -q | xargs -r docker rm
	-docker images "$(APP_NAME)" -q | xargs -r docker rmi

docker-run: docker-clean docker-build
	docker run --env-file .env  \
	--network $(DOCKER_NETWORK) -p $(PORT):$(PORT) \
	-d --name $(APP_NAME) $(APP_NAME):$(VERSION)

aligo-install:
	go install github.com/essentialkaos/aligo/v2@latest

aligo-check:
	aligo check ./...

aligo-view:
	aligo view ./...
