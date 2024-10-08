package helm

import (
	"context"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/f-rambo/ocean/internal/biz"
	"github.com/f-rambo/ocean/internal/conf"
	"github.com/f-rambo/ocean/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	releasePkg "helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
)

type AppConstructRepo struct {
	log *log.Helper
	c   *conf.Bootstrap
}

func NewAppConstructRepo(c *conf.Bootstrap, logger log.Logger) biz.AppConstruct {
	return &AppConstructRepo{
		log: log.NewHelper(logger),
		c:   c,
	}
}

func (r *AppConstructRepo) GetAppVersionChartInfomation(ctx context.Context, appVersion *biz.AppVersion) error {
	appPath, err := utils.GetPackageStorePathByNames(biz.AppPackageName)
	if err != nil {
		return err
	}
	charInfo, err := GetLocalChartInfo(appVersion.AppName, appPath, appVersion.Chart)
	if err != nil {
		return err
	}
	charInfoMetadata, err := json.Marshal(charInfo.Metadata)
	if err != nil {
		return err
	}
	appVersion.Name = charInfo.Name
	appVersion.Config = charInfo.Config
	appVersion.Readme = charInfo.Readme
	appVersion.Description = charInfo.Description
	appVersion.Metadata = charInfoMetadata
	appVersion.Version = charInfo.Version
	appVersion.AppName = charInfo.Name
	appVersion.Chart = charInfo.Chart
	return nil
}

func (r *AppConstructRepo) DeployingApp(ctx context.Context, appDeployed *biz.DeployApp) error {
	helmPkg, err := NewHelmPkg(r.log, appDeployed.Namespace)
	if err != nil {
		return err
	}
	install, err := helmPkg.NewInstall()
	if err != nil {
		return err
	}
	appPath, err := utils.GetPackageStorePathByNames(biz.AppPackageName)
	if err != nil {
		return err
	}
	chartPath := filepath.Join(appPath, appDeployed.Chart)
	if appDeployed.AppTypeID == biz.AppTypeRepo {
		chartPath = filepath.Join(chartPath, biz.AppPackageRepoName, appDeployed.AppName, appDeployed.Chart)
	}
	install.ReleaseName = appDeployed.ReleaseName
	install.Namespace = appDeployed.Namespace
	install.CreateNamespace = true
	install.GenerateName = true
	install.Version = appDeployed.Version
	install.DryRun = appDeployed.IsTest
	install.Atomic = true
	install.Wait = true
	release, err := helmPkg.RunInstall(ctx, install, chartPath, appDeployed.Config)
	appDeployed.Logs = helmPkg.GetLogs()
	if err != nil {
		return err
	}
	if release != nil {
		appDeployed.ReleaseName = release.Name
		appDeployed.Manifest = strings.TrimSpace(release.Manifest)
		if release.Info != nil {
			appDeployed.State = string(release.Info.Status)
			appDeployed.Notes = release.Info.Notes
		}
		return nil
	}
	appDeployed.State = releasePkg.StatusUnknown.String()
	return nil
}

func (r *AppConstructRepo) UnDeployingApp(ctx context.Context, appDeployed *biz.DeployApp) error {
	helmPkg, err := NewHelmPkg(r.log, appDeployed.Namespace)
	if err != nil {
		return err
	}
	uninstall, err := helmPkg.NewUninstall()
	if err != nil {
		return err
	}
	uninstall.KeepHistory = false
	uninstall.DryRun = appDeployed.IsTest
	uninstall.Wait = true
	resp, err := helmPkg.RunUninstall(uninstall, appDeployed.ReleaseName)
	appDeployed.Logs = helmPkg.GetLogs()
	if err != nil {
		return errors.WithMessage(err, "uninstall fail")
	}
	if resp != nil && resp.Release != nil && resp.Release.Info != nil {
		appDeployed.State = string(resp.Release.Info.Status)
	}
	appDeployed.Notes = resp.Info
	return nil
}

func (r *AppConstructRepo) AddAppRepo(ctx context.Context, helmRepo *biz.AppHelmRepo) (err error) {
	settings := cli.New()
	res, err := repo.NewChartRepository(&repo.Entry{
		Name: helmRepo.Name,
		URL:  helmRepo.Url,
	}, getter.All(settings))
	if err != nil {
		return err
	}
	res.CachePath, err = utils.GetPackageStorePathByNames(biz.AppPackageName, biz.AppPackageRepoName)
	if err != nil {
		return err
	}
	indexFile, err := res.DownloadIndexFile()
	if err != nil {
		return err
	}
	helmRepo.SetIndexPath(indexFile)
	return nil
}

func (r *AppConstructRepo) GetAppDetailByRepo(ctx context.Context, helmRepo *biz.AppHelmRepo, appName, version string) (*biz.App, error) {
	index, err := repo.LoadIndexFile(helmRepo.IndexPath)
	if err != nil {
		return nil, err
	}
	app := &biz.App{
		Name:          appName,
		AppTypeID:     biz.AppTypeRepo,
		AppHelmRepoID: helmRepo.ID,
		Versions:      make([]*biz.AppVersion, 0),
	}
	for chartName, chartVersions := range index.Entries {
		if chartName != appName {
			continue
		}
		for i, chartMatedata := range chartVersions {
			if len(chartMatedata.URLs) == 0 {
				return nil, errors.New("chart urls is empty")
			}
			if app.Icon == "" {
				app.Icon = chartMatedata.Icon
			}
			appVersion := &biz.AppVersion{
				AppName:     chartName,
				Name:        chartMatedata.Name,
				Chart:       chartMatedata.URLs[0],
				Version:     chartMatedata.Version,
				Description: chartMatedata.Description,
			}
			if (version == "" && i == 0) || (version != "" && version == chartMatedata.Version) {
				err = r.GetAppVersionChartInfomation(ctx, appVersion)
				if err != nil {
					return nil, err
				}
			}
			app.AddVersion(appVersion)
		}
	}
	return app, nil
}

func (r *AppConstructRepo) GetAppsByRepo(ctx context.Context, helmRepo *biz.AppHelmRepo) ([]*biz.App, error) {
	index, err := repo.LoadIndexFile(helmRepo.IndexPath)
	if err != nil {
		return nil, err
	}
	apps := make([]*biz.App, 0)
	for chartName, chartVersions := range index.Entries {
		app := &biz.App{
			Name:          chartName,
			AppTypeID:     biz.AppTypeRepo,
			AppHelmRepoID: helmRepo.ID,
			Versions:      make([]*biz.AppVersion, 0),
		}
		app.CreatedAt = helmRepo.CreatedAt
		app.UpdatedAt = helmRepo.UpdatedAt
		for _, chartMatedata := range chartVersions {
			if app.Icon == "" {
				app.Icon = chartMatedata.Icon
			}
			if len(chartMatedata.URLs) == 0 {
				return nil, errors.New("chart urls is empty")
			}
			appVersion := &biz.AppVersion{
				AppName:     chartName,
				Name:        chartMatedata.Name,
				Chart:       chartMatedata.URLs[0],
				Version:     chartMatedata.Version,
				Description: chartMatedata.Description,
				State:       biz.AppTested,
			}
			app.AddVersion(appVersion)
		}
		apps = append(apps, app)
	}
	return apps, nil
}

func (r *AppConstructRepo) DeleteAppChart(ctx context.Context, app *biz.App, versionId int64) (err error) {
	appPath, err := utils.GetPackageStorePathByNames(biz.AppPackageName)
	if err != nil {
		return err
	}
	appIconPath := filepath.Join(appPath, biz.AppPathckageIcon, app.Icon)
	if app.Icon != "" && utils.IsFileExist(appIconPath) && versionId == 0 {
		err = utils.DeleteFile(appIconPath)
		if err != nil {
			return err
		}
	}

	for _, v := range app.Versions {
		chartPath := filepath.Join(appPath, v.Chart)
		if v.Chart != "" && utils.IsFileExist(chartPath) && versionId == 0 {
			err = utils.DeleteFile(chartPath)
			if err != nil {
				return err
			}
		}
		if v.Chart != "" && utils.IsFileExist(chartPath) && versionId == v.ID {
			err = utils.DeleteFile(chartPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
