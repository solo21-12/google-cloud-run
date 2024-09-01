// Code generated by MockGen. DO NOT EDIT.
// Source: Domain/Interfaces/role_interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	Dtos "github.com/google-run-code/Domain/Dtos"
	Models "github.com/google-run-code/Domain/Models"
)

// MockRoleController is a mock of RoleController interface.
type MockRoleController struct {
	ctrl     *gomock.Controller
	recorder *MockRoleControllerMockRecorder
}

// MockRoleControllerMockRecorder is the mock recorder for MockRoleController.
type MockRoleControllerMockRecorder struct {
	mock *MockRoleController
}

// NewMockRoleController creates a new mock instance.
func NewMockRoleController(ctrl *gomock.Controller) *MockRoleController {
	mock := &MockRoleController{ctrl: ctrl}
	mock.recorder = &MockRoleControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleController) EXPECT() *MockRoleControllerMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockRoleController) CreateRole(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateRole", c)
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockRoleControllerMockRecorder) CreateRole(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockRoleController)(nil).CreateRole), c)
}

// DeleteRole mocks base method.
func (m *MockRoleController) DeleteRole(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteRole", c)
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockRoleControllerMockRecorder) DeleteRole(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockRoleController)(nil).DeleteRole), c)
}

// GetAllRoles mocks base method.
func (m *MockRoleController) GetAllRoles(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllRoles", c)
}

// GetAllRoles indicates an expected call of GetAllRoles.
func (mr *MockRoleControllerMockRecorder) GetAllRoles(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRoles", reflect.TypeOf((*MockRoleController)(nil).GetAllRoles), c)
}

// GetRoleById mocks base method.
func (m *MockRoleController) GetRoleById(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetRoleById", c)
}

// GetRoleById indicates an expected call of GetRoleById.
func (mr *MockRoleControllerMockRecorder) GetRoleById(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleById", reflect.TypeOf((*MockRoleController)(nil).GetRoleById), c)
}

// GetRoleUsers mocks base method.
func (m *MockRoleController) GetRoleUsers(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetRoleUsers", c)
}

// GetRoleUsers indicates an expected call of GetRoleUsers.
func (mr *MockRoleControllerMockRecorder) GetRoleUsers(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleUsers", reflect.TypeOf((*MockRoleController)(nil).GetRoleUsers), c)
}

// UpdateRole mocks base method.
func (m *MockRoleController) UpdateRole(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateRole", c)
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockRoleControllerMockRecorder) UpdateRole(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockRoleController)(nil).UpdateRole), c)
}

// MockRoleUseCase is a mock of RoleUseCase interface.
type MockRoleUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockRoleUseCaseMockRecorder
}

// MockRoleUseCaseMockRecorder is the mock recorder for MockRoleUseCase.
type MockRoleUseCaseMockRecorder struct {
	mock *MockRoleUseCase
}

// NewMockRoleUseCase creates a new mock instance.
func NewMockRoleUseCase(ctrl *gomock.Controller) *MockRoleUseCase {
	mock := &MockRoleUseCase{ctrl: ctrl}
	mock.recorder = &MockRoleUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleUseCase) EXPECT() *MockRoleUseCaseMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockRoleUseCase) CreateRole(role Dtos.RoleCreateRequest, ctx *gin.Context) (*Dtos.RoleResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", role, ctx)
	ret0, _ := ret[0].(*Dtos.RoleResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockRoleUseCaseMockRecorder) CreateRole(role, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockRoleUseCase)(nil).CreateRole), role, ctx)
}

// DeleteRole mocks base method.
func (m *MockRoleUseCase) DeleteRole(id string, ctx *gin.Context) *Models.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", id, ctx)
	ret0, _ := ret[0].(*Models.ErrorResponse)
	return ret0
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockRoleUseCaseMockRecorder) DeleteRole(id, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockRoleUseCase)(nil).DeleteRole), id, ctx)
}

// GetAllRoles mocks base method.
func (m *MockRoleUseCase) GetAllRoles(ctx *gin.Context) ([]*Dtos.RoleResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRoles", ctx)
	ret0, _ := ret[0].([]*Dtos.RoleResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// GetAllRoles indicates an expected call of GetAllRoles.
func (mr *MockRoleUseCaseMockRecorder) GetAllRoles(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRoles", reflect.TypeOf((*MockRoleUseCase)(nil).GetAllRoles), ctx)
}

// GetRoleById mocks base method.
func (m *MockRoleUseCase) GetRoleById(id string, ctx *gin.Context) (*Dtos.RoleResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleById", id, ctx)
	ret0, _ := ret[0].(*Dtos.RoleResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// GetRoleById indicates an expected call of GetRoleById.
func (mr *MockRoleUseCaseMockRecorder) GetRoleById(id, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleById", reflect.TypeOf((*MockRoleUseCase)(nil).GetRoleById), id, ctx)
}

// GetRoleUsers mocks base method.
func (m *MockRoleUseCase) GetRoleUsers(id string, ctx *gin.Context) ([]*Dtos.UserResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleUsers", id, ctx)
	ret0, _ := ret[0].([]*Dtos.UserResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// GetRoleUsers indicates an expected call of GetRoleUsers.
func (mr *MockRoleUseCaseMockRecorder) GetRoleUsers(id, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleUsers", reflect.TypeOf((*MockRoleUseCase)(nil).GetRoleUsers), id, ctx)
}

// UpdateRole mocks base method.
func (m *MockRoleUseCase) UpdateRole(id string, role Dtos.RoleUpdateRequest, ctx *gin.Context) (*Dtos.RoleResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", id, role, ctx)
	ret0, _ := ret[0].(*Dtos.RoleResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockRoleUseCaseMockRecorder) UpdateRole(id, role, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockRoleUseCase)(nil).UpdateRole), id, role, ctx)
}

// MockRoleRepository is a mock of RoleRepository interface.
type MockRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoleRepositoryMockRecorder
}

// MockRoleRepositoryMockRecorder is the mock recorder for MockRoleRepository.
type MockRoleRepositoryMockRecorder struct {
	mock *MockRoleRepository
}

// NewMockRoleRepository creates a new mock instance.
func NewMockRoleRepository(ctrl *gomock.Controller) *MockRoleRepository {
	mock := &MockRoleRepository{ctrl: ctrl}
	mock.recorder = &MockRoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleRepository) EXPECT() *MockRoleRepositoryMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockRoleRepository) CreateRole(role Dtos.RoleCreateRequest, ctx *gin.Context) (*Dtos.RoleResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", role, ctx)
	ret0, _ := ret[0].(*Dtos.RoleResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockRoleRepositoryMockRecorder) CreateRole(role, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockRoleRepository)(nil).CreateRole), role, ctx)
}

// DeleteRole mocks base method.
func (m *MockRoleRepository) DeleteRole(id string, ctx *gin.Context) *Models.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", id, ctx)
	ret0, _ := ret[0].(*Models.ErrorResponse)
	return ret0
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockRoleRepositoryMockRecorder) DeleteRole(id, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockRoleRepository)(nil).DeleteRole), id, ctx)
}

// GetAllRoles mocks base method.
func (m *MockRoleRepository) GetAllRoles(ctx *gin.Context) ([]*Dtos.RoleResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRoles", ctx)
	ret0, _ := ret[0].([]*Dtos.RoleResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// GetAllRoles indicates an expected call of GetAllRoles.
func (mr *MockRoleRepositoryMockRecorder) GetAllRoles(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRoles", reflect.TypeOf((*MockRoleRepository)(nil).GetAllRoles), ctx)
}

// GetRoleById mocks base method.
func (m *MockRoleRepository) GetRoleById(id string, ctx *gin.Context) (*Dtos.RoleResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleById", id, ctx)
	ret0, _ := ret[0].(*Dtos.RoleResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// GetRoleById indicates an expected call of GetRoleById.
func (mr *MockRoleRepositoryMockRecorder) GetRoleById(id, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleById", reflect.TypeOf((*MockRoleRepository)(nil).GetRoleById), id, ctx)
}

// GetRoleByNameAndRights mocks base method.
func (m *MockRoleRepository) GetRoleByNameAndRights(role Dtos.RoleCreateRequest, ctx *gin.Context) (*Models.Role, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleByNameAndRights", role, ctx)
	ret0, _ := ret[0].(*Models.Role)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// GetRoleByNameAndRights indicates an expected call of GetRoleByNameAndRights.
func (mr *MockRoleRepositoryMockRecorder) GetRoleByNameAndRights(role, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleByNameAndRights", reflect.TypeOf((*MockRoleRepository)(nil).GetRoleByNameAndRights), role, ctx)
}

// GetRoleUsers mocks base method.
func (m *MockRoleRepository) GetRoleUsers(role *Dtos.RoleResponse, ctx *gin.Context) ([]*Dtos.UserResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleUsers", role, ctx)
	ret0, _ := ret[0].([]*Dtos.UserResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// GetRoleUsers indicates an expected call of GetRoleUsers.
func (mr *MockRoleRepositoryMockRecorder) GetRoleUsers(role, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleUsers", reflect.TypeOf((*MockRoleRepository)(nil).GetRoleUsers), role, ctx)
}

// UpdateRole mocks base method.
func (m *MockRoleRepository) UpdateRole(id string, role Dtos.RoleUpdateRequest, ctx *gin.Context) (*Dtos.RoleResponse, *Models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", id, role, ctx)
	ret0, _ := ret[0].(*Dtos.RoleResponse)
	ret1, _ := ret[1].(*Models.ErrorResponse)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockRoleRepositoryMockRecorder) UpdateRole(id, role, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockRoleRepository)(nil).UpdateRole), id, role, ctx)
}
