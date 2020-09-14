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
