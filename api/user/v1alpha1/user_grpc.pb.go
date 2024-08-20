// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: api/user/v1alpha1/user.proto

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
	UserInterface_SignIn_FullMethodName      = "/user.v1alpha1.UserInterface/SignIn"
	UserInterface_GetUserInfo_FullMethodName = "/user.v1alpha1.UserInterface/GetUserInfo"
	UserInterface_GetUsers_FullMethodName    = "/user.v1alpha1.UserInterface/GetUsers"
	UserInterface_SaveUser_FullMethodName    = "/user.v1alpha1.UserInterface/SaveUser"
	UserInterface_DeleteUser_FullMethodName  = "/user.v1alpha1.UserInterface/DeleteUser"
)

// UserInterfaceClient is the client API for UserInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInterfaceClient interface {
	SignIn(ctx context.Context, in *SignIn, opts ...grpc.CallOption) (*User, error)
	GetUserInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*Users, error)
	SaveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Msg, error)
}

type userInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInterfaceClient(cc grpc.ClientConnInterface) UserInterfaceClient {
	return &userInterfaceClient{cc}
}

func (c *userInterfaceClient) SignIn(ctx context.Context, in *SignIn, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserInterface_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInterfaceClient) GetUserInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserInterface_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInterfaceClient) GetUsers(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, UserInterface_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInterfaceClient) SaveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserInterface_SaveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInterfaceClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, UserInterface_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInterfaceServer is the server API for UserInterface service.
// All implementations must embed UnimplementedUserInterfaceServer
// for forward compatibility
type UserInterfaceServer interface {
	SignIn(context.Context, *SignIn) (*User, error)
	GetUserInfo(context.Context, *emptypb.Empty) (*User, error)
	GetUsers(context.Context, *UsersRequest) (*Users, error)
	SaveUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, *User) (*Msg, error)
	mustEmbedUnimplementedUserInterfaceServer()
}

// UnimplementedUserInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedUserInterfaceServer struct {
}

func (UnimplementedUserInterfaceServer) SignIn(context.Context, *SignIn) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedUserInterfaceServer) GetUserInfo(context.Context, *emptypb.Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserInterfaceServer) GetUsers(context.Context, *UsersRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserInterfaceServer) SaveUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (UnimplementedUserInterfaceServer) DeleteUser(context.Context, *User) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserInterfaceServer) mustEmbedUnimplementedUserInterfaceServer() {}

// UnsafeUserInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInterfaceServer will
// result in compilation errors.
type UnsafeUserInterfaceServer interface {
	mustEmbedUnimplementedUserInterfaceServer()
}

func RegisterUserInterfaceServer(s grpc.ServiceRegistrar, srv UserInterfaceServer) {
	s.RegisterService(&UserInterface_ServiceDesc, srv)
}

func _UserInterface_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterfaceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInterface_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterfaceServer).SignIn(ctx, req.(*SignIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInterface_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterfaceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInterface_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterfaceServer).GetUserInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInterface_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterfaceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInterface_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterfaceServer).GetUsers(ctx, req.(*UsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInterface_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterfaceServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInterface_SaveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterfaceServer).SaveUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInterface_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterfaceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInterface_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterfaceServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInterface_ServiceDesc is the grpc.ServiceDesc for UserInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1alpha1.UserInterface",
	HandlerType: (*UserInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _UserInterface_SignIn_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserInterface_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserInterface_GetUsers_Handler,
		},
		{
			MethodName: "SaveUser",
			Handler:    _UserInterface_SaveUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserInterface_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1alpha1/user.proto",
}
