BIN_DIR  := bin
BIN      := $(BIN_DIR)/trello-cli
CMD      := ./cmd/trello-cli
SPEC     := openapi.json
CLIENT   := internal/trello/client.gen.go

VERSION  ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT   ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE     ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

PKG      := github.com/mrfansi/trello-cli/internal/version
LDFLAGS  := -X $(PKG).Version=$(VERSION) -X $(PKG).Commit=$(COMMIT) -X $(PKG).Date=$(DATE)

GOFLAGS  ?=

.DEFAULT_GOAL := build

.PHONY: help build install run test ci vet fmt lint tools clean gen gen-cmds version snapshot

help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make <target>\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build: ## Build the binary into ./bin/trello-cli
	@mkdir -p $(BIN_DIR)
	go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BIN) $(CMD)
	@echo "built $(BIN) ($(VERSION))"

install: ## Install the binary into $$GOBIN / $$GOPATH/bin
	go install $(GOFLAGS) -ldflags "$(LDFLAGS)" $(CMD)

run: build ## Build and run; pass args after `--`, e.g. make run -- me
	@$(BIN) $(filter-out $@,$(MAKECMDGOALS))

version: ## Print the version that `make build` would embed
	@echo "version=$(VERSION) commit=$(COMMIT) date=$(DATE)"

test: ## Run tests with race detector and coverage
	go test -race -cover ./...

ci: vet test ## Run vet + tests (used by CI)

vet: ## go vet ./...
	go vet ./...

fmt: ## gofmt -s -w .
	gofmt -s -w .

lint: ## golangci-lint run (requires `make tools`)
	@if [ -x .tools/golangci-lint ]; then .tools/golangci-lint run ./...; else go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...; fi

tools: ## Install dev tooling (golangci-lint, goimports) into ./.tools
	@mkdir -p .tools
	GOBIN=$(CURDIR)/.tools go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	GOBIN=$(CURDIR)/.tools go install golang.org/x/tools/cmd/goimports@latest

gen: ## Regenerate the typed Trello client from openapi.json
	cd internal/trello && oapi-codegen -config oapi-config.yaml ../../$(SPEC)
	go run ./tools/dedup ./$(CLIENT)
	go build ./...

gen-cmds: ## Regenerate cobra commands + docs/COMMANDS.md from openapi.json
	go run ./tools/cmdgen $(SPEC) internal/commands/auto --docs docs/COMMANDS.md
	go build ./...

snapshot: ## Build a local goreleaser snapshot (requires `goreleaser`)
	goreleaser release --snapshot --clean

clean: ## Remove ./bin and ./dist
	rm -rf $(BIN_DIR) dist

# Allow `make run -- me` etc. by swallowing extra args.
%:
	@:
