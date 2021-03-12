.PHONY: dev test fmt

.DEFAULT: help
help:
	@echo "make dev"
	@echo "       setup development environment"
	@echo "make test"
	@echo "       run go test"
	@echo "make fmt"
	@echo "       run go fmt"

SUBDIRS := $(wildcard */)
define FOREACH
	for DIR in $(SUBDIRS); do \
  		cd $$DIR && $(1) && cd $(CURDIR); \
  	done
endef

dev:
	@type pre-commit > /dev/null || (echo "ERROR: pre-commit (https://pre-commit.com/) is required."; exit 1)
	pre-commit install

test:
	$(call FOREACH,go test -v)

cover:
	$(call FOREACH,go test -cover -v)

tidy:
	$(call FOREACH,go mod tidy)

fmt:
	gofmt -s -w .
