package main

import (
	"crypto/tls"
	"os"
	"os/signal"
	"syscall"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/SyntinelNyx/syntinel-agent/internal/data"
	"github.com/SyntinelNyx/syntinel-agent/internal/grpc"
	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/SyntinelNyx/syntinel-agent/internal/setup"
)

func main() {
	logger.Info("Starting agent %s...", data.ID)
	setup.CheckCommands()
	setup.CreateDirectories()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	cert, err := tls.X509KeyPair(data.Cert, data.Key)
	if err != nil {
		logger.Fatal("Failed to load TLS key pair: %v", err)
	}

	creds := credentials.NewServerTLSFromCert(&cert)
	server := ggrpc.NewServer(ggrpc.Creds(creds))

	go func() {
		server = grpc.Start(server)
	}()

	<-stop

	logger.Info("Shutting down gracefully...")
	server.GracefulStop()
	logger.Info("Shutdown complete.")
}
