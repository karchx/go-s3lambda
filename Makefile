.PHONY: build

build:
	goos=linux go build -ldflags="-s -w" -o bin/main main.go
