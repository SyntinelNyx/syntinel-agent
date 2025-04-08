package grpc

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/SyntinelNyx/syntinel-agent/internal/proto"
	"github.com/SyntinelNyx/syntinel-agent/internal/actions/shx"
	"github.com/SyntinelNyx/syntinel-agent/internal/actions/trivy"
)

type Agent struct {
	proto.UnimplementedAgentServiceServer
}

func (s *Agent) SendHardwareInfo(ctx context.Context, req *proto.HardwareInfoRequest) (*proto.HardwareInfoResponse, error) {
	logger.Info("Received hardware info: %s", req.JsonData)
	return &proto.HardwareInfoResponse{Message: "Hardware info received successfully"}, nil
}

func StartServer(grpcServer *grpc.Server) *grpc.Server {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatal("Failed to listen: %v", err)
	}

	proto.RegisterAgentServiceServer(grpcServer, &Agent{})

	logger.Info("gRPC server listening on :50051 with TLS...")
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("Failed to serve: %v", err)
	}

	return grpcServer
}

func (a *Agent) BidirectionalStream(stream proto.AgentService_BidirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Recv error: %v", err)
			return err
		}

		log.Printf("Received: action=%s ", req.Name)

		var result string
		switch req.Message {
		case "DeepScan":
			trivy.DeepScan(req.Path)
			result = "pong"
		case "echo":
			result = req.Payload
		default:
			result = "unknown command"
		}

		err = stream.Send(&proto.ActionsResponse{
			Name: req.GetName(),
			Status: "Script Received and Saved",
		})
		if err != nil {
			log.Printf("Send error: %v", err)
			return err
		}
	}
}

// func StartBidirectionalStream(client proto.AgentServiceClient) {
// 	ctx := context.Background()
// 	for {
// 		// Create a stream
// 		stream, err := client.BidirectionalStream(ctx)
// 		if err != nil {
// 			log.Fatalf("Error creating bidirectional stream: %v", err)
// 		}
// 		// defer stream.CloseSend()

// 		// Goroutine to handle incoming messages from the server
// 		go func() {
// 			for {
// 				// Receive a script from the server
// 				req, err := stream.Recv()
// 				if err == io.EOF {
// 					log.Println("Stream closed by server")
// 					break
// 				}
// 				if err != nil {
// 					log.Printf("Error receiving message: %v", err)
// 					continue
// 				}

// 				// Save the script locally
// 				scriptPath := "data/received_files/" + req.GetName()
// 				err = os.WriteFile(scriptPath, req.GetContent(), 0755)
// 				if err != nil {
// 					log.Printf("Error saving script %s: %v", scriptPath, err)
// 					continue
// 				}
// 				log.Printf("Script %s saved successfully", scriptPath)

// 				// Send a response back to the server
// 				err = stream.Send(&proto.ActionsResponse{
// 					Name:   req.GetName(),
// 					Status: "Script Received and Saved",
// 				})
// 				if err != nil {
// 					log.Printf("Error sending response for script %s: %v", req.GetName(), err)
// 				}

// 				// Execute the script
// 				output := shx.RunScript(scriptPath)

// 				// Send a response back to the server
// 				err = stream.Send(&proto.ScriptResponse{
// 					Name:   req.GetName(),
// 					Status: "Successfully ran",
// 					Output: output,
// 				})
// 				if err != nil {
// 					log.Printf("Error sending response for script execution %s: %v", scriptPath, err)
// 				}
// 				// defer stream.CloseSend()
// 			}
// 		}()
// 		// Keep the main loop alive to handle reconnections if needed
// 		// <-ctx.Done()
// 		log.Println("Context canceled, exiting bidirectional stream")
// 		return
// 	}
// }

// func SendTrivyScan(client proto.AgentServiceClient) {
// 	ctx := context.Background()
// 	for {
// 		// Create a stream
// 		stream, err := client.SendTrivyReport(ctx)
// 		if err != nil {
// 			log.Fatalf("Error creating trivy stream: %v", err)
// 		}
// 		// defer stream.CloseSend()

// 		go func() {
// 			for {
// 				// Receive request from the server
// 				req, err := stream.Recv()
// 				if err == io.EOF {
// 					log.Println("Stream closed by server")
// 					break
// 				}
// 				if err != nil {
// 					log.Printf("Error receiving message: %v", err)
// 					continue
// 				}

// 				var scanResult string

// 				if req.GetMessage() == "DeepScan" {

// 					// Perform a deep scan using Trivy
// 					scanResult = trivy.DeepScan(req.GetPath())

// 					log.Printf("Deep scan result: %s", scanResult)
// 				}

// 				// Send the scan result back to the server
// 				err = stream.Send(&proto.TrivyReportResponse{
// 					JsonData: scanResult,
// 					Status:   "Scan Complete",
// 				})
// 				if err != nil {
// 					log.Printf("Error sending scan result: %v", err)
// 				}
// 				log.Printf("Scan result sent successfully")
// 			}
// 		}()

// 		// Keep the main loop alive to handle reconnections if needed
// 		<-ctx.Done()
// 		log.Println("Context canceled, exiting bidirectional stream")
// 		return
// 	}
// }
