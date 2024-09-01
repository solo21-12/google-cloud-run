package usecases_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
	mocks "github.com/google-run-code/Tests/Mocks"
	usecases "github.com/google-run-code/Usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type GroupUsecaseTestSuite struct {
	suite.Suite
	groupRepoMock *mocks.MockGroupRepository
	groupUsecase  interfaces.GroupUseCase
	ctrl          *gomock.Controller
}

func (suite *GroupUsecaseTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.groupRepoMock = mocks.NewMockGroupRepository(suite.ctrl)
	suite.groupUsecase = usecases.NewGroupUseCase(suite.groupRepoMock)
}

func (suite *GroupUsecaseTestSuite) TearDownSuite() {
	suite.ctrl.Finish()
}

func (suite *GroupUsecaseTestSuite) TestCreateGroup_Success() {
	ctx := &gin.Context{}
	groupReq := dtos.GroupCreateRequest{
		Name: "Test Group",
	}
	group := &dtos.GroupResponse{
		GID:  uuid.New().String(),
		Name: groupReq.Name,
	}

	suite.groupRepoMock.EXPECT().
		GetGroupByName(groupReq.Name, ctx).
		Return(nil, models.NotFound("Group not found"))

	suite.groupRepoMock.EXPECT().
		CreateGroup(groupReq, ctx).
		Return(group, nil)

	result, err := suite.groupUsecase.CreateGroup(groupReq, ctx)
	suite.Nil(err)
	suite.Equal(group, result)
}

func (suite *GroupUsecaseTestSuite) TestGetGroupById_Success() {
	ctx := &gin.Context{}
	groupID := "some-group-id"
	expectedGroup := &dtos.GroupResponse{
		GID:  groupID,
		Name: "Test Group",
	}

	suite.groupRepoMock.EXPECT().
		GetGroupById(groupID, ctx).
		Return(expectedGroup, nil)

	result, err := suite.groupUsecase.GetGroupById(groupID, ctx)

	suite.Nil(err)
	suite.Equal(expectedGroup, result)
}

func (suite *GroupUsecaseTestSuite) TestGetGroupUsers_Success() {
	ctx := &gin.Context{}
	groupID := "some-group-id"
	expectedUsers := []dtos.UserResponse{
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

	suite.groupRepoMock.EXPECT().
		GetGroupById(groupID, ctx).
		Return(&dtos.GroupResponse{GID: groupID, Name: "Test Group"}, nil)

	suite.groupRepoMock.EXPECT().
		GetGroupUsers(groupID, ctx).
		Return(expectedUsers, nil)

	result, err := suite.groupUsecase.GetGroupUsers(groupID, ctx)

	suite.Nil(err)
	suite.Equal(expectedUsers, result)
}

func (suite *GroupUsecaseTestSuite) TestUpdateGroup_Success() {
	ctx := &gin.Context{}
	groupID := "some-group-id"
	groupReq := dtos.GroupUpdateRequest{
		Name: "Updated Group",
	}
	updatedGroup := &dtos.GroupResponse{
		GID:  groupID,
		Name: groupReq.Name,
	}

	suite.groupRepoMock.EXPECT().
		GetGroupById(groupID, ctx).
		Return(&dtos.GroupResponse{GID: groupID, Name: "Old Group"}, nil)

	suite.groupRepoMock.EXPECT().
		UpdateGroup(groupID, groupReq, ctx).
		Return(updatedGroup, nil)

	result, err := suite.groupUsecase.UpdateGroup(groupID, groupReq, ctx)

	suite.Nil(err)
	suite.Equal(updatedGroup, result)
}

func (suite *GroupUsecaseTestSuite) TestDeleteGroup_Success() {
	ctx := &gin.Context{}
	groupID := "some-group-id"

	suite.groupRepoMock.EXPECT().
		GetGroupById(groupID, ctx).
		Return(&dtos.GroupResponse{GID: groupID, Name: "Test Group"}, nil)

	suite.groupRepoMock.EXPECT().
		DeleteGroup(groupID, ctx).
		Return(nil)

	err := suite.groupUsecase.DeleteGroup(groupID, ctx)

	suite.Nil(err)
}

func TestGroupUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(GroupUsecaseTestSuite))
}
