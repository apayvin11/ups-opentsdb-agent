.PHONY: build
build:
	go build -o bin/ups-agent cmd/main.go

.DEFAULT_GOAL :=build

.PHONY: test
test:
	go test -timeout 30s ./...