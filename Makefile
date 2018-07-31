# Protobuf parameters
PROTOS=$(shell find proto -name '*.proto')
PROTOS_GO=$(PROTOS:.proto=.pb.go)

# == .PHONY ===================================================================
.PHONY: build lint test protobuf update_chainscript

# == all ======================================================================
all: build

# == build ====================================================================
build:
	go build ./...

# == lint =====================================================================
lint:
	golint ./...

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