BUILD_NAME := peko
BUILD_DESCRIPTION := A small programming language using the laugh of a certain rabbit.
BUILD_VERSION := $(shell git describe --always)
BUILD_DATE := $(shell date +"%m/%d/%Y %H:%M:%S %Z")

LDFLAGS := "-X 'github.com/h4ckedneko/pekolang/cmd.BuildName=$(BUILD_NAME)' \
	-X 'github.com/h4ckedneko/pekolang/cmd.BuildDescription=$(BUILD_DESCRIPTION)' \
	-X 'github.com/h4ckedneko/pekolang/cmd.BuildVersion=$(BUILD_VERSION)' \
	-X 'github.com/h4ckedneko/pekolang/cmd.BuildDate=$(BUILD_DATE)'"

.PHONY: build \
	install \
	uninstall \
	generate \
	test \
	format \
	clean

build:
	@ go build -ldflags=$(LDFLAGS) -o $(BUILD_NAME) main.go

install:
	@ go install -ldflags=$(LDFLAGS) -o $(BUILD_NAME) main.go

uninstall:
	@ go clean -i

generate:
	@ go generate ./...

test:
	@ go test -v ./...

format:
	@ go fmt ./...

clean:
	@ rm -f $(BUILD_NAME)*
