// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
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
	ServiceInterface_Ping_FullMethodName         = "/service.v1alpha1.ServiceInterface/Ping"
	ServiceInterface_List_FullMethodName         = "/service.v1alpha1.ServiceInterface/List"
	ServiceInterface_Save_FullMethodName         = "/service.v1alpha1.ServiceInterface/Save"
	ServiceInterface_Get_FullMethodName          = "/service.v1alpha1.ServiceInterface/Get"
	ServiceInterface_Delete_FullMethodName       = "/service.v1alpha1.ServiceInterface/Delete"
	ServiceInterface_GetWorkflow_FullMethodName  = "/service.v1alpha1.ServiceInterface/GetWorkflow"
	ServiceInterface_SaveWorkflow_FullMethodName = "/service.v1alpha1.ServiceInterface/SaveWorkflow"
)

// ServiceInterfaceClient is the client API for ServiceInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceInterfaceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error)
	List(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Services, error)
	Save(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Msg, error)
	Get(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Service, error)
	Delete(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Msg, error)
	GetWorkflow(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Worklfow, error)
	SaveWorkflow(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Msg, error)
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

func (c *serviceInterfaceClient) List(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Services, error) {
	out := new(Services)
	err := c.cc.Invoke(ctx, ServiceInterface_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) Save(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ServiceInterface_Save_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) Get(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, ServiceInterface_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) Delete(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ServiceInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) GetWorkflow(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Worklfow, error) {
	out := new(Worklfow)
	err := c.cc.Invoke(ctx, ServiceInterface_GetWorkflow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceInterfaceClient) SaveWorkflow(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ServiceInterface_SaveWorkflow_FullMethodName, in, out, opts...)
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
	List(context.Context, *ServiceRequest) (*Services, error)
	Save(context.Context, *Service) (*Msg, error)
	Get(context.Context, *ServiceRequest) (*Service, error)
	Delete(context.Context, *ServiceRequest) (*Msg, error)
	GetWorkflow(context.Context, *ServiceRequest) (*Worklfow, error)
	SaveWorkflow(context.Context, *ServiceRequest) (*Msg, error)
	mustEmbedUnimplementedServiceInterfaceServer()
}

// UnimplementedServiceInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceInterfaceServer struct {
}

func (UnimplementedServiceInterfaceServer) Ping(context.Context, *emptypb.Empty) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedServiceInterfaceServer) List(context.Context, *ServiceRequest) (*Services, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedServiceInterfaceServer) Save(context.Context, *Service) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedServiceInterfaceServer) Get(context.Context, *ServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServiceInterfaceServer) Delete(context.Context, *ServiceRequest) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedServiceInterfaceServer) GetWorkflow(context.Context, *ServiceRequest) (*Worklfow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflow not implemented")
}
func (UnimplementedServiceInterfaceServer) SaveWorkflow(context.Context, *ServiceRequest) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveWorkflow not implemented")
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

func _ServiceInterface_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).List(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).Save(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).Get(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).Delete(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_GetWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).GetWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_GetWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).GetWorkflow(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceInterface_SaveWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInterfaceServer).SaveWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceInterface_SaveWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInterfaceServer).SaveWorkflow(ctx, req.(*ServiceRequest))
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
		{
			MethodName: "List",
			Handler:    _ServiceInterface_List_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _ServiceInterface_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServiceInterface_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ServiceInterface_Delete_Handler,
		},
		{
			MethodName: "GetWorkflow",
			Handler:    _ServiceInterface_GetWorkflow_Handler,
		},
		{
			MethodName: "SaveWorkflow",
			Handler:    _ServiceInterface_SaveWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service/v1alpha1/service.proto",
}
