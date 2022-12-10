NAME=tinyurl

.PHONY: build
build:
	go build -o tmp/$(NAME) main.go

.PHONY: install-tools
install-tools:
	go install github.com/cosmtrek/air@latest

.PHONY: run
run: install-tools
	air
