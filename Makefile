BINARY=winter-is-coming

.PHONY: run build test all

all: build

build:
	go build -o ./build/${BINARY} ./cmd/main.go

test:
	go test -v ./...

run: build
	env $(shell cat ./cmd/environment) ./build/${BINARY}

