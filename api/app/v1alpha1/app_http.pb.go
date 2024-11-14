// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.1
// source: api/app/v1alpha1/app.proto

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

const OperationAppInterfaceAppReleaseList = "/app.v1alpha1.AppInterface/AppReleaseList"
const OperationAppInterfaceCreateAppType = "/app.v1alpha1.AppInterface/CreateAppType"
const OperationAppInterfaceDelete = "/app.v1alpha1.AppInterface/Delete"
const OperationAppInterfaceDeleteAppRelease = "/app.v1alpha1.AppInterface/DeleteAppRelease"
const OperationAppInterfaceDeleteAppType = "/app.v1alpha1.AppInterface/DeleteAppType"
const OperationAppInterfaceDeleteRepo = "/app.v1alpha1.AppInterface/DeleteRepo"
const OperationAppInterfaceGet = "/app.v1alpha1.AppInterface/Get"
const OperationAppInterfaceGetAppDetailByRepo = "/app.v1alpha1.AppInterface/GetAppDetailByRepo"
const OperationAppInterfaceGetAppRelease = "/app.v1alpha1.AppInterface/GetAppRelease"
const OperationAppInterfaceGetAppReleaseResources = "/app.v1alpha1.AppInterface/GetAppReleaseResources"
const OperationAppInterfaceGetAppsByRepo = "/app.v1alpha1.AppInterface/GetAppsByRepo"
const OperationAppInterfaceList = "/app.v1alpha1.AppInterface/List"
const OperationAppInterfaceListAppType = "/app.v1alpha1.AppInterface/ListAppType"
const OperationAppInterfaceListRepo = "/app.v1alpha1.AppInterface/ListRepo"
const OperationAppInterfacePing = "/app.v1alpha1.AppInterface/Ping"
const OperationAppInterfaceSave = "/app.v1alpha1.AppInterface/Save"
const OperationAppInterfaceSaveAppRelease = "/app.v1alpha1.AppInterface/SaveAppRelease"
const OperationAppInterfaceSaveRepo = "/app.v1alpha1.AppInterface/SaveRepo"
const OperationAppInterfaceUploadApp = "/app.v1alpha1.AppInterface/UploadApp"

type AppInterfaceHTTPServer interface {
	AppReleaseList(context.Context, *AppReleaseReq) (*AppReleaseList, error)
	CreateAppType(context.Context, *AppType) (*common.Msg, error)
	Delete(context.Context, *AppReq) (*common.Msg, error)
	DeleteAppRelease(context.Context, *AppReleaseReq) (*common.Msg, error)
	DeleteAppType(context.Context, *AppTypeReq) (*common.Msg, error)
	DeleteRepo(context.Context, *AppRepoReq) (*common.Msg, error)
	Get(context.Context, *AppReq) (*App, error)
	GetAppDetailByRepo(context.Context, *AppRepoReq) (*App, error)
	GetAppRelease(context.Context, *AppReleaseReq) (*AppRelease, error)
	GetAppReleaseResources(context.Context, *AppReleaseReq) (*AppReleasepResources, error)
	GetAppsByRepo(context.Context, *AppRepoReq) (*AppList, error)
	List(context.Context, *AppReq) (*AppList, error)
	ListAppType(context.Context, *emptypb.Empty) (*AppTypeList, error)
	ListRepo(context.Context, *emptypb.Empty) (*AppRepoList, error)
	Ping(context.Context, *emptypb.Empty) (*common.Msg, error)
	Save(context.Context, *App) (*common.Msg, error)
	SaveAppRelease(context.Context, *AppReleaseReq) (*AppRelease, error)
	SaveRepo(context.Context, *AppRepo) (*common.Msg, error)
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
	r.POST("/api/v1alpha1/app/repo", _AppInterface_SaveRepo0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/repo/list", _AppInterface_ListRepo0_HTTP_Handler(srv))
	r.DELETE("/api/v1alpha1/app/repo", _AppInterface_DeleteRepo0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/repo/apps", _AppInterface_GetAppsByRepo0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/repo/app/detail", _AppInterface_GetAppDetailByRepo0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/release", _AppInterface_GetAppRelease0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/release/list", _AppInterface_AppReleaseList0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/app/release/resources", _AppInterface_GetAppReleaseResources0_HTTP_Handler(srv))
	r.POST("/api/v1alpha1/app/release", _AppInterface_SaveAppRelease0_HTTP_Handler(srv))
	r.DELETE("/api/v1alpha1/app/release", _AppInterface_DeleteAppRelease0_HTTP_Handler(srv))
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
		reply := out.(*common.Msg)
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
		reply := out.(*common.Msg)
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
		reply := out.(*common.Msg)
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
		reply := out.(*common.Msg)
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
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_SaveRepo0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppRepo
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceSaveRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveRepo(ctx, req.(*AppRepo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
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
		reply := out.(*AppRepoList)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_DeleteRepo0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppRepoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceDeleteRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRepo(ctx, req.(*AppRepoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_GetAppsByRepo0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppRepoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceGetAppsByRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAppsByRepo(ctx, req.(*AppRepoReq))
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
		var in AppRepoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceGetAppDetailByRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAppDetailByRepo(ctx, req.(*AppRepoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*App)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_GetAppRelease0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppReleaseReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceGetAppRelease)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAppRelease(ctx, req.(*AppReleaseReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppRelease)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_AppReleaseList0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppReleaseReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceAppReleaseList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppReleaseList(ctx, req.(*AppReleaseReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppReleaseList)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_GetAppReleaseResources0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppReleaseReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceGetAppReleaseResources)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAppReleaseResources(ctx, req.(*AppReleaseReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppReleasepResources)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_SaveAppRelease0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppReleaseReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceSaveAppRelease)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveAppRelease(ctx, req.(*AppReleaseReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppRelease)
		return ctx.Result(200, reply)
	}
}

func _AppInterface_DeleteAppRelease0_HTTP_Handler(srv AppInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppReleaseReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppInterfaceDeleteAppRelease)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAppRelease(ctx, req.(*AppReleaseReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

type AppInterfaceHTTPClient interface {
	AppReleaseList(ctx context.Context, req *AppReleaseReq, opts ...http.CallOption) (rsp *AppReleaseList, err error)
	CreateAppType(ctx context.Context, req *AppType, opts ...http.CallOption) (rsp *common.Msg, err error)
	Delete(ctx context.Context, req *AppReq, opts ...http.CallOption) (rsp *common.Msg, err error)
	DeleteAppRelease(ctx context.Context, req *AppReleaseReq, opts ...http.CallOption) (rsp *common.Msg, err error)
	DeleteAppType(ctx context.Context, req *AppTypeReq, opts ...http.CallOption) (rsp *common.Msg, err error)
	DeleteRepo(ctx context.Context, req *AppRepoReq, opts ...http.CallOption) (rsp *common.Msg, err error)
	Get(ctx context.Context, req *AppReq, opts ...http.CallOption) (rsp *App, err error)
	GetAppDetailByRepo(ctx context.Context, req *AppRepoReq, opts ...http.CallOption) (rsp *App, err error)
	GetAppRelease(ctx context.Context, req *AppReleaseReq, opts ...http.CallOption) (rsp *AppRelease, err error)
	GetAppReleaseResources(ctx context.Context, req *AppReleaseReq, opts ...http.CallOption) (rsp *AppReleasepResources, err error)
	GetAppsByRepo(ctx context.Context, req *AppRepoReq, opts ...http.CallOption) (rsp *AppList, err error)
	List(ctx context.Context, req *AppReq, opts ...http.CallOption) (rsp *AppList, err error)
	ListAppType(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *AppTypeList, err error)
	ListRepo(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *AppRepoList, err error)
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *common.Msg, err error)
	Save(ctx context.Context, req *App, opts ...http.CallOption) (rsp *common.Msg, err error)
	SaveAppRelease(ctx context.Context, req *AppReleaseReq, opts ...http.CallOption) (rsp *AppRelease, err error)
	SaveRepo(ctx context.Context, req *AppRepo, opts ...http.CallOption) (rsp *common.Msg, err error)
	UploadApp(ctx context.Context, req *FileUploadRequest, opts ...http.CallOption) (rsp *App, err error)
}

type AppInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewAppInterfaceHTTPClient(client *http.Client) AppInterfaceHTTPClient {
	return &AppInterfaceHTTPClientImpl{client}
}

func (c *AppInterfaceHTTPClientImpl) AppReleaseList(ctx context.Context, in *AppReleaseReq, opts ...http.CallOption) (*AppReleaseList, error) {
	var out AppReleaseList
	pattern := "/api/v1alpha1/app/release/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceAppReleaseList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) CreateAppType(ctx context.Context, in *AppType, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/app/type"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceCreateAppType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) Delete(ctx context.Context, in *AppReq, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/app"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) DeleteAppRelease(ctx context.Context, in *AppReleaseReq, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/app/release"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceDeleteAppRelease))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) DeleteAppType(ctx context.Context, in *AppTypeReq, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/app/type"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceDeleteAppType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) DeleteRepo(ctx context.Context, in *AppRepoReq, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/app/repo"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceDeleteRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
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
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) GetAppDetailByRepo(ctx context.Context, in *AppRepoReq, opts ...http.CallOption) (*App, error) {
	var out App
	pattern := "/api/v1alpha1/app/repo/app/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceGetAppDetailByRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) GetAppRelease(ctx context.Context, in *AppReleaseReq, opts ...http.CallOption) (*AppRelease, error) {
	var out AppRelease
	pattern := "/api/v1alpha1/app/release"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceGetAppRelease))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) GetAppReleaseResources(ctx context.Context, in *AppReleaseReq, opts ...http.CallOption) (*AppReleasepResources, error) {
	var out AppReleasepResources
	pattern := "/api/v1alpha1/app/release/resources"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceGetAppReleaseResources))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) GetAppsByRepo(ctx context.Context, in *AppRepoReq, opts ...http.CallOption) (*AppList, error) {
	var out AppList
	pattern := "/api/v1alpha1/app/repo/apps"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceGetAppsByRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
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
	return &out, nil
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
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) ListRepo(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*AppRepoList, error) {
	var out AppRepoList
	pattern := "/api/v1alpha1/app/repo/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfaceListRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/app/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppInterfacePing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) Save(ctx context.Context, in *App, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/app/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceSave))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) SaveAppRelease(ctx context.Context, in *AppReleaseReq, opts ...http.CallOption) (*AppRelease, error) {
	var out AppRelease
	pattern := "/api/v1alpha1/app/release"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceSaveAppRelease))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppInterfaceHTTPClientImpl) SaveRepo(ctx context.Context, in *AppRepo, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/app/repo"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppInterfaceSaveRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
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
	return &out, nil
}
