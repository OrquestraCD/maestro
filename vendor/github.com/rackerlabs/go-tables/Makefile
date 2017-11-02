GOLIST = $(shell go list ./... | grep -v /vendor/)

all: test

test:
		@test -z "$(gofmt -s -l . | tee /dev/stderr)"
		@test -z "$(golint $(GOLIST) | tee /dev/stderr)"
		@go test -v -race $(GOLIST)
		@go vet $(GOLIST)
