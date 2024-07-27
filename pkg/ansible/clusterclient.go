package ansible

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/f-rambo/ocean/internal/biz"
	"github.com/f-rambo/ocean/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
)

type ClusterConstruct struct {
	log                        *log.Helper
	c                          *conf.Bootstrap
	migrateToBostionHostEnable bool
}

func NewClusterConstruct(c *conf.Bootstrap, logger log.Logger) biz.ClusterConstruct {
	return &ClusterConstruct{
		log: log.NewHelper(logger),
		c:   c,
	}
}

// 初始化集群
func (cc *ClusterConstruct) GenerateInitialCluster(ctx context.Context, cluster *biz.Cluster) error {
	if cluster.Type == biz.ClusterTypeLocal.String() && len(cluster.Nodes) < biz.NodeMinSize.Int() {
		return errors.New("local cluster node size must be greater than 1")
	}
	// 云厂商集群
	if cluster.Type != biz.ClusterTypeLocal.String() {
		nodeGroup := &biz.NodeGroup{
			OSImage:                 "ubuntu 22.04",
			CPU:                     4,
			Memory:                  8,
			SystemDisk:              100,
			InternetMaxBandwidthOut: 100,
			MinSize:                 5,
			MaxSize:                 100,
			InstanceType:            "", // 由云服务商决定
			TargetSize:              5,
		}
		nodeGroup.Name = fmt.Sprintf("cloudproider-%s-cpu-%d-mem-%d-disk-%d", cluster.Type, nodeGroup.CPU, int(nodeGroup.Memory), nodeGroup.SystemDisk)
		cluster.NodeGroups = []*biz.NodeGroup{nodeGroup}
		var i int32
		for i = 0; i < nodeGroup.MinSize; i++ {
			roleName := biz.NodeRoleMaster.String()
			if i > 2 {
				roleName = biz.NodeRoleWorker.String()
			}
			node := &biz.Node{
				Name:      fmt.Sprintf("%s-%d", roleName, i),
				Labels:    "",
				Status:    biz.NodeStatusRunning.Uint8(),
				ClusterID: cluster.ID,
				NodeGroup: nodeGroup,
				Role:      roleName,
			}
			cluster.Nodes = append(cluster.Nodes, node)
		}
		return nil
	}
	// 本地集群
	err := cc.getNodesInformation(ctx, cluster)
	if err != nil {
		return err
	}
	sort.Sort(biz.Nodes(cluster.Nodes))
	masterNum := 0
	workNum := 0
	for _, node := range cluster.Nodes {
		if node.NodeGroup == nil {
			return errors.New("node group is nil")
		}
		if node.NodeGroup.Memory >= 8 && node.NodeGroup.CPU >= 4 && masterNum < 3 {
			node.Role = biz.NodeRoleMaster.String()
			node.Status = biz.NodeStatusCreating.Uint8()
			masterNum++
			continue
		}
		if workNum >= 3 {
			node.Status = biz.NodeStatusUnspecified.Uint8()
			continue
		}
		node.Role = biz.NodeRoleWorker.String()
		node.Status = biz.NodeStatusCreating.Uint8()
		workNum++
	}
	return nil
}

// 获取集群节点信息，配置信息
func (cc *ClusterConstruct) getNodesInformation(ctx context.Context, cluster *biz.Cluster) error {
	playbook := getSystemInformation()
	playbookPath, err := savePlaybook(cc.c.Resource.GetClusterPath(), playbook)
	if err != nil {
		return err
	}
	err = cc.exec(ctx, cluster, cc.c.Resource.GetClusterPath(), playbookPath, nil)
	if err != nil {
		return err
	}
	resultMaps := make([]map[string]interface{}, 0)
	remaining := cluster.Logs
	for {
		startIndex := strings.Index(remaining, StartOutputKey.String())
		if startIndex == -1 {
			break
		}
		endIndex := strings.Index(remaining, EndOutputKey.String())
		if endIndex == -1 {
			break
		}
		startIndex += len(StartOutputKey.String())
		if startIndex >= endIndex {
			break
		}
		result := remaining[startIndex:endIndex]
		if result != "" {
			unescapedResult := strings.ReplaceAll(result, `\"`, `"`)
			resultMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(unescapedResult), &resultMap)
			if err != nil {
				return err
			}
			resultMaps = append(resultMaps, resultMap)
		}
		remaining = remaining[endIndex+len(EndOutputKey.String()):]
	}
	getNodeResult := func(nodeID int64) map[string]interface{} {
		for _, resultMap := range resultMaps {
			if _, ok := resultMap["node_id"]; ok && cast.ToInt64(resultMap["node_id"]) == nodeID {
				return resultMap
			}
		}
		return nil
	}
	for _, node := range cluster.Nodes {
		resultMap := getNodeResult(node.ID)
		nodeGroup := &biz.NodeGroup{}
		for k, v := range resultMap {
			switch k {
			case "gpu_number":
				nodeGroup.GPU = cast.ToInt32(v)
			case "gpu_spec":
				nodeGroup.GpuSpec = cast.ToString(v)
			case "cpu_number":
				nodeGroup.CPU = cast.ToInt32(v)
			case "memory":
				nodeGroup.Memory = cast.ToFloat64(v)
			case "disk":
				nodeGroup.SystemDisk = cast.ToInt32(v)
			case "os_info":
				nodeGroup.OSImage = cast.ToString(v)
			}
		}
		node.NodeGroup = nodeGroup
		node.Kernel = cast.ToString(resultMap["kernel_info"])
		node.Container = cast.ToString(resultMap["container_version"])
		node.InternalIP = cast.ToString(resultMap["ip"])
	}
	nodeGroupMap := make(map[string]*biz.NodeGroup)
	for _, node := range cluster.Nodes {
		nodeGroupName := fmt.Sprintf("gpu-%d-gpu_spec-%s-cpu-%d-mem-%d-disk-%d",
			node.NodeGroup.GPU, node.NodeGroup.GpuSpec, node.NodeGroup.CPU, int(node.NodeGroup.Memory), node.NodeGroup.DataDisk)
		node.NodeGroup.Name = nodeGroupName
		nodeGroupMap[nodeGroupName] = node.NodeGroup
	}
	nodeGroups := make([]*biz.NodeGroup, 0)
	for _, nodeGroup := range nodeGroupMap {
		nodeGroups = append(nodeGroups, nodeGroup)
	}
	cluster.NodeGroups = nodeGroups
	return nil
}

func (cc *ClusterConstruct) GenerateNodeLables(ctx context.Context, cluster *biz.Cluster, nodeGroup *biz.NodeGroup) (lables string, err error) {
	lableMap := make(map[string]string)
	lableMap["cluster"] = cluster.Name
	lableMap["cluster_type"] = cluster.Type
	lableMap["region"] = cluster.Region
	lableMap["nodegroup"] = nodeGroup.Name
	lableMap["nodegroup_type"] = nodeGroup.Type
	lableMap["instance_type"] = nodeGroup.InstanceType
	lablebytes, err := json.Marshal(lableMap)
	if err != nil {
		return "", err
	}
	return string(lablebytes), nil
}

func (cc *ClusterConstruct) MigrateToBostionHost(ctx context.Context, cluster *biz.Cluster) error {
	cc.migrateToBostionHostEnable = true
	defer func() {
		cc.migrateToBostionHostEnable = false
	}()
	oceanResource := cc.c.Resource
	migratePlaybook := getMigratePlaybook()
	databasePath := cc.c.Data.GetDBFilePath()
	pulumiPath := cc.c.Resource.GetPulumiPath()
	migratePlaybook.AddSynchronize("database", databasePath, databasePath)
	migratePlaybook.AddSynchronize("pulumi", pulumiPath, pulumiPath)
	migratePlaybookPath, err := savePlaybook(oceanResource.GetClusterPath(), migratePlaybook)
	if err != nil {
		return err
	}
	err = cc.exec(ctx,
		cluster,
		oceanResource.GetClusterPath(),
		migratePlaybookPath,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func (cc *ClusterConstruct) InstallCluster(ctx context.Context, cluster *biz.Cluster) error {
	serversInitPlaybook := getServerInitPlaybook()
	serversInitPlaybookPath, err := savePlaybook(cc.c.Resource.GetClusterPath(), serversInitPlaybook)
	if err != nil {
		return err
	}
	err = cc.exec(ctx, cluster, cc.c.Resource.GetClusterPath(), serversInitPlaybookPath, nil)
	if err != nil {
		return err
	}
	return cc.kubespray(ctx, cluster, GetClusterPlaybookPath(), nil)
}

func (cc *ClusterConstruct) UnInstallCluster(ctx context.Context, cluster *biz.Cluster) error {
	return cc.kubespray(ctx, cluster, GetResetPlaybookPath(), nil)
}

func (cc *ClusterConstruct) AddNodes(ctx context.Context, cluster *biz.Cluster, nodes []*biz.Node) error {
	for _, node := range nodes {
		log.Info("add node", "name", node.Name, "ip", node.ExternalIP, "role", node.Role)
	}
	return cc.kubespray(ctx, cluster, GetScalePlaybookPath(), nil)
}

func (cc *ClusterConstruct) RemoveNodes(ctx context.Context, cluster *biz.Cluster, nodes []*biz.Node) error {
	for _, node := range nodes {
		log.Info("remove node", "name", node.Name, "ip", node.ExternalIP, "role", node.Role)
		err := cc.kubespray(ctx, cluster, GetRemoveNodePlaybookPath(), map[string]string{"node": node.Name})
		if err != nil {
			return err
		}
	}
	return nil
}

func (cc *ClusterConstruct) kubespray(ctx context.Context, cluster *biz.Cluster, playbook string, env map[string]string) error {
	oceanResource := cc.c.Resource
	kubespray, err := NewKubespray(&oceanResource)
	if err != nil {
		return errors.Wrap(err, "new kubespray error")
	}
	return cc.exec(ctx, cluster, kubespray.GetPackagePath(), playbook, env)
}

func (cc *ClusterConstruct) exec(ctx context.Context, cluster *biz.Cluster, cmdRunDir string, playbook string, env map[string]string) error {
	matedataMap := make(map[string]string)
	servers := make([]Server, 0)
	if cc.migrateToBostionHostEnable {
		servers = []Server{
			{Ip: cluster.BostionHost.ExternalIP, Username: "root", ID: cluster.BostionHost.InstanceID, Role: "bostion"},
		}
	}
	if !cc.migrateToBostionHostEnable {
		for _, node := range cluster.Nodes {
			servers = append(servers, Server{Ip: node.ExternalIP, Username: node.User, ID: cast.ToString(node.ID), Role: node.Role})
			matedataMap["node_id"] = cast.ToString(node.ID)
			matedataMap["node_name"] = node.Name
		}
	}
	g := new(errgroup.Group)
	ansibleLog := make(chan string, 1024)
	g.Go(func() error {
		defer close(ansibleLog)
		return NewGoAnsiblePkg(cc.c.Ansible).
			SetAnsiblePlaybookBinary(cc.c.Resource.GetAnsibleCli()).
			SetLogChan(ansibleLog).
			SetServers(servers...).
			SetCmdRunDir(cmdRunDir).
			SetPlaybooks(playbook).
			SetEnvMap(env).
			ExecPlayBooks(ctx)
	})
	g.Go(func() error {
		for {
			select {
			case log, ok := <-ansibleLog:
				if !ok {
					return nil
				}
				cluster.Logs += log
				cc.log.Info(log)
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})
	err := g.Wait()
	if err != nil {
		return err
	}
	return nil
}
