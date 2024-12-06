package sysinfo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os/user"
	"time"

	"github.com/zcalusic/sysinfo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/SyntinelNyx/syntinel-agent/internal/proto"
)

func SysInfo() string {
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal("Requires superuser privilege")
	}

	var si sysinfo.SysInfo

	si.GetSysInfo()

	// Marshal to JSON
	data, err := json.MarshalIndent(&si, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling hardware info to JSON: %v", err)
	}

	return string(data)
}

func ConnectToServer() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHardwareServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Fetch hardware info
	hardwareInfo := SysInfo()

	// Send to server
	resp, err := client.SendHardwareInfo(ctx, &pb.HardwareInfo{JsonData: hardwareInfo})
	if err != nil {
		log.Fatalf("Error calling SendHardwareInfo: %v", err)
	}

	slog.Info(fmt.Sprintf("Response from server: %s", resp.Message))
}
