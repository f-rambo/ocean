// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.20.3
// source: api/app/v1alpha1/app.proto

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

const OperationAppInterfaceAppTest = "/app.v1alpha1.AppInterface/AppTest"
const OperationAppInterfaceCreateAppType = "/app.v1alpha1.AppInterface/CreateAppType"
const OperationAppInterfaceDelete = "/app.v1alpha1.AppInterface/Delete"
const OperationAppInterfaceDeleteAppType = "/app.v1alpha1.AppInterface/DeleteAppType"
const OperationAppInterfaceDeleteDeployedApp = "/app.v1alpha1.AppInterface/DeleteDeployedApp"
const OperationAppInterfaceDeleteRepo = "/app.v1alpha1.AppInterface/DeleteRepo"
const OperationAppInterfaceDeployApp = "/app.v1alpha1.AppInterface/DeployApp"
const OperationAppInterfaceGet = "/app.v1alpha1.AppInterface/Get"
const OperationAppInterfaceGetAppDeployed = "/app.v1alpha1.AppInterface/GetAppDeployed"
const OperationAppInterfaceGetAppDetailByRepo = "/app.v1alpha1.AppInterface/GetAppDetailByRepo"
const OperationAppInterfaceGetAppsByRepo = "/app.v1alpha1.AppInterface/GetAppsByRepo"
const OperationAppInterfaceGetDeployedAppResources = "/app.v1alpha1.AppInterface/GetDeployedAppResources"
const OperationAppInterfaceList = "/app.v1alpha1.AppInterface/List"
const OperationAppInterfaceListAppType = "/app.v1alpha1.AppInterface/ListAppType"
const OperationAppInterfaceListDeployedApp = "/app.v1alpha1.AppInterface/ListDeployedApp"
const OperationAppInterfaceListRepo = "/app.v1alpha1.AppInterface/ListRepo"
const OperationAppInterfacePing = "/app.v1alpha1.AppInterface/Ping"
const OperationAppInterfaceSave = "/app.v1alpha1.AppInterface/Save"
const OperationAppInterfaceSaveRepo = "/app.v1alpha1.AppInterface/SaveRepo"
const OperationAppInterfaceStopApp = "/app.v1alpha1.AppInterface/StopApp"
const OperationAppInterfaceUploadApp = "/app.v1alpha1.AppInterface/UploadApp"

type AppInterfaceHTTPServer interface {
	AppTest(context.Context, *DeployAppReq) (*DeployApp, error)
	CreateAppType(context.Context, *AppType) (*Msg, error)
	Delete(context.Context, *AppReq) (*Msg, error)
	DeleteAppType(context.Context, *AppTypeReq) (*Msg, error)
	DeleteDeployedApp(context.Context, *DeployAppReq) (*Msg, error)
	DeleteRepo(context.Context, *AppHelmRepoReq) (*Msg, error)
	DeployApp(context.Context, *DeployAppReq) (*DeployApp, error)
	Get(context.Context, *AppReq) (*App, error)
	GetAppDeployed(context.Context, *DeployApp) (*DeployApp, error)
	GetAppDetailByRepo(context.Context, *AppHelmRepoReq) (*App, error)
	GetAppsByRepo(context.Context, *AppHelmRepoReq) (*AppList, error)
	GetDeployedAppResources(context.Context, *DeployAppReq) (*DeployAppResources, error)
	List(context.Context, *AppReq) (*AppList, error)
	ListAppType(context.Context, *emptypb.Empty) (*AppTypeList, error)
	ListDeployedApp(context.Context, *DeployAppReq) (*DeployAppList, error)
	ListRepo(context.Context, *emptypb.Empty) (*AppHelmRepoList, error)
	Ping(context.Context, *emptypb.Empty) (*Msg, error)
	Save(context.Context, *App) (*Msg, error)
	SaveRepo(context.Context, *AppHelmRepo) (*Msg, error)
	StopApp(context.Context, *DeployAppReq) (*Msg, error)
	UploadApp(context.Context, *FileUploadRequest) (*App, error)
}

func RegisterAppInterfaceHTTPServer(s *http.Server, srv AppInterfaceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1alpha1/app/ping", _AppInterface_Ping1_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/app/upload", _AppInterface_UploadApp0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/app/save", _AppInterface_Save1_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app", _AppInterface_Get1_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/list", _AppInterface_List1_HTTP_Handler(srv))
	r.DELETE("/api/v1alpha1/app", _AppInterface_Delete1_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/app/type", _AppInterface_CreateAppType0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/type/list", _AppInterface_ListAppType0_HTTP_Handler(srv))
	r.DELETE("/api/v1alpha1/app/type", _AppInterface_DeleteAppType0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/app/test", _AppInterface_AppTest0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/deploy", _AppInterface_GetAppDeployed0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/app/deploy", _AppInterface_DeployApp0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/deploy/list", _AppInterface_ListDeployedApp0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/app/stop", _AppInterface_StopApp0_HTTP_Handler(srv))
	r.DELETE("/api/v1alpha1/app/deploy", _AppInterface_DeleteDeployedApp0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/deploy/resources", _AppInterface_GetDeployedAppResources0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/app/repo", _AppInterface_SaveRepo0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/repo/list", _AppInterface_ListRepo0_HTTP_Handler(srv))
	r.DELETE("/api/v1alpha1/app/repo", _AppInterface_DeleteRepo0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/repo/apps", _AppInterface_GetAppsByRepo0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/repo/app/detail", _AppInterface_GetAppDetailByRepo0_HTTP_Handler(srv))
}

func _AppInterface_Ping1_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfacePing)
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

func _AppInterface_UploadApp0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FileUploadRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceUploadApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadApp(ctx, req.(*FileUploadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*App)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_Save1_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in App
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceSave)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Save(ctx, req.(*App))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_Get1_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*AppReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*App)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_List1_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*AppReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppList)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_Delete1_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*AppReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_CreateAppType0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppType
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceCreateAppType)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAppType(ctx, req.(*AppType))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_ListAppType0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceListAppType)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAppType(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppTypeList)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_DeleteAppType0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppTypeReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceDeleteAppType)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAppType(ctx, req.(*AppTypeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_AppTest0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeployAppReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceAppTest)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppTest(ctx, req.(*DeployAppReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeployApp)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_GetAppDeployed0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeployApp
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceGetAppDeployed)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAppDeployed(ctx, req.(*DeployApp))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeployApp)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_DeployApp0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeployAppReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceDeployApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeployApp(ctx, req.(*DeployAppReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeployApp)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_ListDeployedApp0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeployAppReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceListDeployedApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDeployedApp(ctx, req.(*DeployAppReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeployAppList)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_StopApp0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeployAppReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceStopApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StopApp(ctx, req.(*DeployAppReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_DeleteDeployedApp0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeployAppReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceDeleteDeployedApp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDeployedApp(ctx, req.(*DeployAppReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_GetDeployedAppResources0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeployAppReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceGetDeployedAppResources)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDeployedAppResources(ctx, req.(*DeployAppReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeployAppResources)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_SaveRepo0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppHelmRepo
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceSaveRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveRepo(ctx, req.(*AppHelmRepo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_ListRepo0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceListRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRepo(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppHelmRepoList)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_DeleteRepo0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppHelmRepoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceDeleteRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRepo(ctx, req.(*AppHelmRepoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_GetAppsByRepo0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppHelmRepoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceGetAppsByRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAppsByRepo(ctx, req.(*AppHelmRepoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppList)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_GetAppDetailByRepo0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppHelmRepoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceGetAppDetailByRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAppDetailByRepo(ctx, req.(*AppHelmRepoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*App)
		return ctx.Result(200, reply)
	}
}

type AppInterfaceHTTPClient interface {
	AppTest(ctx context.Context, req *DeployAppReq, opts ...http.CallOption) (rsp *DeployApp, err error)
	CreateAppType(ctx context.Context, req *AppType, opts ...http.CallOption) (rsp *Msg, err error)
	Delete(ctx context.Context, req *AppReq, opts ...http.CallOption) (rsp *Msg, err error)
	DeleteAppType(ctx context.Context, req *AppTypeReq, opts ...http.CallOption) (rsp *Msg, err error)
	DeleteDeployedApp(ctx context.Context, req *DeployAppReq, opts ...http.CallOption) (rsp *Msg, err error)
	DeleteRepo(ctx context.Context, req *AppHelmRepoReq, opts ...http.CallOption) (rsp *Msg, err error)
	DeployApp(ctx context.Context, req *DeployAppReq, opts ...http.CallOption) (rsp *DeployApp, err error)
	Get(ctx context.Context, req *AppReq, opts ...http.CallOption) (rsp *App, err error)
	GetAppDeployed(ctx context.Context, req *DeployApp, opts ...http.CallOption) (rsp *DeployApp, err error)
	GetAppDetailByRepo(ctx context.Context, req *AppHelmRepoReq, opts ...http.CallOption) (rsp *App, err error)
	GetAppsByRepo(ctx context.Context, req *AppHelmRepoReq, opts ...http.CallOption) (rsp *AppList, err error)
	GetDeployedAppResources(ctx context.Context, req *DeployAppReq, opts ...http.CallOption) (rsp *DeployAppResources, err error)
	List(ctx context.Context, req *AppReq, opts ...http.CallOption) (rsp *AppList, err error)
	ListAppType(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *AppTypeList, err error)
	ListDeployedApp(ctx context.Context, req *DeployAppReq, opts ...http.CallOption) (rsp *DeployAppList, err error)
	ListRepo(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *AppHelmRepoList, err error)
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Msg, err error)
	Save(ctx context.Context, req *App, opts ...http.CallOption) (rsp *Msg, err error)
	SaveRepo(ctx context.Context, req *AppHelmRepo, opts ...http.CallOption) (rsp *Msg, err error)
	StopApp(ctx context.Context, req *DeployAppReq, opts ...http.CallOption) (rsp *Msg, err error)
	UploadApp(ctx context.Context, req *FileUploadRequest, opts ...http.CallOption) (rsp *App, err error)
}

type AppInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewAppInterfaceHTTPClient(client *http.Client) AppInterfaceHTTPClient {
	return &AppInterfaceHTTPClientImpl{client}
}

func (c *AppInterfaceHTTPClientImpl) AppTest(ctx context.Context, in *DeployAppReq, opts ...http.CallOption) (*DeployApp, error) {
	var out DeployApp
	pattern := "/api/v1alpha1/app/test"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceAppTest))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) CreateAppType(ctx context.Context, in *AppType, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/app/type"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceCreateAppType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) Delete(ctx context.Context, in *AppReq, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/app"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) DeleteAppType(ctx context.Context, in *AppTypeReq, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/app/type"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceDeleteAppType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) DeleteDeployedApp(ctx context.Context, in *DeployAppReq, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/app/deploy"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceDeleteDeployedApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) DeleteRepo(ctx context.Context, in *AppHelmRepoReq, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/app/repo"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceDeleteRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) DeployApp(ctx context.Context, in *DeployAppReq, opts ...http.CallOption) (*DeployApp, error) {
	var out DeployApp
	pattern := "/api/v1alpha1/app/deploy"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceDeployApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) Get(ctx context.Context, in *AppReq, opts ...http.CallOption) (*App, error) {
	var out App
	pattern := "/api/v1alpha1/app"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) GetAppDeployed(ctx context.Context, in *DeployApp, opts ...http.CallOption) (*DeployApp, error) {
	var out DeployApp
	pattern := "/api/v1alpha1/app/deploy"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceGetAppDeployed))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) GetAppDetailByRepo(ctx context.Context, in *AppHelmRepoReq, opts ...http.CallOption) (*App, error) {
	var out App
	pattern := "/api/v1alpha1/app/repo/app/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceGetAppDetailByRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) GetAppsByRepo(ctx context.Context, in *AppHelmRepoReq, opts ...http.CallOption) (*AppList, error) {
	var out AppList
	pattern := "/api/v1alpha1/app/repo/apps"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceGetAppsByRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) GetDeployedAppResources(ctx context.Context, in *DeployAppReq, opts ...http.CallOption) (*DeployAppResources, error) {
	var out DeployAppResources
	pattern := "/api/v1alpha1/app/deploy/resources"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceGetDeployedAppResources))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) List(ctx context.Context, in *AppReq, opts ...http.CallOption) (*AppList, error) {
	var out AppList
	pattern := "/api/v1alpha1/app/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) ListAppType(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*AppTypeList, error) {
	var out AppTypeList
	pattern := "/api/v1alpha1/app/type/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceListAppType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) ListDeployedApp(ctx context.Context, in *DeployAppReq, opts ...http.CallOption) (*DeployAppList, error) {
	var out DeployAppList
	pattern := "/api/v1alpha1/app/deploy/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceListDeployedApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) ListRepo(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*AppHelmRepoList, error) {
	var out AppHelmRepoList
	pattern := "/api/v1alpha1/app/repo/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceListRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/app/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfacePing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) Save(ctx context.Context, in *App, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/app/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceSave))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) SaveRepo(ctx context.Context, in *AppHelmRepo, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/app/repo"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceSaveRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) StopApp(ctx context.Context, in *DeployAppReq, opts ...http.CallOption) (*Msg, error) {
	var out Msg
	pattern := "/api/v1alpha1/app/stop"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceStopApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppInterfaceHTTPClientImpl) UploadApp(ctx context.Context, in *FileUploadRequest, opts ...http.CallOption) (*App, error) {
	var out App
	pattern := "/api/v1alpha1/app/upload"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceUploadApp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
