.PHONY: dev test fmt

.DEFAULT: help
help:
	@echo "make dev"
	@echo "       setup development environment"
	@echo "make test"
	@echo "       run go test"
	@echo "make fmt"
	@echo "       run go fmt"

dev:
	@type pre-commit > /dev/null || (echo "ERROR: pre-commit (https://pre-commit.com/) is required."; exit 1)
	pre-commit install

test:
	find . -maxdepth 1 -type d -regex '\./[^.]*$$' | sed -e 's/^\.\///' | xargs go test -v

fmt:
	gofmt -s -w .
