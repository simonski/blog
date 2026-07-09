.DEFAULT_GOAL := help

DEPLOY_TARGET ?= blog.simonski.com:./blog/

.PHONY: help build run deploy

help: ## Show available commands
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS=":.*?## "}; {printf "  %-8s %s\n", $$1, $$2}'

build: ## Build the static site into output/
	blog build

run: ## Serve output/ on http://localhost:8000
	blog server -dir output

deploy: build ## Deploy output/ via rsync
	rsync -avz --delete output/ $(DEPLOY_TARGET)
