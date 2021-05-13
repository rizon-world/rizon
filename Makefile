VERSION := $(shell git describe --tags)
COMMIT := $(shell git log -1 --format='%H')
BUILDTAGS := $(shell uname)
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=rizon \
	  -X github.com/cosmos/cosmos-sdk/version.AppName=rizond \
	  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(BUILDTAGS)" \
	  -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

.PHONY: install protocgen

all: install

install: go.sum
	go install -mod=readonly -ldflags '$(ldflags)' ./cmd/rizond

go.sum: go.mod
	@go mod verify
	@go mod tidy

protocgen:
	bash ./scripts/protocgen.sh
