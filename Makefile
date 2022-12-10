.PHONY: install-tools
install-tools:
	go install github.com/cosmtrek/air@latest

.PHONY: run
run: install-tools
	air
