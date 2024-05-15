.PHONY: build
build:
	go build -o bin/ups-agent cmd/main.go

.DEFAULT_GOAL :=build
