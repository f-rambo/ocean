package interfaces

import (
	"context"

	"github.com/f-rambo/ocean/api/cluster/v1alpha1"
	"github.com/f-rambo/ocean/internal/biz"
	"github.com/f-rambo/ocean/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
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

func (uc *ClusterInterface) StartReconcile(ctx context.Context) error {
	for {
		cluser, err := uc.clusterUc.Watch(ctx)
		if err != nil {
			return err
		}
		if cluser == nil {
			continue
		}
		err = uc.clusterUc.Reconcile(ctx, cluser)
		if err != nil {
			return err
		}
	}
}

func (uc *ClusterInterface) StopReconcile(ctx context.Context) error {
	return nil
}

func (c *ClusterInterface) Ping(ctx context.Context, _ *emptypb.Empty) (*v1alpha1.Msg, error) {
	return &v1alpha1.Msg{Message: "pong"}, nil
}

func (c *ClusterInterface) Get(ctx context.Context, clusterID *v1alpha1.ClusterID) (*v1alpha1.Cluster, error) {
	cluster, err := c.clusterUc.Get(ctx, clusterID.Id)
	if err != nil {
		return nil, err
	}
	data := &v1alpha1.Cluster{}
	if cluster == nil {
		return data, nil
	}
	data = c.bizCLusterToCluster(cluster)
	return data, nil
}

func (c *ClusterInterface) Save(ctx context.Context, clusterArgs *v1alpha1.ClusterArgs) (*v1alpha1.Cluster, error) {
	if clusterArgs.Name == "" {
		return nil, errors.New("cluster name is required")
	}
	if clusterArgs.PublicKey == "" {
		return nil, errors.New("public key is required")
	}
	if clusterArgs.ServerType == "" {
		return nil, errors.New("server type is required")
	}
	if biz.ClusterType(clusterArgs.ServerType) != biz.ClusterTypeLocal {
		if clusterArgs.AccessKeyId == "" {
			return nil, errors.New("access key id is required")
		}
		if clusterArgs.SecretAccessKey == "" {
			return nil, errors.New("secret access key is required")
		}
		if clusterArgs.Region == "" {
			return nil, errors.New("region is required")
		}
	}
	cluster := &biz.Cluster{
		Name:      clusterArgs.Name,
		Type:      clusterArgs.ServerType,
		PublicKey: clusterArgs.PublicKey,
		Region:    clusterArgs.Region,
		AccessID:  clusterArgs.AccessKeyId,
		AccessKey: clusterArgs.SecretAccessKey,
		Nodes:     make([]*biz.Node, 0),
	}
	if biz.ClusterType(clusterArgs.ServerType) == biz.ClusterTypeLocal {
		if len(clusterArgs.Nodes) == 0 {
			return nil, errors.New("at least one node is required")
		}
		for _, node := range clusterArgs.Nodes {
			cluster.Nodes = append(cluster.Nodes, &biz.Node{
				ExternalIP: node.Ip,
				User:       node.User,
			})
		}
	}
	err := c.clusterUc.Save(ctx, cluster)
	if err != nil {
		return nil, err
	}
	// first
	cluster.Status = uint8(biz.ClusterStatucCreating)
	err = c.clusterUc.Apply(ctx, cluster)
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

func (c *ClusterInterface) Delete(ctx context.Context, clusterID *v1alpha1.ClusterID) (*v1alpha1.Msg, error) {
	if clusterID.Id == 0 {
		return nil, errors.New("cluster id is required")
	}
	err := c.clusterUc.Delete(ctx, clusterID.Id)
	if err != nil {
		return nil, err
	}
	return &v1alpha1.Msg{}, nil
}

func (c *ClusterInterface) bizCLusterToCluster(bizCluster *biz.Cluster) *v1alpha1.Cluster {
	cluster := &v1alpha1.Cluster{
		Id:               bizCluster.ID,
		Name:             bizCluster.Name,
		ServerVersion:    bizCluster.ServerVersion,
		ApiServerAddress: bizCluster.ApiServerAddress,
		Config:           bizCluster.Config,
		Addons:           bizCluster.Addons,
		Logs:             bizCluster.Logs,
	}
	for _, node := range bizCluster.Nodes {
		cluster.Nodes = append(cluster.Nodes, c.bizNodeToNode(node))
	}
	return cluster
}

func (c *ClusterInterface) bizNodeToNode(bizNode *biz.Node) *v1alpha1.Node {
	node := &v1alpha1.Node{
		Id:         bizNode.ID,
		Name:       bizNode.Name,
		Labels:     bizNode.Labels,
		OsImage:    bizNode.NodeGroup.OSImage,
		Kernel:     bizNode.Kernel,
		Container:  bizNode.Container,
		Kubelet:    bizNode.Kubelet,
		KubeProxy:  bizNode.KubeProxy,
		InternalIp: bizNode.InternalIP,
		ExternalIp: bizNode.ExternalIP,
		User:       bizNode.User,
		Role:       bizNode.Role,
		ClusterId:  bizNode.ClusterID,
	}
	return node
}
