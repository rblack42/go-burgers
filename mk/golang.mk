# Makefile targets for Go projects

PHONY: all

all: ## Build Go project:
	go run main.go

.PHONY: test

test: ## Test Go project
	go test
