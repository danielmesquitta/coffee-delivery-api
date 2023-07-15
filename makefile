.PHONY: default dev build tidy test docs clean

# Variables
APP_NAME=coffee-delivery-api

# Tasks
default: dev

dev:
	@air
build:
	@go build -o $(APP_NAME) main.go
tidy:
	@go mod tidy
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
