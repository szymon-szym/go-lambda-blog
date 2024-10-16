.PHONY: build local deploy

generate:
	cd lambda_handlers && go generate ./...

build: generate
	rm -rf .aws-sam
	sam build

local: build
	sam local start-api --env-vars parameters.json

deploy: build
	sam deploy