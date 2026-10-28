export CGO_CFLAGS := $(shell go env CGO_CFLAGS) $(shell pkg-config --cflags wayland-client xkbcommon)

.PHONY: run build test

run:
	go run .

build:
	go build .

test:
	go test ./...
