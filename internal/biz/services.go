package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Service struct {
	ID          int64   `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Name        string  `json:"name,omitempty" gorm:"column:name; default:''; NOT NULL"`
	CodeRepo    string  `json:"code_repo,omitempty" gorm:"column:code_repo; default:''; NOT NULL"`   // git repo url
	Replicas    int32   `json:"replicas" gorm:"column:replicas; default:0; NOT NULL"`                // 0 == auto
	CPU         float32 `json:"cpu" gorm:"column:cpu; default:0; NOT NULL"`                          // 0.5=500m
	LimitCpu    float32 `json:"limit_cpu" gorm:"column:limit_cpu; default:0; NOT NULL"`              // 0 == auto
	GPU         float32 `json:"gpu" gorm:"column:gpu; default:0; NOT NULL"`                          // 0 == no gpu
	LimitGPU    float32 `json:"limit_gpu" gorm:"column:limit_gpu; default:0; NOT NULL"`              // GPU !==0 LimitGPU 0 == auto
	Memory      float32 `json:"memory" gorm:"column:memory; default:0; NOT NULL"`                    // 0.5=500m
	LimitMemory float32 `json:"limit_memory" gorm:"column:limit_memory; default:0; NOT NULL"`        // 0 == auto
	Disk        float32 `json:"disk" gorm:"column:disk; default:0; NOT NULL"`                        // 0.5=500m
	LimitDisk   float32 `json:"limit_disk" gorm:"column:limit_disk; default:0; NOT NULL"`            // 0 == auto
	Business    string  `json:"business,omitempty" gorm:"column:business; default:''; NOT NULL"`     // business name
	Technology  string  `json:"technology,omitempty" gorm:"column:technology; default:''; NOT NULL"` // technology name
	WorkflowID  int64   `json:"workflow_id,omitempty" gorm:"column:workflow_id; default:0; NOT NULL"`
	Ports       []Port  `json:"ports" gorm:"-"`
	ProjectID   int64   `json:"project_id,omitempty" gorm:"column:project_id; default:0; NOT NULL"`
	PortsJson   []byte  `json:"ports_json" gorm:"column:ports_json; type:json"`
	gorm.Model
}

type Port struct {
	ID            int64  `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	IngressPath   string `json:"ingress_path" gorm:"column:ingress_path; default:''; NOT NULL"`    // /api/v1
	Protocol      string `json:"protocol" gorm:"column:protocol; default:''; NOT NULL"`            // TCP, UDP
	ContainerPort int32  `json:"container_port" gorm:"column:container_port; default:0; NOT NULL"` // 80
}

type CI struct {
	ID           int64  `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Version      string `json:"version,omitempty" gorm:"column:version; default:''; NOT NULL"`
	Branch       string `json:"branch,omitempty" gorm:"column:branch; default:''; NOT NULL"`
	Tag          string `json:"tag,omitempty" gorm:"column:tag; default:''; NOT NULL"`
	Args         string `json:"args,omitempty" gorm:"column:args; type:json"`
	Description  string `json:"description,omitempty" gorm:"column:description; default:''; NOT NULL"`
	WorkflowName string `json:"workflow_name,omitempty" gorm:"column:workflow_name; default:''; NOT NULL"`
	ServiceID    int64  `json:"service_id,omitempty" gorm:"column:service_id; default:0; NOT NULL"`
	Logs         string `json:"logs" gorm:"-"`
	gorm.Model
}

type CD struct {
	ID        int64  `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	ServiceID int64  `json:"service_id" gorm:"column:service_id; default:0; NOT NULL"`
	Logs      string `json:"logs" gorm:"-"`
	gorm.Model
}

type Worklfow struct {
	ID       int64  `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Name     string `json:"name,omitempty" gorm:"column:name; default:''; NOT NULL"`
	Worklfow []byte `json:"workflow" gorm:"column:workflow; type:json"`
}

type ServicesRepo interface {
	List(ctx context.Context, serviceParam *Service, page, pageSize int) ([]*Service, int64, error)
	Save(ctx context.Context, service *Service) error
	Get(ctx context.Context, id int64) (*Service, error)
	Delete(ctx context.Context, id int64) error
}

type ServicesUseCase struct {
	repo ServicesRepo
	log  *log.Helper
}

func NewServicesUseCase(repo ServicesRepo, logger log.Logger) *ServicesUseCase {
	return &ServicesUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ServicesUseCase) List(ctx context.Context, serviceParam *Service, page, pageSize int) ([]*Service, int64, error) {
	return uc.repo.List(ctx, serviceParam, page, pageSize)
}

func (uc *ServicesUseCase) Save(ctx context.Context, service *Service) error {
	return uc.repo.Save(ctx, service)
}

func (uc *ServicesUseCase) Get(ctx context.Context, id int64) (*Service, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *ServicesUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
