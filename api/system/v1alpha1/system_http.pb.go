// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.1
// source: api/system/v1alpha1/system.proto

package v1alpha1

import (
	context "context"
	common "github.com/f-rambo/ocean/api/common"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSystemInterfaceGetSystem = "/system.v1alpha1.SystemInterface/GetSystem"
const OperationSystemInterfacePing = "/system.v1alpha1.SystemInterface/Ping"

type SystemInterfaceHTTPServer interface {
	GetSystem(context.Context, *emptypb.Empty) (*System, error)
	Ping(context.Context, *emptypb.Empty) (*common.Msg, error)
}

func RegisterSystemInterfaceHTTPServer(s *http.Server, srv SystemInterfaceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1alpha1/system/ping", _SystemInterface_Ping3_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/system", _SystemInterface_GetSystem0_HTTP_Handler(srv))
}

func _SystemInterface_Ping3_HTTP_Handler(srv SystemInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemInterfacePing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _SystemInterface_GetSystem0_HTTP_Handler(srv SystemInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemInterfaceGetSystem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSystem(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*System)
		return ctx.Result(200, reply)
	}
}

type SystemInterfaceHTTPClient interface {
	GetSystem(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *System, err error)
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *common.Msg, err error)
}

type SystemInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewSystemInterfaceHTTPClient(client *http.Client) SystemInterfaceHTTPClient {
	return &SystemInterfaceHTTPClientImpl{client}
}

func (c *SystemInterfaceHTTPClientImpl) GetSystem(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*System, error) {
	var out System
	pattern := "/api/v1alpha1/system"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemInterfaceGetSystem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SystemInterfaceHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/system/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemInterfacePing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
