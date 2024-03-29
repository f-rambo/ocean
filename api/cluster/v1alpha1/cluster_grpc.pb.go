// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/cluster/v1alpha1/cluster.proto

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
	ClusterInterface_Ping_FullMethodName               = "/cluster.v1alpha1.ClusterInterface/Ping"
	ClusterInterface_Get_FullMethodName                = "/cluster.v1alpha1.ClusterInterface/Get"
	ClusterInterface_Save_FullMethodName               = "/cluster.v1alpha1.ClusterInterface/Save"
	ClusterInterface_List_FullMethodName               = "/cluster.v1alpha1.ClusterInterface/List"
	ClusterInterface_Delete_FullMethodName             = "/cluster.v1alpha1.ClusterInterface/Delete"
	ClusterInterface_DeleteNode_FullMethodName         = "/cluster.v1alpha1.ClusterInterface/DeleteNode"
	ClusterInterface_GetClusterMockData_FullMethodName = "/cluster.v1alpha1.ClusterInterface/GetClusterMockData"
	ClusterInterface_CheckClusterConfig_FullMethodName = "/cluster.v1alpha1.ClusterInterface/CheckClusterConfig"
	ClusterInterface_SetUpCluster_FullMethodName       = "/cluster.v1alpha1.ClusterInterface/SetUpCluster"
	ClusterInterface_UninstallCluster_FullMethodName   = "/cluster.v1alpha1.ClusterInterface/UninstallCluster"
	ClusterInterface_AddNode_FullMethodName            = "/cluster.v1alpha1.ClusterInterface/AddNode"
	ClusterInterface_RemoveNode_FullMethodName         = "/cluster.v1alpha1.ClusterInterface/RemoveNode"
)

// ClusterInterfaceClient is the client API for ClusterInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterInterfaceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error)
	Get(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Cluster, error)
	Save(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Cluster, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClusterList, error)
	Delete(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error)
	DeleteNode(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error)
	GetClusterMockData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Cluster, error)
	CheckClusterConfig(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Cluster, error)
	SetUpCluster(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error)
	UninstallCluster(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error)
	AddNode(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error)
	RemoveNode(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error)
}

type clusterInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterInterfaceClient(cc grpc.ClientConnInterface) ClusterInterfaceClient {
	return &clusterInterfaceClient{cc}
}

func (c *clusterInterfaceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ClusterInterface_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) Get(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterInterface_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) Save(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterInterface_Save_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClusterList, error) {
	out := new(ClusterList)
	err := c.cc.Invoke(ctx, ClusterInterface_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) Delete(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ClusterInterface_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) DeleteNode(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ClusterInterface_DeleteNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) GetClusterMockData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterInterface_GetClusterMockData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) CheckClusterConfig(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterInterface_CheckClusterConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) SetUpCluster(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ClusterInterface_SetUpCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) UninstallCluster(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ClusterInterface_UninstallCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) AddNode(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ClusterInterface_AddNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterInterfaceClient) RemoveNode(ctx context.Context, in *ClusterID, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, ClusterInterface_RemoveNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterInterfaceServer is the server API for ClusterInterface service.
// All implementations must embed UnimplementedClusterInterfaceServer
// for forward compatibility
type ClusterInterfaceServer interface {
	Ping(context.Context, *emptypb.Empty) (*Msg, error)
	Get(context.Context, *ClusterID) (*Cluster, error)
	Save(context.Context, *Cluster) (*Cluster, error)
	List(context.Context, *emptypb.Empty) (*ClusterList, error)
	Delete(context.Context, *ClusterID) (*Msg, error)
	DeleteNode(context.Context, *ClusterID) (*Msg, error)
	GetClusterMockData(context.Context, *emptypb.Empty) (*Cluster, error)
	CheckClusterConfig(context.Context, *ClusterID) (*Cluster, error)
	SetUpCluster(context.Context, *ClusterID) (*Msg, error)
	UninstallCluster(context.Context, *ClusterID) (*Msg, error)
	AddNode(context.Context, *ClusterID) (*Msg, error)
	RemoveNode(context.Context, *ClusterID) (*Msg, error)
	mustEmbedUnimplementedClusterInterfaceServer()
}

// UnimplementedClusterInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedClusterInterfaceServer struct {
}

func (UnimplementedClusterInterfaceServer) Ping(context.Context, *emptypb.Empty) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedClusterInterfaceServer) Get(context.Context, *ClusterID) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedClusterInterfaceServer) Save(context.Context, *Cluster) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedClusterInterfaceServer) List(context.Context, *emptypb.Empty) (*ClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedClusterInterfaceServer) Delete(context.Context, *ClusterID) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClusterInterfaceServer) DeleteNode(context.Context, *ClusterID) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (UnimplementedClusterInterfaceServer) GetClusterMockData(context.Context, *emptypb.Empty) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterMockData not implemented")
}
func (UnimplementedClusterInterfaceServer) CheckClusterConfig(context.Context, *ClusterID) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckClusterConfig not implemented")
}
func (UnimplementedClusterInterfaceServer) SetUpCluster(context.Context, *ClusterID) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUpCluster not implemented")
}
func (UnimplementedClusterInterfaceServer) UninstallCluster(context.Context, *ClusterID) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UninstallCluster not implemented")
}
func (UnimplementedClusterInterfaceServer) AddNode(context.Context, *ClusterID) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}
func (UnimplementedClusterInterfaceServer) RemoveNode(context.Context, *ClusterID) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNode not implemented")
}
func (UnimplementedClusterInterfaceServer) mustEmbedUnimplementedClusterInterfaceServer() {}

// UnsafeClusterInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterInterfaceServer will
// result in compilation errors.
type UnsafeClusterInterfaceServer interface {
	mustEmbedUnimplementedClusterInterfaceServer()
}

func RegisterClusterInterfaceServer(s grpc.ServiceRegistrar, srv ClusterInterfaceServer) {
	s.RegisterService(&ClusterInterface_ServiceDesc, srv)
}

func _ClusterInterface_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).Get(ctx, req.(*ClusterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).Save(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).Delete(ctx, req.(*ClusterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_DeleteNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).DeleteNode(ctx, req.(*ClusterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_GetClusterMockData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).GetClusterMockData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_GetClusterMockData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).GetClusterMockData(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_CheckClusterConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).CheckClusterConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_CheckClusterConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).CheckClusterConfig(ctx, req.(*ClusterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_SetUpCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).SetUpCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_SetUpCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).SetUpCluster(ctx, req.(*ClusterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_UninstallCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).UninstallCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_UninstallCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).UninstallCluster(ctx, req.(*ClusterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_AddNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).AddNode(ctx, req.(*ClusterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterInterface_RemoveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInterfaceServer).RemoveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterInterface_RemoveNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInterfaceServer).RemoveNode(ctx, req.(*ClusterID))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterInterface_ServiceDesc is the grpc.ServiceDesc for ClusterInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cluster.v1alpha1.ClusterInterface",
	HandlerType: (*ClusterInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ClusterInterface_Ping_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ClusterInterface_Get_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _ClusterInterface_Save_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ClusterInterface_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClusterInterface_Delete_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _ClusterInterface_DeleteNode_Handler,
		},
		{
			MethodName: "GetClusterMockData",
			Handler:    _ClusterInterface_GetClusterMockData_Handler,
		},
		{
			MethodName: "CheckClusterConfig",
			Handler:    _ClusterInterface_CheckClusterConfig_Handler,
		},
		{
			MethodName: "SetUpCluster",
			Handler:    _ClusterInterface_SetUpCluster_Handler,
		},
		{
			MethodName: "UninstallCluster",
			Handler:    _ClusterInterface_UninstallCluster_Handler,
		},
		{
			MethodName: "AddNode",
			Handler:    _ClusterInterface_AddNode_Handler,
		},
		{
			MethodName: "RemoveNode",
			Handler:    _ClusterInterface_RemoveNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/cluster/v1alpha1/cluster.proto",
}
