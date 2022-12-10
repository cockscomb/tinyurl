NAME=tinyurl

.PHONY: build run
build:
	go build -o bin/$(NAME) main.go

run:
	go run main.go
