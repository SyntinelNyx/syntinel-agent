package sysinfo

import (
	"context"
	"encoding/json"
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

    var result map[string]interface{}
    err := json.Unmarshal([]byte(info), &result)
    assert.NoError(t, err, "SysInfo should return valid JSON")

    _, ok := result["sysinfo"].(map[string]interface{})
    assert.True(t, ok, "SysInfo JSON should contain 'sysinfo' field")

    _, ok = result["node"].(map[string]interface{})
    assert.True(t, ok, "SysInfo JSON should contain 'node' field")

    _, ok = result["os"].(map[string]interface{})
    assert.True(t, ok, "SysInfo JSON should contain 'os' field")

    _, ok = result["kernel"].(map[string]interface{})
    assert.True(t, ok, "SysInfo JSON should contain 'kernel' field")

    _, ok = result["bios"].(map[string]interface{})
    assert.True(t, ok, "SysInfo JSON should contain 'bios' field")

    _, ok = result["cpu"].(map[string]interface{})
    assert.True(t, ok, "SysInfo JSON should contain 'cpu' field")

    _, ok = result["memory"].(map[string]interface{})
    assert.True(t, ok, "SysInfo JSON should contain 'memory' field")

    _, ok = result["storage"].([]interface{})
    assert.True(t, ok, "SysInfo JSON should contain 'storage' field")
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