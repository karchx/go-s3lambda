.PHONY: build

build:
	goos=linux go build -ldflags="-s -w" -o bin/main main.go

.PHONY: deploy_prod

deploy_prod: build
	serverless deploy --stage prod --aws-profile gnu_keneth
