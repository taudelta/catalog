.PHONY all
all:
	build test lint

.PHONY build
build:
	CGO_ENABLED=0 && go build -o main

.PHONY test
test:
	go test ./...

.PHONY lint
lint:
	go vet ./...