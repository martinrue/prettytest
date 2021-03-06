VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Commit=${COMMIT}"
MODFLAGS=-mod=vendor

BINARY=prettytest
PACKAGE=./cmd/prettytest

all: dev

clean:
	rm -fr dist/

dev:
	go build ${MODFLAGS} ${LDFLAGS} -o dist/${BINARY} ${PACKAGE}

dist: darwin linux

darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o dist/${BINARY}-darwin-amd64 ${PACKAGE}

linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o dist/${BINARY}-linux-amd64 ${PACKAGE}

test:
	go test ${MODFLAGS} ./...

.PHONY: all clean dev dist darwin linux test
