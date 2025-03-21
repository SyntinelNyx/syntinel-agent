syntinel-agent:
	@go build ./cmd/syntinel-agent
run:
	@go run ./cmd/syntinel-agent
test:
	@go test ./...
clean:
	@rm ./syntinel-agent
proto:
	protoc --go_out=./ --go_opt=paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    ./internal/proto/bidirectional_comm.proto

	protoc --go_out=./ --go_opt=paths=source_relative \
	--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
	./internal/proto/hardwareinfo.proto
