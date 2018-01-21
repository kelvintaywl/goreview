# !make

include $(CURDIR)/.env

REPO := github.com/kelvintaywl/goreview
IMAGE_TAG ?= latest


.PHONY: dep
dep:
	dep ensure -v

.PHONY: init
init:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/joho/godotenv/cmd/godotenv
	make dep

.PHONY: build
build:
	godotenv go build $(REPO)/cmd/goreview

.PHONY: docker_build
docker_build:
	godotenv docker build --rm -t kelvintaywl/goreview:$(IMAGE_TAG) .

.PHONY: run
run:
	godotenv @echo "TODO: run binary"

.PHONY: docker_run
docker_run:
	godotenv docker run --rm -p 127.0.0.1:$(SERVER_PORT):$(SERVER_PORT) kelvintaywl/goreview:$(IMAGE_TAG)
