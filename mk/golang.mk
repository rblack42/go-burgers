# Makefile targets for Go projects

REPONAME := github.com/$(GITHUB)/$(PROJNAME)

PHONY: all

all: ## Build Go project:
	go run main.go

.PHONY: test

test: ## Test Go project
	go test

.PHONY: init
init: ## Initialize go module
	go mod init $(REPONAME)
