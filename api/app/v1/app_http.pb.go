// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.19.4
// source: api/app/v1/app.proto

package v1

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

const OperationAppServiceGetAppConfig = "/app.v1.AppService/GetAppConfig"
const OperationAppServiceGetApps = "/app.v1.AppService/GetApps"
const OperationAppServiceSaveAppConfig = "/app.v1.AppService/SaveAppConfig"
const OperationAppServiceSaveApps = "/app.v1.AppService/SaveApps"

type AppServiceHTTPServer interface {
	GetAppConfig(context.Context, *GetAppConfigRequest) (*GetAppConfigResponse, error)
	GetApps(context.Context, *emptypb.Empty) (*App, error)
	SaveAppConfig(context.Context, *SaveAppConfigRequest) (*Msg, error)
	SaveApps(context.Context, *App) (*Msg, error)
}

func RegisterAppServiceHTTPServer(s *http.Server, srv AppServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/app/v1/get", _AppService_GetApps0_HTTP_Handler(srv))
	r.PUT("/app/v1/save", _AppService_SaveApps0_HTTP_Handler(srv))
	r.GET("/app/v1/config/get", _AppService_GetAppConfig0_HTTP_Handler(srv))
	r.PUT("/app/v1/config/save", _AppService_SaveAppConfig0_HTTP_Handler(srv))
}

func _AppService_GetApps0_HTTP_Handler(srv AppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppServiceGetApps)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetApps(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*App)
		return ctx.Result(200, reply)
	}
}

func _AppService_SaveApps0_HTTP_Handler(srv AppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in App
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppServiceSaveApps)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveApps(ctx, req.(*App))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _AppService_GetAppConfig0_HTTP_Handler(srv AppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAppConfigRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppServiceGetAppConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAppConfig(ctx, req.(*GetAppConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAppConfigResponse)
		return ctx.Result(200, reply)
	}
}

func _AppService_SaveAppConfig0_HTTP_Handler(srv AppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveAppConfigRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppServiceSaveAppConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveAppConfig(ctx, req.(*SaveAppConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

type AppServiceHTTPClient interface {
	GetAppConfig(ctx context.Context, req *GetAppConfigRequest, opts ...http.CallOption) (rsp *GetAppConfigResponse, err error)
	GetApps(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *App, err error)
	SaveAppConfig(ctx context.Context, req *SaveAppConfigRequest, opts ...http.CallOption) (rsp *Msg, err error)
	SaveApps(ctx context.Context, req *App, opts ...http.CallOption) (rsp *Msg, err error)
}

type AppServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewAppServiceHTTPClient(client *http.Client) AppServiceHTTPClient {
	return &AppServiceHTTPClientImpl{client}
}

func (c *AppServiceHTTPClientImpl) GetAppConfig(ctx context.Context, in *GetAppConfigRequest, opts ...http.CallOption) (*GetAppConfigResponse, error) {
	var out GetAppConfigResponse
	pattern := "/app/v1/config/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppServiceGetAppConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppServiceHTTPClientImpl) GetApps(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*App, error) {
	var out App
	pattern := "/app/v1/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppServiceGetApps))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppServiceHTTPClientImpl) SaveAppConfig(ctx context.Context, in *SaveAppConfigRequest, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/app/v1/config/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppServiceSaveAppConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppServiceHTTPClientImpl) SaveApps(ctx context.Context, in *App, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/app/v1/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppServiceSaveApps))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
