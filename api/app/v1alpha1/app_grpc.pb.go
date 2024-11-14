// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: api/app/v1alpha1/app.proto

package v1alpha1

import (
	context "context"
	common "github.com/f-rambo/ocean/api/common"
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
	AppInterface_Ping_FullMethodName                   = "/app.v1alpha1.AppInterface/Ping"
	AppInterface_UploadApp_FullMethodName              = "/app.v1alpha1.AppInterface/UploadApp"
	AppInterface_Save_FullMethodName                   = "/app.v1alpha1.AppInterface/Save"
	AppInterface_Get_FullMethodName                    = "/app.v1alpha1.AppInterface/Get"
	AppInterface_List_FullMethodName                   = "/app.v1alpha1.AppInterface/List"
	AppInterface_Delete_FullMethodName                 = "/app.v1alpha1.AppInterface/Delete"
	AppInterface_CreateAppType_FullMethodName          = "/app.v1alpha1.AppInterface/CreateAppType"
	AppInterface_ListAppType_FullMethodName            = "/app.v1alpha1.AppInterface/ListAppType"
	AppInterface_DeleteAppType_FullMethodName          = "/app.v1alpha1.AppInterface/DeleteAppType"
	AppInterface_SaveRepo_FullMethodName               = "/app.v1alpha1.AppInterface/SaveRepo"
	AppInterface_ListRepo_FullMethodName               = "/app.v1alpha1.AppInterface/ListRepo"
	AppInterface_DeleteRepo_FullMethodName             = "/app.v1alpha1.AppInterface/DeleteRepo"
	AppInterface_GetAppsByRepo_FullMethodName          = "/app.v1alpha1.AppInterface/GetAppsByRepo"
	AppInterface_GetAppDetailByRepo_FullMethodName     = "/app.v1alpha1.AppInterface/GetAppDetailByRepo"
	AppInterface_GetAppRelease_FullMethodName          = "/app.v1alpha1.AppInterface/GetAppRelease"
	AppInterface_AppReleaseList_FullMethodName         = "/app.v1alpha1.AppInterface/AppReleaseList"
	AppInterface_GetAppReleaseResources_FullMethodName = "/app.v1alpha1.AppInterface/GetAppReleaseResources"
	AppInterface_SaveAppRelease_FullMethodName         = "/app.v1alpha1.AppInterface/SaveAppRelease"
	AppInterface_DeleteAppRelease_FullMethodName       = "/app.v1alpha1.AppInterface/DeleteAppRelease"
)

// AppInterfaceClient is the client API for AppInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppInterfaceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.Msg, error)
	UploadApp(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*App, error)
	Save(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Msg, error)
	Get(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*App, error)
	List(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppList, error)
	Delete(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*common.Msg, error)
	CreateAppType(ctx context.Context, in *AppType, opts ...grpc.CallOption) (*common.Msg, error)
	ListAppType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AppTypeList, error)
	DeleteAppType(ctx context.Context, in *AppTypeReq, opts ...grpc.CallOption) (*common.Msg, error)
	SaveRepo(ctx context.Context, in *AppRepo, opts ...grpc.CallOption) (*common.Msg, error)
	ListRepo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AppRepoList, error)
	DeleteRepo(ctx context.Context, in *AppRepoReq, opts ...grpc.CallOption) (*common.Msg, error)
	GetAppsByRepo(ctx context.Context, in *AppRepoReq, opts ...grpc.CallOption) (*AppList, error)
	GetAppDetailByRepo(ctx context.Context, in *AppRepoReq, opts ...grpc.CallOption) (*App, error)
	GetAppRelease(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*AppRelease, error)
	AppReleaseList(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*AppReleaseList, error)
	GetAppReleaseResources(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*AppReleasepResources, error)
	SaveAppRelease(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*AppRelease, error)
	DeleteAppRelease(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*common.Msg, error)
}

type appInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppInterfaceClient(cc grpc.ClientConnInterface) AppInterfaceClient {
	return &appInterfaceClient{cc}
}

func (c *appInterfaceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Msg)
	err := c.cc.Invoke(ctx, AppInterface_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) UploadApp(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*App, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(App)
	err := c.cc.Invoke(ctx, AppInterface_UploadApp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) Save(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Msg)
	err := c.cc.Invoke(ctx, AppInterface_Save_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) Get(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*App, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(App)
	err := c.cc.Invoke(ctx, AppInterface_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) List(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*AppList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppList)
	err := c.cc.Invoke(ctx, AppInterface_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) Delete(ctx context.Context, in *AppReq, opts ...grpc.CallOption) (*common.Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Msg)
	err := c.cc.Invoke(ctx, AppInterface_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) CreateAppType(ctx context.Context, in *AppType, opts ...grpc.CallOption) (*common.Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Msg)
	err := c.cc.Invoke(ctx, AppInterface_CreateAppType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) ListAppType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AppTypeList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppTypeList)
	err := c.cc.Invoke(ctx, AppInterface_ListAppType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) DeleteAppType(ctx context.Context, in *AppTypeReq, opts ...grpc.CallOption) (*common.Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Msg)
	err := c.cc.Invoke(ctx, AppInterface_DeleteAppType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) SaveRepo(ctx context.Context, in *AppRepo, opts ...grpc.CallOption) (*common.Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Msg)
	err := c.cc.Invoke(ctx, AppInterface_SaveRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) ListRepo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AppRepoList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppRepoList)
	err := c.cc.Invoke(ctx, AppInterface_ListRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) DeleteRepo(ctx context.Context, in *AppRepoReq, opts ...grpc.CallOption) (*common.Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Msg)
	err := c.cc.Invoke(ctx, AppInterface_DeleteRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) GetAppsByRepo(ctx context.Context, in *AppRepoReq, opts ...grpc.CallOption) (*AppList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppList)
	err := c.cc.Invoke(ctx, AppInterface_GetAppsByRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) GetAppDetailByRepo(ctx context.Context, in *AppRepoReq, opts ...grpc.CallOption) (*App, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(App)
	err := c.cc.Invoke(ctx, AppInterface_GetAppDetailByRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) GetAppRelease(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*AppRelease, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppRelease)
	err := c.cc.Invoke(ctx, AppInterface_GetAppRelease_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) AppReleaseList(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*AppReleaseList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppReleaseList)
	err := c.cc.Invoke(ctx, AppInterface_AppReleaseList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) GetAppReleaseResources(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*AppReleasepResources, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppReleasepResources)
	err := c.cc.Invoke(ctx, AppInterface_GetAppReleaseResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) SaveAppRelease(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*AppRelease, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppRelease)
	err := c.cc.Invoke(ctx, AppInterface_SaveAppRelease_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInterfaceClient) DeleteAppRelease(ctx context.Context, in *AppReleaseReq, opts ...grpc.CallOption) (*common.Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Msg)
	err := c.cc.Invoke(ctx, AppInterface_DeleteAppRelease_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppInterfaceServer is the server API for AppInterface service.
// All implementations must embed UnimplementedAppInterfaceServer
// for forward compatibility.
type AppInterfaceServer interface {
	Ping(context.Context, *emptypb.Empty) (*common.Msg, error)
	UploadApp(context.Context, *FileUploadRequest) (*App, error)
	Save(context.Context, *App) (*common.Msg, error)
	Get(context.Context, *AppReq) (*App, error)
	List(context.Context, *AppReq) (*AppList, error)
	Delete(context.Context, *AppReq) (*common.Msg, error)
	CreateAppType(context.Context, *AppType) (*common.Msg, error)
	ListAppType(context.Context, *emptypb.Empty) (*AppTypeList, error)
	DeleteAppType(context.Context, *AppTypeReq) (*common.Msg, error)
	SaveRepo(context.Context, *AppRepo) (*common.Msg, error)
	ListRepo(context.Context, *emptypb.Empty) (*AppRepoList, error)
	DeleteRepo(context.Context, *AppRepoReq) (*common.Msg, error)
	GetAppsByRepo(context.Context, *AppRepoReq) (*AppList, error)
	GetAppDetailByRepo(context.Context, *AppRepoReq) (*App, error)
	GetAppRelease(context.Context, *AppReleaseReq) (*AppRelease, error)
	AppReleaseList(context.Context, *AppReleaseReq) (*AppReleaseList, error)
	GetAppReleaseResources(context.Context, *AppReleaseReq) (*AppReleasepResources, error)
	SaveAppRelease(context.Context, *AppReleaseReq) (*AppRelease, error)
	DeleteAppRelease(context.Context, *AppReleaseReq) (*common.Msg, error)
	mustEmbedUnimplementedAppInterfaceServer()
}

// UnimplementedAppInterfaceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAppInterfaceServer struct{}

func (UnimplementedAppInterfaceServer) Ping(context.Context, *emptypb.Empty) (*common.Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAppInterfaceServer) UploadApp(context.Context, *FileUploadRequest) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadApp not implemented")
}
func (UnimplementedAppInterfaceServer) Save(context.Context, *App) (*common.Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedAppInterfaceServer) Get(context.Context, *AppReq) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAppInterfaceServer) List(context.Context, *AppReq) (*AppList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAppInterfaceServer) Delete(context.Context, *AppReq) (*common.Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAppInterfaceServer) CreateAppType(context.Context, *AppType) (*common.Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppType not implemented")
}
func (UnimplementedAppInterfaceServer) ListAppType(context.Context, *emptypb.Empty) (*AppTypeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppType not implemented")
}
func (UnimplementedAppInterfaceServer) DeleteAppType(context.Context, *AppTypeReq) (*common.Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppType not implemented")
}
func (UnimplementedAppInterfaceServer) SaveRepo(context.Context, *AppRepo) (*common.Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveRepo not implemented")
}
func (UnimplementedAppInterfaceServer) ListRepo(context.Context, *emptypb.Empty) (*AppRepoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepo not implemented")
}
func (UnimplementedAppInterfaceServer) DeleteRepo(context.Context, *AppRepoReq) (*common.Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepo not implemented")
}
func (UnimplementedAppInterfaceServer) GetAppsByRepo(context.Context, *AppRepoReq) (*AppList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppsByRepo not implemented")
}
func (UnimplementedAppInterfaceServer) GetAppDetailByRepo(context.Context, *AppRepoReq) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppDetailByRepo not implemented")
}
func (UnimplementedAppInterfaceServer) GetAppRelease(context.Context, *AppReleaseReq) (*AppRelease, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRelease not implemented")
}
func (UnimplementedAppInterfaceServer) AppReleaseList(context.Context, *AppReleaseReq) (*AppReleaseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppReleaseList not implemented")
}
func (UnimplementedAppInterfaceServer) GetAppReleaseResources(context.Context, *AppReleaseReq) (*AppReleasepResources, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppReleaseResources not implemented")
}
func (UnimplementedAppInterfaceServer) SaveAppRelease(context.Context, *AppReleaseReq) (*AppRelease, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveAppRelease not implemented")
}
func (UnimplementedAppInterfaceServer) DeleteAppRelease(context.Context, *AppReleaseReq) (*common.Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppRelease not implemented")
}
func (UnimplementedAppInterfaceServer) mustEmbedUnimplementedAppInterfaceServer() {}
func (UnimplementedAppInterfaceServer) testEmbeddedByValue()                      {}

// UnsafeAppInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppInterfaceServer will
// result in compilation errors.
type UnsafeAppInterfaceServer interface {
	mustEmbedUnimplementedAppInterfaceServer()
}

func RegisterAppInterfaceServer(s grpc.ServiceRegistrar, srv AppInterfaceServer) {
	// If the following call pancis, it indicates UnimplementedAppInterfaceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AppInterface_ServiceDesc, srv)
}

func _AppInterface_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_UploadApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).UploadApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_UploadApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).UploadApp(ctx, req.(*FileUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).Save(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).Get(ctx, req.(*AppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).List(ctx, req.(*AppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).Delete(ctx, req.(*AppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_CreateAppType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).CreateAppType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_CreateAppType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).CreateAppType(ctx, req.(*AppType))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_ListAppType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).ListAppType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_ListAppType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).ListAppType(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_DeleteAppType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).DeleteAppType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_DeleteAppType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).DeleteAppType(ctx, req.(*AppTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_SaveRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRepo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).SaveRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_SaveRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).SaveRepo(ctx, req.(*AppRepo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_ListRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).ListRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_ListRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).ListRepo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_DeleteRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).DeleteRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_DeleteRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).DeleteRepo(ctx, req.(*AppRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_GetAppsByRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).GetAppsByRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_GetAppsByRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).GetAppsByRepo(ctx, req.(*AppRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_GetAppDetailByRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRepoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).GetAppDetailByRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_GetAppDetailByRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).GetAppDetailByRepo(ctx, req.(*AppRepoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_GetAppRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).GetAppRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_GetAppRelease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).GetAppRelease(ctx, req.(*AppReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_AppReleaseList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).AppReleaseList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_AppReleaseList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).AppReleaseList(ctx, req.(*AppReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_GetAppReleaseResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).GetAppReleaseResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_GetAppReleaseResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).GetAppReleaseResources(ctx, req.(*AppReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_SaveAppRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).SaveAppRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_SaveAppRelease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).SaveAppRelease(ctx, req.(*AppReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInterface_DeleteAppRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInterfaceServer).DeleteAppRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppInterface_DeleteAppRelease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInterfaceServer).DeleteAppRelease(ctx, req.(*AppReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AppInterface_ServiceDesc is the grpc.ServiceDesc for AppInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.v1alpha1.AppInterface",
	HandlerType: (*AppInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AppInterface_Ping_Handler,
		},
		{
			MethodName: "UploadApp",
			Handler:    _AppInterface_UploadApp_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _AppInterface_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AppInterface_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AppInterface_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AppInterface_Delete_Handler,
		},
		{
			MethodName: "CreateAppType",
			Handler:    _AppInterface_CreateAppType_Handler,
		},
		{
			MethodName: "ListAppType",
			Handler:    _AppInterface_ListAppType_Handler,
		},
		{
			MethodName: "DeleteAppType",
			Handler:    _AppInterface_DeleteAppType_Handler,
		},
		{
			MethodName: "SaveRepo",
			Handler:    _AppInterface_SaveRepo_Handler,
		},
		{
			MethodName: "ListRepo",
			Handler:    _AppInterface_ListRepo_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _AppInterface_DeleteRepo_Handler,
		},
		{
			MethodName: "GetAppsByRepo",
			Handler:    _AppInterface_GetAppsByRepo_Handler,
		},
		{
			MethodName: "GetAppDetailByRepo",
			Handler:    _AppInterface_GetAppDetailByRepo_Handler,
		},
		{
			MethodName: "GetAppRelease",
			Handler:    _AppInterface_GetAppRelease_Handler,
		},
		{
			MethodName: "AppReleaseList",
			Handler:    _AppInterface_AppReleaseList_Handler,
		},
		{
			MethodName: "GetAppReleaseResources",
			Handler:    _AppInterface_GetAppReleaseResources_Handler,
		},
		{
			MethodName: "SaveAppRelease",
			Handler:    _AppInterface_SaveAppRelease_Handler,
		},
		{
			MethodName: "DeleteAppRelease",
			Handler:    _AppInterface_DeleteAppRelease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/app/v1alpha1/app.proto",
}
