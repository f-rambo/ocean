// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: api/project/v1alpha1/project.proto

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
	ProjectService_Ping_FullMethodName               = "/project.v1alpha1.ProjectService/Ping"
	ProjectService_Save_FullMethodName               = "/project.v1alpha1.ProjectService/Save"
	ProjectService_Get_FullMethodName                = "/project.v1alpha1.ProjectService/Get"
	ProjectService_List_FullMethodName               = "/project.v1alpha1.ProjectService/List"
	ProjectService_Delete_FullMethodName             = "/project.v1alpha1.ProjectService/Delete"
	ProjectService_GetProjectMockData_FullMethodName = "/project.v1alpha1.ProjectService/GetProjectMockData"
	ProjectService_Enable_FullMethodName             = "/project.v1alpha1.ProjectService/Enable"
	ProjectService_Disable_FullMethodName            = "/project.v1alpha1.ProjectService/Disable"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error)
	Save(ctx context.Context, in *Project, opts ...grpc.CallOption) (*Msg, error)
	Get(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*Project, error)
	List(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*ProjectList, error)
	Delete(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*Msg, error)
	GetProjectMockData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Project, error)
	Enable(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*Msg, error)
	Disable(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*Msg, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ProjectService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Save(ctx context.Context, in *Project, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ProjectService_Save_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Get(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, ProjectService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) List(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*ProjectList, error) {
	out := new(ProjectList)
	err := c.cc.Invoke(ctx, ProjectService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Delete(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ProjectService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProjectMockData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, ProjectService_GetProjectMockData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Enable(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ProjectService_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Disable(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ProjectService_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	Ping(context.Context, *emptypb.Empty) (*Msg, error)
	Save(context.Context, *Project) (*Msg, error)
	Get(context.Context, *ProjectReq) (*Project, error)
	List(context.Context, *ProjectReq) (*ProjectList, error)
	Delete(context.Context, *ProjectReq) (*Msg, error)
	GetProjectMockData(context.Context, *emptypb.Empty) (*Project, error)
	Enable(context.Context, *ProjectReq) (*Msg, error)
	Disable(context.Context, *ProjectReq) (*Msg, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) Ping(context.Context, *emptypb.Empty) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedProjectServiceServer) Save(context.Context, *Project) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedProjectServiceServer) Get(context.Context, *ProjectReq) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProjectServiceServer) List(context.Context, *ProjectReq) (*ProjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProjectServiceServer) Delete(context.Context, *ProjectReq) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProjectServiceServer) GetProjectMockData(context.Context, *emptypb.Empty) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectMockData not implemented")
}
func (UnimplementedProjectServiceServer) Enable(context.Context, *ProjectReq) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedProjectServiceServer) Disable(context.Context, *ProjectReq) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Project)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Save(ctx, req.(*Project))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Get(ctx, req.(*ProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).List(ctx, req.(*ProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Delete(ctx, req.(*ProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProjectMockData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProjectMockData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetProjectMockData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProjectMockData(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Enable(ctx, req.(*ProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Disable(ctx, req.(*ProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.v1alpha1.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ProjectService_Ping_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _ProjectService_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProjectService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProjectService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProjectService_Delete_Handler,
		},
		{
			MethodName: "GetProjectMockData",
			Handler:    _ProjectService_GetProjectMockData_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _ProjectService_Enable_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _ProjectService_Disable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/project/v1alpha1/project.proto",
}
