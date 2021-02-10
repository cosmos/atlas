#!/usr/bin/make -f

COMMIT  := $(shell git rev-parse HEAD)
VERSION := $(shell echo $(shell git describe --always) | sed 's/^v//')
BRANCH  := $(shell git rev-parse --abbrev-ref HEAD)

###############################################################################
#                               Build / Install                               #
###############################################################################

LD_FLAGS = -X github.com/cosmos/atlas/cmd.Version=$(VERSION) \
	-X github.com/cosmos/atlas/cmd.Commit=$(COMMIT) \
	-X github.com/cosmos/atlas/cmd.Branch=$(BRANCH)

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building atlas binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/atlas.exe .
else
	@echo "building atlas binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/atlas .
endif

install: go.sum
	@echo "installing atlas binary..."
	@go install -mod=readonly $(BUILD_FLAGS) .

.PHONY: install build

###############################################################################
#                                    Docs                                     #
###############################################################################

update-swagger-docs:
	@swag init -g server/service.go -o ./docs/api --generatedTime=false

verify-clean-swagger-docs: update-swagger-docs
	@git diff --exit-code docs/api

.PHONY: update-swagger-docs verify-clean-swagger-docs

###############################################################################
#                                 Migrations                                  #
###############################################################################

migrate-down:
	@migrate -database ${ATLAS_DATABASE_URL} -path db/migrations down

migrate-up:
	@migrate -database ${ATLAS_DATABASE_URL} -path db/migrations up

migrate: migrate-down migrate-up

.PHONY: migrate-down migrate-up migrate

###############################################################################
#                                    Tests                                    #
###############################################################################

export ATLAS_MIGRATIONS_DIR ?= $(shell pwd)/db/migrations
export ATLAS_TEST_DATABASE_URL ?= host=localhost port=6432 dbname=postgres user=postgres password=postgres sslmode=disable

test:
	@docker-compose down
	@docker-compose up -d
	@bash -c 'while ! nc -z localhost 6432; do sleep 1; done;'
	@go test -p 1 -v -coverprofile=profile.cov --timeout=20m ./...
	@docker-compose down

test-ci:
	@go test -p 1 -v -coverprofile=profile.cov --timeout=20m ./...

lint:
	@golangci-lint run --timeout 10m

.PHONY: test-docker-db test test-ci lint

publish: 
	@atlas init
	@docker run -v $(shell pwd):/workspace --workdir /workspace interchainio/atlas testKey ./atlas.toml true
.PHONY: publish
