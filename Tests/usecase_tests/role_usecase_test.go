package usecases_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	mocks "github.com/google-run-code/Tests/Mocks"
	usecases "github.com/google-run-code/Usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type RoleUsecaseTestSuite struct {
	suite.Suite
	roleRepoMock *mocks.MockRoleRepository
	userRepoMock *mocks.MockUserRepository
	roleUsecase  interfaces.RoleUseCase
	ctrl         *gomock.Controller
}

func (suite *RoleUsecaseTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.roleRepoMock = mocks.NewMockRoleRepository(suite.ctrl)
	suite.userRepoMock = mocks.NewMockUserRepository(suite.ctrl)
	suite.roleUsecase = usecases.NewRoleUseCase(suite.roleRepoMock, suite.userRepoMock)
}

func (suite *RoleUsecaseTestSuite) TearDownSuite() {
	suite.ctrl.Finish()
}

func (suite *RoleUsecaseTestSuite) TestCreateRole_Success() {
	ctx := &gin.Context{}
	roleReq := dtos.RoleCreateRequest{
		Name: "Test Role",
	}
	role := &dtos.RoleResponse{
		UID:    uuid.New().String(),
		Name:   roleReq.Name,
		Rights: roleReq.Rights,
	}

	suite.roleRepoMock.EXPECT().
		GetRoleByNameAndRights(roleReq, ctx).
		Return(nil, nil)

	suite.roleRepoMock.EXPECT().
		CreateRole(roleReq, ctx).
		Return(role, nil)

	result, err := suite.roleUsecase.CreateRole(roleReq, ctx)
	suite.Nil(err)
	suite.Equal(role, result)
}

func (suite *RoleUsecaseTestSuite) TestGetRoleById_Success() {
	ctx := &gin.Context{}
	roleID := "some-role-id"
	expectedRole := &dtos.RoleResponse{
		UID:  roleID,
		Name: "Test Role",
	}

	suite.roleRepoMock.EXPECT().
		GetRoleById(roleID, ctx).
		Return(expectedRole, nil)

	result, err := suite.roleUsecase.GetRoleById(roleID, ctx)

	suite.Nil(err)
	suite.Equal(expectedRole, result)
}

func (suite *RoleUsecaseTestSuite) TestUpdateRole_Success() {
	ctx := &gin.Context{}
	roleID := "some-role-id"
	roleReq := dtos.RoleUpdateRequest{
		Name: "Updated Role",
	}
	updatedRole := &dtos.RoleResponse{
		UID:    roleID,
		Name:   roleReq.Name,
		Rights: roleReq.Rights,
	}

	suite.roleRepoMock.EXPECT().
		GetRoleById(roleID, ctx).
		Return(&dtos.RoleResponse{UID: roleID, Name: "Old Role"}, nil)

	suite.roleRepoMock.EXPECT().
		UpdateRole(roleID, roleReq, ctx).
		Return(updatedRole, nil)

	result, err := suite.roleUsecase.UpdateRole(roleID, roleReq, ctx)

	suite.Nil(err)
	suite.Equal(updatedRole, result)
}

func (suite *RoleUsecaseTestSuite) TestDeleteRole_Success() {
	ctx := &gin.Context{}
	roleID := "some-role-id"

	suite.roleRepoMock.EXPECT().
		GetRoleById(roleID, ctx).
		Return(&dtos.RoleResponse{UID: roleID, Name: "Test Role"}, nil)

	suite.roleRepoMock.EXPECT().
		DeleteRole(roleID, ctx).
		Return(nil)

	err := suite.roleUsecase.DeleteRole(roleID, ctx)

	suite.Nil(err)
}

func (suite *RoleUsecaseTestSuite) TestGetRoleUsers_Success() {
	ctx := &gin.Context{}
	roleID := "some-role-id"
	expectedUsers := []*dtos.UserResponse{
		{
			UID:    uuid.New().String(),
			Name:   "User 1",
			Email:  "user1@example.com",
			Status: 1,
		},
		{
			UID:    uuid.New().String(),
			Name:   "User 2",
			Email:  "user2@example.com",
			Status: 1,
		},
	}

	suite.roleRepoMock.EXPECT().
		GetRoleById(roleID, ctx).
		Return(&dtos.RoleResponse{UID: roleID, Name: "Test Role"}, nil)

	suite.roleRepoMock.EXPECT().
		GetRoleUsers(&dtos.RoleResponse{UID: roleID, Name: "Test Role"}, ctx).
		Return(expectedUsers, nil)

	result, err := suite.roleUsecase.GetRoleUsers(roleID, ctx)

	suite.Nil(err)
	suite.Equal(expectedUsers, result)
}

func TestRoleUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoleUsecaseTestSuite))
}
