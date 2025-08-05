export CGO_ENABLED=1
export GO111MODULE=on
#export GOOS=linux
#export GOARCH=amd64

GOPRIVATE := $(shell go env GOPRIVATE)
GOSUMDB := $(shell go env GOSUMDB)
GOPROXY := $(shell go env GOPROXY)

test:
	go test -race; \
	go test -bench=.
