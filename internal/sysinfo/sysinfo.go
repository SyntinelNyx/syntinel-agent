package sysinfo

import (
	"context"
	"encoding/json"
	"os/user"
	"time"

	"github.com/zcalusic/sysinfo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/SyntinelNyx/syntinel-agent/internal/data"
	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	pb "github.com/SyntinelNyx/syntinel-agent/internal/proto"
)

func SysInfo() string {
	current, err := user.Current()
	if err != nil {
		logger.Fatal("error getting current user: %v", err)
	}

	if current.Uid != "0" {
		logger.Fatal("requires superuser privilege")
	}

	var si sysinfo.SysInfo
	si.GetSysInfo()

	data, err := json.MarshalIndent(&si, "", "  ")
	if err != nil {
		logger.Fatal("Error marshaling hardware info to JSON: %v", err)
	}

	return string(data)
}

func ConnectToServer(conn *grpc.ClientConn) {
	var err error
	if conn == nil {
		creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "api.syntinel.dev")
		if err != nil {
			logger.Fatal("failed to load credentials: %v", err)
		}

		conn, err = grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
		if err != nil {
			logger.Fatal("Failed to connect: %v", err)
		}
		defer conn.Close()
	}

	client := pb.NewHardwareServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	hardwareInfo := SysInfo()

	resp, err := client.SendHardwareInfo(ctx, &pb.HardwareInfo{JsonData: hardwareInfo})
	if err != nil {
		logger.Fatal("Error calling SendHardwareInfo: %v", err)
	}

	logger.Info("Response from server: %s", resp.Message)
}
