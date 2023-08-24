NAME ?= go-common
VERSION ?= 0.1.1

.PHONY: version test coverage

version:
	@echo $(VERSION)

test:
	@echo "Running unit tests..."
	@go get github.com/jstemmer/go-junit-report/v2
	@go test ./... -v -cover -coverprofile=c.out 2>&1 | go-junit-report > report.xml

coverage:
	@echo "Coverage..."
	@go tool cover -html=c.out -o coverage.html
	@go get github.com/t-yuki/gocover-cobertura@v0.0.0-20180217150009-aaee18c8195c
	@gocover-cobertura < c.out > coverage.xml
