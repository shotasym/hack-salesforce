SHELL := /bin/bash

BINARY_NAME := hack-salesforce
BINARY_FILE := $(GOPATH)/bin/$(BINARY_NAME)

.PHONY: build
build:
	go build -o $(BINARY_FILE) -v ./src

.PHONY: test
test:
	go test -v ./src

.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_FILE)