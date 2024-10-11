syntinel-agent:
	@go build ./cmd/syntinel-agent
run:
	@go run ./cmd/syntinel-agent
test:
	@go test ./...
clean:
	@rm ./syntinel-agent