// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.1
// source: api/cloud/v1alpha1/cloud.proto

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

const OperationCloudInterfaceAddKubeletServiceAndSettingKubeadmConfig = "/cloud.v1alpha1.CloudInterface/AddKubeletServiceAndSettingKubeadmConfig"
const OperationCloudInterfaceCloseFirewall = "/cloud.v1alpha1.CloudInterface/CloseFirewall"
const OperationCloudInterfaceCloseSwap = "/cloud.v1alpha1.CloudInterface/CloseSwap"
const OperationCloudInterfaceInstallKubeadmKubeletCriO = "/cloud.v1alpha1.CloudInterface/InstallKubeadmKubeletCriO"
const OperationCloudInterfaceKubeadmInit = "/cloud.v1alpha1.CloudInterface/KubeadmInit"
const OperationCloudInterfaceKubeadmJoin = "/cloud.v1alpha1.CloudInterface/KubeadmJoin"
const OperationCloudInterfaceKubeadmReset = "/cloud.v1alpha1.CloudInterface/KubeadmReset"
const OperationCloudInterfaceKubeadmUpgrade = "/cloud.v1alpha1.CloudInterface/KubeadmUpgrade"
const OperationCloudInterfacePing = "/cloud.v1alpha1.CloudInterface/Ping"
const OperationCloudInterfaceSetingIpv4Forward = "/cloud.v1alpha1.CloudInterface/SetingIpv4Forward"

type CloudInterfaceHTTPServer interface {
	// AddKubeletServiceAndSettingKubeadmConfig AddKubeletServiceAndSettingKubeadmConfig
	AddKubeletServiceAndSettingKubeadmConfig(context.Context, *Cloud) (*common.Msg, error)
	// CloseFirewall CloseFirewall
	CloseFirewall(context.Context, *emptypb.Empty) (*common.Msg, error)
	// CloseSwap CloseSwap
	CloseSwap(context.Context, *emptypb.Empty) (*common.Msg, error)
	// InstallKubeadmKubeletCriO InstallKubeadmKubeletCriO
	InstallKubeadmKubeletCriO(context.Context, *Cloud) (*common.Msg, error)
	// KubeadmInit KubeadmInit
	KubeadmInit(context.Context, *Cloud) (*common.Msg, error)
	// KubeadmJoin KubeadmJoin
	KubeadmJoin(context.Context, *Cloud) (*common.Msg, error)
	// KubeadmReset KubeadmReset
	KubeadmReset(context.Context, *Cloud) (*common.Msg, error)
	// KubeadmUpgrade KubeadmUpgrade
	KubeadmUpgrade(context.Context, *Cloud) (*common.Msg, error)
	Ping(context.Context, *emptypb.Empty) (*common.Msg, error)
	// SetingIpv4Forward SetingIpv4Forward
	SetingIpv4Forward(context.Context, *emptypb.Empty) (*common.Msg, error)
}

func RegisterCloudInterfaceHTTPServer(s *http.Server, srv CloudInterfaceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1alpha1/cloud/ping", _CloudInterface_Ping4_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/cloud/install_kubeadm_kubelet_crio", _CloudInterface_InstallKubeadmKubeletCriO0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/cloud/add_kubelet_service_and_setting_kubeadm_config", _CloudInterface_AddKubeletServiceAndSettingKubeadmConfig0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/cloud/kubeadm_join", _CloudInterface_KubeadmJoin0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/cloud/kubeadm_init", _CloudInterface_KubeadmInit0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/cloud/kubeadm_reset", _CloudInterface_KubeadmReset0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/cloud/kubeadm_upgrade", _CloudInterface_KubeadmUpgrade0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/cloud/setting_ipv4_forward", _CloudInterface_SetingIpv4Forward0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/cloud/close_swap", _CloudInterface_CloseSwap0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/cloud/close_firewall", _CloudInterface_CloseFirewall0_HTTP_Handler(srv))
}

func _CloudInterface_Ping4_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfacePing)
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

func _CloudInterface_InstallKubeadmKubeletCriO0_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Cloud
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfaceInstallKubeadmKubeletCriO)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.InstallKubeadmKubeletCriO(ctx, req.(*Cloud))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _CloudInterface_AddKubeletServiceAndSettingKubeadmConfig0_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Cloud
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfaceAddKubeletServiceAndSettingKubeadmConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddKubeletServiceAndSettingKubeadmConfig(ctx, req.(*Cloud))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _CloudInterface_KubeadmJoin0_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Cloud
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfaceKubeadmJoin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KubeadmJoin(ctx, req.(*Cloud))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _CloudInterface_KubeadmInit0_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Cloud
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfaceKubeadmInit)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KubeadmInit(ctx, req.(*Cloud))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _CloudInterface_KubeadmReset0_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Cloud
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfaceKubeadmReset)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KubeadmReset(ctx, req.(*Cloud))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _CloudInterface_KubeadmUpgrade0_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Cloud
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfaceKubeadmUpgrade)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KubeadmUpgrade(ctx, req.(*Cloud))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _CloudInterface_SetingIpv4Forward0_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfaceSetingIpv4Forward)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetingIpv4Forward(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _CloudInterface_CloseSwap0_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfaceCloseSwap)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CloseSwap(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _CloudInterface_CloseFirewall0_HTTP_Handler(srv CloudInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCloudInterfaceCloseFirewall)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CloseFirewall(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

type CloudInterfaceHTTPClient interface {
	AddKubeletServiceAndSettingKubeadmConfig(ctx context.Context, req *Cloud, opts ...http.CallOption) (rsp *common.Msg, err error)
	CloseFirewall(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *common.Msg, err error)
	CloseSwap(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *common.Msg, err error)
	InstallKubeadmKubeletCriO(ctx context.Context, req *Cloud, opts ...http.CallOption) (rsp *common.Msg, err error)
	KubeadmInit(ctx context.Context, req *Cloud, opts ...http.CallOption) (rsp *common.Msg, err error)
	KubeadmJoin(ctx context.Context, req *Cloud, opts ...http.CallOption) (rsp *common.Msg, err error)
	KubeadmReset(ctx context.Context, req *Cloud, opts ...http.CallOption) (rsp *common.Msg, err error)
	KubeadmUpgrade(ctx context.Context, req *Cloud, opts ...http.CallOption) (rsp *common.Msg, err error)
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *common.Msg, err error)
	SetingIpv4Forward(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *common.Msg, err error)
}

type CloudInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewCloudInterfaceHTTPClient(client *http.Client) CloudInterfaceHTTPClient {
	return &CloudInterfaceHTTPClientImpl{client}
}

func (c *CloudInterfaceHTTPClientImpl) AddKubeletServiceAndSettingKubeadmConfig(ctx context.Context, in *Cloud, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/add_kubelet_service_and_setting_kubeadm_config"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCloudInterfaceAddKubeletServiceAndSettingKubeadmConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CloudInterfaceHTTPClientImpl) CloseFirewall(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/close_firewall"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCloudInterfaceCloseFirewall))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CloudInterfaceHTTPClientImpl) CloseSwap(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/close_swap"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCloudInterfaceCloseSwap))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CloudInterfaceHTTPClientImpl) InstallKubeadmKubeletCriO(ctx context.Context, in *Cloud, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/install_kubeadm_kubelet_crio"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCloudInterfaceInstallKubeadmKubeletCriO))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CloudInterfaceHTTPClientImpl) KubeadmInit(ctx context.Context, in *Cloud, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/kubeadm_init"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCloudInterfaceKubeadmInit))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CloudInterfaceHTTPClientImpl) KubeadmJoin(ctx context.Context, in *Cloud, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/kubeadm_join"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCloudInterfaceKubeadmJoin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CloudInterfaceHTTPClientImpl) KubeadmReset(ctx context.Context, in *Cloud, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/kubeadm_reset"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCloudInterfaceKubeadmReset))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CloudInterfaceHTTPClientImpl) KubeadmUpgrade(ctx context.Context, in *Cloud, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/kubeadm_upgrade"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCloudInterfaceKubeadmUpgrade))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CloudInterfaceHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCloudInterfacePing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CloudInterfaceHTTPClientImpl) SetingIpv4Forward(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cloud/setting_ipv4_forward"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCloudInterfaceSetingIpv4Forward))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
