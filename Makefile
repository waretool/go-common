NAME ?= go-common
VERSION ?= 0.2.0

.PHONY: version test coverage

version:
	@echo $(VERSION)

test:
	@echo "Running unit tests..."
	@go test ./... -cover
