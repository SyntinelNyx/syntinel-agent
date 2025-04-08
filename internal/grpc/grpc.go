package grpc

import (
	"net"

	"google.golang.org/grpc"

	"github.com/SyntinelNyx/syntinel-agent/internal/grpc/control"
	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	controlpb "github.com/SyntinelNyx/syntinel-agent/internal/proto/controlpb"
)

func Start(server *grpc.Server) *grpc.Server {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatal("Failed to listen: %v", err)
	}

	controlpb.RegisterAgentServiceServer(server, &control.Agent{})
	logger.Info("Agent is listening on :50051 with TLS")

	if err := server.Serve(listener); err != nil {
		logger.Fatal("Failed to serve: %v", err)
	}

	return server
}
