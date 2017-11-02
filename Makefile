GOLIST = $(shell go list ./... | grep -v /vendor/)
COMMAND_NAME = maestro
SUPPORTED_SYSTEMS = linux windows darwin
RELEASE = $(shell git describe --always --tags)

ifeq ($(OS), Windows_NT)
	BINARY_NAME = $(COMMAND_NAME).exe
	INSTALLPRE = c:\windows\system32
else
	BINARY_NAME = $(COMMAND_NAME)
	INSTALLPRE = /usr/local
endif

LDFLAGS = "-X main.Version=${RELEASE}"

all: build

build: bin/maestro
bin/maestro:
		@echo "Building binary..."
		go build -ldflags ${LDFLAGS} -o bin/${BINARY_NAME}

test:
		@test -z "$(gofmt -s -l . | tee /dev/stderr)"
		@test -z "$(golint $(GOLIST) | tee /dev/stderr)"
		@go test -race -test.v $(GOLIST)
		@go vet $(GOLIST)

clean:
		@echo "Cleaning up..."
ifeq ($(OS),Windows_NT)
	powershell.exe -Command "if(Test-path .\bin ){ rm .\bin -Recurse -Force}"
else
	rm -rf bin
endif

rebuild: clean build

install:
		@echo "Installing to $(INSTALLPRE)..."
		cp bin/$(BINARY_NAME) $(INSTALLPRE)/bin/

cross_compile: linux windows darwin

linux: bin/linux/maestro
bin/linux/maestro:
		@mkdir -p bin/linux
		GOOS=linux go build -v -o bin/linux/$(COMMAND_NAME)

windows: bin/windows/maestro.exe
bin/windows/maestro.exe:
		@mkdir -p bin/windows
		GOOS=windows go build -v -o bin/windows/$(COMMAND_NAME).exe

darwin: bin/darwin/maestro
bin/darwin/maestro:
		@mkdir -p bin/darwin
		GOOS=darwin go build -v -o bin/darwin/$(COMMAND_NAME)

circleci_package: cross_compile
		./release.sh

.PHONY: linux windows darwin cross_compile rebuild circleci_package
