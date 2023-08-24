NAME ?= go-common
VERSION ?= 0.1.1

.PHONY: version test coverage

version:
	@echo $(VERSION)

test:
	@echo "Running unit tests..."
	@go test ./... -cover
