package grpc

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/SyntinelNyx/syntinel-agent/internal/sysinfo"
	"github.com/SyntinelNyx/syntinel-agent/internal/data"
	pb "github.com/SyntinelNyx/syntinel-agent/internal/proto"
)

func ConnectToServer() (pb.HardwareServiceClient, context.Context) {
	// Create tls based credential
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "api.syntinel.dev")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHardwareServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client, ctx
}

func SendHardwareInfo() {
	client, ctx := ConnectToServer()
	
	// Fetch hardware info
	hardwareInfo := sysinfo.SysInfo()

	// Send to server
	resp, err := client.SendHardwareInfo(ctx, &pb.HardwareInfo{JsonData: hardwareInfo})
	if err != nil {
		log.Fatalf("Error calling SendHardwareInfo: %v", err)
	}

	slog.Info(fmt.Sprintf("Response from server: %s", resp.Message))
}