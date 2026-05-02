BIN := trello-cli
CLIENT := internal/trello/client.gen.go
SPEC := openapi.json

.PHONY: build gen vet test clean install

build:
	go build -o $(BIN) ./cmd/trello-cli

gen:
	cd internal/trello && oapi-codegen -config oapi-config.yaml ../../$(SPEC)
	go run ./tools/dedup ./$(CLIENT)
	go build ./...

vet:
	go vet ./...

test:
	go test -race ./...

clean:
	rm -f $(BIN)

install:
	go install ./cmd/trello-cli
