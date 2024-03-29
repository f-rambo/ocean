// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/service/v1alpha1/service.proto

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ServiceInterface_Ping_FullMethodName = "/service.v1alpha1.ServiceInterface/Ping"
)

// ServiceInterfaceClient is the client API for ServiceInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceInterfaceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error)
}

type serviceInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceInterfaceClient(cc grpc.ClientConnInterface) ServiceInterfaceClient {
	return &serviceInterfaceClient{cc}
}

func (c *serviceInterfaceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ServiceInterface_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceInterfaceServer is the server API for ServiceInterface service.
// All implementations must embed UnimplementedServiceInterfaceServer
// for forward compatibility
type ServiceInterfaceServer interface {
	Ping(context.Context, *emptypb.Empty) (*Msg, error)
	mustEmbedUnimplementedServiceInterfaceServer()
}

// UnimplementedServiceInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceInterfaceServer struct {
}

func (UnimplementedServiceInterfaceServer) Ping(context.Context, *emptypb.Empty) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedServiceInterfaceServer) mustEmbedUnimplementedServiceInterfaceServer() {}

// UnsafeServiceInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceInterfaceServer will
// result in compilation errors.
type UnsafeServiceInterfaceServer interface {
	mustEmbedUnimplementedServiceInterfaceServer()
}

func RegisterServiceInterfaceServer(s grpc.ServiceRegistrar, srv ServiceInterfaceServer) {
	s.RegisterService(&ServiceInterface_ServiceDesc, srv)
}

func _ServiceInterface_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceInterface_ServiceDesc is the grpc.ServiceDesc for ServiceInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.v1alpha1.ServiceInterface",
	HandlerType: (*ServiceInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ServiceInterface_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service/v1alpha1/service.proto",
}
