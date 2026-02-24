.DEFAULT_GOAL := all

.PHONY: all build run deploy

all: build run

build:
	go run ./blog.go

run:
	python3 -m http.server 8000 --bind 127.0.0.1 --directory output

deploy:
	scp -r output/* blog.simonski.com:
