// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: api/app/v1alpha1/app.proto

package v1alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe3, 0x03,
	0x0a, 0x0a, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x1a, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x70, 0x70, 0x73, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x5f, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x7d,
	0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x49, 0x44, 0x7d, 0x12, 0x4d, 0x0a, 0x04,
	0x53, 0x61, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x44, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x1a, 0x12, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x63, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x32, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x1a, 0x27, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x49, 0x44, 0x7d,
	0x12, 0x61, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x44, 0x1a,
	0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d,
	0x73, 0x67, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x2a, 0x27, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x44, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2f, 0x7b, 0x61, 0x70, 0x70,
	0x49, 0x44, 0x7d, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_app_v1alpha1_app_proto_goTypes = []interface{}{
	(*ClusterID)(nil), // 0: app.v1alpha1.ClusterID
	(*AppID)(nil),     // 1: app.v1alpha1.AppID
	(*App)(nil),       // 2: app.v1alpha1.App
	(*Apps)(nil),      // 3: app.v1alpha1.Apps
	(*Msg)(nil),       // 4: app.v1alpha1.Msg
}
var file_api_app_v1alpha1_app_proto_depIdxs = []int32{
	0, // 0: app.v1alpha1.AppService.GetApps:input_type -> app.v1alpha1.ClusterID
	1, // 1: app.v1alpha1.AppService.GetApp:input_type -> app.v1alpha1.AppID
	2, // 2: app.v1alpha1.AppService.Save:input_type -> app.v1alpha1.App
	1, // 3: app.v1alpha1.AppService.Apply:input_type -> app.v1alpha1.AppID
	1, // 4: app.v1alpha1.AppService.Delete:input_type -> app.v1alpha1.AppID
	3, // 5: app.v1alpha1.AppService.GetApps:output_type -> app.v1alpha1.Apps
	2, // 6: app.v1alpha1.AppService.GetApp:output_type -> app.v1alpha1.App
	1, // 7: app.v1alpha1.AppService.Save:output_type -> app.v1alpha1.AppID
	4, // 8: app.v1alpha1.AppService.Apply:output_type -> app.v1alpha1.Msg
	4, // 9: app.v1alpha1.AppService.Delete:output_type -> app.v1alpha1.Msg
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_app_v1alpha1_app_proto_init() }
func file_api_app_v1alpha1_app_proto_init() {
	if File_api_app_v1alpha1_app_proto != nil {
		return
	}
	file_api_app_v1alpha1_error_reason_proto_init()
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