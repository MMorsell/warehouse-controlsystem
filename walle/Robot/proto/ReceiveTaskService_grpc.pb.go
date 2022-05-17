// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: Robot/proto/ReceiveTaskService.proto

package proto

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

// ReceiveTaskServiceClient is the client API for ReceiveTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReceiveTaskServiceClient interface {
	//function for Hive to call,
	ReceiveTask(ctx context.Context, in *Instructions, opts ...grpc.CallOption) (*HasReceivedTask, error)
}

type receiveTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReceiveTaskServiceClient(cc grpc.ClientConnInterface) ReceiveTaskServiceClient {
	return &receiveTaskServiceClient{cc}
}

func (c *receiveTaskServiceClient) ReceiveTask(ctx context.Context, in *Instructions, opts ...grpc.CallOption) (*HasReceivedTask, error) {
	out := new(HasReceivedTask)
	err := c.cc.Invoke(ctx, "/main.ReceiveTaskService/ReceiveTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiveTaskServiceServer is the server API for ReceiveTaskService service.
// All implementations must embed UnimplementedReceiveTaskServiceServer
// for forward compatibility
type ReceiveTaskServiceServer interface {
	//function for Hive to call,
	ReceiveTask(context.Context, *Instructions) (*HasReceivedTask, error)
	mustEmbedUnimplementedReceiveTaskServiceServer()
}

// UnimplementedReceiveTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReceiveTaskServiceServer struct {
}

func (UnimplementedReceiveTaskServiceServer) ReceiveTask(context.Context, *Instructions) (*HasReceivedTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveTask not implemented")
}
func (UnimplementedReceiveTaskServiceServer) mustEmbedUnimplementedReceiveTaskServiceServer() {}

// UnsafeReceiveTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReceiveTaskServiceServer will
// result in compilation errors.
type UnsafeReceiveTaskServiceServer interface {
	mustEmbedUnimplementedReceiveTaskServiceServer()
}

func RegisterReceiveTaskServiceServer(s grpc.ServiceRegistrar, srv ReceiveTaskServiceServer) {
	s.RegisterService(&ReceiveTaskService_ServiceDesc, srv)
}

func _ReceiveTaskService_ReceiveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instructions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiveTaskServiceServer).ReceiveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ReceiveTaskService/ReceiveTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiveTaskServiceServer).ReceiveTask(ctx, req.(*Instructions))
	}
	return interceptor(ctx, in, info, handler)
}

// ReceiveTaskService_ServiceDesc is the grpc.ServiceDesc for ReceiveTaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReceiveTaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ReceiveTaskService",
	HandlerType: (*ReceiveTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveTask",
			Handler:    _ReceiveTaskService_ReceiveTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Robot/proto/ReceiveTaskService.proto",
}