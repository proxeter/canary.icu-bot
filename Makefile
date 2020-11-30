default: build

.PHONY: build
build:
	go build -o ./target/bot

.PHONY: lint
lint:
	golint
