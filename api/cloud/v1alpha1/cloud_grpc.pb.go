// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: api/cloud/v1alpha1/cloud.proto

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
	CloudInterface_Ping_FullMethodName = "/cloud.v1alpha1.CloudInterface/Ping"
)

// CloudInterfaceClient is the client API for CloudInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudInterfaceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error)
}

type cloudInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudInterfaceClient(cc grpc.ClientConnInterface) CloudInterfaceClient {
	return &cloudInterfaceClient{cc}
}

func (c *cloudInterfaceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Msg)
	err := c.cc.Invoke(ctx, CloudInterface_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudInterfaceServer is the server API for CloudInterface service.
// All implementations must embed UnimplementedCloudInterfaceServer
// for forward compatibility.
type CloudInterfaceServer interface {
	Ping(context.Context, *emptypb.Empty) (*Msg, error)
	mustEmbedUnimplementedCloudInterfaceServer()
}

// UnimplementedCloudInterfaceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCloudInterfaceServer struct{}

func (UnimplementedCloudInterfaceServer) Ping(context.Context, *emptypb.Empty) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCloudInterfaceServer) mustEmbedUnimplementedCloudInterfaceServer() {}
func (UnimplementedCloudInterfaceServer) testEmbeddedByValue()                        {}

// UnsafeCloudInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudInterfaceServer will
// result in compilation errors.
type UnsafeCloudInterfaceServer interface {
	mustEmbedUnimplementedCloudInterfaceServer()
}

func RegisterCloudInterfaceServer(s grpc.ServiceRegistrar, srv CloudInterfaceServer) {
	// If the following call pancis, it indicates UnimplementedCloudInterfaceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CloudInterface_ServiceDesc, srv)
}

func _CloudInterface_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudInterfaceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudInterface_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudInterfaceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudInterface_ServiceDesc is the grpc.ServiceDesc for CloudInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.v1alpha1.CloudInterface",
	HandlerType: (*CloudInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CloudInterface_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/cloud/v1alpha1/cloud.proto",
}
