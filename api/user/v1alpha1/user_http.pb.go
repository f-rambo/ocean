// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.25.1
// source: api/user/v1alpha1/user.proto

package v1alpha1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserInterfaceGetUserInfo = "/user.v1alpha1.UserInterface/GetUserInfo"
const OperationUserInterfaceSignIn = "/user.v1alpha1.UserInterface/SignIn"
const OperationUserInterfaceSignOut = "/user.v1alpha1.UserInterface/SignOut"

type UserInterfaceHTTPServer interface {
	GetUserInfo(context.Context, *emptypb.Empty) (*User, error)
	SignIn(context.Context, *SignIn) (*User, error)
	SignOut(context.Context, *emptypb.Empty) (*Msg, error)
}

func RegisterUserInterfaceHTTPServer(s *http.Server, srv UserInterfaceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1alpha1/user/signin", _UserInterface_SignIn0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/user/signout", _UserInterface_SignOut0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/user", _UserInterface_GetUserInfo0_HTTP_Handler(srv))
}

func _UserInterface_SignIn0_HTTP_Handler(srv UserInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignIn
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserInterfaceSignIn)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignIn(ctx, req.(*SignIn))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*User)
		return ctx.Result(200, reply)
	}
}

func _UserInterface_SignOut0_HTTP_Handler(srv UserInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserInterfaceSignOut)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignOut(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _UserInterface_GetUserInfo0_HTTP_Handler(srv UserInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserInterfaceGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*User)
		return ctx.Result(200, reply)
	}
}

type UserInterfaceHTTPClient interface {
	GetUserInfo(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *User, err error)
	SignIn(ctx context.Context, req *SignIn, opts ...http.CallOption) (rsp *User, err error)
	SignOut(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Msg, err error)
}

type UserInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserInterfaceHTTPClient(client *http.Client) UserInterfaceHTTPClient {
	return &UserInterfaceHTTPClientImpl{client}
}

func (c *UserInterfaceHTTPClientImpl) GetUserInfo(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*User, error) {
	var out User
	pattern := "/api/v1alpha1/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserInterfaceGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserInterfaceHTTPClientImpl) SignIn(ctx context.Context, in *SignIn, opts ...http.CallOption) (*User, error) {
	var out User
	pattern := "/api/v1alpha1/user/signin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserInterfaceSignIn))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserInterfaceHTTPClientImpl) SignOut(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/user/signout"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserInterfaceSignOut))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}