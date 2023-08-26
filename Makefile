NAME ?= go-common
VERSION ?= 0.4.1

.PHONY: version test coverage

version:
	@echo $(VERSION)

test:
	@echo "Running unit tests..."
	@go test ./... -coverprofile cover.out
	@echo "Coverage details:"
	@go tool cover -func cover.out
