// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: api/service/v1alpha1/message.proto

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

type ServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string    `protobuf:"bytes,1,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Page      int64     `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64     `protobuf:"varint,3,opt,name=pageSize,json=page_size,proto3" json:"pageSize,omitempty"`
	ProjectId int64     `protobuf:"varint,4,opt,name=projectId,json=project_id,proto3" json:"projectId,omitempty"`
	Id        int64     `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	WfArgs    string    `protobuf:"bytes,8,opt,name=wfArgs,json=wf_args,proto3" json:"wfArgs,omitempty"`
	Workflow  *Worklfow `protobuf:"bytes,9,opt,name=workflow,proto3" json:"workflow,omitempty"`
	WfType    string    `protobuf:"bytes,10,opt,name=wfType,json=wf_type,proto3" json:"wfType,omitempty"`
}

func (x *ServiceRequest) Reset() {
	*x = ServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRequest) ProtoMessage() {}

func (x *ServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRequest.ProtoReflect.Descriptor instead.
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ServiceRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ServiceRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *ServiceRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ServiceRequest) GetWfArgs() string {
	if x != nil {
		return x.WfArgs
	}
	return ""
}

func (x *ServiceRequest) GetWorkflow() *Worklfow {
	if x != nil {
		return x.Workflow
	}
	return nil
}

func (x *ServiceRequest) GetWfType() string {
	if x != nil {
		return x.WfType
	}
	return ""
}

type Services struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Service `protobuf:"bytes,1,rep,name=Services,json=services,proto3" json:"Services,omitempty"`
	Total    int64      `protobuf:"varint,2,opt,name=Total,json=total,proto3" json:"Total,omitempty"`
}

func (x *Services) Reset() {
	*x = Services{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Services) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Services) ProtoMessage() {}

func (x *Services) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Services.ProtoReflect.Descriptor instead.
func (*Services) Descriptor() ([]byte, []int) {
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{1}
}

func (x *Services) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Services) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64   `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	CodeRepo    string  `protobuf:"bytes,3,opt,name=CodeRepo,json=code_repo,proto3" json:"CodeRepo,omitempty"`
	Replicas    int32   `protobuf:"varint,7,opt,name=Replicas,json=replicas,proto3" json:"Replicas,omitempty"`
	CPU         float32 `protobuf:"fixed32,8,opt,name=CPU,json=cpu,proto3" json:"CPU,omitempty"`
	LimitCpu    float32 `protobuf:"fixed32,9,opt,name=LimitCpu,json=limit_cpu,proto3" json:"LimitCpu,omitempty"`
	GPU         float32 `protobuf:"fixed32,10,opt,name=GPU,json=gpu,proto3" json:"GPU,omitempty"`
	LimitGPU    float32 `protobuf:"fixed32,11,opt,name=LimitGPU,json=limit_gpu,proto3" json:"LimitGPU,omitempty"`
	Memory      float32 `protobuf:"fixed32,12,opt,name=Memory,json=memory,proto3" json:"Memory,omitempty"`
	LimitMemory float32 `protobuf:"fixed32,13,opt,name=LimitMemory,json=limit_memory,proto3" json:"LimitMemory,omitempty"`
	Disk        float32 `protobuf:"fixed32,14,opt,name=Disk,json=disk,proto3" json:"Disk,omitempty"`
	LimitDisk   float32 `protobuf:"fixed32,15,opt,name=LimitDisk,json=limit_disk,proto3" json:"LimitDisk,omitempty"`
	WorkflowID  int64   `protobuf:"varint,16,opt,name=WorkflowID,json=workflow_id,proto3" json:"WorkflowID,omitempty"`
	Ports       []*Port `protobuf:"bytes,17,rep,name=Ports,json=ports,proto3" json:"Ports,omitempty"`
	ProjectID   int64   `protobuf:"varint,20,opt,name=ProjectID,json=project_id,proto3" json:"ProjectID,omitempty"`
	Business    string  `protobuf:"bytes,21,opt,name=Business,json=business,proto3" json:"Business,omitempty"`
	Technology  string  `protobuf:"bytes,22,opt,name=Technology,json=technology,proto3" json:"Technology,omitempty"`
	ProjectName string  `protobuf:"bytes,23,opt,name=project_name,proto3" json:"project_name,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[2]
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
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{2}
}

func (x *Service) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetCodeRepo() string {
	if x != nil {
		return x.CodeRepo
	}
	return ""
}

func (x *Service) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *Service) GetCPU() float32 {
	if x != nil {
		return x.CPU
	}
	return 0
}

func (x *Service) GetLimitCpu() float32 {
	if x != nil {
		return x.LimitCpu
	}
	return 0
}

func (x *Service) GetGPU() float32 {
	if x != nil {
		return x.GPU
	}
	return 0
}

func (x *Service) GetLimitGPU() float32 {
	if x != nil {
		return x.LimitGPU
	}
	return 0
}

func (x *Service) GetMemory() float32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Service) GetLimitMemory() float32 {
	if x != nil {
		return x.LimitMemory
	}
	return 0
}

func (x *Service) GetDisk() float32 {
	if x != nil {
		return x.Disk
	}
	return 0
}

func (x *Service) GetLimitDisk() float32 {
	if x != nil {
		return x.LimitDisk
	}
	return 0
}

func (x *Service) GetWorkflowID() int64 {
	if x != nil {
		return x.WorkflowID
	}
	return 0
}

func (x *Service) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Service) GetProjectID() int64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *Service) GetBusiness() string {
	if x != nil {
		return x.Business
	}
	return ""
}

func (x *Service) GetTechnology() string {
	if x != nil {
		return x.Technology
	}
	return ""
}

func (x *Service) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int64  `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	IngressPath   string `protobuf:"bytes,2,opt,name=IngressPath,json=ingress_path,proto3" json:"IngressPath,omitempty"`
	ContainerPort int32  `protobuf:"varint,3,opt,name=ContainerPort,json=container_port,proto3" json:"ContainerPort,omitempty"`
	Protocol      string `protobuf:"bytes,4,opt,name=Protocol,json=protocol,proto3" json:"Protocol,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{3}
}

func (x *Port) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Port) GetIngressPath() string {
	if x != nil {
		return x.IngressPath
	}
	return ""
}

func (x *Port) GetContainerPort() int32 {
	if x != nil {
		return x.ContainerPort
	}
	return 0
}

func (x *Port) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

type CI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int64  `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Version      string `protobuf:"bytes,2,opt,name=Version,json=version,proto3" json:"Version,omitempty"`
	Branch       string `protobuf:"bytes,3,opt,name=Branch,json=branch,proto3" json:"Branch,omitempty"`
	Tag          string `protobuf:"bytes,4,opt,name=Tag,json=tag,proto3" json:"Tag,omitempty"`
	Args         string `protobuf:"bytes,5,opt,name=Args,json=args,proto3" json:"Args,omitempty"`
	Description  string `protobuf:"bytes,6,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	WorkflowName string `protobuf:"bytes,7,opt,name=WorkflowName,json=workflow_name,proto3" json:"WorkflowName,omitempty"`
	ServiceID    int64  `protobuf:"varint,8,opt,name=ServiceID,json=service_id,proto3" json:"ServiceID,omitempty"`
}

func (x *CI) Reset() {
	*x = CI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CI) ProtoMessage() {}

func (x *CI) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CI.ProtoReflect.Descriptor instead.
func (*CI) Descriptor() ([]byte, []int) {
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{4}
}

func (x *CI) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CI) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CI) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *CI) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *CI) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *CI) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CI) GetWorkflowName() string {
	if x != nil {
		return x.WorkflowName
	}
	return ""
}

func (x *CI) GetServiceID() int64 {
	if x != nil {
		return x.ServiceID
	}
	return 0
}

type CD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64 `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	ServiceID int64 `protobuf:"varint,2,opt,name=ServiceID,json=service_id,proto3" json:"ServiceID,omitempty"`
}

func (x *CD) Reset() {
	*x = CD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CD) ProtoMessage() {}

func (x *CD) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CD.ProtoReflect.Descriptor instead.
func (*CD) Descriptor() ([]byte, []int) {
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{5}
}

func (x *CD) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CD) GetServiceID() int64 {
	if x != nil {
		return x.ServiceID
	}
	return 0
}

type Worklfow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64         `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Name        string        `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Description string        `protobuf:"bytes,3,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	Steps       []*WfStep     `protobuf:"bytes,4,rep,name=Steps,json=steps,proto3" json:"Steps,omitempty"`
	Templates   []*WfTemplate `protobuf:"bytes,5,rep,name=Templates,json=templates,proto3" json:"Templates,omitempty"`
}

func (x *Worklfow) Reset() {
	*x = Worklfow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Worklfow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Worklfow) ProtoMessage() {}

func (x *Worklfow) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Worklfow.ProtoReflect.Descriptor instead.
func (*Worklfow) Descriptor() ([]byte, []int) {
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{6}
}

func (x *Worklfow) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Worklfow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Worklfow) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Worklfow) GetSteps() []*WfStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *Worklfow) GetTemplates() []*WfTemplate {
	if x != nil {
		return x.Templates
	}
	return nil
}

type WfStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tasks   []*WfTask `protobuf:"bytes,2,rep,name=Tasks,json=tasks,proto3" json:"Tasks,omitempty"`
	Default bool      `protobuf:"varint,3,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *WfStep) Reset() {
	*x = WfStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WfStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WfStep) ProtoMessage() {}

func (x *WfStep) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WfStep.ProtoReflect.Descriptor instead.
func (*WfStep) Descriptor() ([]byte, []int) {
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{7}
}

func (x *WfStep) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WfStep) GetTasks() []*WfTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *WfStep) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

type WfTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TemplateName string `protobuf:"bytes,2,opt,name=templateName,json=template_name,proto3" json:"templateName,omitempty"`
	Default      bool   `protobuf:"varint,3,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *WfTask) Reset() {
	*x = WfTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WfTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WfTask) ProtoMessage() {}

func (x *WfTask) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WfTask.ProtoReflect.Descriptor instead.
func (*WfTask) Descriptor() ([]byte, []int) {
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{8}
}

func (x *WfTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WfTask) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *WfTask) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

type WfTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image    string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Command  []string `protobuf:"bytes,3,rep,name=command,proto3" json:"command,omitempty"`
	Args     []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	Source   string   `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	IsScript bool     `protobuf:"varint,6,opt,name=IsScript,json=is_script,proto3" json:"IsScript,omitempty"`
}

func (x *WfTemplate) Reset() {
	*x = WfTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_v1alpha1_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WfTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WfTemplate) ProtoMessage() {}

func (x *WfTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_v1alpha1_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WfTemplate.ProtoReflect.Descriptor instead.
func (*WfTemplate) Descriptor() ([]byte, []int) {
	return file_api_service_v1alpha1_message_proto_rawDescGZIP(), []int{9}
}

func (x *WfTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WfTemplate) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *WfTemplate) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *WfTemplate) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *WfTemplate) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *WfTemplate) GetIsScript() bool {
	if x != nil {
		return x.IsScript
	}
	return false
}

var File_api_service_v1alpha1_message_proto protoreflect.FileDescriptor

var file_api_service_v1alpha1_message_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0xee, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x06, 0x77, 0x66, 0x41, 0x72, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77,
	0x66, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x66, 0x6f, 0x77, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x17,
	0x0a, 0x06, 0x77, 0x66, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x77, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x57, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x80, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x08, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x50, 0x55,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x1b, 0x0a, 0x08, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x43, 0x70, 0x75, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x70, 0x75, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x50, 0x55, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x67, 0x70, 0x75, 0x12, 0x1b, 0x0a, 0x08, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x47, 0x50, 0x55, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x5f, 0x67, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x21, 0x0a, 0x0b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x09, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44,
	0x69, 0x73, 0x6b, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x49, 0x44, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x25,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x22, 0xd2, 0x01, 0x0a, 0x02, 0x43, 0x49, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x61,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x41, 0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x02, 0x43, 0x44, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x09,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x22, 0xbc, 0x01, 0x0a, 0x08,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x66, 0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e,
	0x0a, 0x05, 0x53, 0x74, 0x65, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x57, 0x66, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x3a,
	0x0a, 0x09, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x66, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x06, 0x57, 0x66,
	0x53, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x66, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x22, 0x5b, 0x0a, 0x06, 0x57, 0x66, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22,
	0x99, 0x01, 0x0a, 0x0a, 0x57, 0x66, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x08, 0x49, 0x73, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x3b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_service_v1alpha1_message_proto_rawDescOnce sync.Once
	file_api_service_v1alpha1_message_proto_rawDescData = file_api_service_v1alpha1_message_proto_rawDesc
)

func file_api_service_v1alpha1_message_proto_rawDescGZIP() []byte {
	file_api_service_v1alpha1_message_proto_rawDescOnce.Do(func() {
		file_api_service_v1alpha1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_service_v1alpha1_message_proto_rawDescData)
	})
	return file_api_service_v1alpha1_message_proto_rawDescData
}

var file_api_service_v1alpha1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_service_v1alpha1_message_proto_goTypes = []interface{}{
	(*ServiceRequest)(nil), // 0: service.v1alpha1.ServiceRequest
	(*Services)(nil),       // 1: service.v1alpha1.Services
	(*Service)(nil),        // 2: service.v1alpha1.Service
	(*Port)(nil),           // 3: service.v1alpha1.Port
	(*CI)(nil),             // 4: service.v1alpha1.CI
	(*CD)(nil),             // 5: service.v1alpha1.CD
	(*Worklfow)(nil),       // 6: service.v1alpha1.Worklfow
	(*WfStep)(nil),         // 7: service.v1alpha1.WfStep
	(*WfTask)(nil),         // 8: service.v1alpha1.WfTask
	(*WfTemplate)(nil),     // 9: service.v1alpha1.WfTemplate
}
var file_api_service_v1alpha1_message_proto_depIdxs = []int32{
	6, // 0: service.v1alpha1.ServiceRequest.workflow:type_name -> service.v1alpha1.Worklfow
	2, // 1: service.v1alpha1.Services.Services:type_name -> service.v1alpha1.Service
	3, // 2: service.v1alpha1.Service.Ports:type_name -> service.v1alpha1.Port
	7, // 3: service.v1alpha1.Worklfow.Steps:type_name -> service.v1alpha1.WfStep
	9, // 4: service.v1alpha1.Worklfow.Templates:type_name -> service.v1alpha1.WfTemplate
	8, // 5: service.v1alpha1.WfStep.Tasks:type_name -> service.v1alpha1.WfTask
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_service_v1alpha1_message_proto_init() }
func file_api_service_v1alpha1_message_proto_init() {
	if File_api_service_v1alpha1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_service_v1alpha1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRequest); i {
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
		file_api_service_v1alpha1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Services); i {
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
		file_api_service_v1alpha1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_service_v1alpha1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_api_service_v1alpha1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CI); i {
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
		file_api_service_v1alpha1_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CD); i {
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
		file_api_service_v1alpha1_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Worklfow); i {
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
		file_api_service_v1alpha1_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WfStep); i {
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
		file_api_service_v1alpha1_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WfTask); i {
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
		file_api_service_v1alpha1_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WfTemplate); i {
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
			RawDescriptor: file_api_service_v1alpha1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_service_v1alpha1_message_proto_goTypes,
		DependencyIndexes: file_api_service_v1alpha1_message_proto_depIdxs,
		MessageInfos:      file_api_service_v1alpha1_message_proto_msgTypes,
	}.Build()
	File_api_service_v1alpha1_message_proto = out.File
	file_api_service_v1alpha1_message_proto_rawDesc = nil
	file_api_service_v1alpha1_message_proto_goTypes = nil
	file_api_service_v1alpha1_message_proto_depIdxs = nil
}
