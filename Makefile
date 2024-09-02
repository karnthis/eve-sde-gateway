all: lint test install

build: go.sum check-go-version

install: go.sum check-go-version
	go install ./...

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

.PHONY: all install install-debug run \
	go-mod-cache draw-deps clean build format \
	test test-all test-build test-cover test-unit test-race \
	test-sim-import-export check-go-version\

# Add check to make sure we are using the proper Go version before proceeding with anything
check-go-version:
	@if ! go version | grep -q "go1.2[0-9]"; then \
		echo "\033[0;31mERROR:\033[0m Go version 1.20+ is required for compiling. It looks like you are using" "$(shell go version) \nPlease download Go version 1.20+ and retry. Thank you!"; \
		exit 1; \
	fi
