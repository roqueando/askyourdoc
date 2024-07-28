.PHONY: build run

build:
	go build -o build/ayd-$(module) ./cmd/ayd-$(module)

run:
	@./build/ayd-$(module)

test:
	@go test ./... -v

