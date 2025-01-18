build:
	@go build -o ./bin/fs ./cmd

run: build
	@./bin/fs

test:
	@go test ./... -v
