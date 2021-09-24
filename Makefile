BINDIR ?= $(GOPATH)/bin
VERSION := $(shell git describe --tags)
COMMIT := $(shell git log -1 --format='%H')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')
LEDGER_ENABLED ?= true

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  UNAME_S = $(shell uname -s)
  ifeq ($(UNAME_S),OpenBSD)
    $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
  else
    GCC = $(shell command -v gcc 2> /dev/null)
    ifeq ($(GCC),)
      $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  endif
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=rizon \
	  -X github.com/cosmos/cosmos-sdk/version.AppName=rizond \
	  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
	  -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

.PHONY: install protocgen update-swagger-docs

all: install

update-swagger-docs: statik
	$(BINDIR)/statik -src=client/docs/swagger-ui -dest=client/docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
        exit 1;\
    else \
        echo "\033[92mSwagger docs are in sync\033[0m";\
    fi

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/rizond

go.sum: go.mod
	@go mod verify
	@go mod tidy

protocgen:
	bash ./scripts/protocgen.sh
