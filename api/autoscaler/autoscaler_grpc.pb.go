//
//Copyright 2022 The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/autoscaler/autoscaler.proto

package autoscaler

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AutoscalerService_NodeGroups_FullMethodName                  = "/autoscaler.AutoscalerService/NodeGroups"
	AutoscalerService_NodeGroupForNode_FullMethodName            = "/autoscaler.AutoscalerService/NodeGroupForNode"
	AutoscalerService_PricingNodePrice_FullMethodName            = "/autoscaler.AutoscalerService/PricingNodePrice"
	AutoscalerService_PricingPodPrice_FullMethodName             = "/autoscaler.AutoscalerService/PricingPodPrice"
	AutoscalerService_GPULabel_FullMethodName                    = "/autoscaler.AutoscalerService/GPULabel"
	AutoscalerService_GetAvailableGPUTypes_FullMethodName        = "/autoscaler.AutoscalerService/GetAvailableGPUTypes"
	AutoscalerService_Cleanup_FullMethodName                     = "/autoscaler.AutoscalerService/Cleanup"
	AutoscalerService_Refresh_FullMethodName                     = "/autoscaler.AutoscalerService/Refresh"
	AutoscalerService_NodeGroupTargetSize_FullMethodName         = "/autoscaler.AutoscalerService/NodeGroupTargetSize"
	AutoscalerService_NodeGroupIncreaseSize_FullMethodName       = "/autoscaler.AutoscalerService/NodeGroupIncreaseSize"
	AutoscalerService_NodeGroupDeleteNodes_FullMethodName        = "/autoscaler.AutoscalerService/NodeGroupDeleteNodes"
	AutoscalerService_NodeGroupDecreaseTargetSize_FullMethodName = "/autoscaler.AutoscalerService/NodeGroupDecreaseTargetSize"
	AutoscalerService_NodeGroupNodes_FullMethodName              = "/autoscaler.AutoscalerService/NodeGroupNodes"
	AutoscalerService_NodeGroupTemplateNodeInfo_FullMethodName   = "/autoscaler.AutoscalerService/NodeGroupTemplateNodeInfo"
	AutoscalerService_NodeGroupGetOptions_FullMethodName         = "/autoscaler.AutoscalerService/NodeGroupGetOptions"
)

// AutoscalerServiceClient is the client API for AutoscalerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutoscalerServiceClient interface {
	// NodeGroups returns all node groups configured for this cloud provider.
	NodeGroups(ctx context.Context, in *NodeGroupsRequest, opts ...grpc.CallOption) (*NodeGroupsResponse, error)
	// NodeGroupForNode returns the node group for the given node.
	// The node group id is an empty string if the node should not
	// be processed by cluster autoscaler.
	NodeGroupForNode(ctx context.Context, in *NodeGroupForNodeRequest, opts ...grpc.CallOption) (*NodeGroupForNodeResponse, error)
	// PricingNodePrice returns a theoretical minimum price of running a node for
	// a given period of time on a perfectly matching machine.
	// Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	PricingNodePrice(ctx context.Context, in *PricingNodePriceRequest, opts ...grpc.CallOption) (*PricingNodePriceResponse, error)
	// PricingPodPrice returns a theoretical minimum price of running a pod for a given
	// period of time on a perfectly matching machine.
	// Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	PricingPodPrice(ctx context.Context, in *PricingPodPriceRequest, opts ...grpc.CallOption) (*PricingPodPriceResponse, error)
	// GPULabel returns the label added to nodes with GPU resource.
	GPULabel(ctx context.Context, in *GPULabelRequest, opts ...grpc.CallOption) (*GPULabelResponse, error)
	// GetAvailableGPUTypes return all available GPU types cloud provider supports.
	GetAvailableGPUTypes(ctx context.Context, in *GetAvailableGPUTypesRequest, opts ...grpc.CallOption) (*GetAvailableGPUTypesResponse, error)
	// Cleanup cleans up open resources before the cloud provider is destroyed, i.e. go routines etc.
	Cleanup(ctx context.Context, in *CleanupRequest, opts ...grpc.CallOption) (*CleanupResponse, error)
	// Refresh is called before every main loop and can be used to dynamically update cloud provider state.
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
	// NodeGroupTargetSize returns the current target size of the node group. It is possible
	// that the number of nodes in Kubernetes is different at the moment but should be equal
	// to the size of a node group once everything stabilizes (new nodes finish startup and
	// registration or removed nodes are deleted completely).
	NodeGroupTargetSize(ctx context.Context, in *NodeGroupTargetSizeRequest, opts ...grpc.CallOption) (*NodeGroupTargetSizeResponse, error)
	// NodeGroupIncreaseSize increases the size of the node group. To delete a node you need
	// to explicitly name it and use NodeGroupDeleteNodes. This function should wait until
	// node group size is updated.
	NodeGroupIncreaseSize(ctx context.Context, in *NodeGroupIncreaseSizeRequest, opts ...grpc.CallOption) (*NodeGroupIncreaseSizeResponse, error)
	// NodeGroupDeleteNodes deletes nodes from this node group (and also decreasing the size
	// of the node group with that). Error is returned either on failure or if the given node
	// doesn't belong to this node group. This function should wait until node group size is updated.
	NodeGroupDeleteNodes(ctx context.Context, in *NodeGroupDeleteNodesRequest, opts ...grpc.CallOption) (*NodeGroupDeleteNodesResponse, error)
	// NodeGroupDecreaseTargetSize decreases the target size of the node group. This function
	// doesn't permit to delete any existing node and can be used only to reduce the request
	// for new nodes that have not been yet fulfilled. Delta should be negative. It is assumed
	// that cloud provider will not delete the existing nodes if the size when there is an option
	// to just decrease the target.
	NodeGroupDecreaseTargetSize(ctx context.Context, in *NodeGroupDecreaseTargetSizeRequest, opts ...grpc.CallOption) (*NodeGroupDecreaseTargetSizeResponse, error)
	// NodeGroupNodes returns a list of all nodes that belong to this node group.
	NodeGroupNodes(ctx context.Context, in *NodeGroupNodesRequest, opts ...grpc.CallOption) (*NodeGroupNodesResponse, error)
	// NodeGroupTemplateNodeInfo returns a structure of an empty (as if just started) node,
	// with all of the labels, capacity and allocatable information. This will be used in
	// scale-up simulations to predict what would a new node look like if a node group was expanded.
	// Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	NodeGroupTemplateNodeInfo(ctx context.Context, in *NodeGroupTemplateNodeInfoRequest, opts ...grpc.CallOption) (*NodeGroupTemplateNodeInfoResponse, error)
	// GetOptions returns NodeGroupAutoscalingOptions that should be used for this particular
	// NodeGroup.
	// Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	NodeGroupGetOptions(ctx context.Context, in *NodeGroupAutoscalingOptionsRequest, opts ...grpc.CallOption) (*NodeGroupAutoscalingOptionsResponse, error)
}

type autoscalerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutoscalerServiceClient(cc grpc.ClientConnInterface) AutoscalerServiceClient {
	return &autoscalerServiceClient{cc}
}

func (c *autoscalerServiceClient) NodeGroups(ctx context.Context, in *NodeGroupsRequest, opts ...grpc.CallOption) (*NodeGroupsResponse, error) {
	out := new(NodeGroupsResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_NodeGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) NodeGroupForNode(ctx context.Context, in *NodeGroupForNodeRequest, opts ...grpc.CallOption) (*NodeGroupForNodeResponse, error) {
	out := new(NodeGroupForNodeResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_NodeGroupForNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) PricingNodePrice(ctx context.Context, in *PricingNodePriceRequest, opts ...grpc.CallOption) (*PricingNodePriceResponse, error) {
	out := new(PricingNodePriceResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_PricingNodePrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) PricingPodPrice(ctx context.Context, in *PricingPodPriceRequest, opts ...grpc.CallOption) (*PricingPodPriceResponse, error) {
	out := new(PricingPodPriceResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_PricingPodPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) GPULabel(ctx context.Context, in *GPULabelRequest, opts ...grpc.CallOption) (*GPULabelResponse, error) {
	out := new(GPULabelResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_GPULabel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) GetAvailableGPUTypes(ctx context.Context, in *GetAvailableGPUTypesRequest, opts ...grpc.CallOption) (*GetAvailableGPUTypesResponse, error) {
	out := new(GetAvailableGPUTypesResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_GetAvailableGPUTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) Cleanup(ctx context.Context, in *CleanupRequest, opts ...grpc.CallOption) (*CleanupResponse, error) {
	out := new(CleanupResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_Cleanup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) NodeGroupTargetSize(ctx context.Context, in *NodeGroupTargetSizeRequest, opts ...grpc.CallOption) (*NodeGroupTargetSizeResponse, error) {
	out := new(NodeGroupTargetSizeResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_NodeGroupTargetSize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) NodeGroupIncreaseSize(ctx context.Context, in *NodeGroupIncreaseSizeRequest, opts ...grpc.CallOption) (*NodeGroupIncreaseSizeResponse, error) {
	out := new(NodeGroupIncreaseSizeResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_NodeGroupIncreaseSize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) NodeGroupDeleteNodes(ctx context.Context, in *NodeGroupDeleteNodesRequest, opts ...grpc.CallOption) (*NodeGroupDeleteNodesResponse, error) {
	out := new(NodeGroupDeleteNodesResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_NodeGroupDeleteNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) NodeGroupDecreaseTargetSize(ctx context.Context, in *NodeGroupDecreaseTargetSizeRequest, opts ...grpc.CallOption) (*NodeGroupDecreaseTargetSizeResponse, error) {
	out := new(NodeGroupDecreaseTargetSizeResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_NodeGroupDecreaseTargetSize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) NodeGroupNodes(ctx context.Context, in *NodeGroupNodesRequest, opts ...grpc.CallOption) (*NodeGroupNodesResponse, error) {
	out := new(NodeGroupNodesResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_NodeGroupNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) NodeGroupTemplateNodeInfo(ctx context.Context, in *NodeGroupTemplateNodeInfoRequest, opts ...grpc.CallOption) (*NodeGroupTemplateNodeInfoResponse, error) {
	out := new(NodeGroupTemplateNodeInfoResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_NodeGroupTemplateNodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoscalerServiceClient) NodeGroupGetOptions(ctx context.Context, in *NodeGroupAutoscalingOptionsRequest, opts ...grpc.CallOption) (*NodeGroupAutoscalingOptionsResponse, error) {
	out := new(NodeGroupAutoscalingOptionsResponse)
	err := c.cc.Invoke(ctx, AutoscalerService_NodeGroupGetOptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoscalerServiceServer is the server API for AutoscalerService service.
// All implementations must embed UnimplementedAutoscalerServiceServer
// for forward compatibility
type AutoscalerServiceServer interface {
	// NodeGroups returns all node groups configured for this cloud provider.
	NodeGroups(context.Context, *NodeGroupsRequest) (*NodeGroupsResponse, error)
	// NodeGroupForNode returns the node group for the given node.
	// The node group id is an empty string if the node should not
	// be processed by cluster autoscaler.
	NodeGroupForNode(context.Context, *NodeGroupForNodeRequest) (*NodeGroupForNodeResponse, error)
	// PricingNodePrice returns a theoretical minimum price of running a node for
	// a given period of time on a perfectly matching machine.
	// Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	PricingNodePrice(context.Context, *PricingNodePriceRequest) (*PricingNodePriceResponse, error)
	// PricingPodPrice returns a theoretical minimum price of running a pod for a given
	// period of time on a perfectly matching machine.
	// Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	PricingPodPrice(context.Context, *PricingPodPriceRequest) (*PricingPodPriceResponse, error)
	// GPULabel returns the label added to nodes with GPU resource.
	GPULabel(context.Context, *GPULabelRequest) (*GPULabelResponse, error)
	// GetAvailableGPUTypes return all available GPU types cloud provider supports.
	GetAvailableGPUTypes(context.Context, *GetAvailableGPUTypesRequest) (*GetAvailableGPUTypesResponse, error)
	// Cleanup cleans up open resources before the cloud provider is destroyed, i.e. go routines etc.
	Cleanup(context.Context, *CleanupRequest) (*CleanupResponse, error)
	// Refresh is called before every main loop and can be used to dynamically update cloud provider state.
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	// NodeGroupTargetSize returns the current target size of the node group. It is possible
	// that the number of nodes in Kubernetes is different at the moment but should be equal
	// to the size of a node group once everything stabilizes (new nodes finish startup and
	// registration or removed nodes are deleted completely).
	NodeGroupTargetSize(context.Context, *NodeGroupTargetSizeRequest) (*NodeGroupTargetSizeResponse, error)
	// NodeGroupIncreaseSize increases the size of the node group. To delete a node you need
	// to explicitly name it and use NodeGroupDeleteNodes. This function should wait until
	// node group size is updated.
	NodeGroupIncreaseSize(context.Context, *NodeGroupIncreaseSizeRequest) (*NodeGroupIncreaseSizeResponse, error)
	// NodeGroupDeleteNodes deletes nodes from this node group (and also decreasing the size
	// of the node group with that). Error is returned either on failure or if the given node
	// doesn't belong to this node group. This function should wait until node group size is updated.
	NodeGroupDeleteNodes(context.Context, *NodeGroupDeleteNodesRequest) (*NodeGroupDeleteNodesResponse, error)
	// NodeGroupDecreaseTargetSize decreases the target size of the node group. This function
	// doesn't permit to delete any existing node and can be used only to reduce the request
	// for new nodes that have not been yet fulfilled. Delta should be negative. It is assumed
	// that cloud provider will not delete the existing nodes if the size when there is an option
	// to just decrease the target.
	NodeGroupDecreaseTargetSize(context.Context, *NodeGroupDecreaseTargetSizeRequest) (*NodeGroupDecreaseTargetSizeResponse, error)
	// NodeGroupNodes returns a list of all nodes that belong to this node group.
	NodeGroupNodes(context.Context, *NodeGroupNodesRequest) (*NodeGroupNodesResponse, error)
	// NodeGroupTemplateNodeInfo returns a structure of an empty (as if just started) node,
	// with all of the labels, capacity and allocatable information. This will be used in
	// scale-up simulations to predict what would a new node look like if a node group was expanded.
	// Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	NodeGroupTemplateNodeInfo(context.Context, *NodeGroupTemplateNodeInfoRequest) (*NodeGroupTemplateNodeInfoResponse, error)
	// GetOptions returns NodeGroupAutoscalingOptions that should be used for this particular
	// NodeGroup.
	// Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	NodeGroupGetOptions(context.Context, *NodeGroupAutoscalingOptionsRequest) (*NodeGroupAutoscalingOptionsResponse, error)
	mustEmbedUnimplementedAutoscalerServiceServer()
}

// UnimplementedAutoscalerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAutoscalerServiceServer struct {
}

func (UnimplementedAutoscalerServiceServer) NodeGroups(context.Context, *NodeGroupsRequest) (*NodeGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroups not implemented")
}
func (UnimplementedAutoscalerServiceServer) NodeGroupForNode(context.Context, *NodeGroupForNodeRequest) (*NodeGroupForNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupForNode not implemented")
}
func (UnimplementedAutoscalerServiceServer) PricingNodePrice(context.Context, *PricingNodePriceRequest) (*PricingNodePriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PricingNodePrice not implemented")
}
func (UnimplementedAutoscalerServiceServer) PricingPodPrice(context.Context, *PricingPodPriceRequest) (*PricingPodPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PricingPodPrice not implemented")
}
func (UnimplementedAutoscalerServiceServer) GPULabel(context.Context, *GPULabelRequest) (*GPULabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GPULabel not implemented")
}
func (UnimplementedAutoscalerServiceServer) GetAvailableGPUTypes(context.Context, *GetAvailableGPUTypesRequest) (*GetAvailableGPUTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableGPUTypes not implemented")
}
func (UnimplementedAutoscalerServiceServer) Cleanup(context.Context, *CleanupRequest) (*CleanupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cleanup not implemented")
}
func (UnimplementedAutoscalerServiceServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAutoscalerServiceServer) NodeGroupTargetSize(context.Context, *NodeGroupTargetSizeRequest) (*NodeGroupTargetSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupTargetSize not implemented")
}
func (UnimplementedAutoscalerServiceServer) NodeGroupIncreaseSize(context.Context, *NodeGroupIncreaseSizeRequest) (*NodeGroupIncreaseSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupIncreaseSize not implemented")
}
func (UnimplementedAutoscalerServiceServer) NodeGroupDeleteNodes(context.Context, *NodeGroupDeleteNodesRequest) (*NodeGroupDeleteNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupDeleteNodes not implemented")
}
func (UnimplementedAutoscalerServiceServer) NodeGroupDecreaseTargetSize(context.Context, *NodeGroupDecreaseTargetSizeRequest) (*NodeGroupDecreaseTargetSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupDecreaseTargetSize not implemented")
}
func (UnimplementedAutoscalerServiceServer) NodeGroupNodes(context.Context, *NodeGroupNodesRequest) (*NodeGroupNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupNodes not implemented")
}
func (UnimplementedAutoscalerServiceServer) NodeGroupTemplateNodeInfo(context.Context, *NodeGroupTemplateNodeInfoRequest) (*NodeGroupTemplateNodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupTemplateNodeInfo not implemented")
}
func (UnimplementedAutoscalerServiceServer) NodeGroupGetOptions(context.Context, *NodeGroupAutoscalingOptionsRequest) (*NodeGroupAutoscalingOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupGetOptions not implemented")
}
func (UnimplementedAutoscalerServiceServer) mustEmbedUnimplementedAutoscalerServiceServer() {}

// UnsafeAutoscalerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutoscalerServiceServer will
// result in compilation errors.
type UnsafeAutoscalerServiceServer interface {
	mustEmbedUnimplementedAutoscalerServiceServer()
}

func RegisterAutoscalerServiceServer(s grpc.ServiceRegistrar, srv AutoscalerServiceServer) {
	s.RegisterService(&AutoscalerService_ServiceDesc, srv)
}

func _AutoscalerService_NodeGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).NodeGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_NodeGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).NodeGroups(ctx, req.(*NodeGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_NodeGroupForNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupForNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).NodeGroupForNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_NodeGroupForNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).NodeGroupForNode(ctx, req.(*NodeGroupForNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_PricingNodePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PricingNodePriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).PricingNodePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_PricingNodePrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).PricingNodePrice(ctx, req.(*PricingNodePriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_PricingPodPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PricingPodPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).PricingPodPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_PricingPodPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).PricingPodPrice(ctx, req.(*PricingPodPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_GPULabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GPULabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).GPULabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_GPULabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).GPULabel(ctx, req.(*GPULabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_GetAvailableGPUTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableGPUTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).GetAvailableGPUTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_GetAvailableGPUTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).GetAvailableGPUTypes(ctx, req.(*GetAvailableGPUTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_Cleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).Cleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_Cleanup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).Cleanup(ctx, req.(*CleanupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_NodeGroupTargetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupTargetSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).NodeGroupTargetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_NodeGroupTargetSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).NodeGroupTargetSize(ctx, req.(*NodeGroupTargetSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_NodeGroupIncreaseSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupIncreaseSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).NodeGroupIncreaseSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_NodeGroupIncreaseSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).NodeGroupIncreaseSize(ctx, req.(*NodeGroupIncreaseSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_NodeGroupDeleteNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupDeleteNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).NodeGroupDeleteNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_NodeGroupDeleteNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).NodeGroupDeleteNodes(ctx, req.(*NodeGroupDeleteNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_NodeGroupDecreaseTargetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupDecreaseTargetSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).NodeGroupDecreaseTargetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_NodeGroupDecreaseTargetSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).NodeGroupDecreaseTargetSize(ctx, req.(*NodeGroupDecreaseTargetSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_NodeGroupNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).NodeGroupNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_NodeGroupNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).NodeGroupNodes(ctx, req.(*NodeGroupNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_NodeGroupTemplateNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupTemplateNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).NodeGroupTemplateNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_NodeGroupTemplateNodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).NodeGroupTemplateNodeInfo(ctx, req.(*NodeGroupTemplateNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoscalerService_NodeGroupGetOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupAutoscalingOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoscalerServiceServer).NodeGroupGetOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoscalerService_NodeGroupGetOptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoscalerServiceServer).NodeGroupGetOptions(ctx, req.(*NodeGroupAutoscalingOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AutoscalerService_ServiceDesc is the grpc.ServiceDesc for AutoscalerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutoscalerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "autoscaler.AutoscalerService",
	HandlerType: (*AutoscalerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NodeGroups",
			Handler:    _AutoscalerService_NodeGroups_Handler,
		},
		{
			MethodName: "NodeGroupForNode",
			Handler:    _AutoscalerService_NodeGroupForNode_Handler,
		},
		{
			MethodName: "PricingNodePrice",
			Handler:    _AutoscalerService_PricingNodePrice_Handler,
		},
		{
			MethodName: "PricingPodPrice",
			Handler:    _AutoscalerService_PricingPodPrice_Handler,
		},
		{
			MethodName: "GPULabel",
			Handler:    _AutoscalerService_GPULabel_Handler,
		},
		{
			MethodName: "GetAvailableGPUTypes",
			Handler:    _AutoscalerService_GetAvailableGPUTypes_Handler,
		},
		{
			MethodName: "Cleanup",
			Handler:    _AutoscalerService_Cleanup_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _AutoscalerService_Refresh_Handler,
		},
		{
			MethodName: "NodeGroupTargetSize",
			Handler:    _AutoscalerService_NodeGroupTargetSize_Handler,
		},
		{
			MethodName: "NodeGroupIncreaseSize",
			Handler:    _AutoscalerService_NodeGroupIncreaseSize_Handler,
		},
		{
			MethodName: "NodeGroupDeleteNodes",
			Handler:    _AutoscalerService_NodeGroupDeleteNodes_Handler,
		},
		{
			MethodName: "NodeGroupDecreaseTargetSize",
			Handler:    _AutoscalerService_NodeGroupDecreaseTargetSize_Handler,
		},
		{
			MethodName: "NodeGroupNodes",
			Handler:    _AutoscalerService_NodeGroupNodes_Handler,
		},
		{
			MethodName: "NodeGroupTemplateNodeInfo",
			Handler:    _AutoscalerService_NodeGroupTemplateNodeInfo_Handler,
		},
		{
			MethodName: "NodeGroupGetOptions",
			Handler:    _AutoscalerService_NodeGroupGetOptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/autoscaler/autoscaler.proto",
}
