// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: internal/repository/clusterruntime/api/project/message.proto

package project

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

type CreateNamespaceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *CreateNamespaceReq) Reset() {
	*x = CreateNamespaceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_repository_clusterruntime_api_project_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceReq) ProtoMessage() {}

func (x *CreateNamespaceReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_repository_clusterruntime_api_project_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceReq.ProtoReflect.Descriptor instead.
func (*CreateNamespaceReq) Descriptor() ([]byte, []int) {
	return file_internal_repository_clusterruntime_api_project_message_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNamespaceReq) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type Namesapces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *Namesapces) Reset() {
	*x = Namesapces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_repository_clusterruntime_api_project_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namesapces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namesapces) ProtoMessage() {}

func (x *Namesapces) ProtoReflect() protoreflect.Message {
	mi := &file_internal_repository_clusterruntime_api_project_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namesapces.ProtoReflect.Descriptor instead.
func (*Namesapces) Descriptor() ([]byte, []int) {
	return file_internal_repository_clusterruntime_api_project_message_proto_rawDescGZIP(), []int{1}
}

func (x *Namesapces) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

var File_internal_repository_clusterruntime_api_project_message_proto protoreflect.FileDescriptor

var file_internal_repository_clusterruntime_api_project_message_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x2c,
	0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x61, 0x70, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x42, 0x52, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x2d, 0x72, 0x61, 0x6d,
	0x62, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6f, 0x70, 0x69, 0x6c, 0x6f, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_repository_clusterruntime_api_project_message_proto_rawDescOnce sync.Once
	file_internal_repository_clusterruntime_api_project_message_proto_rawDescData = file_internal_repository_clusterruntime_api_project_message_proto_rawDesc
)

func file_internal_repository_clusterruntime_api_project_message_proto_rawDescGZIP() []byte {
	file_internal_repository_clusterruntime_api_project_message_proto_rawDescOnce.Do(func() {
		file_internal_repository_clusterruntime_api_project_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_repository_clusterruntime_api_project_message_proto_rawDescData)
	})
	return file_internal_repository_clusterruntime_api_project_message_proto_rawDescData
}

var file_internal_repository_clusterruntime_api_project_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_repository_clusterruntime_api_project_message_proto_goTypes = []any{
	(*CreateNamespaceReq)(nil), // 0: clusterruntime.api.project.CreateNamespaceReq
	(*Namesapces)(nil),         // 1: clusterruntime.api.project.Namesapces
}
var file_internal_repository_clusterruntime_api_project_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_repository_clusterruntime_api_project_message_proto_init() }
func file_internal_repository_clusterruntime_api_project_message_proto_init() {
	if File_internal_repository_clusterruntime_api_project_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_repository_clusterruntime_api_project_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateNamespaceReq); i {
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
		file_internal_repository_clusterruntime_api_project_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Namesapces); i {
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
			RawDescriptor: file_internal_repository_clusterruntime_api_project_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_repository_clusterruntime_api_project_message_proto_goTypes,
		DependencyIndexes: file_internal_repository_clusterruntime_api_project_message_proto_depIdxs,
		MessageInfos:      file_internal_repository_clusterruntime_api_project_message_proto_msgTypes,
	}.Build()
	File_internal_repository_clusterruntime_api_project_message_proto = out.File
	file_internal_repository_clusterruntime_api_project_message_proto_rawDesc = nil
	file_internal_repository_clusterruntime_api_project_message_proto_goTypes = nil
	file_internal_repository_clusterruntime_api_project_message_proto_depIdxs = nil
}
