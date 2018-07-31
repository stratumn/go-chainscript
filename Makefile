# Protobuf parameters
PROTOS=$(shell find proto -name '*.proto')
PROTOS_GO=$(PROTOS:.proto=.pb.go)

# Linter parameters
GO_LINT_CMD=golangci-lint
GO_LINT=$(GO_LINT_CMD) run --build-tags="lint" --deadline=4m --disable="ineffassign" --disable="gas" --tests=false --skip-files=".*\\.pb.go"

# == .PHONY ===================================================================
.PHONY: dep golangcilint deps build lint test protobuf update_chainscript

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

# == protobuf =================================================================
protobuf: $(PROTOS_GO)

%.pb.go: %.proto
	protoc --go_out=. -Iproto $<

# == update_chainscript =======================================================
update_chainscript:
	git subtree pull --prefix proto git@github.com:stratumn/chainscript.git master --squash