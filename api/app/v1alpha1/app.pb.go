// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/app/v1alpha1/app.proto

package v1alpha1

import (
	common "github.com/f-rambo/ocean/api/common"
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

var File_api_app_v1alpha1_app_proto protoreflect.FileDescriptor

var file_api_app_v1alpha1_app_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe0, 0x0f, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d,
	0x73, 0x67, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x64, 0x0a, 0x09, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x70, 0x70, 0x12,
	0x1f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x70, 0x70, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x49, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65,
	0x12, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x70, 0x70, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x73, 0x67,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x61, 0x76, 0x65, 0x12, 0x49, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x70, 0x70, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x12, 0x53,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x73, 0x67,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x12, 0x56, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x73, 0x67,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x65, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d,
	0x73, 0x67, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x61, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41,
	0x70, 0x70, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x12, 0x64, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70,
	0x1a, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x65, 0x0a, 0x09, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x12, 0x71, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x64, 0x41, 0x70, 0x70, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x25,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x70, 0x41, 0x70, 0x70,
	0x12, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x5e, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x41, 0x70,
	0x70, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x83, 0x01, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x55, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x19,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01,
	0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x66, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70,
	0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x57, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x12,
	0x1c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x70, 0x70, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x18, 0x2a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x69, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x73, 0x42, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x1c, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x48, 0x65,
	0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x12, 0x70, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x1c, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x48, 0x65,
	0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x22, 0x29, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_app_v1alpha1_app_proto_goTypes = []any{
	(*emptypb.Empty)(nil),      // 0: google.protobuf.Empty
	(*FileUploadRequest)(nil),  // 1: app.v1alpha1.FileUploadRequest
	(*App)(nil),                // 2: app.v1alpha1.App
	(*AppReq)(nil),             // 3: app.v1alpha1.AppReq
	(*AppType)(nil),            // 4: app.v1alpha1.AppType
	(*AppTypeReq)(nil),         // 5: app.v1alpha1.AppTypeReq
	(*DeployAppReq)(nil),       // 6: app.v1alpha1.DeployAppReq
	(*DeployApp)(nil),          // 7: app.v1alpha1.DeployApp
	(*AppHelmRepo)(nil),        // 8: app.v1alpha1.AppHelmRepo
	(*AppHelmRepoReq)(nil),     // 9: app.v1alpha1.AppHelmRepoReq
	(*common.Msg)(nil),         // 10: common.Msg
	(*AppList)(nil),            // 11: app.v1alpha1.AppList
	(*AppTypeList)(nil),        // 12: app.v1alpha1.AppTypeList
	(*DeployAppList)(nil),      // 13: app.v1alpha1.DeployAppList
	(*DeployAppResources)(nil), // 14: app.v1alpha1.DeployAppResources
	(*AppHelmRepoList)(nil),    // 15: app.v1alpha1.AppHelmRepoList
}
var file_api_app_v1alpha1_app_proto_depIdxs = []int32{
	0,  // 0: app.v1alpha1.AppInterface.Ping:input_type -> google.protobuf.Empty
	1,  // 1: app.v1alpha1.AppInterface.UploadApp:input_type -> app.v1alpha1.FileUploadRequest
	2,  // 2: app.v1alpha1.AppInterface.Save:input_type -> app.v1alpha1.App
	3,  // 3: app.v1alpha1.AppInterface.Get:input_type -> app.v1alpha1.AppReq
	3,  // 4: app.v1alpha1.AppInterface.List:input_type -> app.v1alpha1.AppReq
	3,  // 5: app.v1alpha1.AppInterface.Delete:input_type -> app.v1alpha1.AppReq
	4,  // 6: app.v1alpha1.AppInterface.CreateAppType:input_type -> app.v1alpha1.AppType
	0,  // 7: app.v1alpha1.AppInterface.ListAppType:input_type -> google.protobuf.Empty
	5,  // 8: app.v1alpha1.AppInterface.DeleteAppType:input_type -> app.v1alpha1.AppTypeReq
	6,  // 9: app.v1alpha1.AppInterface.AppTest:input_type -> app.v1alpha1.DeployAppReq
	7,  // 10: app.v1alpha1.AppInterface.GetAppDeployed:input_type -> app.v1alpha1.DeployApp
	6,  // 11: app.v1alpha1.AppInterface.DeployApp:input_type -> app.v1alpha1.DeployAppReq
	6,  // 12: app.v1alpha1.AppInterface.ListDeployedApp:input_type -> app.v1alpha1.DeployAppReq
	6,  // 13: app.v1alpha1.AppInterface.StopApp:input_type -> app.v1alpha1.DeployAppReq
	6,  // 14: app.v1alpha1.AppInterface.DeleteDeployedApp:input_type -> app.v1alpha1.DeployAppReq
	6,  // 15: app.v1alpha1.AppInterface.GetDeployedAppResources:input_type -> app.v1alpha1.DeployAppReq
	8,  // 16: app.v1alpha1.AppInterface.SaveRepo:input_type -> app.v1alpha1.AppHelmRepo
	0,  // 17: app.v1alpha1.AppInterface.ListRepo:input_type -> google.protobuf.Empty
	9,  // 18: app.v1alpha1.AppInterface.DeleteRepo:input_type -> app.v1alpha1.AppHelmRepoReq
	9,  // 19: app.v1alpha1.AppInterface.GetAppsByRepo:input_type -> app.v1alpha1.AppHelmRepoReq
	9,  // 20: app.v1alpha1.AppInterface.GetAppDetailByRepo:input_type -> app.v1alpha1.AppHelmRepoReq
	10, // 21: app.v1alpha1.AppInterface.Ping:output_type -> common.Msg
	2,  // 22: app.v1alpha1.AppInterface.UploadApp:output_type -> app.v1alpha1.App
	10, // 23: app.v1alpha1.AppInterface.Save:output_type -> common.Msg
	2,  // 24: app.v1alpha1.AppInterface.Get:output_type -> app.v1alpha1.App
	11, // 25: app.v1alpha1.AppInterface.List:output_type -> app.v1alpha1.AppList
	10, // 26: app.v1alpha1.AppInterface.Delete:output_type -> common.Msg
	10, // 27: app.v1alpha1.AppInterface.CreateAppType:output_type -> common.Msg
	12, // 28: app.v1alpha1.AppInterface.ListAppType:output_type -> app.v1alpha1.AppTypeList
	10, // 29: app.v1alpha1.AppInterface.DeleteAppType:output_type -> common.Msg
	7,  // 30: app.v1alpha1.AppInterface.AppTest:output_type -> app.v1alpha1.DeployApp
	7,  // 31: app.v1alpha1.AppInterface.GetAppDeployed:output_type -> app.v1alpha1.DeployApp
	7,  // 32: app.v1alpha1.AppInterface.DeployApp:output_type -> app.v1alpha1.DeployApp
	13, // 33: app.v1alpha1.AppInterface.ListDeployedApp:output_type -> app.v1alpha1.DeployAppList
	10, // 34: app.v1alpha1.AppInterface.StopApp:output_type -> common.Msg
	10, // 35: app.v1alpha1.AppInterface.DeleteDeployedApp:output_type -> common.Msg
	14, // 36: app.v1alpha1.AppInterface.GetDeployedAppResources:output_type -> app.v1alpha1.DeployAppResources
	10, // 37: app.v1alpha1.AppInterface.SaveRepo:output_type -> common.Msg
	15, // 38: app.v1alpha1.AppInterface.ListRepo:output_type -> app.v1alpha1.AppHelmRepoList
	10, // 39: app.v1alpha1.AppInterface.DeleteRepo:output_type -> common.Msg
	11, // 40: app.v1alpha1.AppInterface.GetAppsByRepo:output_type -> app.v1alpha1.AppList
	2,  // 41: app.v1alpha1.AppInterface.GetAppDetailByRepo:output_type -> app.v1alpha1.App
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_app_v1alpha1_app_proto_init() }
func file_api_app_v1alpha1_app_proto_init() {
	if File_api_app_v1alpha1_app_proto != nil {
		return
	}
	file_api_app_v1alpha1_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_app_v1alpha1_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_app_v1alpha1_app_proto_goTypes,
		DependencyIndexes: file_api_app_v1alpha1_app_proto_depIdxs,
	}.Build()
	File_api_app_v1alpha1_app_proto = out.File
	file_api_app_v1alpha1_app_proto_rawDesc = nil
	file_api_app_v1alpha1_app_proto_goTypes = nil
	file_api_app_v1alpha1_app_proto_depIdxs = nil
}
