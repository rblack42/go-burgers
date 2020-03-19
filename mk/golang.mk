# Makefile targets for Go projects

PHONY: all

all: ## Build Go project:
	go run main.go

.PHONY: test

test: ## Test Go project
	go test

.PHONY: init
init: ## INitialize go module
	go mod init github.com/rblack42/$(PROJNAME)
