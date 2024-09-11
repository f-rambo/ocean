// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: api/system/v1alpha1/system.proto

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SystemInterface_Ping_FullMethodName      = "/system.v1alpha1.SystemInterface/Ping"
	SystemInterface_GetSystem_FullMethodName = "/system.v1alpha1.SystemInterface/GetSystem"
	SystemInterface_GetLogs_FullMethodName   = "/system.v1alpha1.SystemInterface/GetLogs"
)

// SystemInterfaceClient is the client API for SystemInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemInterfaceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error)
	GetSystem(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*System, error)
	GetLogs(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[LogRequest, LogResponse], error)
}

type systemInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemInterfaceClient(cc grpc.ClientConnInterface) SystemInterfaceClient {
	return &systemInterfaceClient{cc}
}

func (c *systemInterfaceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Msg)
	err := c.cc.Invoke(ctx, SystemInterface_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemInterfaceClient) GetSystem(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*System, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(System)
	err := c.cc.Invoke(ctx, SystemInterface_GetSystem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemInterfaceClient) GetLogs(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[LogRequest, LogResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SystemInterface_ServiceDesc.Streams[0], SystemInterface_GetLogs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[LogRequest, LogResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SystemInterface_GetLogsClient = grpc.BidiStreamingClient[LogRequest, LogResponse]

// SystemInterfaceServer is the server API for SystemInterface service.
// All implementations must embed UnimplementedSystemInterfaceServer
// for forward compatibility.
type SystemInterfaceServer interface {
	Ping(context.Context, *emptypb.Empty) (*Msg, error)
	GetSystem(context.Context, *emptypb.Empty) (*System, error)
	GetLogs(grpc.BidiStreamingServer[LogRequest, LogResponse]) error
	mustEmbedUnimplementedSystemInterfaceServer()
}

// UnimplementedSystemInterfaceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSystemInterfaceServer struct{}

func (UnimplementedSystemInterfaceServer) Ping(context.Context, *emptypb.Empty) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSystemInterfaceServer) GetSystem(context.Context, *emptypb.Empty) (*System, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystem not implemented")
}
func (UnimplementedSystemInterfaceServer) GetLogs(grpc.BidiStreamingServer[LogRequest, LogResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}
func (UnimplementedSystemInterfaceServer) mustEmbedUnimplementedSystemInterfaceServer() {}
func (UnimplementedSystemInterfaceServer) testEmbeddedByValue()                         {}

// UnsafeSystemInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemInterfaceServer will
// result in compilation errors.
type UnsafeSystemInterfaceServer interface {
	mustEmbedUnimplementedSystemInterfaceServer()
}

func RegisterSystemInterfaceServer(s grpc.ServiceRegistrar, srv SystemInterfaceServer) {
	// If the following call pancis, it indicates UnimplementedSystemInterfaceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SystemInterface_ServiceDesc, srv)
}

func _SystemInterface_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemInterfaceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemInterface_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemInterfaceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemInterface_GetSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemInterfaceServer).GetSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemInterface_GetSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemInterfaceServer).GetSystem(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemInterface_GetLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SystemInterfaceServer).GetLogs(&grpc.GenericServerStream[LogRequest, LogResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SystemInterface_GetLogsServer = grpc.BidiStreamingServer[LogRequest, LogResponse]

// SystemInterface_ServiceDesc is the grpc.ServiceDesc for SystemInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system.v1alpha1.SystemInterface",
	HandlerType: (*SystemInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SystemInterface_Ping_Handler,
		},
		{
			MethodName: "GetSystem",
			Handler:    _SystemInterface_GetSystem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLogs",
			Handler:       _SystemInterface_GetLogs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/system/v1alpha1/system.proto",
}
