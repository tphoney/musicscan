// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tphoney/musicscan/internal/store (interfaces: AlbumStore,ArtistStore,MemberStore,ProjectStore,SystemStore,UserStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/tphoney/musicscan/types"
)

// MockAlbumStore is a mock of AlbumStore interface.
type MockAlbumStore struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumStoreMockRecorder
}

// MockAlbumStoreMockRecorder is the mock recorder for MockAlbumStore.
type MockAlbumStoreMockRecorder struct {
	mock *MockAlbumStore
}

// NewMockAlbumStore creates a new mock instance.
func NewMockAlbumStore(ctrl *gomock.Controller) *MockAlbumStore {
	mock := &MockAlbumStore{ctrl: ctrl}
	mock.recorder = &MockAlbumStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlbumStore) EXPECT() *MockAlbumStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAlbumStore) Create(arg0 context.Context, arg1 *types.Album) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAlbumStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAlbumStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockAlbumStore) Delete(arg0 context.Context, arg1 *types.Album) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAlbumStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAlbumStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockAlbumStore) Find(arg0 context.Context, arg1 int64) (*types.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockAlbumStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAlbumStore)(nil).Find), arg0, arg1)
}

// FindByName mocks base method.
func (m *MockAlbumStore) FindByName(arg0 context.Context, arg1 int64, arg2 string) (*types.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockAlbumStoreMockRecorder) FindByName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockAlbumStore)(nil).FindByName), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockAlbumStore) List(arg0 context.Context, arg1 int64, arg2 types.Params) ([]*types.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAlbumStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAlbumStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockAlbumStore) Update(arg0 context.Context, arg1 *types.Album) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAlbumStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAlbumStore)(nil).Update), arg0, arg1)
}

// MockArtistStore is a mock of ArtistStore interface.
type MockArtistStore struct {
	ctrl     *gomock.Controller
	recorder *MockArtistStoreMockRecorder
}

// MockArtistStoreMockRecorder is the mock recorder for MockArtistStore.
type MockArtistStoreMockRecorder struct {
	mock *MockArtistStore
}

// NewMockArtistStore creates a new mock instance.
func NewMockArtistStore(ctrl *gomock.Controller) *MockArtistStore {
	mock := &MockArtistStore{ctrl: ctrl}
	mock.recorder = &MockArtistStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArtistStore) EXPECT() *MockArtistStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArtistStore) Create(arg0 context.Context, arg1 *types.Artist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockArtistStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArtistStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockArtistStore) Delete(arg0 context.Context, arg1 *types.Artist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockArtistStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArtistStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockArtistStore) Find(arg0 context.Context, arg1 int64) (*types.Artist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.Artist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockArtistStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockArtistStore)(nil).Find), arg0, arg1)
}

// FindByName mocks base method.
func (m *MockArtistStore) FindByName(arg0 context.Context, arg1 string) (*types.Artist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0, arg1)
	ret0, _ := ret[0].(*types.Artist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockArtistStoreMockRecorder) FindByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockArtistStore)(nil).FindByName), arg0, arg1)
}

// List mocks base method.
func (m *MockArtistStore) List(arg0 context.Context, arg1 int64, arg2 types.Params) ([]*types.Artist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.Artist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockArtistStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArtistStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockArtistStore) Update(arg0 context.Context, arg1 *types.Artist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArtistStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArtistStore)(nil).Update), arg0, arg1)
}

// MockMemberStore is a mock of MemberStore interface.
type MockMemberStore struct {
	ctrl     *gomock.Controller
	recorder *MockMemberStoreMockRecorder
}

// MockMemberStoreMockRecorder is the mock recorder for MockMemberStore.
type MockMemberStoreMockRecorder struct {
	mock *MockMemberStore
}

// NewMockMemberStore creates a new mock instance.
func NewMockMemberStore(ctrl *gomock.Controller) *MockMemberStore {
	mock := &MockMemberStore{ctrl: ctrl}
	mock.recorder = &MockMemberStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemberStore) EXPECT() *MockMemberStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMemberStore) Create(arg0 context.Context, arg1 *types.Membership) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMemberStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMemberStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockMemberStore) Delete(arg0 context.Context, arg1, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMemberStoreMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMemberStore)(nil).Delete), arg0, arg1, arg2)
}

// Find mocks base method.
func (m *MockMemberStore) Find(arg0 context.Context, arg1, arg2 int64) (*types.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMemberStoreMockRecorder) Find(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMemberStore)(nil).Find), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockMemberStore) List(arg0 context.Context, arg1 int64, arg2 types.Params) ([]*types.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMemberStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMemberStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockMemberStore) Update(arg0 context.Context, arg1 *types.Membership) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMemberStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMemberStore)(nil).Update), arg0, arg1)
}

// MockProjectStore is a mock of ProjectStore interface.
type MockProjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockProjectStoreMockRecorder
}

// MockProjectStoreMockRecorder is the mock recorder for MockProjectStore.
type MockProjectStoreMockRecorder struct {
	mock *MockProjectStore
}

// NewMockProjectStore creates a new mock instance.
func NewMockProjectStore(ctrl *gomock.Controller) *MockProjectStore {
	mock := &MockProjectStore{ctrl: ctrl}
	mock.recorder = &MockProjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectStore) EXPECT() *MockProjectStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProjectStore) Create(arg0 context.Context, arg1 *types.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockProjectStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProjectStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockProjectStore) Delete(arg0 context.Context, arg1 *types.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProjectStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProjectStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockProjectStore) Find(arg0 context.Context, arg1 int64) (*types.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockProjectStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockProjectStore)(nil).Find), arg0, arg1)
}

// FindToken mocks base method.
func (m *MockProjectStore) FindToken(arg0 context.Context, arg1 string) (*types.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindToken", arg0, arg1)
	ret0, _ := ret[0].(*types.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindToken indicates an expected call of FindToken.
func (mr *MockProjectStoreMockRecorder) FindToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindToken", reflect.TypeOf((*MockProjectStore)(nil).FindToken), arg0, arg1)
}

// List mocks base method.
func (m *MockProjectStore) List(arg0 context.Context, arg1 int64, arg2 types.Params) ([]*types.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProjectStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProjectStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockProjectStore) Update(arg0 context.Context, arg1 *types.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProjectStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProjectStore)(nil).Update), arg0, arg1)
}

// MockSystemStore is a mock of SystemStore interface.
type MockSystemStore struct {
	ctrl     *gomock.Controller
	recorder *MockSystemStoreMockRecorder
}

// MockSystemStoreMockRecorder is the mock recorder for MockSystemStore.
type MockSystemStoreMockRecorder struct {
	mock *MockSystemStore
}

// NewMockSystemStore creates a new mock instance.
func NewMockSystemStore(ctrl *gomock.Controller) *MockSystemStore {
	mock := &MockSystemStore{ctrl: ctrl}
	mock.recorder = &MockSystemStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemStore) EXPECT() *MockSystemStoreMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockSystemStore) Config(arg0 context.Context) *types.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config", arg0)
	ret0, _ := ret[0].(*types.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockSystemStoreMockRecorder) Config(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockSystemStore)(nil).Config), arg0)
}

// MockUserStore is a mock of UserStore interface.
type MockUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMockRecorder
}

// MockUserStoreMockRecorder is the mock recorder for MockUserStore.
type MockUserStoreMockRecorder struct {
	mock *MockUserStore
}

// NewMockUserStore creates a new mock instance.
func NewMockUserStore(ctrl *gomock.Controller) *MockUserStore {
	mock := &MockUserStore{ctrl: ctrl}
	mock.recorder = &MockUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStore) EXPECT() *MockUserStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockUserStore) Count(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockUserStoreMockRecorder) Count(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockUserStore)(nil).Count), arg0)
}

// Create mocks base method.
func (m *MockUserStore) Create(arg0 context.Context, arg1 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockUserStore) Delete(arg0 context.Context, arg1 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockUserStore) Find(arg0 context.Context, arg1 int64) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserStore)(nil).Find), arg0, arg1)
}

// FindEmail mocks base method.
func (m *MockUserStore) FindEmail(arg0 context.Context, arg1 string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEmail", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEmail indicates an expected call of FindEmail.
func (mr *MockUserStoreMockRecorder) FindEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEmail", reflect.TypeOf((*MockUserStore)(nil).FindEmail), arg0, arg1)
}

// FindKey mocks base method.
func (m *MockUserStore) FindKey(arg0 context.Context, arg1 string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindKey", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindKey indicates an expected call of FindKey.
func (mr *MockUserStoreMockRecorder) FindKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKey", reflect.TypeOf((*MockUserStore)(nil).FindKey), arg0, arg1)
}

// List mocks base method.
func (m *MockUserStore) List(arg0 context.Context, arg1 types.Params) ([]*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserStore)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserStore) Update(arg0 context.Context, arg1 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserStore)(nil).Update), arg0, arg1)
}
