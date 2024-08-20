// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.27.1
// source: api/app/v1alpha1/error_reason.proto

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

type ErrorReason int32

const (
	ErrorReason_SUCCEED ErrorReason = 0
	ErrorReason_FAILED  ErrorReason = 1
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "SUCCEED",
		1: "FAILED",
	}
	ErrorReason_value = map[string]int32{
		"SUCCEED": 0,
		"FAILED":  1,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_app_v1alpha1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_app_v1alpha1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_app_v1alpha1_error_reason_proto_rawDescGZIP(), []int{0}
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason  ErrorReason `protobuf:"varint,1,opt,name=Reason,proto3,enum=app.v1alpha1.ErrorReason" json:"Reason,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_v1alpha1_error_reason_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_v1alpha1_error_reason_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_api_app_v1alpha1_error_reason_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetReason() ErrorReason {
	if x != nil {
		return x.Reason
	}
	return ErrorReason_SUCCEED
}

func (x *Msg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_app_v1alpha1_error_reason_proto protoreflect.FileDescriptor

var file_api_app_v1alpha1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x22, 0x52, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x31, 0x0a, 0x06, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x26, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x42,
	0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x3b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_app_v1alpha1_error_reason_proto_rawDescOnce sync.Once
	file_api_app_v1alpha1_error_reason_proto_rawDescData = file_api_app_v1alpha1_error_reason_proto_rawDesc
)

func file_api_app_v1alpha1_error_reason_proto_rawDescGZIP() []byte {
	file_api_app_v1alpha1_error_reason_proto_rawDescOnce.Do(func() {
		file_api_app_v1alpha1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_app_v1alpha1_error_reason_proto_rawDescData)
	})
	return file_api_app_v1alpha1_error_reason_proto_rawDescData
}

var file_api_app_v1alpha1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_app_v1alpha1_error_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_app_v1alpha1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: app.v1alpha1.ErrorReason
	(*Msg)(nil),      // 1: app.v1alpha1.Msg
}
var file_api_app_v1alpha1_error_reason_proto_depIdxs = []int32{
	0, // 0: app.v1alpha1.Msg.Reason:type_name -> app.v1alpha1.ErrorReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_app_v1alpha1_error_reason_proto_init() }
func file_api_app_v1alpha1_error_reason_proto_init() {
	if File_api_app_v1alpha1_error_reason_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_app_v1alpha1_error_reason_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
			RawDescriptor: file_api_app_v1alpha1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_app_v1alpha1_error_reason_proto_goTypes,
		DependencyIndexes: file_api_app_v1alpha1_error_reason_proto_depIdxs,
		EnumInfos:         file_api_app_v1alpha1_error_reason_proto_enumTypes,
		MessageInfos:      file_api_app_v1alpha1_error_reason_proto_msgTypes,
	}.Build()
	File_api_app_v1alpha1_error_reason_proto = out.File
	file_api_app_v1alpha1_error_reason_proto_rawDesc = nil
	file_api_app_v1alpha1_error_reason_proto_goTypes = nil
	file_api_app_v1alpha1_error_reason_proto_depIdxs = nil
}
