package sysinfo

import (
    "context"
    "net"
    "testing"

    "github.com/stretchr/testify/assert"
    "google.golang.org/grpc"
    "google.golang.org/grpc/test/bufconn"

	pb "github.com/SyntinelNyx/syntinel-agent/internal/proto"
)

func TestSysInfo(t *testing.T) {
    info := SysInfo()
    assert.NotEmpty(t, info, "SysInfo should return non-empty JSON string")
}

func TestConnectToServer(t *testing.T) {
    // Create a listener for the mock server
    listener := bufconn.Listen(1024 * 1024)
    server := grpc.NewServer()
    pb.RegisterHardwareServiceServer(server, &mockHardwareServiceServer{})

    errChan := make(chan error, 1)
    go func() {
        errChan <- server.Serve(listener)
    }()
    defer func() {
        if err := <-errChan; err != nil {
            t.Fatalf("Server exited with error: %v", err)
        }
    }()
    defer server.Stop()

    // Create a client connection to the mock server
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer(listener)), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()

    // Call the function to test
    ConnectToServer(conn)
}

func bufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
    return func(context.Context, string) (net.Conn, error) {
        return listener.Dial()
    }
}

type mockHardwareServiceServer struct {
    pb.UnimplementedHardwareServiceServer
}

func (s *mockHardwareServiceServer) SendHardwareInfo(ctx context.Context, req *pb.HardwareInfo) (*pb.Response, error) {
        return &pb.Response{Message: "Success"}, nil
    }