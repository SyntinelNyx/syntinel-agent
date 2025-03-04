syntinel-agent:
	@go build ./cmd/syntinel-agent
run:
	@go run ./cmd/syntinel-agent
test:
	@go test ./... -v
clean:
	@rm ./syntinel-agent
proto:
	protoc --go_out=./ --go_opt=paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    ./internal/proto/hardware_info.proto
