package main

import (
	// "github.com/SyntinelNyx/syntinel-agent/internal/grpc"
	"github.com/SyntinelNyx/syntinel-agent/internal/trivy"
)

func main() {
	// client := grpc.InitConnectToServer()

	// grpc.StartBidirectionalStream(client)

	trivy.DeepScan("/")

	// grpc.heartbeat(client)
}
