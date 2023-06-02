LOCAL_BIN:=$(CURDIR)/bin

LINTER:= $(LOCAL_BIN)/golangci-lint
LINTER_VERSION:=1.52.2

OAPI_CODEGEN:=$(LOCAL_BIN)/oapi-codegen
SPEC:=$(CURDIR)/api/api.yaml
PACKAGE_NAME:=api
GENERATED_PATH:=$(CURDIR)/internal/app/api
GENERATED_API:=api.gen.go

.PHONY: help
help:
	@echo "This is a helper makefile for go-service"
	@echo "Targets:"
	@echo "    generate:    regenerate all generated files"
	@echo "    generate-clear:    regenerate all generated files and clear bin"
	@echo "    lint         run linters"
	@echo "    lint-clear   run linters and clear bin"

.PHONY: init
init:
	@echo "Installing dependencies..."
	@go mod download

.PHONY: generate
generate: $(OAPI_CODEGEN)
	@echo "Generating code..."
	@$(OAPI_CODEGEN) -generate gin,types -package $(PACKAGE_NAME) -o $(GENERATED_PATH)/$(GENERATED_API) $(SPEC)
	@sed -i '' 's/Id/ID/g' $(GENERATED_PATH)/$(GENERATED_API)
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

