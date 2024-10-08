// Code generated by MockGen. DO NOT EDIT.
// Source: internal/biz/cluster.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	biz "github.com/f-rambo/ocean/internal/biz"
	gomock "github.com/golang/mock/gomock"
)

// MockClusterRepo is a mock of ClusterRepo interface.
type MockClusterRepo struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRepoMockRecorder
}

// MockClusterRepoMockRecorder is the mock recorder for MockClusterRepo.
type MockClusterRepoMockRecorder struct {
	mock *MockClusterRepo
}

// NewMockClusterRepo creates a new mock instance.
func NewMockClusterRepo(ctrl *gomock.Controller) *MockClusterRepo {
	mock := &MockClusterRepo{ctrl: ctrl}
	mock.recorder = &MockClusterRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRepo) EXPECT() *MockClusterRepoMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockClusterRepo) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterRepo)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockClusterRepo) Get(arg0 context.Context, arg1 int64) (*biz.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*biz.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClusterRepoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterRepo)(nil).Get), arg0, arg1)
}

// GetByName mocks base method.
func (m *MockClusterRepo) GetByName(arg0 context.Context, arg1 string) (*biz.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0, arg1)
	ret0, _ := ret[0].(*biz.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockClusterRepoMockRecorder) GetByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockClusterRepo)(nil).GetByName), arg0, arg1)
}

// List mocks base method.
func (m *MockClusterRepo) List(arg0 context.Context, arg1 *biz.Cluster) ([]*biz.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*biz.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockClusterRepoMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterRepo)(nil).List), arg0, arg1)
}

// Put mocks base method.
func (m *MockClusterRepo) Put(ctx context.Context, cluster *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockClusterRepoMockRecorder) Put(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockClusterRepo)(nil).Put), ctx, cluster)
}

// Save mocks base method.
func (m *MockClusterRepo) Save(arg0 context.Context, arg1 *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockClusterRepoMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockClusterRepo)(nil).Save), arg0, arg1)
}

// Watch mocks base method.
func (m *MockClusterRepo) Watch(ctx context.Context) (*biz.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", ctx)
	ret0, _ := ret[0].(*biz.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockClusterRepoMockRecorder) Watch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterRepo)(nil).Watch), ctx)
}

// MockClusterInfrastructure is a mock of ClusterInfrastructure interface.
type MockClusterInfrastructure struct {
	ctrl     *gomock.Controller
	recorder *MockClusterInfrastructureMockRecorder
}

// MockClusterInfrastructureMockRecorder is the mock recorder for MockClusterInfrastructure.
type MockClusterInfrastructureMockRecorder struct {
	mock *MockClusterInfrastructure
}

// NewMockClusterInfrastructure creates a new mock instance.
func NewMockClusterInfrastructure(ctrl *gomock.Controller) *MockClusterInfrastructure {
	mock := &MockClusterInfrastructure{ctrl: ctrl}
	mock.recorder = &MockClusterInfrastructureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterInfrastructure) EXPECT() *MockClusterInfrastructureMockRecorder {
	return m.recorder
}

// AddNodes mocks base method.
func (m *MockClusterInfrastructure) AddNodes(arg0 context.Context, arg1 *biz.Cluster, arg2 []*biz.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNodes indicates an expected call of AddNodes.
func (mr *MockClusterInfrastructureMockRecorder) AddNodes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNodes", reflect.TypeOf((*MockClusterInfrastructure)(nil).AddNodes), arg0, arg1, arg2)
}

// DistributeDaemonApp mocks base method.
func (m *MockClusterInfrastructure) DistributeDaemonApp(arg0 context.Context, arg1 *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DistributeDaemonApp", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistributeDaemonApp indicates an expected call of DistributeDaemonApp.
func (mr *MockClusterInfrastructureMockRecorder) DistributeDaemonApp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeDaemonApp", reflect.TypeOf((*MockClusterInfrastructure)(nil).DistributeDaemonApp), arg0, arg1)
}

// GetNodesSystemInfo mocks base method.
func (m *MockClusterInfrastructure) GetNodesSystemInfo(arg0 context.Context, arg1 *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodesSystemInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetNodesSystemInfo indicates an expected call of GetNodesSystemInfo.
func (mr *MockClusterInfrastructureMockRecorder) GetNodesSystemInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodesSystemInfo", reflect.TypeOf((*MockClusterInfrastructure)(nil).GetNodesSystemInfo), arg0, arg1)
}

// GetRegions mocks base method.
func (m *MockClusterInfrastructure) GetRegions(arg0 context.Context, arg1 *biz.Cluster) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegions", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegions indicates an expected call of GetRegions.
func (mr *MockClusterInfrastructureMockRecorder) GetRegions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegions", reflect.TypeOf((*MockClusterInfrastructure)(nil).GetRegions), arg0, arg1)
}

// Install mocks base method.
func (m *MockClusterInfrastructure) Install(arg0 context.Context, arg1 *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockClusterInfrastructureMockRecorder) Install(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockClusterInfrastructure)(nil).Install), arg0, arg1)
}

// MigrateToBostionHost mocks base method.
func (m *MockClusterInfrastructure) MigrateToBostionHost(arg0 context.Context, arg1 *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateToBostionHost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MigrateToBostionHost indicates an expected call of MigrateToBostionHost.
func (mr *MockClusterInfrastructureMockRecorder) MigrateToBostionHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateToBostionHost", reflect.TypeOf((*MockClusterInfrastructure)(nil).MigrateToBostionHost), arg0, arg1)
}

// RemoveNodes mocks base method.
func (m *MockClusterInfrastructure) RemoveNodes(arg0 context.Context, arg1 *biz.Cluster, arg2 []*biz.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNodes indicates an expected call of RemoveNodes.
func (mr *MockClusterInfrastructureMockRecorder) RemoveNodes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNodes", reflect.TypeOf((*MockClusterInfrastructure)(nil).RemoveNodes), arg0, arg1, arg2)
}

// Start mocks base method.
func (m *MockClusterInfrastructure) Start(arg0 context.Context, arg1 *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockClusterInfrastructureMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockClusterInfrastructure)(nil).Start), arg0, arg1)
}

// Stop mocks base method.
func (m *MockClusterInfrastructure) Stop(arg0 context.Context, arg1 *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockClusterInfrastructureMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockClusterInfrastructure)(nil).Stop), arg0, arg1)
}

// UnInstall mocks base method.
func (m *MockClusterInfrastructure) UnInstall(arg0 context.Context, arg1 *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnInstall", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnInstall indicates an expected call of UnInstall.
func (mr *MockClusterInfrastructureMockRecorder) UnInstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnInstall", reflect.TypeOf((*MockClusterInfrastructure)(nil).UnInstall), arg0, arg1)
}

// MockClusterRuntime is a mock of ClusterRuntime interface.
type MockClusterRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRuntimeMockRecorder
}

// MockClusterRuntimeMockRecorder is the mock recorder for MockClusterRuntime.
type MockClusterRuntimeMockRecorder struct {
	mock *MockClusterRuntime
}

// NewMockClusterRuntime creates a new mock instance.
func NewMockClusterRuntime(ctrl *gomock.Controller) *MockClusterRuntime {
	mock := &MockClusterRuntime{ctrl: ctrl}
	mock.recorder = &MockClusterRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRuntime) EXPECT() *MockClusterRuntimeMockRecorder {
	return m.recorder
}

// CurrentCluster mocks base method.
func (m *MockClusterRuntime) CurrentCluster(arg0 context.Context, arg1 *biz.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentCluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CurrentCluster indicates an expected call of CurrentCluster.
func (mr *MockClusterRuntimeMockRecorder) CurrentCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentCluster", reflect.TypeOf((*MockClusterRuntime)(nil).CurrentCluster), arg0, arg1)
}
