package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/SyntinelNyx/syntinel-agent/internal/grpc"
	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading .env file")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	creds, err := credentials.NewServerTLSFromFile(os.Getenv("TLS_CERT_PATH"), os.Getenv("TLS_KEY_PATH"))
	if err != nil {
		logger.Fatal("failed to create credentials: %v", err)
	}

	grpcServer := ggrpc.NewServer(ggrpc.Creds(creds))
	go func() {
		grpc.StartServer(grpcServer)
		fmt.Println("gRPC server started")
	}()

	<-stop
	logger.Info("Shutting down gracefully...")
	
	grpcServer.GracefulStop()

	logger.Info("Shutdown complete.")
}
