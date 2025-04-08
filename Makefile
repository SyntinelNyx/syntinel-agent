id ?= dev

syntinel-agent:
	@go clean && go build -o syntinel-agent-$(id) ./cmd/syntinel-agent 
run:
	@go run ./cmd/syntinel-agent
test:
	@go test ./... -v
test_coverage:
	@go test -cover ./...
clean:
	@rm ./syntinel-agent
proto:
	protoc --go_out=. --go-grpc_out=. internal/proto/control.proto
