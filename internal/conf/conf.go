package conf

import (
	"os"
	"strconv"
	"strings"
)

type Bootstrap struct {
	Env      string   `json:"env,omitempty"`
	Server   Server   `json:"server,omitempty"`
	Data     Data     `json:"data,omitempty"`
	ETCD     ETCD     `json:"etcd,omitempty"`
	Log      Log      `json:"log,omitempty"`
	Auth     Auth     `json:"auth,omitempty"`
	Resource Resource `json:"resource,omitempty"`
	Business any      `json:"business,omitempty"`
	Ansible  any      `json:"ansible,omitempty"`
}

type ETCD struct {
	Username  string   `json:"username,omitempty"`
	Password  string   `json:"password,omitempty"`
	Endpoints []string `json:"endpoints,omitempty"`
}

func (e ETCD) GetETCDEndpoints() []string {
	endpoints := os.Getenv("ETCD_ENDPOINTS")
	endpointArr := strings.Split(endpoints, ",")
	if len(endpointArr) > 0 {
		return endpointArr
	}
	return e.Endpoints
}

func (e ETCD) GetUsername() string {
	username := os.Getenv("ETCD_USERNAME")
	if username != "" {
		return username
	}
	return e.Username
}

func (e ETCD) GetPassword() string {
	password := os.Getenv("ETCD_PASSWORD")
	if password != "" {
		return password
	}
	return e.Password
}

type Env string

const (
	EnvLocal       Env = "local"
	EnvBostionHost Env = "bostionhost"
	EnvCluster     Env = "cluster"
)

func (e Env) String() string {
	return string(e)
}

func (b Bootstrap) GetEnv() string {
	env := os.Getenv("ENV")
	if env != "" {
		return env
	}
	return b.Env
}

type Resource struct {
	App          string `json:"app,omitempty"`           // app chart package
	Icon         string `json:"icon,omitempty"`          // app icon img
	Repo         string `json:"repo,omitempty"`          // app repo
	Cluster      string `json:"cluster,omitempty"`       // cluster setting
	KubesprayUrl string `json:"kubespray_url,omitempty"` // kubespray url
	AnsibleCli   string `json:"ansible_cli,omitempty"`   // ansible cli
	PulumiPath   string `json:"pulumi_path,omitempty"`   // pulumi path
}

func (r Resource) GetClusterPath() string {
	clusterPath := os.Getenv("CLUSTER_PATH")
	if clusterPath == "" {
		clusterPath = r.Cluster
	}
	return clusterPath
}

func (r Resource) GetAppPath() string {
	appPath := os.Getenv("APP_PATH")
	if appPath == "" {
		appPath = r.App
	}
	return appPath
}

func (r Resource) GetIconPath() string {
	iconPath := os.Getenv("ICON_PATH")
	if iconPath == "" {
		iconPath = r.Icon
	}
	return iconPath
}

func (r Resource) GetRepoPath() string {
	repoPath := os.Getenv("REPO_PATH")
	if repoPath == "" {
		repoPath = r.Repo
	}
	return repoPath
}

func (r Resource) GetKubesprayUrl() string {
	kubesprayUrl := os.Getenv("KUBESPRAY_URL")
	if kubesprayUrl == "" {
		kubesprayUrl = r.KubesprayUrl
	}
	return kubesprayUrl
}

func (r Resource) GetAnsibleCli() string {
	ansibleCli := os.Getenv("ANSIBLE_CLI")
	if ansibleCli == "" {
		ansibleCli = r.AnsibleCli
	}
	return ansibleCli
}

func (r Resource) GetPulumiPath() string {
	pulumiPath := os.Getenv("PULUMI_PATH")
	if pulumiPath == "" {
		pulumiPath = r.PulumiPath
	}
	return pulumiPath
}

type Server struct {
	Name   string `json:"name,omitempty"`
	HTTP   HTTP   `json:"http,omitempty"`
	GRPC   GRPC   `json:"grpc,omitempty"`
	STATIC STATIC `json:"static,omitempty"`
}

type HTTP struct {
	Network string `json:"network,omitempty"`
	Addr    string `json:"addr,omitempty"`
}

func (s Server) GetName() string {
	name := os.Getenv("SERVER_NAME")
	if name != "" {
		return name
	}
	return s.Name
}

func (h HTTP) GetNetwork() string {
	newWork := os.Getenv("HTTP_NETWORK")
	if newWork != "" {
		return newWork
	}
	return h.Network
}

func (h HTTP) GetAddr() string {
	addr := os.Getenv("HTTP_ADDR")
	if addr != "" {
		return addr
	}
	return h.Addr
}

type GRPC struct {
	Network string `json:"network,omitempty"`
	Addr    string `json:"addr,omitempty"`
}

func (g GRPC) GetNetwork() string {
	netWork := os.Getenv("GRPC_NETWORK")
	if netWork != "" {
		return netWork
	}
	return g.Network
}

func (g GRPC) GetAddr() string {
	addr := os.Getenv("GRPC_ADDR")
	if addr != "" {
		return addr
	}
	return g.Addr
}

type STATIC struct {
	Network string `json:"network,omitempty"`
	Addr    string `json:"addr,omitempty"`
}

func (s STATIC) GetNetwork() string {
	netWork := os.Getenv("STATIC_NETWORK")
	if netWork != "" {
		return netWork
	}
	return s.Network
}

func (s STATIC) GetAddr() string {
	addr := os.Getenv("STATIC_ADDR")
	if addr != "" {
		return addr
	}
	return s.Addr
}

type Data struct {
	Driver     string `json:"driver,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Host       string `json:"host,omitempty"`
	Port       int32  `json:"port,omitempty"`
	Database   string `json:"database,omitempty"`
	DBFilePath string `json:"dbfilepath,omitempty"`
}

func (d Data) GetDriver() string {
	driver := os.Getenv("DATABASE_DRIVER")
	if driver != "" {
		return driver
	}
	return d.Driver
}

func (d Data) GetUsername() string {
	username := os.Getenv("DATABASE_USERNAME")
	if username != "" {
		return username
	}
	return d.Username
}

func (d Data) GetPassword() string {
	password := os.Getenv("DATABASE_PASSWORD")
	if password != "" {
		return password
	}
	return d.Password
}

func (d Data) GetHost() string {
	host := os.Getenv("DATABASE_HOST")
	if host != "" {
		return host
	}
	return d.Host
}

func (d Data) GetPort() int32 {
	port := os.Getenv("DATABASE_PORT")
	if port != "" {
		portInt, _ := strconv.Atoi(port)
		return int32(portInt)
	}
	return d.Port
}

func (d Data) GetDatabase() string {
	database := os.Getenv("DATABASE_DATABASE")
	if database != "" {
		return database
	}
	return d.Database
}

func (d Data) GetDBFilePath() string {
	dbFilePath := os.Getenv("DATABASE_DBFILEPATH")
	if dbFilePath != "" {
		return dbFilePath
	}
	return d.DBFilePath
}

type Log struct {
	Path       string `json:"path,omitempty"`
	Filename   string `json:"filename,omitempty"`
	MaxSize    int32  `json:"max_size,omitempty"`
	MaxAge     int32  `json:"max_age,omitempty"`
	MaxBackups int32  `json:"max_backups,omitempty"`
	Compress   bool   `json:"compress,omitempty"`
	LocalTime  bool   `json:"local_time,omitempty"`
}

func (l Log) GetPath() string {
	path := os.Getenv("LOG_PATH")
	if path != "" {
		return path
	}
	return l.Path
}

type Auth struct {
	Exp      int32  `json:"exp,omitempty"`
	Key      string `json:"key,omitempty"`
	Email    string `json:"email,omitempty"`
	PassWord string `json:"password,omitempty"`
}

func (a Auth) GetExp() int32 {
	exp := os.Getenv("AUTH_EXP")
	if exp != "" {
		expInt, _ := strconv.Atoi(exp)
		return int32(expInt)
	}
	return a.Exp
}

func (a Auth) GetKey() string {
	key := os.Getenv("AUTH_KEY")
	if key != "" {
		return key
	}
	return a.Key
}

func (a Auth) GetEmail() string {
	email := os.Getenv("AUTH_EMAIL")
	if email != "" {
		return email
	}
	return a.Email
}

func (a Auth) GetPassWord() string {
	password := os.Getenv("AUTH_PASSWORD")
	if password != "" {
		return password
	}
	return a.PassWord
}
