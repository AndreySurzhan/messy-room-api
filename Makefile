LOCAL_BIN:=$(CURDIR)/bin
SPEC:=$(CURDIR)/api/api.yaml

.PHONY: help
help:
	@echo "This is a helper makefile for go-service"
	@echo "Targets:"
	@echo "    generate:    regenerate all generated files"
	@echo "    init         download go dependencies"
	@echo "    lint         run linters"

.PHONY: init
init:
	@echo "Installing dependencies..."
	@go mod download

.PHONY: generate
generate: install-oapi-codegen generate-server clear-bin

install-oapi-codegen:
	@echo "Installing generators..."
	@GOBIN=${LOCAL_BIN} go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

generate-server:
	@echo "Generating code..."
	@${LOCAL_BIN}/oapi-codegen -generate gin -package api -o $(CURDIR)/internal/app/api/api.gen.go ${SPEC}
	@${LOCAL_BIN}/oapi-codegen -generate types -package api -o $(CURDIR)/internal/app/api/types.gen.go ${SPEC}
	@${LOCAL_BIN}/oapi-codegen -generate spec -package api -o $(CURDIR)/internal/app/api/spec.gen.go ${SPEC}

clear-bin:
	@echo "Clearing bin..."
	@rm -rf ${LOCAL_BIN}

install-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN) v1.50.1

.PHONY: lint
lint: install-golangci-lint run-lint clear-bin

.PHONY: run-lint
run-lint: install-golangci-lint
	@echo "Running linters..."
	@${LOCAL_BIN}/golangci-lint run --timeout=5m