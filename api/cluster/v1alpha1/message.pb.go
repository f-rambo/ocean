// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/cluster/v1alpha1/message.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ServerType      string      `protobuf:"bytes,2,opt,name=server_type,proto3" json:"server_type,omitempty"`
	PublicKey       string      `protobuf:"bytes,3,opt,name=public_key,proto3" json:"public_key,omitempty"`
	Region          string      `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	AccessKeyId     string      `protobuf:"bytes,5,opt,name=access_key_id,proto3" json:"access_key_id,omitempty"`
	SecretAccessKey string      `protobuf:"bytes,6,opt,name=secret_access_key,proto3" json:"secret_access_key,omitempty"`
	Nodes           []*NodeArgs `protobuf:"bytes,7,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ClusterArgs) Reset() {
	*x = ClusterArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterArgs) ProtoMessage() {}

func (x *ClusterArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterArgs.ProtoReflect.Descriptor instead.
func (*ClusterArgs) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterArgs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterArgs) GetServerType() string {
	if x != nil {
		return x.ServerType
	}
	return ""
}

func (x *ClusterArgs) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *ClusterArgs) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ClusterArgs) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *ClusterArgs) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

func (x *ClusterArgs) GetNodes() []*NodeArgs {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type NodeArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *NodeArgs) Reset() {
	*x = NodeArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeArgs) ProtoMessage() {}

func (x *NodeArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeArgs.ProtoReflect.Descriptor instead.
func (*NodeArgs) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{1}
}

func (x *NodeArgs) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *NodeArgs) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type ClusterID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeId int64 `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *ClusterID) Reset() {
	*x = ClusterID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterID) ProtoMessage() {}

func (x *ClusterID) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterID.ProtoReflect.Descriptor instead.
func (*ClusterID) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ClusterID) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

type ClusterList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *ClusterList) Reset() {
	*x = ClusterList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterList) ProtoMessage() {}

func (x *ClusterList) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterList.ProtoReflect.Descriptor instead.
func (*ClusterList) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterList) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ServerVersion    string  `protobuf:"bytes,3,opt,name=server_version,proto3" json:"server_version,omitempty"`
	ApiServerAddress string  `protobuf:"bytes,4,opt,name=api_server_address,proto3" json:"api_server_address,omitempty"`
	Config           string  `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	Addons           string  `protobuf:"bytes,6,opt,name=addons,proto3" json:"addons,omitempty"`
	AddonsConfig     string  `protobuf:"bytes,7,opt,name=addons_config,proto3" json:"addons_config,omitempty"`
	State            string  `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
	Nodes            []*Node `protobuf:"bytes,9,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Logs             string  `protobuf:"bytes,10,opt,name=logs,proto3" json:"logs,omitempty"`
	IsCurrentCluster bool    `protobuf:"varint,11,opt,name=is_current_cluster,proto3" json:"is_current_cluster,omitempty"` // 是否是当前集群状态
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{4}
}

func (x *Cluster) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetServerVersion() string {
	if x != nil {
		return x.ServerVersion
	}
	return ""
}

func (x *Cluster) GetApiServerAddress() string {
	if x != nil {
		return x.ApiServerAddress
	}
	return ""
}

func (x *Cluster) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Cluster) GetAddons() string {
	if x != nil {
		return x.Addons
	}
	return ""
}

func (x *Cluster) GetAddonsConfig() string {
	if x != nil {
		return x.AddonsConfig
	}
	return ""
}

func (x *Cluster) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Cluster) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Cluster) GetLogs() string {
	if x != nil {
		return x.Logs
	}
	return ""
}

func (x *Cluster) GetIsCurrentCluster() bool {
	if x != nil {
		return x.IsCurrentCluster
	}
	return false
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels      string `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
	Annotations string `protobuf:"bytes,4,opt,name=annotations,proto3" json:"annotations,omitempty"`
	OsImage     string `protobuf:"bytes,5,opt,name=os_image,proto3" json:"os_image,omitempty"`
	Kernel      string `protobuf:"bytes,6,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Container   string `protobuf:"bytes,7,opt,name=container,proto3" json:"container,omitempty"`
	Kubelet     string `protobuf:"bytes,8,opt,name=kubelet,proto3" json:"kubelet,omitempty"`
	KubeProxy   string `protobuf:"bytes,9,opt,name=kube_proxy,proto3" json:"kube_proxy,omitempty"`
	InternalIp  string `protobuf:"bytes,10,opt,name=internal_ip,proto3" json:"internal_ip,omitempty"`
	ExternalIp  string `protobuf:"bytes,11,opt,name=external_ip,proto3" json:"external_ip,omitempty"`
	User        string `protobuf:"bytes,12,opt,name=user,proto3" json:"user,omitempty"`
	Role        string `protobuf:"bytes,15,opt,name=role,proto3" json:"role,omitempty"`
	State       string `protobuf:"bytes,16,opt,name=state,proto3" json:"state,omitempty"`
	ClusterId   int64  `protobuf:"varint,17,opt,name=cluster_id,proto3" json:"cluster_id,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{5}
}

func (x *Node) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

func (x *Node) GetAnnotations() string {
	if x != nil {
		return x.Annotations
	}
	return ""
}

func (x *Node) GetOsImage() string {
	if x != nil {
		return x.OsImage
	}
	return ""
}

func (x *Node) GetKernel() string {
	if x != nil {
		return x.Kernel
	}
	return ""
}

func (x *Node) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *Node) GetKubelet() string {
	if x != nil {
		return x.Kubelet
	}
	return ""
}

func (x *Node) GetKubeProxy() string {
	if x != nil {
		return x.KubeProxy
	}
	return ""
}

func (x *Node) GetInternalIp() string {
	if x != nil {
		return x.InternalIp
	}
	return ""
}

func (x *Node) GetExternalIp() string {
	if x != nil {
		return x.ExternalIp
	}
	return ""
}

func (x *Node) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Node) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Node) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Node) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

type CheckBostionHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arch                               string `protobuf:"bytes,1,opt,name=arch,proto3" json:"arch,omitempty"`
	OceanVersion                       string `protobuf:"bytes,2,opt,name=ocean_version,proto3" json:"ocean_version,omitempty"`
	ShipVersion                        string `protobuf:"bytes,3,opt,name=ship_version,proto3" json:"ship_version,omitempty"`
	OceanDataTarGzPackagePath          string `protobuf:"bytes,4,opt,name=ocean_data_tar_gz_package_path,proto3" json:"ocean_data_tar_gz_package_path,omitempty"`
	OceanDataTarGzPackageSha256SumPath string `protobuf:"bytes,5,opt,name=ocean_data_tar_gz_package_sha256sum_path,proto3" json:"ocean_data_tar_gz_package_sha256sum_path,omitempty"`
	OceanPath                          string `protobuf:"bytes,6,opt,name=ocean_path,proto3" json:"ocean_path,omitempty"`
	ShipPath                           string `protobuf:"bytes,7,opt,name=ship_path,proto3" json:"ship_path,omitempty"`
}

func (x *CheckBostionHostRequest) Reset() {
	*x = CheckBostionHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckBostionHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBostionHostRequest) ProtoMessage() {}

func (x *CheckBostionHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBostionHostRequest.ProtoReflect.Descriptor instead.
func (*CheckBostionHostRequest) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{6}
}

func (x *CheckBostionHostRequest) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *CheckBostionHostRequest) GetOceanVersion() string {
	if x != nil {
		return x.OceanVersion
	}
	return ""
}

func (x *CheckBostionHostRequest) GetShipVersion() string {
	if x != nil {
		return x.ShipVersion
	}
	return ""
}

func (x *CheckBostionHostRequest) GetOceanDataTarGzPackagePath() string {
	if x != nil {
		return x.OceanDataTarGzPackagePath
	}
	return ""
}

func (x *CheckBostionHostRequest) GetOceanDataTarGzPackageSha256SumPath() string {
	if x != nil {
		return x.OceanDataTarGzPackageSha256SumPath
	}
	return ""
}

func (x *CheckBostionHostRequest) GetOceanPath() string {
	if x != nil {
		return x.OceanPath
	}
	return ""
}

func (x *CheckBostionHostRequest) GetShipPath() string {
	if x != nil {
		return x.ShipPath
	}
	return ""
}

type ClusterLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId int64 `protobuf:"varint,1,opt,name=cluster_id,proto3" json:"cluster_id,omitempty"`
}

func (x *ClusterLogsRequest) Reset() {
	*x = ClusterLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterLogsRequest) ProtoMessage() {}

func (x *ClusterLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterLogsRequest.ProtoReflect.Descriptor instead.
func (*ClusterLogsRequest) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{7}
}

func (x *ClusterLogsRequest) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

type ClusterLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs string `protobuf:"bytes,1,opt,name=logs,proto3" json:"logs,omitempty"`
}

func (x *ClusterLogsResponse) Reset() {
	*x = ClusterLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterLogsResponse) ProtoMessage() {}

func (x *ClusterLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterLogsResponse.ProtoReflect.Descriptor instead.
func (*ClusterLogsResponse) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{8}
}

func (x *ClusterLogsResponse) GetLogs() string {
	if x != nil {
		return x.Logs
	}
	return ""
}

var File_api_cluster_v1alpha1_message_proto protoreflect.FileDescriptor

var file_api_cluster_v1alpha1_message_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x81, 0x02, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41,
	0x72, 0x67, 0x73, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x4e, 0x6f,
	0x64, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x09, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x22, 0x44, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0xe3, 0x02, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e,
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x70, 0x69, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x12, 0x24,
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x12,
	0x69, 0x73, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x92, 0x03, 0x0a,
	0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x73, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x73, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69,
	0x70, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x22, 0xd9, 0x02, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x68, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x1e, 0x6f,
	0x63, 0x65, 0x61, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x7a,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x1e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x74, 0x61, 0x72, 0x5f, 0x67, 0x7a, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x5a, 0x0a, 0x28, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x7a, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x73, 0x75, 0x6d, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x28, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x7a, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x73, 0x75, 0x6d, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1e, 0x0a, 0x0a, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x22, 0x34, 0x0a,
	0x12, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x1f,
	0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cluster_v1alpha1_message_proto_rawDescOnce sync.Once
	file_api_cluster_v1alpha1_message_proto_rawDescData = file_api_cluster_v1alpha1_message_proto_rawDesc
)

func file_api_cluster_v1alpha1_message_proto_rawDescGZIP() []byte {
	file_api_cluster_v1alpha1_message_proto_rawDescOnce.Do(func() {
		file_api_cluster_v1alpha1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cluster_v1alpha1_message_proto_rawDescData)
	})
	return file_api_cluster_v1alpha1_message_proto_rawDescData
}

var file_api_cluster_v1alpha1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_cluster_v1alpha1_message_proto_goTypes = []any{
	(*ClusterArgs)(nil),             // 0: cluster.v1alpha1.ClusterArgs
	(*NodeArgs)(nil),                // 1: cluster.v1alpha1.NodeArgs
	(*ClusterID)(nil),               // 2: cluster.v1alpha1.ClusterID
	(*ClusterList)(nil),             // 3: cluster.v1alpha1.ClusterList
	(*Cluster)(nil),                 // 4: cluster.v1alpha1.Cluster
	(*Node)(nil),                    // 5: cluster.v1alpha1.Node
	(*CheckBostionHostRequest)(nil), // 6: cluster.v1alpha1.CheckBostionHostRequest
	(*ClusterLogsRequest)(nil),      // 7: cluster.v1alpha1.ClusterLogsRequest
	(*ClusterLogsResponse)(nil),     // 8: cluster.v1alpha1.ClusterLogsResponse
}
var file_api_cluster_v1alpha1_message_proto_depIdxs = []int32{
	1, // 0: cluster.v1alpha1.ClusterArgs.nodes:type_name -> cluster.v1alpha1.NodeArgs
	4, // 1: cluster.v1alpha1.ClusterList.clusters:type_name -> cluster.v1alpha1.Cluster
	5, // 2: cluster.v1alpha1.Cluster.nodes:type_name -> cluster.v1alpha1.Node
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_cluster_v1alpha1_message_proto_init() }
func file_api_cluster_v1alpha1_message_proto_init() {
	if File_api_cluster_v1alpha1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_cluster_v1alpha1_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ClusterArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cluster_v1alpha1_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NodeArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cluster_v1alpha1_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ClusterID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cluster_v1alpha1_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ClusterList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cluster_v1alpha1_message_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Cluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cluster_v1alpha1_message_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cluster_v1alpha1_message_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CheckBostionHostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cluster_v1alpha1_message_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ClusterLogsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cluster_v1alpha1_message_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ClusterLogsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cluster_v1alpha1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_cluster_v1alpha1_message_proto_goTypes,
		DependencyIndexes: file_api_cluster_v1alpha1_message_proto_depIdxs,
		MessageInfos:      file_api_cluster_v1alpha1_message_proto_msgTypes,
	}.Build()
	File_api_cluster_v1alpha1_message_proto = out.File
	file_api_cluster_v1alpha1_message_proto_rawDesc = nil
	file_api_cluster_v1alpha1_message_proto_goTypes = nil
	file_api_cluster_v1alpha1_message_proto_depIdxs = nil
}
