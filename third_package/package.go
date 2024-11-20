package third_package

import (
	"github.com/f-rambo/cloud-copilot/third_package/argoworkflows"
	"github.com/f-rambo/cloud-copilot/third_package/githubapi"
	"github.com/f-rambo/cloud-copilot/third_package/helm"
	infrastructure "github.com/f-rambo/cloud-copilot/third_package/infrastructure"
	"github.com/f-rambo/cloud-copilot/third_package/kubernetes"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	argoworkflows.NewWorkflowRepo,
	helm.NewAppConstructRepo,
	kubernetes.NewAppDeployedResource,
	kubernetes.NewClusterRuntime,
	kubernetes.NewProjectClient,
	githubapi.NewUserClient,
	infrastructure.NewClusterInfrastructure,
)
