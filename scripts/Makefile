APP_NAME := damsharaz

.PHONY: build run clean tidy test

build:
	cd .. && mkdir -p bin && go build -o bin/$(APP_NAME) ./cmd

run:
	cd .. && go run ./cmd

clean:
	cd .. && rm -rf bin

tidy:
	cd .. && go mod tidy

test:
	cd .. && go test ./...
