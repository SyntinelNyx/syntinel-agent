// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.5
// source: internal/proto/control.proto

package controlpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AgentService_Control_FullMethodName = "/control.AgentService/Control"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	Control(ctx context.Context, opts ...grpc.CallOption) (AgentService_ControlClient, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) Control(ctx context.Context, opts ...grpc.CallOption) (AgentService_ControlClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[0], AgentService_Control_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceControlClient{stream}
	return x, nil
}

type AgentService_ControlClient interface {
	Send(*ControlMessage) error
	Recv() (*ControlResponse, error)
	grpc.ClientStream
}

type agentServiceControlClient struct {
	grpc.ClientStream
}

func (x *agentServiceControlClient) Send(m *ControlMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentServiceControlClient) Recv() (*ControlResponse, error) {
	m := new(ControlResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	Control(AgentService_ControlServer) error
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) Control(AgentService_ControlServer) error {
	return status.Errorf(codes.Unimplemented, "method Control not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_Control_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServiceServer).Control(&agentServiceControlServer{stream})
}

type AgentService_ControlServer interface {
	Send(*ControlResponse) error
	Recv() (*ControlMessage, error)
	grpc.ServerStream
}

type agentServiceControlServer struct {
	grpc.ServerStream
}

func (x *agentServiceControlServer) Send(m *ControlResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentServiceControlServer) Recv() (*ControlMessage, error) {
	m := new(ControlMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "control.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Control",
			Handler:       _AgentService_Control_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/proto/control.proto",
}
