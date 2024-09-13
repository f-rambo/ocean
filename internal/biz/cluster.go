package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/f-rambo/ocean/internal/conf"
	"github.com/f-rambo/ocean/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
)

const (
	ClusterPackageName = "cluster"
)

type Cluster struct {
	ID                   int64         `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Name                 string        `json:"name" gorm:"column:name; default:''; NOT NULL"` // *
	CloudID              string        `json:"cloud_id" gorm:"column:cloud_id; default:''; NOT NULL"`
	Connections          string        `json:"connections" gorm:"column:connections; default:''; NOT NULL"`
	CertificateAuthority string        `json:"certificate_authority" gorm:"column:certificate_authority; default:''; NOT NULL"`
	Version              string        `json:"version" gorm:"column:version; default:''; NOT NULL"`
	ApiServerAddress     string        `json:"api_server_address" gorm:"column:api_server_address; default:''; NOT NULL"`
	Config               string        `json:"config" gorm:"column:config; default:''; NOT NULL;"`
	Addons               string        `json:"addons" gorm:"column:addons; default:''; NOT NULL;"`
	AddonsConfig         string        `json:"addons_config" gorm:"column:addons_config; default:''; NOT NULL;"`
	Status               ClusterStatus `json:"status" gorm:"column:status; default:0; NOT NULL;"`
	Type                 ClusterType   `json:"type" gorm:"column:type; default:''; NOT NULL;"` //*  aws google cloud azure alicloud local
	KubeConfig           string        `json:"kube_config" gorm:"column:kube_config; default:''; NOT NULL; type:json"`
	KeyPair              string        `json:"key_pair" gorm:"column:key_pair; default:''; NOT NULL;"`
	PublicKey            string        `json:"public_key" gorm:"column:public_key; default:''; NOT NULL;"` // *
	PrivateKey           string        `json:"private_key" gorm:"column:private_key; default:''; NOT NULL;"`
	Region               string        `json:"region" gorm:"column:region; default:''; NOT NULL;"` // *
	VpcID                string        `json:"vpc_id" gorm:"column:vpc_id; default:''; NOT NULL;"`
	VpcCidr              string        `json:"vpc_cidr" gorm:"column:vpc_cidr; default:''; NOT NULL;"`
	EipID                string        `json:"eip_id" gorm:"column:eip_id; default:''; NOT NULL;"`
	NatGatewayID         string        `json:"nat_gateway_id" gorm:"column:nat_gateway_id; default:''; NOT NULL;"`
	ResourceGroupID      string        `json:"resource_group_id" gorm:"column:resource_group_id; default:''; NOT NULL;"`
	SecurityGroupIDs     string        `json:"security_group_ids" gorm:"column:security_group_ids; default:''; NOT NULL;"`
	ExternalIP           string        `json:"external_ip" gorm:"column:external_ip; default:''; NOT NULL;"`
	AccessID             string        `json:"access_id" gorm:"column:access_id; default:''; NOT NULL;"`   // *
	AccessKey            string        `json:"access_key" gorm:"column:access_key; default:''; NOT NULL;"` // *
	LoadBalancerID       string        `json:"load_balancer_id" gorm:"column:load_balancer_id; default:''; NOT NULL;"`
	BostionHost          *BostionHost  `json:"bostion_host" gorm:"-"`
	Nodes                []*Node       `json:"nodes" gorm:"-"`
	NodeGroups           []*NodeGroup  `json:"node_groups" gorm:"-"`
	log                  *log.Helper   `json:"-"`
	gorm.Model
}

type NodeGroup struct {
	ID               string        `json:"id" gorm:"column:id;primaryKey; NOT NULL"`
	Name             string        `json:"name" gorm:"column:name; default:''; NOT NULL"`
	CloudNoodGroupID string        `json:"cloud_nodegroup_id" gorm:"column:cloud_nodegroup_id; default:''; NOT NULL"`
	Type             NodeGroupType `json:"type" gorm:"column:type; default:''; NOT NULL;"`
	InstanceType     string        `json:"instance_type" gorm:"column:instance_type; default:''; NOT NULL"`
	Image            string        `json:"image" gorm:"column:image; default:''; NOT NULL"`
	OS               string        `json:"os" gorm:"column:os; default:''; NOT NULL"`        //*
	ARCH             string        `json:"arch" gorm:"column:arch; default:''; NOT NULL"`    //*
	CPU              int32         `json:"cpu" gorm:"column:cpu; default:0; NOT NULL"`       //*
	Memory           float64       `json:"memory" gorm:"column:memory; default:0; NOT NULL"` //*
	GPU              int32         `json:"gpu" gorm:"column:gpu; default:0; NOT NULL"`       //*
	NodeInitScript   string        `json:"cloud_init_script" gorm:"column:cloud_init_script; default:''; NOT NULL"`
	MinSize          int32         `json:"min_size" gorm:"column:min_size; default:0; NOT NULL"`
	MaxSize          int32         `json:"max_size" gorm:"column:max_size; default:0; NOT NULL"`
	TargetSize       int32         `json:"target_size" gorm:"column:target_size; default:0; NOT NULL"`
	SystemDisk       int32         `json:"system_disk" gorm:"column:system_disk; default:0; NOT NULL"`
	DataDisk         int32         `json:"data_disk" gorm:"column:data_disk; default:0; NOT NULL"`
	ClusterID        int64         `json:"cluster_id" gorm:"column:cluster_id; default:0; NOT NULL"`
}

type Node struct {
	ID                      int64      `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	InstanceID              string     `json:"instance_id" gorm:"column:instance_id; default:''; NOT NULL"`
	Name                    string     `json:"name" gorm:"column:name; default:''; NOT NULL"`
	Labels                  string     `json:"labels" gorm:"column:labels; default:''; NOT NULL"` // map[string]string json
	Kernel                  string     `json:"kernel" gorm:"column:kernel; default:''; NOT NULL"`
	ContainerRuntime        string     `json:"container_runtime" gorm:"column:container_runtime; default:''; NOT NULL"`
	Kubelet                 string     `json:"kubelet" gorm:"column:kubelet; default:''; NOT NULL"`
	KubeProxy               string     `json:"kube_proxy" gorm:"column:kube_proxy; default:''; NOT NULL"`
	SshPort                 int32      `json:"ssh_port" gorm:"column:ssh_port; default:0; NOT NULL"`
	InternalIP              string     `json:"internal_ip" gorm:"column:internal_ip; default:''; NOT NULL"`
	ExternalIP              string     `json:"external_ip" gorm:"column:external_ip; default:''; NOT NULL"`
	User                    string     `json:"user" gorm:"column:user; default:''; NOT NULL"`
	Role                    NodeRole   `json:"role" gorm:"column:role; default:''; NOT NULL;"` // master worker edge
	Status                  NodeStatus `json:"status" gorm:"column:status; default:0; NOT NULL;"`
	ErrorInfo               string     `json:"error_info" gorm:"column:error_info; default:''; NOT NULL"`
	Zone                    string     `json:"zone" gorm:"column:zone; default:''; NOT NULL"`
	SubnetId                string     `json:"subnet_id" gorm:"column:subnet_id; default:''; NOT NULL"`
	SubnetCidr              string     `json:"subnet_cidr" gorm:"column:subnet_cidr; default:''; NOT NULL"`
	PublicKey               string     `json:"public_key" gorm:"column:public_key; default:''; NOT NULL;"`
	GpuSpec                 string     `json:"gpu_spec" gorm:"column:gpu_spec; default:''; NOT NULL"`
	SystemDisk              int32      `json:"system_disk" gorm:"column:system_disk; default:0; NOT NULL"`
	DataDisk                int32      `json:"data_disk" gorm:"column:data_disk; default:0; NOT NULL"`
	NodePrice               float64    `json:"node_price" gorm:"column:node_price; default:0; NOT NULL;"` // 节点价格
	PodPrice                float64    `json:"pod_price" gorm:"column:pod_price; default:0; NOT NULL;"`   // 节点上pod的价格
	InternetMaxBandwidthOut int32      `json:"internet_max_bandwidth_out" gorm:"column:internet_max_bandwidth_out; default:0; NOT NULL"`
	ClusterID               int64      `json:"cluster_id" gorm:"column:cluster_id; default:0; NOT NULL"`
	NodeGroupID             string     `json:"node_group_id" gorm:"column:node_group_id; default:''; NOT NULL"`
	gorm.Model
}

type BostionHost struct {
	ID           int64   `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	InstanceType string  `json:"instance_type" gorm:"column:instance_type; default:''; NOT NULL"`
	InstanceID   string  `json:"instance_id" gorm:"column:instance_id; default:''; NOT NULL"`
	User         string  `json:"user" gorm:"column:user; default:''; NOT NULL"`
	ImageID      string  `json:"image_id" gorm:"column:image_id; default:''; NOT NULL"`
	Image        string  `json:"image" gorm:"column:image; default:''; NOT NULL"`
	OS           string  `json:"os" gorm:"column:os; default:''; NOT NULL"`
	ARCH         string  `json:"arch" gorm:"column:arch; default:''; NOT NULL"`
	Hostname     string  `json:"hostname" gorm:"column:hostname; default:''; NOT NULL"`
	ExternalIP   string  `json:"external_ip" gorm:"column:external_ip; default:''; NOT NULL"`
	InternalIP   string  `json:"internal_ip" gorm:"column:internal_ip; default:''; NOT NULL"`
	SshPort      int32   `json:"ssh_port" gorm:"column:ssh_port; default:0; NOT NULL"`
	PrivateIP    string  `json:"private_ip" gorm:"column:private_ip; default:''; NOT NULL"`
	ClusterID    int64   `json:"cluster_id" gorm:"column:cluster_id; default:0; NOT NULL"`
	CPU          int32   `json:"cpu" gorm:"column:cpu; default:0; NOT NULL"`
	Memory       float64 `json:"memory" gorm:"column:memory; default:0; NOT NULL"`
	gorm.Model
}

type ClusterRepo interface {
	Save(context.Context, *Cluster) error
	Get(context.Context, int64) (*Cluster, error)
	GetByName(context.Context, string) (*Cluster, error)
	List(context.Context, *Cluster) ([]*Cluster, error)
	Delete(context.Context, int64) error
	Put(ctx context.Context, cluster *Cluster) error
	Watch(ctx context.Context) (*Cluster, error)
}

type ClusterInfrastructure interface {
	Start(context.Context, *Cluster) error
	Stop(context.Context, *Cluster) error
	GetRegions(context.Context, *Cluster) ([]string, error)
	MigrateToBostionHost(context.Context, *Cluster) error
	Install(context.Context, *Cluster) error
	UnInstall(context.Context, *Cluster) error
	AddNodes(context.Context, *Cluster, []*Node) error
	RemoveNodes(context.Context, *Cluster, []*Node) error
}

type ClusterRuntime interface {
	CurrentCluster(context.Context, *Cluster) error
}

var ErrClusterNotFound error = errors.New("cluster not found")

type ClusterType string

const (
	ClusterTypeLocal       ClusterType = "local"
	ClusterTypeAWSEc2      ClusterType = "aws_ec2"
	ClusterTypeAWSEks      ClusterType = "aws_eks"
	ClusterTypeAliCloudEcs ClusterType = "alicloud_ecs"
	ClusterTypeAliCloudAks ClusterType = "alicloud_aks"
)

func (c ClusterType) String() string {
	return string(c)
}

func (c ClusterType) IsCloud() bool {
	return c != ClusterTypeLocal
}

func (c ClusterType) IsIntegratedCloud() bool {
	return c == ClusterTypeAWSEks || c == ClusterTypeAliCloudAks
}

func ClusterTypes() []ClusterType {
	return []ClusterType{
		ClusterTypeLocal,
		ClusterTypeAWSEc2,
		ClusterTypeAWSEks,
		ClusterTypeAliCloudEcs,
		ClusterTypeAliCloudAks,
	}
}

type ClusterStatus uint8

func (s ClusterStatus) Uint8() uint8 {
	return uint8(s)
}

const (
	ClusterStatusUnspecified ClusterStatus = 0
	ClusterStatusRunning     ClusterStatus = 1
	ClusterStatusDeleted     ClusterStatus = 2
	ClusterStatucCreating    ClusterStatus = 3
)

var (
	ClusterStatusName = map[uint8]string{
		0: "unspecified",
		1: "running",
		2: "deleted",
		3: "creating",
	}
	ClusterStatusValue = map[string]uint8{
		"unspecified": 0,
		"running":     1,
		"deleted":     2,
		"creating":    3,
	}
)

type NodeGroups []*NodeGroup

func (n NodeGroups) Len() int {
	return len(n)
}

func (n NodeGroups) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// 从小到大排序
func (n NodeGroups) Less(i, j int) bool {
	if n[i] == nil || n[j] == nil {
		return false
	}
	if n[i].Memory == n[j].Memory {
		return n[i].CPU < n[j].CPU
	}
	return n[i].Memory < n[j].Memory
}

type NodeGroupType string

func (n NodeGroupType) String() string {
	return string(n)
}

const (
	NodeGroupTypeNormal          NodeGroupType = "normal"
	NodeGroupTypeHighComputation NodeGroupType = "highComputation"
	NodeGroupTypeGPUAcceleraterd NodeGroupType = "gpuAcceleraterd"
	NodeGroupTypeHighMemory      NodeGroupType = "highMemory"
	NodeGroupTypeLargeHardDisk   NodeGroupType = "largeHardDisk"
)

type NodeSize int32

func (n NodeSize) Int32() int32 {
	return int32(n)
}

func (n NodeSize) Int() int {
	return int(n)
}

const (
	NodeMinSize NodeSize = 5
)

type NodeRole string

const (
	NodeRoleMaster NodeRole = "master"
	NodeRoleWorker NodeRole = "worker"
	NodeRoleEdge   NodeRole = "edge"
)

func (n NodeRole) String() string {
	return string(n)
}

type NodeStatus uint8

const (
	// an Unspecified instanceState means the actual instance status is undefined (nil).
	NodeStatusUnspecified NodeStatus = 0
	// NodeStatusRunning means instance is running.
	NodeStatusRunning NodeStatus = 1
	// NodeStatusCreating means instance is being created.
	NodeStatusCreating NodeStatus = 2
	// NodeStatusDeleting means instance is being deleted.
	NodeStatusDeleting NodeStatus = 3
)

func (s NodeStatus) Uint8() uint8 {
	return uint8(s)
}

var (
	NodeStatusName = map[uint8]string{
		0: "unspecified",
		1: "instanceRunning",
		2: "instanceCreating",
		3: "instanceDeleting",
	}
	NodeStatusValue = map[string]uint8{
		"unspecified":      0,
		"instanceRunning":  1,
		"instanceCreating": 2,
		"instanceDeleting": 3,
	}
)

func (c *Cluster) generateNodeLables(nodeGroup *NodeGroup) string {
	lableMap := make(map[string]string)
	lableMap["cluster"] = c.Name
	lableMap["cluster_type"] = c.Type.String()
	lableMap["region"] = c.Region
	lableMap["nodegroup"] = nodeGroup.Name
	lableMap["nodegroup_type"] = nodeGroup.Type.String()
	lableMap["instance_type"] = nodeGroup.InstanceType
	lablebytes, _ := json.Marshal(lableMap)
	return string(lablebytes)
}

func (c *Cluster) logPath() string {
	return fmt.Sprintf("logs/cluster-%d.log", c.ID)
}

func (c *Cluster) Write(content []byte) (int, error) {
	if c.log == nil {
		logger := log.With(log.NewStdLogger(&lumberjack.Logger{
			Filename:  c.logPath(),
			MaxAge:    int(7),
			LocalTime: true,
		}), "ts", log.DefaultTimestamp)
		c.log = log.NewHelper(logger)
	}
	c.log.Info(string(content))
	return len(content), nil
}

func isClusterEmpty(c *Cluster) bool {
	if c == nil {
		return true
	}
	if c.ID == 0 {
		return true
	}
	return false
}

func (c *Cluster) IsDeleteed() bool {
	return c.DeletedAt.Valid
}

func (c *Cluster) GetNode(nodeId int64) *Node {
	for _, node := range c.Nodes {
		if node.ID == nodeId {
			return node
		}
	}
	return nil
}

func (ng *NodeGroup) SetTargetSize(size int32) {
	ng.TargetSize = size
}

func (c *Cluster) GetType() ClusterType {
	return ClusterType(c.Type)
}

func (n *Node) GetStatus() NodeStatus {
	return NodeStatus(n.Status)
}

func (c *Cluster) GetNodeGroup(nodeGroupId string) *NodeGroup {
	for _, nodeGroup := range c.NodeGroups {
		if nodeGroup.ID == nodeGroupId {
			return nodeGroup
		}
	}
	return nil
}

func (c *Cluster) GenerateNodeGroupName(nodeGroup *NodeGroup) string {
	return strings.Join([]string{c.Name, nodeGroup.Name, nodeGroup.Type.String()}, "-")
}

func (c *Cluster) NewNodeGroup() *NodeGroup {
	return &NodeGroup{
		ID: utils.GetRandomString(),
	}
}

type ClusterUsecase struct {
	clusterRepo           ClusterRepo
	clusterInfrastructure ClusterInfrastructure
	clusterRuntime        ClusterRuntime
	conf                  *conf.Bootstrap
	log                   *log.Helper
}

func NewClusterUseCase(conf *conf.Bootstrap, clusterRepo ClusterRepo, clusterInfrastructure ClusterInfrastructure, clusterRuntime ClusterRuntime, logger log.Logger) *ClusterUsecase {
	c := &ClusterUsecase{
		clusterRepo:           clusterRepo,
		clusterInfrastructure: clusterInfrastructure,
		clusterRuntime:        clusterRuntime,
		conf:                  conf,
		log:                   log.NewHelper(logger),
	}
	return c
}

func (uc *ClusterUsecase) Get(ctx context.Context, id int64) (*Cluster, error) {
	cluster, err := uc.clusterRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func (uc *ClusterUsecase) List(ctx context.Context) ([]*Cluster, error) {
	return uc.clusterRepo.List(ctx, nil)
}

func (uc *ClusterUsecase) Delete(ctx context.Context, clusterID int64) error {
	cluster, err := uc.clusterRepo.Get(ctx, clusterID)
	if err != nil {
		return err
	}
	if isClusterEmpty(cluster) {
		return nil
	}
	err = uc.clusterRepo.Delete(ctx, clusterID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ClusterUsecase) Save(ctx context.Context, cluster *Cluster) error {
	data, err := uc.clusterRepo.GetByName(ctx, cluster.Name)
	if err != nil {
		return err
	}
	if !isClusterEmpty(cluster) && cluster.ID != data.ID {
		return errors.New("cluster name already exists")
	}
	for _, node := range cluster.Nodes {
		if node.Name == "" {
			node.Name = fmt.Sprintf("%s-%s", cluster.Name, utils.GetRandomString())
		}
	}
	err = uc.clusterRepo.Save(ctx, cluster)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ClusterUsecase) GetRegions(ctx context.Context, cluster *Cluster) ([]string, error) {
	return uc.clusterInfrastructure.GetRegions(ctx, cluster)
}

// 获取当前集群最新信息
func (uc *ClusterUsecase) GetCurrentCluster(ctx context.Context) (*Cluster, error) {
	cluster := &Cluster{}
	err := uc.clusterRuntime.CurrentCluster(ctx, cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// 根据nodegroup增加节点
func (uc *ClusterUsecase) NodeGroupIncreaseSize(ctx context.Context, cluster *Cluster, nodeGroup *NodeGroup, size int32) error {
	for i := 0; i < int(size); i++ {
		node := &Node{
			Name:        fmt.Sprintf("%s-%s", cluster.Name, utils.GetRandomString()),
			Role:        NodeRoleWorker,
			Status:      NodeStatusCreating,
			ClusterID:   cluster.ID,
			NodeGroupID: nodeGroup.ID,
		}
		cluster.Nodes = append(cluster.Nodes, node)
	}
	return uc.Apply(ctx, cluster)
}

// 删除节点
func (uc *ClusterUsecase) DeleteNodes(ctx context.Context, cluster *Cluster, nodes []*Node) error {
	for _, node := range nodes {
		for i, n := range cluster.Nodes {
			if n.ID == node.ID {
				cluster.Nodes = append(cluster.Nodes[:i], cluster.Nodes[i+1:]...)
				break
			}
		}
	}
	return uc.Apply(ctx, cluster)
}

// 预测一个节点配置，也就是根据当前节点组目前还可以配置的节点
func (uc *ClusterUsecase) NodeGroupTemplateNodeInfo(ctx context.Context, cluster *Cluster, nodeGroup *NodeGroup) (*Node, error) {
	return &Node{
		Name:        fmt.Sprintf("%s-%s", cluster.Name, utils.GetRandomString()),
		Role:        NodeRoleWorker,
		Status:      NodeStatusCreating,
		ClusterID:   cluster.ID,
		NodeGroupID: nodeGroup.ID,
		Labels:      cluster.generateNodeLables(nodeGroup),
	}, nil
}

// 在云提供商销毁前清理打开的资源，例如协程等
func (uc *ClusterUsecase) Cleanup(ctx context.Context) error {
	return nil
}

// 在每个主循环前调用，用于动态更新云提供商状态
func (uc *ClusterUsecase) Refresh(ctx context.Context) error {
	// 获取当前集群状态更新状态
	cluster := &Cluster{}
	err := uc.clusterRuntime.CurrentCluster(ctx, cluster)
	if err != nil {
		return err
	}
	cluster, err = uc.clusterRepo.GetByName(ctx, cluster.Name)
	if err != nil {
		return err
	}
	for _, v := range cluster.Nodes {
		for _, currentNode := range cluster.Nodes {
			if v.Name == currentNode.Name {
				v.Status = currentNode.Status
				break
			}
		}
	}
	err = uc.Save(ctx, cluster)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ClusterUsecase) Apply(ctx context.Context, cluster *Cluster) error {
	return uc.clusterRepo.Put(ctx, cluster)
}

func (uc *ClusterUsecase) Watch(ctx context.Context) (*Cluster, error) {
	return uc.clusterRepo.Watch(ctx)
}

func (uc *ClusterUsecase) Reconcile(ctx context.Context, cluster *Cluster) (err error) {
	defer func() {
		if err != nil {
			return
		}
		err = uc.clusterRepo.Save(ctx, cluster)
	}()
	if cluster.IsDeleteed() {
		err = uc.clusterInfrastructure.UnInstall(ctx, cluster)
		if err != nil {
			return err
		}
		err = uc.clusterInfrastructure.Stop(ctx, cluster)
		if err != nil {
			return err
		}
		return nil
	}
	err = uc.clusterRuntime.CurrentCluster(ctx, cluster)
	if errors.Is(err, ErrClusterNotFound) {
		err = uc.clusterInfrastructure.Start(ctx, cluster)
		if err != nil {
			return err
		}
		if uc.conf.Server.GetEnv() == conf.EnvLocal {
			err = uc.clusterInfrastructure.MigrateToBostionHost(ctx, cluster)
			if err != nil {
				return err
			}
			return nil
		}
		err = uc.clusterInfrastructure.Install(ctx, cluster)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	err = uc.clusterInfrastructure.Start(ctx, cluster)
	if err != nil {
		return err
	}
	err = uc.handlerAddNode(ctx, cluster)
	if err != nil {
		return err
	}
	err = uc.handlerRemoveNode(ctx, cluster)
	if err != nil {
		return err
	}
	return
}

func (uc *ClusterUsecase) handlerAddNode(ctx context.Context, cluster *Cluster) error {
	addNodes := make([]*Node, 0)
	for _, node := range cluster.Nodes {
		if node.Status == NodeStatusCreating {
			addNodes = append(addNodes, node)
		}
	}
	err := uc.clusterInfrastructure.AddNodes(ctx, cluster, addNodes)
	if err != nil {
		return err
	}
	for _, node := range cluster.Nodes {
		for _, n := range addNodes {
			if node.Name == n.Name {
				node.Status = NodeStatusRunning
			}
		}
	}
	return nil
}

func (uc *ClusterUsecase) handlerRemoveNode(ctx context.Context, cluster *Cluster) error {
	removeNodes := make([]*Node, 0)
	for _, node := range cluster.Nodes {
		if node.Status == NodeStatusDeleting {
			removeNodes = append(removeNodes, node)
		}
	}
	err := uc.clusterInfrastructure.RemoveNodes(ctx, cluster, removeNodes)
	if err != nil {
		return err
	}
	newNodes := make([]*Node, 0)
	for _, node := range cluster.Nodes {
		ok := false
		for _, n := range removeNodes {
			if node.Name == n.Name {
				ok = true
			}
		}
		if !ok {
			newNodes = append(newNodes, node)
		}
	}
	cluster.Nodes = newNodes
	return nil
}
