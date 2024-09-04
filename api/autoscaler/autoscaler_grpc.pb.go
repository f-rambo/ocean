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
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: api/autoscaler/autoscaler.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CloudProvider_NodeGroups_FullMethodName                  = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/NodeGroups"
	CloudProvider_NodeGroupForNode_FullMethodName            = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/NodeGroupForNode"
	CloudProvider_PricingNodePrice_FullMethodName            = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/PricingNodePrice"
	CloudProvider_PricingPodPrice_FullMethodName             = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/PricingPodPrice"
	CloudProvider_GPULabel_FullMethodName                    = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/GPULabel"
	CloudProvider_GetAvailableGPUTypes_FullMethodName        = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/GetAvailableGPUTypes"
	CloudProvider_Cleanup_FullMethodName                     = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/Cleanup"
	CloudProvider_Refresh_FullMethodName                     = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/Refresh"
	CloudProvider_NodeGroupTargetSize_FullMethodName         = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/NodeGroupTargetSize"
	CloudProvider_NodeGroupIncreaseSize_FullMethodName       = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/NodeGroupIncreaseSize"
	CloudProvider_NodeGroupDeleteNodes_FullMethodName        = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/NodeGroupDeleteNodes"
	CloudProvider_NodeGroupDecreaseTargetSize_FullMethodName = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/NodeGroupDecreaseTargetSize"
	CloudProvider_NodeGroupNodes_FullMethodName              = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/NodeGroupNodes"
	CloudProvider_NodeGroupTemplateNodeInfo_FullMethodName   = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/NodeGroupTemplateNodeInfo"
	CloudProvider_NodeGroupGetOptions_FullMethodName         = "/clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider/NodeGroupGetOptions"
)

// CloudProviderClient is the client API for CloudProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudProviderClient interface {
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

type cloudProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudProviderClient(cc grpc.ClientConnInterface) CloudProviderClient {
	return &cloudProviderClient{cc}
}

func (c *cloudProviderClient) NodeGroups(ctx context.Context, in *NodeGroupsRequest, opts ...grpc.CallOption) (*NodeGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGroupsResponse)
	err := c.cc.Invoke(ctx, CloudProvider_NodeGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) NodeGroupForNode(ctx context.Context, in *NodeGroupForNodeRequest, opts ...grpc.CallOption) (*NodeGroupForNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGroupForNodeResponse)
	err := c.cc.Invoke(ctx, CloudProvider_NodeGroupForNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) PricingNodePrice(ctx context.Context, in *PricingNodePriceRequest, opts ...grpc.CallOption) (*PricingNodePriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PricingNodePriceResponse)
	err := c.cc.Invoke(ctx, CloudProvider_PricingNodePrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) PricingPodPrice(ctx context.Context, in *PricingPodPriceRequest, opts ...grpc.CallOption) (*PricingPodPriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PricingPodPriceResponse)
	err := c.cc.Invoke(ctx, CloudProvider_PricingPodPrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) GPULabel(ctx context.Context, in *GPULabelRequest, opts ...grpc.CallOption) (*GPULabelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GPULabelResponse)
	err := c.cc.Invoke(ctx, CloudProvider_GPULabel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) GetAvailableGPUTypes(ctx context.Context, in *GetAvailableGPUTypesRequest, opts ...grpc.CallOption) (*GetAvailableGPUTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAvailableGPUTypesResponse)
	err := c.cc.Invoke(ctx, CloudProvider_GetAvailableGPUTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) Cleanup(ctx context.Context, in *CleanupRequest, opts ...grpc.CallOption) (*CleanupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CleanupResponse)
	err := c.cc.Invoke(ctx, CloudProvider_Cleanup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, CloudProvider_Refresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) NodeGroupTargetSize(ctx context.Context, in *NodeGroupTargetSizeRequest, opts ...grpc.CallOption) (*NodeGroupTargetSizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGroupTargetSizeResponse)
	err := c.cc.Invoke(ctx, CloudProvider_NodeGroupTargetSize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) NodeGroupIncreaseSize(ctx context.Context, in *NodeGroupIncreaseSizeRequest, opts ...grpc.CallOption) (*NodeGroupIncreaseSizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGroupIncreaseSizeResponse)
	err := c.cc.Invoke(ctx, CloudProvider_NodeGroupIncreaseSize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) NodeGroupDeleteNodes(ctx context.Context, in *NodeGroupDeleteNodesRequest, opts ...grpc.CallOption) (*NodeGroupDeleteNodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGroupDeleteNodesResponse)
	err := c.cc.Invoke(ctx, CloudProvider_NodeGroupDeleteNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) NodeGroupDecreaseTargetSize(ctx context.Context, in *NodeGroupDecreaseTargetSizeRequest, opts ...grpc.CallOption) (*NodeGroupDecreaseTargetSizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGroupDecreaseTargetSizeResponse)
	err := c.cc.Invoke(ctx, CloudProvider_NodeGroupDecreaseTargetSize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) NodeGroupNodes(ctx context.Context, in *NodeGroupNodesRequest, opts ...grpc.CallOption) (*NodeGroupNodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGroupNodesResponse)
	err := c.cc.Invoke(ctx, CloudProvider_NodeGroupNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) NodeGroupTemplateNodeInfo(ctx context.Context, in *NodeGroupTemplateNodeInfoRequest, opts ...grpc.CallOption) (*NodeGroupTemplateNodeInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGroupTemplateNodeInfoResponse)
	err := c.cc.Invoke(ctx, CloudProvider_NodeGroupTemplateNodeInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderClient) NodeGroupGetOptions(ctx context.Context, in *NodeGroupAutoscalingOptionsRequest, opts ...grpc.CallOption) (*NodeGroupAutoscalingOptionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGroupAutoscalingOptionsResponse)
	err := c.cc.Invoke(ctx, CloudProvider_NodeGroupGetOptions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudProviderServer is the server API for CloudProvider service.
// All implementations must embed UnimplementedCloudProviderServer
// for forward compatibility.
type CloudProviderServer interface {
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
	mustEmbedUnimplementedCloudProviderServer()
}

// UnimplementedCloudProviderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCloudProviderServer struct{}

func (UnimplementedCloudProviderServer) NodeGroups(context.Context, *NodeGroupsRequest) (*NodeGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroups not implemented")
}
func (UnimplementedCloudProviderServer) NodeGroupForNode(context.Context, *NodeGroupForNodeRequest) (*NodeGroupForNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupForNode not implemented")
}
func (UnimplementedCloudProviderServer) PricingNodePrice(context.Context, *PricingNodePriceRequest) (*PricingNodePriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PricingNodePrice not implemented")
}
func (UnimplementedCloudProviderServer) PricingPodPrice(context.Context, *PricingPodPriceRequest) (*PricingPodPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PricingPodPrice not implemented")
}
func (UnimplementedCloudProviderServer) GPULabel(context.Context, *GPULabelRequest) (*GPULabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GPULabel not implemented")
}
func (UnimplementedCloudProviderServer) GetAvailableGPUTypes(context.Context, *GetAvailableGPUTypesRequest) (*GetAvailableGPUTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableGPUTypes not implemented")
}
func (UnimplementedCloudProviderServer) Cleanup(context.Context, *CleanupRequest) (*CleanupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cleanup not implemented")
}
func (UnimplementedCloudProviderServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedCloudProviderServer) NodeGroupTargetSize(context.Context, *NodeGroupTargetSizeRequest) (*NodeGroupTargetSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupTargetSize not implemented")
}
func (UnimplementedCloudProviderServer) NodeGroupIncreaseSize(context.Context, *NodeGroupIncreaseSizeRequest) (*NodeGroupIncreaseSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupIncreaseSize not implemented")
}
func (UnimplementedCloudProviderServer) NodeGroupDeleteNodes(context.Context, *NodeGroupDeleteNodesRequest) (*NodeGroupDeleteNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupDeleteNodes not implemented")
}
func (UnimplementedCloudProviderServer) NodeGroupDecreaseTargetSize(context.Context, *NodeGroupDecreaseTargetSizeRequest) (*NodeGroupDecreaseTargetSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupDecreaseTargetSize not implemented")
}
func (UnimplementedCloudProviderServer) NodeGroupNodes(context.Context, *NodeGroupNodesRequest) (*NodeGroupNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupNodes not implemented")
}
func (UnimplementedCloudProviderServer) NodeGroupTemplateNodeInfo(context.Context, *NodeGroupTemplateNodeInfoRequest) (*NodeGroupTemplateNodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupTemplateNodeInfo not implemented")
}
func (UnimplementedCloudProviderServer) NodeGroupGetOptions(context.Context, *NodeGroupAutoscalingOptionsRequest) (*NodeGroupAutoscalingOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupGetOptions not implemented")
}
func (UnimplementedCloudProviderServer) mustEmbedUnimplementedCloudProviderServer() {}
func (UnimplementedCloudProviderServer) testEmbeddedByValue()                       {}

// UnsafeCloudProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudProviderServer will
// result in compilation errors.
type UnsafeCloudProviderServer interface {
	mustEmbedUnimplementedCloudProviderServer()
}

func RegisterCloudProviderServer(s grpc.ServiceRegistrar, srv CloudProviderServer) {
	// If the following call pancis, it indicates UnimplementedCloudProviderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CloudProvider_ServiceDesc, srv)
}

func _CloudProvider_NodeGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).NodeGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_NodeGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).NodeGroups(ctx, req.(*NodeGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_NodeGroupForNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupForNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).NodeGroupForNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_NodeGroupForNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).NodeGroupForNode(ctx, req.(*NodeGroupForNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_PricingNodePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PricingNodePriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).PricingNodePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_PricingNodePrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).PricingNodePrice(ctx, req.(*PricingNodePriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_PricingPodPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PricingPodPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).PricingPodPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_PricingPodPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).PricingPodPrice(ctx, req.(*PricingPodPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_GPULabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GPULabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).GPULabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_GPULabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).GPULabel(ctx, req.(*GPULabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_GetAvailableGPUTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableGPUTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).GetAvailableGPUTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_GetAvailableGPUTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).GetAvailableGPUTypes(ctx, req.(*GetAvailableGPUTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_Cleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).Cleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_Cleanup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).Cleanup(ctx, req.(*CleanupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_NodeGroupTargetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupTargetSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).NodeGroupTargetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_NodeGroupTargetSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).NodeGroupTargetSize(ctx, req.(*NodeGroupTargetSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_NodeGroupIncreaseSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupIncreaseSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).NodeGroupIncreaseSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_NodeGroupIncreaseSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).NodeGroupIncreaseSize(ctx, req.(*NodeGroupIncreaseSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_NodeGroupDeleteNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupDeleteNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).NodeGroupDeleteNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_NodeGroupDeleteNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).NodeGroupDeleteNodes(ctx, req.(*NodeGroupDeleteNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_NodeGroupDecreaseTargetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupDecreaseTargetSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).NodeGroupDecreaseTargetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_NodeGroupDecreaseTargetSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).NodeGroupDecreaseTargetSize(ctx, req.(*NodeGroupDecreaseTargetSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_NodeGroupNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).NodeGroupNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_NodeGroupNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).NodeGroupNodes(ctx, req.(*NodeGroupNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_NodeGroupTemplateNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupTemplateNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).NodeGroupTemplateNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_NodeGroupTemplateNodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).NodeGroupTemplateNodeInfo(ctx, req.(*NodeGroupTemplateNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProvider_NodeGroupGetOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupAutoscalingOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServer).NodeGroupGetOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudProvider_NodeGroupGetOptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServer).NodeGroupGetOptions(ctx, req.(*NodeGroupAutoscalingOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudProvider_ServiceDesc is the grpc.ServiceDesc for CloudProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clusterautoscaler.cloudprovider.v1.externalgrpc.CloudProvider",
	HandlerType: (*CloudProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NodeGroups",
			Handler:    _CloudProvider_NodeGroups_Handler,
		},
		{
			MethodName: "NodeGroupForNode",
			Handler:    _CloudProvider_NodeGroupForNode_Handler,
		},
		{
			MethodName: "PricingNodePrice",
			Handler:    _CloudProvider_PricingNodePrice_Handler,
		},
		{
			MethodName: "PricingPodPrice",
			Handler:    _CloudProvider_PricingPodPrice_Handler,
		},
		{
			MethodName: "GPULabel",
			Handler:    _CloudProvider_GPULabel_Handler,
		},
		{
			MethodName: "GetAvailableGPUTypes",
			Handler:    _CloudProvider_GetAvailableGPUTypes_Handler,
		},
		{
			MethodName: "Cleanup",
			Handler:    _CloudProvider_Cleanup_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _CloudProvider_Refresh_Handler,
		},
		{
			MethodName: "NodeGroupTargetSize",
			Handler:    _CloudProvider_NodeGroupTargetSize_Handler,
		},
		{
			MethodName: "NodeGroupIncreaseSize",
			Handler:    _CloudProvider_NodeGroupIncreaseSize_Handler,
		},
		{
			MethodName: "NodeGroupDeleteNodes",
			Handler:    _CloudProvider_NodeGroupDeleteNodes_Handler,
		},
		{
			MethodName: "NodeGroupDecreaseTargetSize",
			Handler:    _CloudProvider_NodeGroupDecreaseTargetSize_Handler,
		},
		{
			MethodName: "NodeGroupNodes",
			Handler:    _CloudProvider_NodeGroupNodes_Handler,
		},
		{
			MethodName: "NodeGroupTemplateNodeInfo",
			Handler:    _CloudProvider_NodeGroupTemplateNodeInfo_Handler,
		},
		{
			MethodName: "NodeGroupGetOptions",
			Handler:    _CloudProvider_NodeGroupGetOptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/autoscaler/autoscaler.proto",
}
