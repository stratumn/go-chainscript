# Protobuf parameters
PROTOS=$(shell find proto -name '*.proto')
PROTOS_GO=$(PROTOS:.proto=.pb.go)

# Linter parameters
GO_LINT_CMD=golangci-lint
GO_LINT=$(GO_LINT_CMD) run --build-tags="lint" --deadline=4m --disable="ineffassign" --disable="gas" --tests=false --skip-files=".*\\.pb.go"

# Test parameters
COVERAGE_FILE=coverage.txt

# Test data
TESTDATA_FILE=./samples/go-samples.json

# == .PHONY ===================================================================
.PHONY: dep golangcilint deps build lint test coverage protobuf update_chainscript docker testdata_generate testdata_validate

# == all ======================================================================
all: build

# == deps =====================================================================
dep:
	go get -u github.com/golang/dep/cmd/dep

golangcilint:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

deps: dep golangcilint
	dep ensure

# == build ====================================================================
build:
	go build ./...

# == lint =====================================================================
lint:
	$(GO_LINT) ./...

# == test =====================================================================
test:
	go test ./...

# == coverage =================================================================
coverage: $(COVERAGE_FILE)

$(COVERAGE_FILE):
	go test ./... -coverprofile=$(COVERAGE_FILE) -covermode=atomic

# == protobuf =================================================================
protobuf: $(PROTOS_GO)

%.pb.go: %.proto
	protoc --go_out=. -Iproto $<

# == update_chainscript =======================================================
update_chainscript: SHELL:=/bin/bash
update_chainscript:
	@if [[ -n $$(git diff --stat) ]]; then \
		echo "Chainscript cannot be updated: you have unstaged changes in your working tree. Please commit or stash your current work and retry."; \
	elif [[ -n $$(git status -s) ]]; then \
		echo "Chainscript cannot be updated: you have staged changes. Please commit or stash your current work and retry."; \
	else \
		rm -r ./proto; \
		git add ./proto; \
		git commit -m "removing old chainscript subtree"; \
		git subtree add --prefix proto git@github.com:stratumn/chainscript.git master --squash; \
	fi

# == docker ===================================================================
docker:
	docker build -t stratumn/go-chainscript:latest .

# == testdata_generate ========================================================
testdata_generate:
	go run cmd/*.go generate $(TESTDATA_FILE)

# == testdata_validate ========================================================
testdata_validate:
	go run cmd/*.go validate $(TESTDATA_FILE)