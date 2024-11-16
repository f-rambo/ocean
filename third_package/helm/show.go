package helm

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/action"
	pkgChart "helm.sh/helm/v3/pkg/chart"
)

type ChartInfo struct {
	Name        string
	Config      string
	Readme      string
	Description string
	Metadata    pkgChart.Metadata
	Version     string
	AppName     string
	Chart       string
}

func GetLocalChartInfo(chart string) (*ChartInfo, error) {
	readme, err := action.NewShow(action.ShowReadme).Run(chart)
	if err != nil {
		return nil, errors.WithMessage(err, "show readme fail")
	}
	chartYaml, err := action.NewShow(action.ShowChart).Run(chart)
	if err != nil {
		return nil, errors.WithMessage(err, "show chart fail")
	}
	valuesYaml, err := action.NewShow(action.ShowValues).Run(chart)
	if err != nil {
		return nil, errors.WithMessage(err, "show values fail")
	}
	chartMateData := &pkgChart.Metadata{}
	err = yaml.Unmarshal([]byte(chartYaml), chartMateData)
	if err != nil {
		return nil, errors.WithMessage(err, "unmarshal chart yaml fail")
	}
	chartInfo := &ChartInfo{
		Name:        chartMateData.Name,
		Config:      valuesYaml,
		Readme:      readme,
		Description: chartMateData.Description,
		Metadata:    *chartMateData,
		Version:     chartMateData.Version,
		AppName:     chartMateData.Name,
		Chart:       chart,
	}
	return chartInfo, nil
}
