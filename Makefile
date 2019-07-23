BIN=phpunisher-bin
CMD=./cmd/phpunisher

GIT_HASH=`git rev-parse --short HEAD`
BUILD_DATE=`date +%FT%T%z`

LDFLAGS=-X main.GitHash=${GIT_HASH} -X main.BuildDate=${BUILD_DATE}
LDFLAGS_REL=-w -s ${LDFLAGS}

.PHONY: clean build

build:
	go build -ldflags "${LDFLAGS}" -o "${BIN}" "${CMD}"

release:
	go build -ldflags "${LDFLAGS_REL}" -o "${BIN}" "${CMD}"

vet:
	go vet

clean:
	[ -f "${BIN}" ] && rm "${BIN}"

