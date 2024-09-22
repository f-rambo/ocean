package interfaces

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	"github.com/f-rambo/ocean/api/cluster/v1alpha1"
	systemv1alpha1 "github.com/f-rambo/ocean/api/system/v1alpha1"
	"github.com/f-rambo/ocean/internal/biz"
	"github.com/f-rambo/ocean/internal/conf"
	"github.com/f-rambo/ocean/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ClusterInterface struct {
	v1alpha1.UnimplementedClusterInterfaceServer
	clusterUc *biz.ClusterUsecase
	projectUc *biz.ProjectUsecase
	appUc     *biz.AppUsecase
	c         *conf.Bootstrap
	log       *log.Helper
}

func NewClusterInterface(clusterUc *biz.ClusterUsecase, projectUc *biz.ProjectUsecase, appUc *biz.AppUsecase, c *conf.Bootstrap, logger log.Logger) *ClusterInterface {
	return &ClusterInterface{
		clusterUc: clusterUc,
		projectUc: projectUc,
		appUc:     appUc,
		c:         c,
		log:       log.NewHelper(logger),
	}
}

func (uc *ClusterInterface) StartReconcile(ctx context.Context) (err error) {
	uc.log.Info("start watching reconcile...")
	for {
		cluster, err := uc.clusterUc.Watch(ctx)
		if err != nil {
			return err
		}
		if cluster == nil {
			continue
		}
		err = uc.clusterUc.Reconcile(ctx, cluster)
		if err != nil {
			return err
		}
	}
}

func (uc *ClusterInterface) StopReconcile(ctx context.Context) error {
	uc.log.Info("stop watching reconcile...")
	return nil
}

func (c *ClusterInterface) Ping(ctx context.Context, _ *emptypb.Empty) (*v1alpha1.Msg, error) {
	return &v1alpha1.Msg{Message: "pong"}, nil
}

func (c *ClusterInterface) Get(ctx context.Context, clusterID *v1alpha1.ClusterArgs) (*v1alpha1.Cluster, error) {
	cluster, err := c.clusterUc.Get(ctx, clusterID.Id)
	if err != nil {
		return nil, err
	}
	data := &v1alpha1.Cluster{}
	if cluster == nil {
		return data, nil
	}
	return c.bizCLusterToCluster(cluster), nil
}

func (c *ClusterInterface) Save(ctx context.Context, clusterArgs *v1alpha1.ClusterArgs) (clusterApiRes *v1alpha1.Cluster, err error) {
	if clusterArgs.Name == "" {
		return nil, errors.New("cluster name is required")
	}
	if clusterArgs.PublicKey == "" {
		return nil, errors.New("public key is required")
	}
	if clusterArgs.Type == "" {
		return nil, errors.New("server type is required")
	}
	if biz.ClusterType(clusterArgs.Type) != biz.ClusterTypeLocal {
		if clusterArgs.AccessKeyId == "" {
			return nil, errors.New("access key id is required")
		}
		if clusterArgs.SecretAccessKey == "" {
			return nil, errors.New("secret access key is required")
		}
	}
	var cluster *biz.Cluster
	if clusterArgs.Id != 0 {
		cluster, err = c.clusterUc.Get(ctx, clusterArgs.Id)
		if err != nil {
			return nil, err
		}
		if cluster == nil {
			return nil, errors.New("cluster not found")
		}
		cluster.Name = clusterArgs.Name
		cluster.Type = biz.ClusterType(clusterArgs.Type)
		cluster.PublicKey = clusterArgs.PublicKey
		cluster.Region = clusterArgs.Region
		cluster.AccessID = clusterArgs.AccessKeyId
		cluster.AccessKey = clusterArgs.SecretAccessKey
		for _, node := range cluster.Nodes {
			if node.ID == 0 {
				continue
			}
			for _, nodeArg := range clusterArgs.Nodes {
				if node.ID == nodeArg.Id {
					node.InternalIP = nodeArg.Ip
					node.User = nodeArg.User
					node.Role = biz.NodeRole(nodeArg.Role)
				}
			}
		}
		for _, nodeArg := range clusterArgs.Nodes {
			if nodeArg.Id == 0 {
				cluster.Nodes = append(cluster.Nodes, &biz.Node{
					ID:         nodeArg.Id,
					InternalIP: nodeArg.Ip,
					User:       nodeArg.User,
					Role:       biz.NodeRole(nodeArg.Role),
				})
			}
		}
	} else {
		cluster = &biz.Cluster{
			ID:        clusterArgs.Id,
			Name:      clusterArgs.Name,
			Type:      biz.ClusterType(clusterArgs.Type),
			PublicKey: clusterArgs.PublicKey,
			Region:    clusterArgs.Region,
			AccessID:  clusterArgs.AccessKeyId,
			AccessKey: clusterArgs.SecretAccessKey,
			Nodes:     make([]*biz.Node, 0),
		}
		for _, node := range clusterArgs.Nodes {
			cluster.Nodes = append(cluster.Nodes, &biz.Node{
				ID:         node.Id,
				InternalIP: node.Ip,
				User:       node.User,
				Role:       biz.NodeRole(node.Role),
			})
		}
	}
	err = c.clusterUc.Save(ctx, cluster)
	if err != nil {
		return nil, err
	}
	return c.bizCLusterToCluster(cluster), nil
}

func (c *ClusterInterface) List(ctx context.Context, _ *emptypb.Empty) (*v1alpha1.ClusterList, error) {
	data := &v1alpha1.ClusterList{}
	clusters, err := c.clusterUc.List(ctx)
	if err != nil {
		return nil, err
	}
	if len(clusters) == 0 {
		return data, nil
	}
	for _, v := range clusters {
		data.Clusters = append(data.Clusters, c.bizCLusterToCluster(v))
	}
	return data, nil
}

func (c *ClusterInterface) Delete(ctx context.Context, clusterID *v1alpha1.ClusterArgs) (*v1alpha1.Msg, error) {
	if clusterID.Id == 0 {
		return nil, errors.New("cluster id is required")
	}
	err := c.clusterUc.Delete(ctx, clusterID.Id)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{}, nil
}

func (c *ClusterInterface) StartCluster(ctx context.Context, clusterID *v1alpha1.ClusterArgs) (*v1alpha1.Msg, error) {
	if clusterID == nil || clusterID.Id == 0 {
		return nil, errors.New("cluster id is required")
	}
	cluster, err := c.clusterUc.Get(ctx, clusterID.Id)
	if err != nil {
		return nil, err
	}
	if cluster.Type == biz.ClusterTypeLocal && len(cluster.Nodes) == 0 {
		return nil, errors.New("at least one node is required")
	}
	err = c.clusterUc.Apply(ctx, cluster)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{}, nil
}

// check bostion host data and resources
func (c *ClusterInterface) CheckBostionHost(ctx context.Context, req *v1alpha1.CheckBostionHostRequest) (*v1alpha1.Msg, error) {
	if req.Arch == "" {
		return nil, errors.New("arch is required")
	}
	if req.OceanVersion == "" {
		return nil, errors.New("ocean version is required")
	}
	if req.ShipVersion == "" {
		return nil, errors.New("ship version is required")
	}
	if req.OceanDataTarGzPackagePath == "" {
		return nil, errors.New("ocean data tar gz package path is required")
	}
	if req.OceanDataTarGzPackageSha256SumPath == "" {
		return nil, errors.New("ocean data tar gz package sha256sum path is required")
	}
	if req.OceanPath == "" {
		return nil, errors.New("ocean path is required")
	}
	if req.ShipPath == "" {
		return nil, errors.New("ship path is required")
	}
	if req.Arch != runtime.GOOS {
		return nil, errors.New("arch is wrong")
	}
	// check ocean data tar gz package
	if ok := utils.IsFileExist(req.OceanDataTarGzPackagePath); !ok {
		return nil, errors.New("ocean data tar gz package is not exist")
	}
	// check ship
	if ok := utils.IsFileExist(req.ShipPath); !ok {
		return nil, errors.New("ship is not exist")
	}
	// check ocean data tar gz package sha256sum
	output, err := exec.Command("sudo", "sha256sum", "-c", req.OceanDataTarGzPackageSha256SumPath).CombinedOutput()
	if err != nil {
		return nil, errors.New(string(output))
	}
	// check ocean
	if ok := utils.IsFileExist(req.OceanPath); !ok {
		return nil, errors.New("ocean is not exist")
	}
	return &v1alpha1.Msg{}, nil
}

// get regions
func (c *ClusterInterface) GetRegions(ctx context.Context, clusterID *v1alpha1.ClusterArgs) (*v1alpha1.Regions, error) {
	if clusterID == nil || clusterID.Id == 0 {
		return nil, errors.New("cluster id is required")
	}
	cluster, err := c.clusterUc.Get(ctx, clusterID.Id)
	if err != nil {
		return nil, err
	}
	regions, err := c.clusterUc.GetRegions(ctx, cluster)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Regions{Regions: regions}, nil
}

// get logs
func (c *ClusterInterface) GetLogs(stream v1alpha1.ClusterInterface_GetLogsServer) error {
	i := 0
	for {
		ctx, cancel := context.WithCancel(stream.Context())
		defer cancel()

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if i > 0 {
			c.log.Info("repeat message, don't need to process")
			continue
		}
		i++
		if req.TailLines == 0 {
			req.TailLines = 30
		}
		clusterName := c.c.Server.GetClusterName()
		if req.ClusterName != clusterName {
			return errors.New("cluster name mismatch")
		}

		clusterLogPath, err := utils.GetLogFilePath(c.c.Server.Name)
		if err != nil {
			return err
		}
		if ok := utils.IsFileExist(clusterLogPath); !ok {
			return errors.New("cluster log does not exist")
		}

		file, err := os.Open(clusterLogPath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Read initial lines if TailLines is specified
		if req.TailLines > 0 {
			initialLogs, err := utils.ReadLastNLines(file, int(req.TailLines))
			if err != nil {
				return err
			}
			err = stream.Send(&v1alpha1.ClusterLogsResponse{Logs: initialLogs})
			if err != nil {
				return err
			}
		}

		// Move to the end of the file
		_, err = file.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}

		// Start watching for new logs
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return err
		}
		defer watcher.Close()

		err = watcher.Add(clusterLogPath)
		if err != nil {
			return err
		}

		// get ship logs
		shipLogContentChan := make(chan string)
		defer close(shipLogContentChan)
		cluster, err := c.clusterUc.Get(ctx, req.ClusterId)
		if err != nil {
			return err
		}
		if cluster != nil {
			for _, node := range cluster.Nodes {
				err = c.getShipLogContent(ctx, shipLogContentChan, node.InternalIP, node.SshPort)
				if err != nil {
					return err
				}
			}
		}

		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						newLogs, err := readNewLines(file)
						if err != nil {
							return
						}
						if newLogs != "" {
							err = stream.Send(&v1alpha1.ClusterLogsResponse{Logs: newLogs})
							if err != nil {
								return
							}
						}
					}
				case shipLogContent, ok := <-shipLogContentChan:
					if !ok {
						c.log.Info("Ship GetLogs stream closed by ship content")
						return
					}
					err = stream.Send(&v1alpha1.ClusterLogsResponse{Logs: shipLogContent})
					if err != nil {
						c.log.Errorf("Error sending ship log message: %v", err)
						return
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					c.log.Errorf("error watching log file: %v", err)
				case <-ctx.Done():
					c.log.Info("GetLogs stream closed by client")
					return
				}
			}
		}()
	}
}

func (c *ClusterInterface) getShipLogContent(ctx context.Context, contentChan chan string, nodeIp string, nodePort int32) error {
	conn, err := grpc.DialInsecure(
		ctx, // with cancel
		grpc.WithEndpoint(fmt.Sprintf("%s:%d", nodeIp, nodePort)),
	)
	if err != nil {
		return err
	}
	client := systemv1alpha1.NewSystemInterfaceClient(conn)
	stream, err := client.GetLogs(ctx)
	if err != nil {
		return err
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				c.log.Errorf("Error receiving ship log message: %v", err)
				return
			}
			contentChan <- msg.Log
		}
	}()

	err = stream.Send(&systemv1alpha1.LogRequest{
		TailLines: 30,
	})
	if err != nil {
		return err
	}
	return nil
}

func readNewLines(file *os.File) (string, error) {
	currentPos, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		return "", err
	}

	newContent, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	if len(newContent) > 0 {
		_, err = file.Seek(currentPos+int64(len(newContent)), io.SeekStart)
		if err != nil {
			return "", err
		}
		return string(newContent), nil
	}

	return "", nil
}

func (c *ClusterInterface) bizCLusterToCluster(bizCluster *biz.Cluster) *v1alpha1.Cluster {
	nodes := make([]*v1alpha1.Node, 0)
	for _, v := range bizCluster.Nodes {
		if v == nil {
			continue
		}
		nodes = append(nodes, c.bizNodeToNode(v))
	}
	nodeGroups := make([]*v1alpha1.NodeGroup, 0)
	for _, v := range bizCluster.NodeGroups {
		if v == nil {
			continue
		}
		nodeGroups = append(nodeGroups, c.bizNodeGroupToNodeGroup(v))
	}
	var bostionHost *v1alpha1.BostionHost
	if bizCluster.BostionHost != nil {
		bostionHost = c.bizBostionHostToBostionHost(bizCluster.BostionHost)
	}
	return &v1alpha1.Cluster{
		Id:                   bizCluster.ID,
		Name:                 bizCluster.Name,
		CloudId:              bizCluster.CloudID,
		Connections:          bizCluster.Connections,
		CertificateAuthority: bizCluster.CertificateAuthority,
		Version:              bizCluster.Version,
		ApiServerAddress:     bizCluster.ApiServerAddress,
		Config:               bizCluster.Config,
		Addons:               bizCluster.Addons,
		AddonsConfig:         bizCluster.AddonsConfig,
		Status:               uint32(bizCluster.Status),
		Type:                 string(bizCluster.Type),
		KubeConfig:           bizCluster.KubeConfig,
		KeyPair:              bizCluster.KeyPair,
		PublicKey:            bizCluster.PublicKey,
		PrivateKey:           bizCluster.PrivateKey,
		Region:               bizCluster.Region,
		VpcId:                bizCluster.VpcID,
		VpcCidr:              bizCluster.VpcCidr,
		EipId:                bizCluster.EipID,
		NatGatewayId:         bizCluster.NatGatewayID,
		ResourceGroupId:      bizCluster.ResourceGroupID,
		SecurityGroupIds:     bizCluster.SecurityGroupIDs,
		ExternalIp:           bizCluster.ExternalIP,
		AccessId:             bizCluster.AccessID,
		AccessKey:            bizCluster.AccessKey,
		LoadBalancerId:       bizCluster.LoadBalancerID,
		Nodes:                nodes,
		NodeGroups:           nodeGroups,
		BostionHost:          bostionHost,
	}
}

func (c *ClusterInterface) bizNodeToNode(bizNode *biz.Node) *v1alpha1.Node {
	return &v1alpha1.Node{
		Id:                      bizNode.ID,
		InstanceId:              bizNode.InstanceID,
		Name:                    bizNode.Name,
		Labels:                  bizNode.Labels,
		Kernel:                  bizNode.Kernel,
		ContainerRuntime:        bizNode.ContainerRuntime,
		Kubelet:                 bizNode.Kubelet,
		KubeProxy:               bizNode.KubeProxy,
		SshPort:                 int32(bizNode.SshPort),
		InternalIp:              bizNode.InternalIP,
		ExternalIp:              bizNode.ExternalIP,
		User:                    bizNode.User,
		Role:                    string(bizNode.Role),
		Status:                  string(bizNode.Status),
		ErrorInfo:               bizNode.ErrorInfo,
		Zone:                    bizNode.Zone,
		SubnetId:                bizNode.SubnetId,
		SubnetCidr:              bizNode.SubnetCidr,
		PublicKey:               bizNode.PublicKey,
		GpuSpec:                 bizNode.GpuSpec,
		SystemDisk:              int32(bizNode.SystemDisk),
		DataDisk:                int32(bizNode.DataDisk),
		NodePrice:               float32(bizNode.NodePrice),
		PodPrice:                float32(bizNode.PodPrice),
		InternetMaxBandwidthOut: int32(bizNode.InternetMaxBandwidthOut),
		ClusterId:               bizNode.ClusterID,
		NodeGroupId:             bizNode.NodeGroupID,
	}
}

func (c *ClusterInterface) bizNodeGroupToNodeGroup(bizNodeGroup *biz.NodeGroup) *v1alpha1.NodeGroup {
	return &v1alpha1.NodeGroup{
		Id:               bizNodeGroup.ID,
		Name:             bizNodeGroup.Name,
		CloudNodegroupId: bizNodeGroup.CloudNoodGroupID,
		Type:             string(bizNodeGroup.Type),
		InstanceType:     bizNodeGroup.InstanceType,
		Image:            bizNodeGroup.Image,
		Os:               bizNodeGroup.OS,
		Arch:             bizNodeGroup.ARCH,
		Cpu:              int32(bizNodeGroup.CPU),
		Memory:           float32(bizNodeGroup.Memory),
		Gpu:              int32(bizNodeGroup.GPU),
		NodeInitScript:   bizNodeGroup.NodeInitScript,
		MinSize:          int32(bizNodeGroup.MinSize),
		MaxSize:          int32(bizNodeGroup.MaxSize),
		TargetSize:       int32(bizNodeGroup.TargetSize),
		SystemDisk:       int32(bizNodeGroup.SystemDisk),
		DataDisk:         int32(bizNodeGroup.DataDisk),
		ClusterId:        int64(bizNodeGroup.ClusterID),
	}
}

func (c *ClusterInterface) bizBostionHostToBostionHost(bizBostionHost *biz.BostionHost) *v1alpha1.BostionHost {
	return &v1alpha1.BostionHost{
		Id:           bizBostionHost.ID,
		InstanceType: bizBostionHost.InstanceType,
		InstanceId:   bizBostionHost.InstanceID,
		User:         bizBostionHost.User,
		ImageId:      bizBostionHost.ImageID,
		Image:        bizBostionHost.Image,
		Os:           bizBostionHost.OS,
		Arch:         bizBostionHost.ARCH,
		Hostname:     bizBostionHost.Hostname,
		ExternalIp:   bizBostionHost.ExternalIP,
		InternalIp:   bizBostionHost.InternalIP,
		SshPort:      int32(bizBostionHost.SshPort),
		PrivateIp:    bizBostionHost.PrivateIP,
		ClusterId:    int64(bizBostionHost.ClusterID),
		Cpu:          int32(bizBostionHost.CPU),
		Memory:       float32(bizBostionHost.Memory),
	}
}
