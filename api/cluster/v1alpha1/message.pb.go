// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
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
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterID) ProtoMessage() {}

func (x *ClusterID) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ClusterID.ProtoReflect.Descriptor instead.
func (*ClusterID) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{0}
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
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterList) ProtoMessage() {}

func (x *ClusterList) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ClusterList.ProtoReflect.Descriptor instead.
func (*ClusterList) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{1}
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
	State            string  `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Nodes            []*Node `protobuf:"bytes,8,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Logs             string  `protobuf:"bytes,9,opt,name=logs,proto3" json:"logs,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{2}
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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels       string `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
	Annotations  string `protobuf:"bytes,4,opt,name=annotations,proto3" json:"annotations,omitempty"`
	OsImage      string `protobuf:"bytes,5,opt,name=os_image,proto3" json:"os_image,omitempty"`
	Kernel       string `protobuf:"bytes,6,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Container    string `protobuf:"bytes,7,opt,name=container,proto3" json:"container,omitempty"`
	Kubelet      string `protobuf:"bytes,8,opt,name=kubelet,proto3" json:"kubelet,omitempty"`
	KubeProxy    string `protobuf:"bytes,9,opt,name=kube_proxy,proto3" json:"kube_proxy,omitempty"`
	InternalIp   string `protobuf:"bytes,10,opt,name=internal_ip,proto3" json:"internal_ip,omitempty"`
	ExternalIp   string `protobuf:"bytes,11,opt,name=external_ip,proto3" json:"external_ip,omitempty"`
	User         string `protobuf:"bytes,12,opt,name=user,proto3" json:"user,omitempty"`
	Password     string `protobuf:"bytes,13,opt,name=password,proto3" json:"password,omitempty"`
	SudoPassword string `protobuf:"bytes,14,opt,name=sudo_password,proto3" json:"sudo_password,omitempty"`
	Role         string `protobuf:"bytes,15,opt,name=role,proto3" json:"role,omitempty"`
	State        string `protobuf:"bytes,16,opt,name=state,proto3" json:"state,omitempty"`
	ClusterId    int64  `protobuf:"varint,17,opt,name=cluster_id,proto3" json:"cluster_id,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cluster_v1alpha1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_api_cluster_v1alpha1_message_proto_rawDescGZIP(), []int{3}
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

func (x *Node) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Node) GetSudoPassword() string {
	if x != nil {
		return x.SudoPassword
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

var File_api_cluster_v1alpha1_message_proto protoreflect.FileDescriptor

var file_api_cluster_v1alpha1_message_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x34, 0x0a, 0x09, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x0b,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x22, 0x8d, 0x02, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x2c, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x22, 0xd4, 0x03, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x73, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x73, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6b,
	0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x75,
	0x62, 0x65, 0x6c, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x69, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x75, 0x64,
	0x6f, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x75, 0x64, 0x6f, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_api_cluster_v1alpha1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_cluster_v1alpha1_message_proto_goTypes = []interface{}{
	(*ClusterID)(nil),   // 0: cluster.v1alpha1.ClusterID
	(*ClusterList)(nil), // 1: cluster.v1alpha1.ClusterList
	(*Cluster)(nil),     // 2: cluster.v1alpha1.Cluster
	(*Node)(nil),        // 3: cluster.v1alpha1.Node
}
var file_api_cluster_v1alpha1_message_proto_depIdxs = []int32{
	2, // 0: cluster.v1alpha1.ClusterList.clusters:type_name -> cluster.v1alpha1.Cluster
	3, // 1: cluster.v1alpha1.Cluster.nodes:type_name -> cluster.v1alpha1.Node
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_cluster_v1alpha1_message_proto_init() }
func file_api_cluster_v1alpha1_message_proto_init() {
	if File_api_cluster_v1alpha1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_cluster_v1alpha1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_cluster_v1alpha1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_cluster_v1alpha1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_cluster_v1alpha1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cluster_v1alpha1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
