// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/cloud/v1alpha1/cloud.proto

package v1alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_cloud_v1alpha1_cloud_proto protoreflect.FileDescriptor

var file_api_cloud_v1alpha1_cloud_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x87, 0x09, 0x0a, 0x0e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x84,
	0x01, 0x0a, 0x19, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x4b, 0x75, 0x62, 0x65, 0x61, 0x64,
	0x6d, 0x4b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x43, 0x72, 0x69, 0x4f, 0x12, 0x15, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35,
	0x3a, 0x01, 0x2a, 0x22, 0x30, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x5f, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x64, 0x6d, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74,
	0x5f, 0x63, 0x72, 0x69, 0x6f, 0x12, 0xa5, 0x01, 0x0a, 0x28, 0x41, 0x64, 0x64, 0x4b, 0x75, 0x62,
	0x65, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x75, 0x62, 0x65, 0x61, 0x64, 0x6d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x15, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x4d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x47, 0x3a, 0x01, 0x2a, 0x22, 0x42, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x64, 0x64, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6b,
	0x75, 0x62, 0x65, 0x61, 0x64, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x66, 0x0a,
	0x0b, 0x4b, 0x75, 0x62, 0x65, 0x61, 0x64, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x64, 0x6d,
	0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x66, 0x0a, 0x0b, 0x4b, 0x75, 0x62, 0x65, 0x61, 0x64, 0x6d,
	0x49, 0x6e, 0x69, 0x74, 0x12, 0x15, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x1a, 0x13, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x64, 0x6d, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x68, 0x0a,
	0x0c, 0x4b, 0x75, 0x62, 0x65, 0x61, 0x64, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x15, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x64,
	0x6d, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x6c, 0x0a, 0x0e, 0x4b, 0x75, 0x62, 0x65, 0x61,
	0x64, 0x6d, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a,
	0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x64, 0x6d, 0x5f, 0x75, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x75, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x70, 0x76, 0x34, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a,
	0x01, 0x2a, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x70, 0x76, 0x34, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x63, 0x0a, 0x09,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x77, 0x61, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01,
	0x2a, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x73, 0x77, 0x61,
	0x70, 0x12, 0x6b, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x42, 0x1d,
	0x5a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_cloud_v1alpha1_cloud_proto_goTypes = []any{
	(*emptypb.Empty)(nil), // 0: google.protobuf.Empty
	(*Cloud)(nil),         // 1: cloud.v1alpha1.Cloud
	(*Msg)(nil),           // 2: cloud.v1alpha1.Msg
}
var file_api_cloud_v1alpha1_cloud_proto_depIdxs = []int32{
	0,  // 0: cloud.v1alpha1.CloudInterface.Ping:input_type -> google.protobuf.Empty
	1,  // 1: cloud.v1alpha1.CloudInterface.InstallKubeadmKubeletCriO:input_type -> cloud.v1alpha1.Cloud
	1,  // 2: cloud.v1alpha1.CloudInterface.AddKubeletServiceAndSettingKubeadmConfig:input_type -> cloud.v1alpha1.Cloud
	1,  // 3: cloud.v1alpha1.CloudInterface.KubeadmJoin:input_type -> cloud.v1alpha1.Cloud
	1,  // 4: cloud.v1alpha1.CloudInterface.KubeadmInit:input_type -> cloud.v1alpha1.Cloud
	1,  // 5: cloud.v1alpha1.CloudInterface.KubeadmReset:input_type -> cloud.v1alpha1.Cloud
	1,  // 6: cloud.v1alpha1.CloudInterface.KubeadmUpgrade:input_type -> cloud.v1alpha1.Cloud
	0,  // 7: cloud.v1alpha1.CloudInterface.SetingIpv4Forward:input_type -> google.protobuf.Empty
	0,  // 8: cloud.v1alpha1.CloudInterface.CloseSwap:input_type -> google.protobuf.Empty
	0,  // 9: cloud.v1alpha1.CloudInterface.CloseFirewall:input_type -> google.protobuf.Empty
	2,  // 10: cloud.v1alpha1.CloudInterface.Ping:output_type -> cloud.v1alpha1.Msg
	2,  // 11: cloud.v1alpha1.CloudInterface.InstallKubeadmKubeletCriO:output_type -> cloud.v1alpha1.Msg
	2,  // 12: cloud.v1alpha1.CloudInterface.AddKubeletServiceAndSettingKubeadmConfig:output_type -> cloud.v1alpha1.Msg
	2,  // 13: cloud.v1alpha1.CloudInterface.KubeadmJoin:output_type -> cloud.v1alpha1.Msg
	2,  // 14: cloud.v1alpha1.CloudInterface.KubeadmInit:output_type -> cloud.v1alpha1.Msg
	2,  // 15: cloud.v1alpha1.CloudInterface.KubeadmReset:output_type -> cloud.v1alpha1.Msg
	2,  // 16: cloud.v1alpha1.CloudInterface.KubeadmUpgrade:output_type -> cloud.v1alpha1.Msg
	2,  // 17: cloud.v1alpha1.CloudInterface.SetingIpv4Forward:output_type -> cloud.v1alpha1.Msg
	2,  // 18: cloud.v1alpha1.CloudInterface.CloseSwap:output_type -> cloud.v1alpha1.Msg
	2,  // 19: cloud.v1alpha1.CloudInterface.CloseFirewall:output_type -> cloud.v1alpha1.Msg
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_cloud_v1alpha1_cloud_proto_init() }
func file_api_cloud_v1alpha1_cloud_proto_init() {
	if File_api_cloud_v1alpha1_cloud_proto != nil {
		return
	}
	file_api_cloud_v1alpha1_error_reason_proto_init()
	file_api_cloud_v1alpha1_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cloud_v1alpha1_cloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_cloud_v1alpha1_cloud_proto_goTypes,
		DependencyIndexes: file_api_cloud_v1alpha1_cloud_proto_depIdxs,
	}.Build()
	File_api_cloud_v1alpha1_cloud_proto = out.File
	file_api_cloud_v1alpha1_cloud_proto_rawDesc = nil
	file_api_cloud_v1alpha1_cloud_proto_goTypes = nil
	file_api_cloud_v1alpha1_cloud_proto_depIdxs = nil
}
