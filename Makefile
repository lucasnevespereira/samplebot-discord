run:
	GOFLAGS=-mod=mod go run cmd/main.go -t $(BOT_TOKEN)

build:
	GOFLAGS=-mod=mod go build -o bin/sample-bot-discord cmd/main.go

exec:
	./bin/sample-bot-discord -t $(BOT_TOKEN)
