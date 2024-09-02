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

type UserUsecaseTestSuite struct {
	suite.Suite
	userRepoMock      *mocks.MockUserRepository
	roleRepoMock      *mocks.MockRoleRepository
	groupRepoMock     *mocks.MockGroupRepository
	emailService      *mocks.MockEmailService
	userUsecase       interfaces.UserUseCase
	userUsecaseMocker *mocks.MockUserUseCase
	ctrl              *gomock.Controller
}

func (suite *UserUsecaseTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.userRepoMock = mocks.NewMockUserRepository(suite.ctrl)
	suite.roleRepoMock = mocks.NewMockRoleRepository(suite.ctrl)
	suite.groupRepoMock = mocks.NewMockGroupRepository(suite.ctrl)
	suite.emailService = mocks.NewMockEmailService(suite.ctrl)
	suite.userUsecaseMocker = mocks.NewMockUserUseCase(suite.ctrl)

	suite.userUsecase = usecases.NewUserUseCase(
		suite.userRepoMock,
		suite.emailService,
		suite.roleRepoMock,
		suite.groupRepoMock,
	)
}

func (suite *UserUsecaseTestSuite) TearDownSuite() {
	suite.ctrl.Finish()
}

func (suite *UserUsecaseTestSuite) TestCreateUser_Success() {
	ctx := &gin.Context{}
	userReq := dtos.UserCreateRequest{
		Name:   "Test User",
		Email:  "test@example.com",
		Status: 1,
		RoleId: "some-role-uid",
	}
	user := &dtos.UserResponse{
		UID:    uuid.New().String(),
		Name:   userReq.Name,
		Email:  userReq.Email,
		Status: userReq.Status,
	}

	userID := &dtos.UserResponseSingle{
		UID:    user.UID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
	
	suite.userRepoMock.EXPECT().
		GetUserById(gomock.Any(), ctx).
		Return(userID, nil)

	suite.userRepoMock.EXPECT().
		GetUserByEmail(userReq.Email, ctx).
		Return(nil, models.NotFound("User not found"))

	suite.emailService.EXPECT().
		IsValidEmail(userReq.Email).
		Return(true)

	suite.userRepoMock.EXPECT().
		CreateUser(userReq, ctx).
		Return(user, nil)


	result, err := suite.userUsecase.CreateUser(userReq, ctx)
	suite.Nil(err)
	suite.Equal(userID, result)
}



func (suite *UserUsecaseTestSuite) TestGetUserByID_Success() {
	ctx := &gin.Context{}
	UserUID := "some-user-id"
	expectedUser := &dtos.UserResponseSingle{
		UID:    uuid.New().String(),
		Name:   "Test User",
		Email:  "test@example.com",
		Status: 1,
	}

	suite.userRepoMock.EXPECT().
		GetUserById(UserUID, ctx).
		Return(expectedUser, nil)

	result, err := suite.userUsecase.GetUserById(UserUID, ctx)

	suite.Nil(err)
	suite.Equal(expectedUser, result)
}

func (suite *UserUsecaseTestSuite) TestSearchUsers_Success() {
	ctx := &gin.Context{}
	searchFields := dtos.SearchFields{
		Name: "Test",
	}
	expectedUsers := []*dtos.UserResponseAll{
		{
			UID:    uuid.New().String(),
			Name:   "Test User 1",
			Email:  "user1@example.com",
			Status: 1,
		},
		{
			UID:    uuid.New().String(),
			Name:   "Test User 2",
			Email:  "user2@example.com",
			Status: 1,
		},
	}

	suite.userRepoMock.EXPECT().
		SearchUsers(searchFields, ctx).
		Return(expectedUsers, nil)

	result, err := suite.userUsecase.SearchUsers(searchFields, ctx)

	suite.Nil(err)
	suite.Equal(expectedUsers, result)
}

func (suite *UserUsecaseTestSuite) TestDeleteUser_Success() {
	ctx := &gin.Context{}
	UserUID := "some-uuid"

	existingUser := &dtos.UserResponseSingle{
		UID:  UserUID,
		Name: "Test User",
	}

	suite.userRepoMock.EXPECT().
		GetUserById(UserUID, ctx).
		Return(existingUser, nil)

	suite.userRepoMock.EXPECT().
		DeleteUser(UserUID, ctx).
		Return(nil)

	err := suite.userUsecase.DeleteUser(UserUID, ctx)

	suite.Nil(err)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
