# !make

include $(CURDIR)/.env

TESTPKGS = $(shell go list ./... | grep -v cmd | grep -v test | grep -v vendor | grep -v script | grep -v examples)

REPO := github.com/kelvintaywl/goreview
HEROKU_APP_NAME := kelvintaywl/goreview
IMAGE_TAG ?= latest


.PHONY: dep
dep:
	dep ensure -v

.PHONY: init
init:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/joho/godotenv/cmd/godotenv
	go get -u github.com/go-playground/overalls
	make dep

.PHONY: test
test:
	godotenv go test -v $(TESTPKGS)

.PHONY: coverage
coverage:
	godotenv overalls \
	-project github.com/kelvintaywl/goreview \
	-covermode atomic \
	-ignore=cmd,examples,vendor, \
	-concurrency 1
	@mv overalls.coverprofile coverage.txt

.PHONY: build
build:
	godotenv go build $(REPO)/cmd/goreview

.PHONY: docker_build
docker_build:
	godotenv docker build --rm -t registry.heroku.com/$(HEROKU_APP_NAME)/web:$(IMAGE_TAG) .

.PHONY: run
run:
	godotenv @echo "TODO: run binary"

.PHONY: docker_run
docker_run:
	godotenv docker run --rm -p 127.0.0.1:$(SERVER_PORT):$(SERVER_PORT) -e GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) registry.heroku.com/$(HEROKU_APP_NAME)/web:$(IMAGE_TAG)
