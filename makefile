format:
	@go fmt ./...

run:
	@go run cmd/server/main.go

build:
	@go build -o rill cmd/server/main.go