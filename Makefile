.PHONY: default test build

OS := $(shell uname)
VERSION ?= 1.0.0
APPNAME := ecommerce

# target #

default: unit-test integration-test build run

build:
	mkdir -p bin
	@echo "Setup ecommerce"
ifeq ($(OS), Linux)
	@echo "Build ecommerce..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o ./bin/$(APPNAME) main.go
else
	@echo "Build $(OS)"
	go build -ldflags "-s -w -X main.Version=$(VERSION)" -o ./bin/$(APPNAME) main.go
endif
ifeq ($(OS) ,Darwin)
	@echo "Build ecommerce..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/$(APPNAME) main.go
endif
ifeq ($(OS),Windows_NT)
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(APPNAME).exe main.go
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

# Test Packages

unit-test:
	@go test -count=1 -v --cover ./... -tags="unit"

integration-test:
 	@go test -count=1 -v --cover -tags="integration" -p 1 ./... --env-path=.env

run:
	./bin/ecommerce