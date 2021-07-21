BIN=phpunisher-bin
CMD=./cmd/phpunisher
COVER=test.cover

GO_HTML_COV             := ./coverage.html
GOLANG_DOCKER_IMAGE     := golang:1.16
CC_TEST_REPORTER_ID	:= ${CC_TEST_REPORTER_ID}
CC_PREFIX		:= github.com/mottaquikarim/esquerydsl

GIT_HASH=`git rev-parse --short HEAD`
BUILD_DATE=`date +%FT%T%z`

LDFLAGS=-X main.GitHash=${GIT_HASH} -X main.BuildDate=${BUILD_DATE}
LDFLAGS_REL=-w -s ${LDFLAGS}

.PHONY: clean build

build: vet
	go build -ldflags "${LDFLAGS}" -o "${BIN}" "${CMD}"

release: vet
	go build -ldflags "${LDFLAGS_REL}" -o "${BIN}" "${CMD}"

vet:
	go vet ./...

test:
	go test -race -count 1 -v -coverprofile="${COVER}" ./...

test-cover: test
	go tool cover -func="${COVER}"

test-ci:
	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go test ./... -coverprofile=${COVER}
	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go tool cover -html=${COVER} -o ${GO_HTML_COV}

_before-cc:
	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} \
		/bin/bash -c \
		"curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter"

	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} chmod +x ./cc-test-reporter

	docker run -w /app -v ${ROOT}:/app \
		 -e CC_TEST_REPORTER_ID=${CC_TEST_REPORTER_ID} \
		${GOLANG_DOCKER_IMAGE} ./cc-test-reporter before-build

_after-cc:
	docker run -w /app -v ${ROOT}:/app \
		-e CC_TEST_REPORTER_ID=${CC_TEST_REPORTER_ID} \
		${GOLANG_DOCKER_IMAGE} ./cc-test-reporter after-build --prefix ${PREFIX}

cover-ci: _before-cc test-ci _after-cc

lint:
	golangci-lint run

clean:
	[ -f "${BIN}" ] && rm "${BIN}"
	[ -f "${COVER}" ] && rm "${COVER}"
