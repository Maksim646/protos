PROTO_FILES := proto/link/link.proto
GEN_DIR := ./gen/go

# Проверка наличия protoc
PROTOC ?= protoc
PROTOC_VERSION_EXPECTED := 25  # Замените на ожидаемую версию

# Проверка наличия protoc-gen-go и protoc-gen-go-grpc
PROTOC_GEN_GO ?= protoc-gen-go
PROTOC_GEN_GO_GRPC ?= protoc-gen-go-grpc

all: generate

generate: check_protoc check_plugins
	protoc -I proto $(PROTO_FILES) --go_out=$(GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative
	@echo "Successfully generated Go code"

clean:
	rm -rf $(GEN_DIR)
	@echo "Cleaned generated Go code"

check_protoc:
	@which $(PROTOC) > /dev/null 2>&1 || (echo "Error: protoc not found. Please install it." && exit 1)

check_plugins:
	@which $(PROTOC_GEN_GO) > /dev/null 2>&1 || (echo "Error: protoc-gen-go not found. Please install it with 'go install google.golang.org/protobuf/cmd/protoc-gen-go@latest'" && exit 1)
	@which $(PROTOC_GEN_GO_GRPC) > /dev/null 2>&1 || (echo "Error: protoc-gen-go-grpc not found. Please install it with 'go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest'" && exit 1)

.PHONY: all generate clean check_protoc check_plugins