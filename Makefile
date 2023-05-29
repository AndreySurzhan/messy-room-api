LOCAL_BIN:=$(CURDIR)/bin

LINTER:= $(LOCAL_BIN)/golangci-lint
LINTER_VERSION:=1.52.2

OAPI_CODEGEN:=$(LOCAL_BIN)/oapi-codegen
SPEC:=$(CURDIR)/api/api.yaml
PACKAGE_NAME:=api
GENERATED_PATH:=$(CURDIR)/internal/app/api

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
generate: $(OAPI_CODEGEN)
	@echo "Generating code..."
	@$(OAPI_CODEGEN) -generate gin -package $(PACKAGE_NAME) -o $(GENERATED_PATH)/api.gen.go $(SPEC)
	@$(OAPI_CODEGEN) -generate types -package $(PACKAGE_NAME) -o $(GENERATED_PATH)/types.gen.go $(SPEC)
	@$(OAPI_CODEGEN) -generate spec -package $(PACKAGE_NAME) -o $(GENERATED_PATH)/spec.gen.go $(SPEC)

.PHONY: generate-clear
generate-clear: generate clear-bin

$(OAPI_CODEGEN):
	@echo "Installing oapi-codegen..."
	@GOBIN=$(LOCAL_BIN) go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

.PHONY: clear-bin
clear-bin:
	@echo "Clearing bin..."
	@rm -rf $(LOCAL_BIN)

$(LINTER):
	@echo "Installing golangci-lint..."
	@curl sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN) v$(LINTER_VERSION)

.PHONY: lint
lint: $(LINTER)
	@echo "Running linters..."
	@$(LOCAL_BIN)/golangci-lint run --timeout=5m --config=.golangci.yaml

.PHONY: lint-clear
lint-clear: lint clear-bin

