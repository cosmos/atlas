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
#                                 Migrations                                  #
###############################################################################

migrate:
	@migrate -database ${ATLAS_DATABASE_URL} -path db/migrations down
	@migrate -database ${ATLAS_DATABASE_URL} -path db/migrations up

.PHONY: migrate

###############################################################################
#                                    Tests                                    #
###############################################################################

export ATLAS_MIGRATIONS_DIR ?= $(shell pwd)/db/migrations
export ATLAS_TEST_DATABASE_URL ?= "host=localhost port=6432 dbname=postgres user=postgres password=postgres sslmode=disable"

test:
	@docker-compose down
	@docker-compose up -d
	@go test -v -coverprofile=profile.cov --timeout=20m ./...

test-ci:
	@go test -v -coverprofile=profile.cov --timeout=20m ./...

lint:
	@golangci-lint run --timeout 10m

.PHONY: test-docker-db test test-ci lint
