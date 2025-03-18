package grpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/SyntinelNyx/syntinel-agent/internal/data"
	"github.com/SyntinelNyx/syntinel-agent/internal/proto"
	pb "github.com/SyntinelNyx/syntinel-agent/internal/proto"
	"github.com/SyntinelNyx/syntinel-agent/internal/sysinfo"
)

func InitConnectToServer() pb.AgentServiceClient {
	// Create TLS-based credentials
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "api.syntinel.dev")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// Establish a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	// defer conn.Close()

	client := pb.NewAgentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Fetch hardware info
	hardwareInfo := sysinfo.SysInfo()

	// Send to server
	resp, err := client.SendHardwareInfo(ctx, &pb.HardwareInfo{JsonData: hardwareInfo})
	if err != nil {
		log.Fatalf("Error calling SendHardwareInfo: %v", err)
	}

	slog.Info(fmt.Sprintf("Response from server: %s", resp.Message))

    // Keep the connection open until the agent is terminated
    // go func() {
    //     <-context.Background().Done() // Wait for termination signal
    //     conn.Close()                 // Close the connection when exiting
    // }()

    return client
}

func StartBidirectionalStream(client proto.AgentServiceClient) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.BidirectionalStream(ctx)
	if err != nil {
		log.Fatalf("Error creating bidirectional stream: %v", err)
	}

    // directory to save received files
    err = os.MkdirAll("./data/received_files", 0755)
    if err != nil {
        log.Fatalf("Error creating directory: %v", err)
    }

	// Goroutine to handle receiving files from the server
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("Server closed the stream")
				break
			}
			if err != nil {
				log.Fatalf("Error receiving file: %v", err)
			}

			// Save the received file
			filePath := fmt.Sprintf("./data/received_files/%s", resp.Name)
			err = os.WriteFile(filePath, []byte(resp.Content), 0644)
			if err != nil {
				log.Fatalf("Error saving file: %v", err)
			}
			log.Printf("Received file: %s, Status: %s", resp.Name, resp.Status)
		}
	}()

    // Keep sending periodic heartbeats or acknowledgments to the server
    for {
        err = stream.Send(&proto.ReceiveScript{
            Name:    "Heartbeat",
            Content: []byte("Agent is alive and awaiting commands."),
        })
        if err != nil {
            log.Printf("Error sending heartbeat to server: %v", err)
            break
        }
        time.Sleep(10 * time.Second) // Adjust the interval as needed
    }

    // // Close the stream after communication
	// err = stream.CloseSend()
	// if err != nil {
	// 	log.Fatalf("Error closing stream: %v", err)
	// }
}

