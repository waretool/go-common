NAME ?= go-common
VERSION ?= 0.7.0

.PHONY: version tests

version:
	@echo $(VERSION)

tests:
	@echo "Running unit tests..."
	@go test ./... -coverprofile cover.out
	@echo "Coverage details:"
	@go tool cover -func cover.out
