.DEFAULT_GOAL := build

clean:
	go clean

fmt: clean
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build