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
	"github.com/SyntinelNyx/syntinel-agent/internal/shx"
	"github.com/SyntinelNyx/syntinel-agent/internal/sysinfo"
	"github.com/SyntinelNyx/syntinel-agent/internal/trivy"
)

func InitConnectToServer() pb.AgentServiceClient {
	// Create TLS-based credentials
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "api.syntinel.dev")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// Establish a connection to the server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	// defer conn.Close()

	client := pb.NewAgentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Directory to save received files
	err = os.MkdirAll("./data/received_files", 0755)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	// Fetch hardware info
	hardwareInfo := sysinfo.SysInfo()

	// Send to server
	resp, err := client.SendHardwareInfo(ctx, &pb.HardwareInfoRequest{JsonData: hardwareInfo})
	if err != nil {
		log.Fatalf("Error calling SendHardwareInfo: %v", err)
	}

	slog.Info(fmt.Sprintf("Response from server: %s", resp.Message))

	return client
}

func StartBidirectionalStream(client proto.AgentServiceClient) {
	ctx := context.Background()
	for {
		// Create a stream
		stream, err := client.BidirectionalStream(ctx)
		if err != nil {
			log.Fatalf("Error creating bidirectional stream: %v", err)
		}
		defer stream.CloseSend()

		// Goroutine to handle incoming messages from the server
		go func() {
			for {
				// Receive a script from the server
				req, err := stream.Recv()
				if err == io.EOF {
					log.Println("Stream closed by server")
					break
				}
				if err != nil {
					log.Printf("Error receiving message: %v", err)
					continue
				}

				// Save the script locally
				scriptPath := "data/received_files/" + req.GetName()
				err = os.WriteFile(scriptPath, req.GetContent(), 0755)
				if err != nil {
					log.Printf("Error saving script %s: %v", scriptPath, err)
					continue
				}
				log.Printf("Script %s saved successfully", scriptPath)

				// Send a response back to the server
				err = stream.Send(&proto.ScriptResponse{
					Name:   req.GetName(),
					Status: "Script Received and Saved",
				})
				if err != nil {
					log.Printf("Error sending response for script %s: %v", req.GetName(), err)
				}

				// Execute the script
				shx.RunScript(scriptPath)

			}
		}()
		// Keep the main loop alive to handle reconnections if needed
		<-ctx.Done()
		log.Println("Context canceled, exiting bidirectional stream")
		return
	}
}

func SendTrivyScan(client proto.AgentServiceClient) {
	ctx := context.Background()
	for {
		// Create a stream
		stream, err := client.SendTrivyReport(ctx)
		if err != nil {
			log.Fatalf("Error creating trivy stream: %v", err)
		}
		defer stream.CloseSend()

		go func() {
			for {
				// Receive request from the server
				req, err := stream.Recv()
				if err == io.EOF {
					log.Println("Stream closed by server")
					break
				}
				if err != nil {
					log.Printf("Error receiving message: %v", err)
					continue
				}

				var scanResult string

				if req.GetMessage() == "DeepScan" {

					// Perform a deep scan using Trivy
					scanResult := trivy.DeepScan(req.GetPath())
					if err != nil {
						log.Printf("Error during deep scan: %v", err)
						continue
					}
					log.Printf("Deep scan result: %s", scanResult)
				}

				// Send the scan result back to the server
				err = stream.Send(&proto.TrivyReportResponse{
					JsonData: scanResult,
					Status:  "Scan Complete",
				})
				if err != nil {
					log.Printf("Error sending scan result: %v", err)
				}
				log.Printf("Scan result sent successfully")
			}
		}()

		// Keep the main loop alive to handle reconnections if needed
		<-ctx.Done()
		log.Println("Context canceled, exiting bidirectional stream")
		return
	}
}
