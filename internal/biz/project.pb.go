// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: internal/biz/project.proto

package biz

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"gorm.io/gorm"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProjectStatus int32

const (
	ProjectStatus_PROJECT_INIT    ProjectStatus = 0
	ProjectStatus_PROJECT_RUNNING ProjectStatus = 1
	ProjectStatus_PROJECT_STOPPED ProjectStatus = 2
)

// Enum value maps for ProjectStatus.
var (
	ProjectStatus_name = map[int32]string{
		0: "PROJECT_INIT",
		1: "PROJECT_RUNNING",
		2: "PROJECT_STOPPED",
	}
	ProjectStatus_value = map[string]int32{
		"PROJECT_INIT":    0,
		"PROJECT_RUNNING": 1,
		"PROJECT_STOPPED": 2,
	}
)

func (x ProjectStatus) Enum() *ProjectStatus {
	p := new(ProjectStatus)
	*p = x
	return p
}

func (x ProjectStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_biz_project_proto_enumTypes[0].Descriptor()
}

func (ProjectStatus) Type() protoreflect.EnumType {
	return &file_internal_biz_project_proto_enumTypes[0]
}

func (x ProjectStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectStatus.Descriptor instead.
func (ProjectStatus) EnumDescriptor() ([]byte, []int) {
	return file_internal_biz_project_proto_rawDescGZIP(), []int{0}
}

type Technology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @goimport: "gorm.io/gorm"
	// @gofield: gorm.Model
	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"column:id;primaryKey;AUTO_INCREMENT"`                                       // @gotags: gorm:"column:id;primaryKey;AUTO_INCREMENT"
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"column:name; default:''; NOT NULL"`                                      // @gotags: gorm:"column:name; default:''; NOT NULL"
	BusinessId int64  `protobuf:"varint,3,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty" gorm:"column:business_id; default:0; NOT NULL"` // @gotags: gorm:"column:business_id; default:0; NOT NULL"
	gorm.Model
}

func (x *Technology) Reset() {
	*x = Technology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_biz_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Technology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Technology) ProtoMessage() {}

func (x *Technology) ProtoReflect() protoreflect.Message {
	mi := &file_internal_biz_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Technology.ProtoReflect.Descriptor instead.
func (*Technology) Descriptor() ([]byte, []int) {
	return file_internal_biz_project_proto_rawDescGZIP(), []int{0}
}

func (x *Technology) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Technology) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Technology) GetBusinessId() int64 {
	if x != nil {
		return x.BusinessId
	}
	return 0
}

type Business struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @goimport: "gorm.io/gorm"
	// @gofield: gorm.Model
	Id          int64         `json:"id,omitempty" gorm:"column:id;primaryKey;AUTO_INCREMENT" protobuf:"varint,1,opt,name=id,proto3"`                                   // @gotags: gorm:"column:id;primaryKey;AUTO_INCREMENT"
	Name        string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"column:name; default:''; NOT NULL"`                                  // @gotags: gorm:"column:name; default:''; NOT NULL"
	ProjectId   int64         `json:"project_id,omitempty" gorm:"column:project_id; default:0; NOT NULL" protobuf:"varint,3,opt,name=project_id,json=projectId,proto3"` // @gotags: gorm:"column:project_id; default:0; NOT NULL"
	Technologys []*Technology `protobuf:"bytes,4,rep,name=technologys,proto3" json:"technologys,omitempty" gorm:"-"`
	gorm.Model                // @gotags: gorm:"-"
}

func (x *Business) Reset() {
	*x = Business{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_biz_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Business) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Business) ProtoMessage() {}

func (x *Business) ProtoReflect() protoreflect.Message {
	mi := &file_internal_biz_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Business.ProtoReflect.Descriptor instead.
func (*Business) Descriptor() ([]byte, []int) {
	return file_internal_biz_project_proto_rawDescGZIP(), []int{1}
}

func (x *Business) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Business) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Business) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *Business) GetTechnologys() []*Technology {
	if x != nil {
		return x.Technologys
	}
	return nil
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @goimport: "gorm.io/gorm"
	// @gofield: gorm.Model
	Id                 int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"column:id;primaryKey;AUTO_INCREMENT"`                                                                       // @gotags: gorm:"column:id;primaryKey;AUTO_INCREMENT"
	Name               string        `json:"name,omitempty" gorm:"column:name; default:''; NOT NULL" protobuf:"bytes,2,opt,name=name,proto3"`                                                                      // @gotags: gorm:"column:name; default:''; NOT NULL"
	Description        string        `gorm:"column:namespace; default:''; NOT NULL" protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`                                                   // @gotags: gorm:"column:namespace; default:''; NOT NULL"
	ClusterId          int64         `protobuf:"varint,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" gorm:"column:cluster_id; default:0; NOT NULL"`                                     // @gotags: gorm:"column:cluster_id; default:0; NOT NULL"
	Namespace          string        `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty" gorm:"column:namespace; default:''; NOT NULL"`                                                       // @gotags: gorm:"column:namespace; default:''; NOT NULL"
	Status             ProjectStatus `protobuf:"varint,6,opt,name=status,proto3,enum=biz.project.ProjectStatus" json:"status,omitempty" gorm:"column:status; default:0; NOT NULL"`                                 // @gotags: gorm:"column:status; default:0; NOT NULL"
	Business           []*Business   `protobuf:"bytes,7,rep,name=business,proto3" json:"business,omitempty" gorm:"-"`                                                                                              // @gotags: gorm:"-"
	BusinessTechnology string        `json:"business_technology,omitempty" gorm:"column:business_technology; default:''; NOT NULL" protobuf:"bytes,8,opt,name=business_technology,json=businessTechnology,proto3"` // @gotags: gorm:"column:business_technology; default:''; NOT NULL"
	gorm.Model
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_biz_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_internal_biz_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_internal_biz_project_proto_rawDescGZIP(), []int{2}
}

func (x *Project) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *Project) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Project) GetStatus() ProjectStatus {
	if x != nil {
		return x.Status
	}
	return ProjectStatus_PROJECT_INIT
}

func (x *Project) GetBusiness() []*Business {
	if x != nil {
		return x.Business
	}
	return nil
}

func (x *Project) GetBusinessTechnology() string {
	if x != nil {
		return x.BusinessTechnology
	}
	return ""
}

var File_internal_biz_project_proto protoreflect.FileDescriptor

var file_internal_biz_project_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x69,
	0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x51, 0x0a, 0x0a, 0x54, 0x65, 0x63,
	0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0x88, 0x01, 0x0a,
	0x08, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0b,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x0b, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x73, 0x22, 0xa4, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a,
	0x13, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2a, 0x4b,
	0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x55, 0x4e,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43,
	0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x02, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x2d, 0x72, 0x61, 0x6d, 0x62,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6f, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x69, 0x7a, 0x3b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_biz_project_proto_rawDescOnce sync.Once
	file_internal_biz_project_proto_rawDescData = file_internal_biz_project_proto_rawDesc
)

func file_internal_biz_project_proto_rawDescGZIP() []byte {
	file_internal_biz_project_proto_rawDescOnce.Do(func() {
		file_internal_biz_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_biz_project_proto_rawDescData)
	})
	return file_internal_biz_project_proto_rawDescData
}

var file_internal_biz_project_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_biz_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_biz_project_proto_goTypes = []any{
	(ProjectStatus)(0), // 0: biz.project.ProjectStatus
	(*Technology)(nil), // 1: biz.project.Technology
	(*Business)(nil),   // 2: biz.project.Business
	(*Project)(nil),    // 3: biz.project.Project
}
var file_internal_biz_project_proto_depIdxs = []int32{
	1, // 0: biz.project.Business.technologys:type_name -> biz.project.Technology
	0, // 1: biz.project.Project.status:type_name -> biz.project.ProjectStatus
	2, // 2: biz.project.Project.business:type_name -> biz.project.Business
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_biz_project_proto_init() }
func file_internal_biz_project_proto_init() {
	if File_internal_biz_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_biz_project_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Technology); i {
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
		file_internal_biz_project_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Business); i {
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
		file_internal_biz_project_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Project); i {
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
			RawDescriptor: file_internal_biz_project_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_biz_project_proto_goTypes,
		DependencyIndexes: file_internal_biz_project_proto_depIdxs,
		EnumInfos:         file_internal_biz_project_proto_enumTypes,
		MessageInfos:      file_internal_biz_project_proto_msgTypes,
	}.Build()
	File_internal_biz_project_proto = out.File
	file_internal_biz_project_proto_rawDesc = nil
	file_internal_biz_project_proto_goTypes = nil
	file_internal_biz_project_proto_depIdxs = nil
}
