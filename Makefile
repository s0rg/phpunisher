SHELL=/bin/bash

BIN=bin/phpunisher
CMD=./cmd/phpunisher
COP=test.coverage

GIT_TAG=`git describe --abbrev=0 2>/dev/null || echo -n "no-tag"`
GIT_HASH=`git rev-parse --short HEAD 2>/dev/null || echo -n "no-git"`
BUILD_AT=`date +%FT%T%z`

LDFLAGS=-w -s -X main.gitHash=${GIT_HASH} -X main.buildDate=${BUILD_AT} -X main.gitVersion=${GIT_TAG}

export CGO_ENABLED=0
export GOARCH=amd64

.PHONY: build

build: vet
	go build -ldflags "${LDFLAGS}" -o "${BIN}" "${CMD}"

vet:
	go vet ./...

test: vet
	CGO_ENABLED=1 go test -race -count 1 -v -tags=test -coverprofile="${COP}" ./...

test-cover: test
	go tool cover -func="${COP}"

lint:
	golangci-lint run

clean:
	[ -f "${BIN}" ] && rm "${BIN}"
	[ -f "${COP}" ] && rm "${COP}"
