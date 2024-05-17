// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.20.3
// source: api/service/v1alpha1/service.proto

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

const OperationServiceInterfaceCommitWorklfow = "/service.v1alpha1.ServiceInterface/CommitWorklfow"
const OperationServiceInterfaceDelete = "/service.v1alpha1.ServiceInterface/Delete"
const OperationServiceInterfaceGet = "/service.v1alpha1.ServiceInterface/Get"
const OperationServiceInterfaceGetServiceCis = "/service.v1alpha1.ServiceInterface/GetServiceCis"
const OperationServiceInterfaceGetWorkflow = "/service.v1alpha1.ServiceInterface/GetWorkflow"
const OperationServiceInterfaceList = "/service.v1alpha1.ServiceInterface/List"
const OperationServiceInterfacePing = "/service.v1alpha1.ServiceInterface/Ping"
const OperationServiceInterfaceSave = "/service.v1alpha1.ServiceInterface/Save"
const OperationServiceInterfaceSaveWorkflow = "/service.v1alpha1.ServiceInterface/SaveWorkflow"

type ServiceInterfaceHTTPServer interface {
	CommitWorklfow(context.Context, *ServiceRequest) (*Msg, error)
	Delete(context.Context, *ServiceRequest) (*Msg, error)
	Get(context.Context, *ServiceRequest) (*Service, error)
	GetServiceCis(context.Context, *CIsRequest) (*CIsResult, error)
	GetWorkflow(context.Context, *ServiceRequest) (*Worklfow, error)
	List(context.Context, *ServiceRequest) (*Services, error)
	Ping(context.Context, *emptypb.Empty) (*Msg, error)
	Save(context.Context, *Service) (*Msg, error)
	SaveWorkflow(context.Context, *ServiceRequest) (*Msg, error)
}

func RegisterServiceInterfaceHTTPServer(s *http.Server, srv ServiceInterfaceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1alpha1/service/ping", _ServiceInterface_Ping3_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/service/list", _ServiceInterface_List3_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/service/save", _ServiceInterface_Save3_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/service/get", _ServiceInterface_Get3_HTTP_Handler(srv))
	r.DELETE("/api/v1alpha1/service/delete", _ServiceInterface_Delete3_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/service/workflow", _ServiceInterface_GetWorkflow0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/service/workflow", _ServiceInterface_SaveWorkflow0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/service/commit", _ServiceInterface_CommitWorklfow0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/service/cis", _ServiceInterface_GetServiceCis0_HTTP_Handler(srv))
}

func _ServiceInterface_Ping3_HTTP_Handler(srv ServiceInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceInterfacePing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _ServiceInterface_List3_HTTP_Handler(srv ServiceInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ServiceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceInterfaceList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Services)
		return ctx.Result(200, reply)
	}
}

func _ServiceInterface_Save3_HTTP_Handler(srv ServiceInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Service
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceInterfaceSave)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Save(ctx, req.(*Service))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _ServiceInterface_Get3_HTTP_Handler(srv ServiceInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ServiceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceInterfaceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*ServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Service)
		return ctx.Result(200, reply)
	}
}

func _ServiceInterface_Delete3_HTTP_Handler(srv ServiceInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ServiceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceInterfaceDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*ServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _ServiceInterface_GetWorkflow0_HTTP_Handler(srv ServiceInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ServiceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceInterfaceGetWorkflow)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWorkflow(ctx, req.(*ServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Worklfow)
		return ctx.Result(200, reply)
	}
}

func _ServiceInterface_SaveWorkflow0_HTTP_Handler(srv ServiceInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ServiceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceInterfaceSaveWorkflow)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveWorkflow(ctx, req.(*ServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _ServiceInterface_CommitWorklfow0_HTTP_Handler(srv ServiceInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ServiceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceInterfaceCommitWorklfow)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommitWorklfow(ctx, req.(*ServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _ServiceInterface_GetServiceCis0_HTTP_Handler(srv ServiceInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CIsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceInterfaceGetServiceCis)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetServiceCis(ctx, req.(*CIsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CIsResult)
		return ctx.Result(200, reply)
	}
}

type ServiceInterfaceHTTPClient interface {
	CommitWorklfow(ctx context.Context, req *ServiceRequest, opts ...http.CallOption) (rsp *Msg, err error)
	Delete(ctx context.Context, req *ServiceRequest, opts ...http.CallOption) (rsp *Msg, err error)
	Get(ctx context.Context, req *ServiceRequest, opts ...http.CallOption) (rsp *Service, err error)
	GetServiceCis(ctx context.Context, req *CIsRequest, opts ...http.CallOption) (rsp *CIsResult, err error)
	GetWorkflow(ctx context.Context, req *ServiceRequest, opts ...http.CallOption) (rsp *Worklfow, err error)
	List(ctx context.Context, req *ServiceRequest, opts ...http.CallOption) (rsp *Services, err error)
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Msg, err error)
	Save(ctx context.Context, req *Service, opts ...http.CallOption) (rsp *Msg, err error)
	SaveWorkflow(ctx context.Context, req *ServiceRequest, opts ...http.CallOption) (rsp *Msg, err error)
}

type ServiceInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewServiceInterfaceHTTPClient(client *http.Client) ServiceInterfaceHTTPClient {
	return &ServiceInterfaceHTTPClientImpl{client}
}

func (c *ServiceInterfaceHTTPClientImpl) CommitWorklfow(ctx context.Context, in *ServiceRequest, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/service/commit"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceInterfaceCommitWorklfow))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceInterfaceHTTPClientImpl) Delete(ctx context.Context, in *ServiceRequest, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/service/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceInterfaceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceInterfaceHTTPClientImpl) Get(ctx context.Context, in *ServiceRequest, opts ...http.CallOption) (*Service, error) {
	var out Service
	pattern := "/api/v1alpha1/service/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceInterfaceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceInterfaceHTTPClientImpl) GetServiceCis(ctx context.Context, in *CIsRequest, opts ...http.CallOption) (*CIsResult, error) {
	var out CIsResult
	pattern := "/api/v1alpha1/service/cis"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceInterfaceGetServiceCis))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceInterfaceHTTPClientImpl) GetWorkflow(ctx context.Context, in *ServiceRequest, opts ...http.CallOption) (*Worklfow, error) {
	var out Worklfow
	pattern := "/api/v1alpha1/service/workflow"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceInterfaceGetWorkflow))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceInterfaceHTTPClientImpl) List(ctx context.Context, in *ServiceRequest, opts ...http.CallOption) (*Services, error) {
	var out Services
	pattern := "/api/v1alpha1/service/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceInterfaceList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceInterfaceHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/service/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceInterfacePing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceInterfaceHTTPClientImpl) Save(ctx context.Context, in *Service, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/service/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceInterfaceSave))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceInterfaceHTTPClientImpl) SaveWorkflow(ctx context.Context, in *ServiceRequest, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/service/workflow"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceInterfaceSaveWorkflow))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
