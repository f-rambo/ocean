// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: internal/conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Env int32

const (
	Env_EnvLocal       Env = 0
	Env_EnvBostionHost Env = 1
	Env_EnvCluster     Env = 2
)

// Enum value maps for Env.
var (
	Env_name = map[int32]string{
		0: "EnvLocal",
		1: "EnvBostionHost",
		2: "EnvCluster",
	}
	Env_value = map[string]int32{
		"EnvLocal":       0,
		"EnvBostionHost": 1,
		"EnvCluster":     2,
	}
)

func (x Env) Enum() *Env {
	p := new(Env)
	*p = x
	return p
}

func (x Env) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Env) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_conf_conf_proto_enumTypes[0].Descriptor()
}

func (Env) Type() protoreflect.EnumType {
	return &file_internal_conf_conf_proto_enumTypes[0]
}

func (x Env) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Env.Descriptor instead.
func (Env) EnumDescriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

type ClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Env  Env    `protobuf:"varint,1,opt,name=env,proto3,enum=Env" json:"env,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ClusterConfig) Reset() {
	*x = ClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfig) ProtoMessage() {}

func (x *ClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterConfig.ProtoReflect.Descriptor instead.
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterConfig) GetEnv() Env {
	if x != nil {
		return x.Env
	}
	return Env_EnvLocal
}

func (x *ClusterConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type HTTPServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout int64  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *HTTPServer) Reset() {
	*x = HTTPServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPServer) ProtoMessage() {}

func (x *HTTPServer) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPServer.ProtoReflect.Descriptor instead.
func (*HTTPServer) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *HTTPServer) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *HTTPServer) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *HTTPServer) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type GRPCServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout int64  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *GRPCServer) Reset() {
	*x = GRPCServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCServer) ProtoMessage() {}

func (x *GRPCServer) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCServer.ProtoReflect.Descriptor instead.
func (*GRPCServer) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *GRPCServer) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *GRPCServer) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *GRPCServer) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string      `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Http    *HTTPServer `protobuf:"bytes,3,opt,name=http,proto3" json:"http,omitempty"`
	Grpc    *GRPCServer `protobuf:"bytes,4,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Server) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Server) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Server) GetHttp() *HTTPServer {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *GRPCServer {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Port    int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Addr    string `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout int64  `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Service) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Service) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Service) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Host     string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port     int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Database) Reset() {
	*x = Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{5}
}

func (x *Database) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Database) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Database) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Database) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Database) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxSize    int32 `protobuf:"varint,1,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	MaxBackups int32 `protobuf:"varint,2,opt,name=max_backups,json=maxBackups,proto3" json:"max_backups,omitempty"`
	MaxAge     int32 `protobuf:"varint,3,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{6}
}

func (x *Log) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *Log) GetMaxBackups() int32 {
	if x != nil {
		return x.MaxBackups
	}
	return 0
}

func (x *Log) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exp           int32  `protobuf:"varint,1,opt,name=exp,proto3" json:"exp,omitempty"`
	Key           string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	AdminEmail    string `protobuf:"bytes,3,opt,name=admin_email,json=adminEmail,proto3" json:"admin_email,omitempty"`
	AdminPassword string `protobuf:"bytes,4,opt,name=admin_password,json=adminPassword,proto3" json:"admin_password,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{7}
}

func (x *Auth) GetExp() int32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *Auth) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Auth) GetAdminEmail() string {
	if x != nil {
		return x.AdminEmail
	}
	return ""
}

func (x *Auth) GetAdminPassword() string {
	if x != nil {
		return x.AdminPassword
	}
	return ""
}

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster  *ClusterConfig `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Server   *Server        `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Services []*Service     `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
	Data     *Database      `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Log      *Log           `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
	Auth     *Auth          `protobuf:"bytes,6,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{8}
}

func (x *Bootstrap) GetCluster() *ClusterConfig {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Bootstrap) GetData() *Database {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *Bootstrap) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

var File_internal_conf_conf_proto protoreflect.FileDescriptor

var file_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0d,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a,
	0x03, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x04, 0x2e, 0x45, 0x6e, 0x76,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x54, 0x0a,
	0x0a, 0x48, 0x54, 0x54, 0x50, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x22, 0x54, 0x0a, 0x0a, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x78, 0x0a, 0x06, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x04, 0x68, 0x74,
	0x74, 0x70, 0x12, 0x1f, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x22, 0x79, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x86,
	0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5a, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x6d, 0x61, 0x78, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61,
	0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78,
	0x41, 0x67, 0x65, 0x22, 0x72, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x78, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xce, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x19, 0x0a,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x2a, 0x37, 0x0a, 0x03, 0x45, 0x6e, 0x76, 0x12,
	0x0c, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x45, 0x6e, 0x76, 0x42, 0x6f, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x6e, 0x76, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x10,
	0x02, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x2d, 0x72, 0x61, 0x6d, 0x62, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x63, 0x6f,
	0x70, 0x69, 0x6c, 0x6f, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_conf_conf_proto_rawDescOnce sync.Once
	file_internal_conf_conf_proto_rawDescData = file_internal_conf_conf_proto_rawDesc
)

func file_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_conf_conf_proto_rawDescData)
	})
	return file_internal_conf_conf_proto_rawDescData
}

var file_internal_conf_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_conf_conf_proto_goTypes = []any{
	(Env)(0),              // 0: Env
	(*ClusterConfig)(nil), // 1: ClusterConfig
	(*HTTPServer)(nil),    // 2: HTTPServer
	(*GRPCServer)(nil),    // 3: GRPCServer
	(*Server)(nil),        // 4: Server
	(*Service)(nil),       // 5: Service
	(*Database)(nil),      // 6: Database
	(*Log)(nil),           // 7: Log
	(*Auth)(nil),          // 8: Auth
	(*Bootstrap)(nil),     // 9: Bootstrap
}
var file_internal_conf_conf_proto_depIdxs = []int32{
	0, // 0: ClusterConfig.env:type_name -> Env
	2, // 1: Server.http:type_name -> HTTPServer
	3, // 2: Server.grpc:type_name -> GRPCServer
	1, // 3: Bootstrap.cluster:type_name -> ClusterConfig
	4, // 4: Bootstrap.server:type_name -> Server
	5, // 5: Bootstrap.services:type_name -> Service
	6, // 6: Bootstrap.data:type_name -> Database
	7, // 7: Bootstrap.log:type_name -> Log
	8, // 8: Bootstrap.auth:type_name -> Auth
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_internal_conf_conf_proto_init() }
func file_internal_conf_conf_proto_init() {
	if File_internal_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_conf_conf_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ClusterConfig); i {
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
		file_internal_conf_conf_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HTTPServer); i {
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
		file_internal_conf_conf_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GRPCServer); i {
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
		file_internal_conf_conf_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Server); i {
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
		file_internal_conf_conf_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Service); i {
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
		file_internal_conf_conf_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Database); i {
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
		file_internal_conf_conf_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Log); i {
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
		file_internal_conf_conf_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Auth); i {
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
		file_internal_conf_conf_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Bootstrap); i {
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
			RawDescriptor: file_internal_conf_conf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_internal_conf_conf_proto_depIdxs,
		EnumInfos:         file_internal_conf_conf_proto_enumTypes,
		MessageInfos:      file_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_internal_conf_conf_proto = out.File
	file_internal_conf_conf_proto_rawDesc = nil
	file_internal_conf_conf_proto_goTypes = nil
	file_internal_conf_conf_proto_depIdxs = nil
}
