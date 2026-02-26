.DEFAULT_GOAL := help

.PHONY: help all build run deploy

help: ## Show available commands
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS=":.*?## "}; {printf "  %-8s %s\n", $$1, $$2}'

all: build run ## Build site and run local server

build: ## Build site output
	go run ./blog.go

run: ## Serve output on http://127.0.0.1:8000
	python3 -m http.server 8000 --bind 127.0.0.1 --directory output

deploy: build ## Deploy output to production host
	scp -r output/* blog.simonski.com:./blog
