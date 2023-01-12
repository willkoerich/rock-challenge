install-mocks:
	@go install github.com/vektra/mockery/v2@latest

build-mocks:
	@~/go/bin/mockery --dir ${PWD}:/internal /internal/mocks --all
	@mockery --recursive --all --dir ./internal --output ./internal/mocks --case snake -r --inpackage --keeptree

test:
	@go test ./..

linter-install:
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	@~/go/bin/golangci-lint run -v -c .code_quality/.golangci.yml --skip-dirs /internal/mocks

fmt:
	@go fmt ./...