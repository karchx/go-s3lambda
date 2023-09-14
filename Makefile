.PHONY: build

build:
	goos=linux go build -o main main.go && zip main.zip main
