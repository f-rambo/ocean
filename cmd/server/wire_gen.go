// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/f-rambo/ocean/internal/biz"
	"github.com/f-rambo/ocean/internal/conf"
	"github.com/f-rambo/ocean/internal/data"
	"github.com/f-rambo/ocean/internal/interfaces"
	"github.com/f-rambo/ocean/internal/server"
	"github.com/f-rambo/ocean/pkg/ansible"
	"github.com/f-rambo/ocean/pkg/argoworkflows"
	"github.com/f-rambo/ocean/pkg/githubapi"
	"github.com/f-rambo/ocean/pkg/helm"
	"github.com/f-rambo/ocean/pkg/kubernetes"
	"github.com/f-rambo/ocean/pkg/pulumi"
	"github.com/f-rambo/ocean/pkg/sailor"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "github.com/joho/godotenv/autoload"
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	clusterRepo := data.NewClusterRepo(dataData, bootstrap, logger)
	infrastructure := pulumi.NewClusterInfrastructure(bootstrap, logger)
	clusterConstruct := ansible.NewClusterConstruct(bootstrap, logger)
	clusterRuntime := kubernetes.NewClusterRuntime(bootstrap, logger)
	clusterUsecase := biz.NewClusterUseCase(clusterRepo, infrastructure, clusterConstruct, clusterRuntime, logger)
	projectRepo := data.NewProjectRepo(dataData, bootstrap, logger)
	clusterPorjectRepo := kubernetes.NewProjectClient(bootstrap, logger)
	projectUsecase := biz.NewProjectUseCase(projectRepo, clusterPorjectRepo, logger)
	appRepo := data.NewAppRepo(dataData, logger)
	sailorRepo := sailor.NewSailorClient(bootstrap, logger)
	appRuntime := kubernetes.NewAppDeployedResource(bootstrap, logger)
	appConstruct := helm.NewAppConstructRepo(bootstrap, logger)
	appUsecase := biz.NewAppUsecase(appRepo, clusterRepo, projectRepo, sailorRepo, appRuntime, appConstruct, logger, bootstrap)
	clusterInterface := interfaces.NewClusterInterface(clusterUsecase, projectUsecase, appUsecase, bootstrap, logger)
	userRepo := data.NewUserRepo(dataData, bootstrap, logger)
	thirdparty := githubapi.NewUserClient(bootstrap, logger)
	userUseCase := biz.NewUseUser(userRepo, thirdparty, logger)
	appInterface := interfaces.NewAppInterface(appUsecase, userUseCase, bootstrap, logger)
	servicesRepo := data.NewServicesRepo(dataData, logger)
	workflowRepo := argoworkflows.NewWorkflowRepo(bootstrap, logger)
	servicesUseCase := biz.NewServicesUseCase(servicesRepo, workflowRepo, logger)
	servicesInterface := interfaces.NewServicesInterface(servicesUseCase, projectUsecase)
	userInterface := interfaces.NewUserInterface(userUseCase, bootstrap)
	projectInterface := interfaces.NewProjectInterface(projectUsecase, appUsecase, clusterUsecase, bootstrap, logger)
	autoscaler := interfaces.NewAutoscaler(clusterUsecase, bootstrap, logger)
	grpcServer := server.NewGRPCServer(bootstrap, clusterInterface, appInterface, servicesInterface, userInterface, projectInterface, autoscaler, logger)
	httpServer := server.NewHTTPServer(bootstrap, clusterInterface, appInterface, servicesInterface, userInterface, projectInterface, logger)
	otherServer := server.NewOtherServer(clusterInterface)
	app := newApp(logger, grpcServer, httpServer, otherServer)
	return app, func() {
		cleanup()
	}, nil
}
