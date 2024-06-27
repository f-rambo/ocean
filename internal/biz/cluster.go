package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/f-rambo/ocean/internal/conf"
	"github.com/f-rambo/ocean/pkg/ansible"
	"github.com/f-rambo/ocean/pkg/kubeclient"
	"github.com/f-rambo/ocean/pkg/pulumiapi"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ClusterTypeLocal    = "local"
	ClusterTypeAWS      = "aws"
	ClusterTypeGoogle   = "google"
	ClusterTypeAzure    = "azure"
	ClusterTypeAliCloud = "alicloud"
)

const (
	ClusterRoleMaster = "master"
	ClusterRoleWorker = "worker"
	ClusterRoleEdge   = "edge"
)

type Cluster struct {
	ID               int64        `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Name             string       `json:"name" gorm:"column:name; default:''; NOT NULL"` // *
	ServerVersion    string       `json:"server_version" gorm:"column:server_version; default:''; NOT NULL"`
	ApiServerAddress string       `json:"api_server_address" gorm:"column:api_server_address; default:''; NOT NULL"`
	Config           string       `json:"config" gorm:"column:config; default:''; NOT NULL;"`
	Addons           string       `json:"addons" gorm:"column:addons; default:''; NOT NULL;"`
	AddonsConfig     string       `json:"addons_config" gorm:"column:addons_config; default:''; NOT NULL;"`
	Status           uint8        `json:"status" gorm:"column:status; default:0; NOT NULL;"`
	Type             string       `json:"type" gorm:"column:type; default:''; NOT NULL;"` //*  aws google cloud azure alicloud local
	KubeConfig       []byte       `json:"kube_config" gorm:"column:kube_config; default:''; NOT NULL; type:json"`
	PublicKey        string       `json:"public_key" gorm:"column:public_key; default:''; NOT NULL;"` // *
	Region           string       `json:"region" gorm:"column:region; default:''; NOT NULL;"`         // *
	VpcID            string       `json:"vpc_id" gorm:"column:vpc_id; default:''; NOT NULL;"`
	ExternalIP       string       `json:"external_ip" gorm:"column:external_ip; default:''; NOT NULL;"`
	AccessID         string       `json:"access_id" gorm:"column:access_id; default:''; NOT NULL;"`   // *
	AccessKey        string       `json:"access_key" gorm:"column:access_key; default:''; NOT NULL;"` // *
	BostionHost      *BostionHost `json:"bostion_host" gorm:"-"`
	Logs             string       `json:"logs" gorm:"-"` // logs data from localfile
	Nodes            []*Node      `json:"nodes" gorm:"-"`
	gorm.Model
}

type Node struct {
	ID                      int64   `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Name                    string  `json:"name" gorm:"column:name; default:''; NOT NULL"`
	CPU                     int     `json:"cpu" gorm:"column:cpu; default:0; NOT NULL"`
	Memory                  float64 `json:"memory" gorm:"column:memory; default:0; NOT NULL"`
	GPU                     int     `json:"gpu" gorm:"column:gpu; default:0; NOT NULL"`
	GpuSpec                 string  `json:"gpu_spec" gorm:"column:gpu_spec; default:''; NOT NULL"`      // 1080ti 2080ti 3090
	SystemDisk              int     `json:"system_disk" gorm:"column:system_disk; default:0; NOT NULL"` // 随着服务释放掉的存储空间
	DataDisk                int     `json:"data_disk" gorm:"column:data_disk; default:0; NOT NULL"`
	InstanceType            string  `json:"instance_type" gorm:"column:instance_type; default:''; NOT NULL"`
	Labels                  string  `json:"labels" gorm:"column:labels; default:''; NOT NULL"`
	Annotations             string  `json:"annotations" gorm:"column:annotations; default:''; NOT NULL"`
	OSImage                 string  `json:"os_image" gorm:"column:os_image; default:''; NOT NULL"`
	Kernel                  string  `json:"kernel" gorm:"column:kernel; default:''; NOT NULL"`
	Container               string  `json:"container" gorm:"column:container; default:''; NOT NULL"`
	Kubelet                 string  `json:"kubelet" gorm:"column:kubelet; default:''; NOT NULL"`
	KubeProxy               string  `json:"kube_proxy" gorm:"column:kube_proxy; default:''; NOT NULL"`
	InternalIP              string  `json:"internal_ip" gorm:"column:internal_ip; default:''; NOT NULL"`
	ExternalIP              string  `json:"external_ip" gorm:"column:external_ip; default:''; NOT NULL"`
	User                    string  `json:"user" gorm:"column:user; default:''; NOT NULL"`
	Password                string  `json:"password" gorm:"column:password; default:''; NOT NULL"`
	SudoPassword            string  `json:"sudo_password" gorm:"column:sudo_password; default:''; NOT NULL"`
	Role                    string  `json:"role" gorm:"column:role; default:''; NOT NULL;"` // master worker edge
	Status                  uint8   `json:"status" gorm:"column:status; default:0; NOT NULL;"`
	InternetMaxBandwidthOut int     `json:"internet_max_bandwidth_out" gorm:"column:internet_max_bandwidth_out; default:0; NOT NULL"`
	NodeInitScript          string  `json:"cloud_init_script" gorm:"column:cloud_init_script; default:''; NOT NULL"`
	ClusterID               int64   `json:"cluster_id" gorm:"column:cluster_id; default:0; NOT NULL"`
	gorm.Model
}

type BostionHost struct {
	ID         int64  `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	InstanceID string `json:"instance_id" gorm:"column:instance_id; default:''; NOT NULL"`
	Hostname   string `json:"hostname" gorm:"column:hostname; default:''; NOT NULL"`
	ExternalIP string `json:"external_ip" gorm:"column:external_ip; default:''; NOT NULL"`
	PublicIP   string `json:"public_ip" gorm:"column:public_ip; default:''; NOT NULL"`
	PrivateIP  string `json:"private_ip" gorm:"column:private_ip; default:''; NOT NULL"`
	ClusterID  int64  `json:"cluster_id" gorm:"column:cluster_id; default:0; NOT NULL"`
	gorm.Model
}

type ansibleArgs struct {
	servers []ansible.Server
	env     map[string]string
}

type pulumiArgs struct {
	delete bool
}

func (c *Cluster) IsEmpty() bool {
	return c.ID == 0
}

func (c *Cluster) GetNode(nodeId int64) *Node {
	for _, node := range c.Nodes {
		if node.ID == nodeId {
			return node
		}
	}
	return nil
}

type ClusterRepo interface {
	Save(context.Context, *Cluster) error
	Get(context.Context, int64) (*Cluster, error)
	GetByName(context.Context, string) (*Cluster, error)
	List(context.Context, *Cluster) ([]*Cluster, error)
	Delete(context.Context, int64) error
	ReadClusterLog(cluster *Cluster) error
	WriteClusterLog(cluster *Cluster) error
}

type ClusterUsecase struct {
	c    *conf.Bootstrap
	repo ClusterRepo
	log  *log.Helper
}

func NewClusterUseCase(c *conf.Bootstrap, repo ClusterRepo, logger log.Logger) *ClusterUsecase {
	return &ClusterUsecase{
		c:    c,
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ClusterUsecase) Get(ctx context.Context, id int64) (*Cluster, error) {
	cluster, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	err = uc.repo.ReadClusterLog(cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func (uc *ClusterUsecase) List(ctx context.Context) ([]*Cluster, error) {
	return uc.repo.List(ctx, nil)
}

func (uc *ClusterUsecase) Apply(ctx context.Context, cluster *Cluster) error {
	data, err := uc.repo.GetByName(ctx, cluster.Name)
	if err != nil {
		return err
	}
	if !data.IsEmpty() {
		return errors.New("cluster name already exists")
	}
	err = uc.repo.Save(ctx, cluster)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ClusterUsecase) Delete(ctx context.Context, clusterID int64) error {
	cluster, err := uc.repo.Get(ctx, clusterID)
	if err != nil {
		return err
	}
	if cluster.IsEmpty() {
		return nil
	}
	err = uc.repo.Delete(ctx, clusterID)
	if err != nil {
		return err
	}
	return nil
}

// 集群信息回调处理
func (uc *ClusterUsecase) Reconcile(ctx context.Context, cluster *Cluster) error {
	data, err := uc.repo.Get(ctx, cluster.ID)
	if err != nil {
		return err
	}
	uc.c.GetArgoWorkflowConfig()
	if data.IsEmpty() {
		cluster.Logs = "start uninstal cluster..."
		err = uc.clusterUninstall(ctx, cluster)
		if err != nil {
			return err
		}
		cluster.Logs = "start delete cluster..."
		err = uc.servers(ctx, cluster, &pulumiArgs{delete: true})
		if err != nil {
			return err
		}
		return nil
	}
	cluster.Logs = "start update cluster..."
	env := uc.c.Ocean.GetEnv()
	if env == conf.EnvLocal && cluster.Type != ClusterTypeLocal {
		err = uc.servers(ctx, cluster, nil)
		if err != nil {
			return err
		}
	}
	if env == conf.EnvBostionHost || cluster.Type == ClusterTypeLocal {
		err = uc.cluster(ctx, cluster)
		if err != nil {
			return err
		}
	}
	if env == conf.EnvCluster {
		// remove node
		for _, node := range cluster.Nodes {
			ok := false
			for _, v := range data.Nodes {
				if v.ID == node.ID {
					ok = true
					break
				}
			}
			if !ok {
				err = uc.removeNode(ctx, cluster, node)
				if err != nil {
					return err
				}
			}
		}
		// add node
		err = uc.addNode(ctx, cluster)
		if err != nil {
			return err
		}
		// update servers
		err = uc.servers(ctx, cluster, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

// 创建/删除服务器
func (uc *ClusterUsecase) servers(ctx context.Context, cluster *Cluster, pulumiArgs *pulumiArgs) error {
	switch cluster.Type {
	case ClusterTypeAliCloud:
		err := uc.alicloudServers(ctx, cluster, pulumiArgs)
		if err != nil {
			return err
		}
	case ClusterTypeAWS:
	default:
		return errors.New("not support cluster type")
	}
	err := uc.migrateToBostionHost(ctx, cluster)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ClusterUsecase) alicloudServers(ctx context.Context, cluster *Cluster, pulumiArgs *pulumiArgs) error {
	args := pulumiapi.AlicloudClusterArgs{
		Name:      cluster.Name,
		PublicKey: cluster.PublicKey,
		Nodes:     make([]pulumiapi.AlicloudNodeArgs, 0),
	}
	for _, node := range cluster.Nodes {
		labels := make(map[string]string)
		if node.Labels != "" {
			err := json.Unmarshal([]byte(node.Labels), &labels)
			if err != nil {
				return err
			}
		}
		args.Nodes = append(args.Nodes, pulumiapi.AlicloudNodeArgs{
			Name:                    node.Name,
			InstanceType:            node.InstanceType,
			CPU:                     node.CPU,
			Memory:                  node.Memory,
			GPU:                     node.GPU,
			GpuSpec:                 node.GpuSpec,
			OSImage:                 node.OSImage,
			InternetMaxBandwidthOut: node.InternetMaxBandwidthOut,
			SystemDisk:              node.SystemDisk,
			DataDisk:                node.DataDisk,
			Labels:                  labels,
			NodeInitScript:          node.NodeInitScript,
		})
	}
	var pulumiFunc pulumiapi.PulumiFunc
	pulumiFunc = pulumiapi.StartAlicloudCluster(args).StartServers
	if pulumiArgs != nil && pulumiArgs.delete {
		pulumiFunc = pulumiapi.StartAlicloudCluster(args).Clear
	}
	g := new(errgroup.Group)
	bostionHost := &BostionHost{}
	pulumiOutput := make(chan string, 1024)
	pulumiOutput <- "starting alicloud servers..."
	g.Go(func() error {
		defer close(pulumiOutput)
		output, err := pulumiapi.NewPulumiAPI(ctx, pulumiOutput).
			ProjectName(pulumiapi.AlicloudProjectName).
			StackName(pulumiapi.AlicloudStackName).
			Plugin(pulumiapi.PulumiPlugin{Kind: "alicloud", Version: "3.56.0"}, pulumiapi.PulumiPlugin{Kind: "kubernetes", Version: "4.12.0"}).
			Env(map[string]string{"ALICLOUD_ACCESS_KEY": cluster.AccessID, "ALICLOUD_SECRET_KEY": cluster.AccessKey, "ALICLOUD_REGION": cluster.Region}).
			RegisterDeployFunc(pulumiFunc).
			Up(ctx)
		if err != nil {
			return err
		}
		outputMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(output), &outputMap)
		if err != nil {
			return err
		}
		for k, v := range outputMap {
			switch k {
			case "vpc_id":
				cluster.VpcID = cast.ToString(v)
			case "external_ip":
				cluster.ExternalIP = cast.ToString(v)
				bostionHost.ExternalIP = cast.ToString(v)
			case "bostion_public_ip":
				bostionHost.PublicIP = cast.ToString(v)
			case "bostion_private_ip":
				bostionHost.PrivateIP = cast.ToString(v)
			case "bostion_id":
				bostionHost.InstanceID = cast.ToString(v)
			case "bostion_hostname":
				bostionHost.Hostname = cast.ToString(v)
			default:
				cluster.Logs += fmt.Sprintf("%s: %s\n", k, v)
			}
		}
		return nil
	})
	g.Go(func() error {
		return uc.clusterLog(ctx, cluster, pulumiOutput)
	})
	cluster.BostionHost = bostionHost
	err := g.Wait()
	if err != nil {
		return err
	}
	pulumiOutput <- "alicloud servers success"
	return nil
}

// 数据迁移到bostion主机
func (uc *ClusterUsecase) migrateToBostionHost(ctx context.Context, cluster *Cluster) error {
	oceanResource := uc.c.GetOceanResource()
	migratePlaybook := ansible.GetMigratePlaybook()
	migratePlaybook.AddSynchronize("database", uc.c.GetOceanData().GetDBFilePath(), "/var/lib/mysql")
	migratePlaybook.AddSynchronize("pulumi", "~/.pulumi", "/root/.pulumi")
	migratePlaybookPath, err := ansible.SavePlaybook(oceanResource.GetClusterPath(), migratePlaybook)
	if err != nil {
		return err
	}
	args := &ansibleArgs{
		servers: []ansible.Server{
			{Ip: cluster.BostionHost.ExternalIP, Username: "root", ID: cluster.BostionHost.InstanceID, Role: "bostion"},
		},
	}
	err = uc.exec(ctx,
		cluster,
		oceanResource.GetClusterPath(),
		migratePlaybookPath,
		args,
	)
	if err != nil {
		return err
	}
	return nil
}

// 集群安装
func (uc *ClusterUsecase) cluster(ctx context.Context, cluster *Cluster) error {
	// todo 设置集群配置文件
	serversInitPlaybook := ansible.GetServerInitPlaybook()
	serversInitPlaybookPath, err := ansible.SavePlaybook(uc.c.GetOceanResource().GetClusterPath(), serversInitPlaybook)
	if err != nil {
		return err
	}
	err = uc.exec(ctx, cluster, uc.c.GetOceanResource().GetClusterPath(), serversInitPlaybookPath, nil)
	if err != nil {
		return err
	}
	return uc.kubespray(ctx, cluster, ansible.GetClusterPlaybookPath(), nil)
}

// 增加节点
func (uc *ClusterUsecase) addNode(ctx context.Context, cluster *Cluster) error {
	return uc.kubespray(ctx, cluster, ansible.GetScalePlaybookPath(), nil)
}

// 移除节点
func (uc *ClusterUsecase) removeNode(ctx context.Context, cluster *Cluster, node *Node) error {
	args := &ansibleArgs{
		env: map[string]string{"node": node.Name},
	}
	return uc.kubespray(ctx, cluster, ansible.GetRemoveNodePlaybookPath(), args)
}

// 卸载集群
func (uc *ClusterUsecase) clusterUninstall(ctx context.Context, cluster *Cluster) error {
	return uc.kubespray(ctx, cluster, ansible.GetResetPlaybookPath(), nil)
}

// 获取集群信息
func (uc *ClusterUsecase) clusterInfomation(ctx context.Context, cluster *Cluster) error {
	clientSet, err := kubeclient.GetKubeClientSet()
	if err != nil {
		return err
	}
	versionInfo, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		return err
	}
	serverAddress := clientSet.Discovery().RESTClient().Get().URL().Host
	cluster.ServerVersion = versionInfo.String()
	cluster.ApiServerAddress = serverAddress
	err = uc.getNodes(ctx, cluster)
	if err != nil {
		return err
	}
	for _, node := range cluster.Nodes {
		if strings.Contains(node.Role, ClusterRoleMaster) {
			cluster.Name = node.Name
			break
		}
	}
	return nil
}

func (uc *ClusterUsecase) getNodes(ctx context.Context, cluster *Cluster) error {
	clientSet, err := kubeclient.GetKubeClientSet()
	if err != nil {
		return err
	}
	nodeList, err := clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, node := range nodeList.Items {
		labelsJson, err := json.Marshal(node.Labels)
		if err != nil {
			return err
		}
		annotationsJson, err := json.Marshal(node.Annotations)
		if err != nil {
			return err
		}
		roles := make([]string, 0)
		if _, ok := node.Labels["node-role.kubernetes.io/master"]; ok {
			roles = append(roles, ClusterRoleMaster)
		}
		if _, ok := node.Labels["node-role.kubernetes.io/control-plane"]; ok {
			roles = append(roles, ClusterRoleMaster)
		}
		if _, ok := node.Labels["node-role.kubernetes.io/edge"]; ok {
			roles = append(roles, ClusterRoleEdge)
		}
		if _, ok := node.Labels["node-role.kubernetes.io/worker"]; ok {
			roles = append(roles, ClusterRoleWorker)
		}
		if len(roles) == 0 {
			roles = append(roles, ClusterRoleWorker)
		}
		roleJson, err := json.Marshal(roles)
		if err != nil {
			return err
		}
		var internalIP string
		var externalIP string
		for _, addr := range node.Status.Addresses {
			if addr.Type == coreV1.NodeInternalIP {
				internalIP = addr.Address
			}
			if addr.Type == coreV1.NodeExternalIP {
				externalIP = addr.Address
			}
		}
		cluster.Nodes = append(cluster.Nodes, &Node{
			Name:        node.Name,
			Labels:      string(labelsJson),
			Annotations: string(annotationsJson),
			OSImage:     node.Status.NodeInfo.OSImage,
			Kernel:      node.Status.NodeInfo.KernelVersion,
			Container:   node.Status.NodeInfo.ContainerRuntimeVersion,
			Kubelet:     node.Status.NodeInfo.KubeletVersion,
			KubeProxy:   node.Status.NodeInfo.KubeProxyVersion,
			InternalIP:  internalIP,
			ExternalIP:  externalIP,
			Role:        string(roleJson),
			ClusterID:   cluster.ID,
		})
	}
	return nil
}

func (uc *ClusterUsecase) kubespray(ctx context.Context, cluster *Cluster, playbook string, args *ansibleArgs) error {
	oceanResource := uc.c.GetOceanResource()
	kubespray, err := ansible.NewKubespray(&oceanResource)
	if err != nil {
		return errors.Wrap(err, "new kubespray error")
	}
	return uc.exec(ctx, cluster, kubespray.GetPackagePath(), playbook, args)
}

func (uc *ClusterUsecase) exec(ctx context.Context, cluster *Cluster, playbook string, cmdRunDir string, args *ansibleArgs) error {
	servers := make([]ansible.Server, 0)
	for _, node := range cluster.Nodes {
		servers = append(servers, ansible.Server{Ip: node.ExternalIP, Username: node.User, ID: cast.ToString(node.ID), Role: node.Role})
	}
	if args != nil && len(args.servers) > 0 {
		servers = args.servers
	}
	env := make(map[string]string)
	if args != nil && len(args.env) > 0 {
		env = args.env
	}
	g := new(errgroup.Group)
	ansibleLog := make(chan string, 1024)
	g.Go(func() error {
		defer close(ansibleLog)
		return ansible.NewGoAnsiblePkg(uc.c).
			SetAnsiblePlaybookBinary(uc.c.GetOceanResource().GetAnsibleCli()).
			SetLogChan(ansibleLog).
			SetServers(servers...).
			SetCmdRunDir(cmdRunDir).
			SetPlaybooks(playbook).
			SetEnvMap(env).
			ExecPlayBooks(ctx)
	})
	g.Go(func() error {
		return uc.clusterLog(ctx, cluster, ansibleLog)
	})
	err := g.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (uc *ClusterUsecase) clusterLog(ctx context.Context, cluster *Cluster, output chan string) (err error) {
	defer func() {
		if err != nil {
			uc.log.Errorf("err: %v", err)
		}
		if r := recover(); r != nil {
			uc.log.Errorf("panic: %v", r)
		}
		if cluster.Logs != "" {
			err := uc.repo.WriteClusterLog(cluster)
			if err != nil {
				uc.log.Errorf("write cluster log error: %v", err)
				return
			}
		}
	}()
	for {
		select {
		case log, ok := <-output:
			if !ok {
				return nil
			}
			cluster.Logs += log
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(3 * time.Second):
			err := uc.repo.WriteClusterLog(cluster)
			if err != nil {
				return err
			}
		}
	}
}
