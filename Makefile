.PHONY: install-tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: run
run: install-tools
	air

.PHONY: wire
wire: install-tools
	wire
