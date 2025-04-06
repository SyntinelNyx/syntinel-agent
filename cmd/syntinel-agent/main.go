package main

import (
	"github.com/SyntinelNyx/syntinel-agent/internal/grpc"
)

func main() {
	client := grpc.InitConnectToServer()

	grpc.StartBidirectionalStream(client)

	grpc.SendTrivyScan(client)

	// grpc.heartbeat(client)
}
