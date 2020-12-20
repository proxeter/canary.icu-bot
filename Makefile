default: build

.PHONY: build
build:
	go build -o ./output/bot ./cmd/icu-bot
