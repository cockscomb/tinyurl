.PHONY: install-tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: run
run: install-tools
	air

.PHONY: test
test:
	go test -v ./...

.PHONY: wire
wire: install-tools
	wire
