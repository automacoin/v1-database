# v1 Database

VERSION=v1.0.0
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

clean:
	rm -f build/*

database: clean
	go build ${LDFLAGS} -o build/v1-database main.go

test:
	ginkgo -r

.PHONY: clean database test