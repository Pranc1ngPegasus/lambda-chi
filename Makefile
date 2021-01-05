.PHONY: build wiregen

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/healthcheck ./cmd/healthcheck
	env GOOS=linux go build -ldflags="-s -w" -o bin/server ./cmd/server

wiregen:
	third_party/bin/wire gen ./...

deploy: wiregen build
	sls deploy --verbose
