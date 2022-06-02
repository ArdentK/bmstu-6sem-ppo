.PHONY: build
build :
		go build -v ./cmd/apiserver

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

cover:
		go test ./... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

.DEFAULT_GOAL := build