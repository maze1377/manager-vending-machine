help:
	@echo "Please use \`make <ROOT>' where <ROOT> is one of"
	@echo "  update-dependencies  to update dependencies"
	@echo "  dependencies         to install the dependencies"
	@echo "  vendingd             to build the main binary for current platform"
	@echo "  test                 to run unittests"
	@echo "  clean                to remove generated files"
	@echo "  lint-fix             format files"
	@echo "  lint             	  run linter"
	@echo "  lint-get             download linter"

clean:
	rm -rf vendingd .test.profile.cov res.json coverage.html

update-dependencies:
	go get -u ./...

dependencies:
	go mod download

vendingd: $(GO_PACKAGES)
	$(GO_VARS) $(GO) build -o="vendingd" -ldflags="$(LD_FLAGS)" --tags="${TAGS}" $(ROOT)/cmd/vending

test: $(GO_PACKAGES)
	$(GO_VARS) $(GO) test $(GO_PACKAGES)  \
	-cover -coverpkg=./... -coverprofile=.test.profile.cov -timeout 30s -v --tags="${TAGS}" && \
	echo -e "\nTesting is passed." && \
	$(GO_VARS) $(GO) tool cover -func .test.profile.cov && \
	$(GO) tool cover -html=.test.profile.cov -o coverage.html

race: $(GO_PACKAGES)
	$(GO_VARS) $(GO) test -race -timeout 30s -v --tags="${TAGS}" $(GO_PACKAGES)

lint-fix: develop-dependencies
	# run go imports for all files
	find . -name \*.go  -exec goimports -w {} \;
	find . -name \*.go  -exec gofumpt -w {} \;
	find . -name \*.go  -exec gci write {} \;
	fieldalignment -fix ./...;

lint-get:
	mkdir -p bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin $(LINTER_VERSION)

lint: dependencies
	./bin/golangci-lint run --config=.golangci.yml

develop-dependencies:
	if [ -z "$$(which mockery)" ]; then go install github.com/vektra/mockery/v2@latest; fi
	if [ -z "$$(which goimports)" ]; then go install golang.org/x/tools/cmd/goimports@latest; fi
	if [ -z "$$(which gocov)" ]; then go install github.com/axw/gocov/gocov@latest; fi
	if [ -z "$$(which gofumpt)" ]; then go install mvdan.cc/gofumpt@latest; fi
	if [ -z "$$(which gci)" ]; then go install github.com/daixiang0/gci@latest; fi
	if [ -z "$$(which fieldalignment)" ]; then go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest; fi

## Project Vars ########################################################################################################
ROOT := github.com/maze1377/manager-vending-machine
.PHONY: help update-dependencies dependencies vendingd test clean

## Commons Vars ########################################################################################################
GO_VARS ?=
GO_PACKAGES := $(shell go list ./... | grep -v /examples/ | grep -v /mocks)
GO ?= go
GIT ?= git
COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= $(shell $(GIT) describe --tags ${COMMIT} 2> /dev/null || echo "$(COMMIT)")
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
LINTER_VERSION ?= v1.52.2
LD_FLAGS := -X $(ROOT).Version=$(VERSION) -X $(ROOT).Commit=$(COMMIT) -X $(ROOT).BuildTime=$(BUILD_TIME)
TAGS := static # Set to `dynamic` for macOS on Apple Silicon
