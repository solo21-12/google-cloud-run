package controllers_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	controllers "github.com/google-run-code/Delivery/Controllers"
	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	mocks "github.com/google-run-code/Tests/Mocks"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	usecase    *mocks.MockUserUseCase
	controller interfaces.UserController
	server     *httptest.Server
	ctrl       *gomock.Controller
}

func (suite *UserControllerTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.usecase = mocks.NewMockUserUseCase(suite.ctrl)

	suite.controller = controllers.NewUserController(suite.usecase)

	router := gin.Default()
	router.POST("/users", suite.controller.CreateUser)
	router.PUT("/users/:id", suite.controller.UpdateUser)
	router.DELETE("/users/:id", suite.controller.DeleteUser)
	router.GET("/users/:id", suite.controller.GetUserById)
	router.GET("/users", suite.controller.GetUsers)
	router.POST("/users/:id/groups", suite.controller.AddUserToGroup)

	suite.server = httptest.NewServer(router)
}

func (suite *UserControllerTestSuite) TearDownSuite() {
	suite.server.Close()
	suite.ctrl.Finish()
}

func (suite *UserControllerTestSuite) TestCreateUser_Success() {
	input := dtos.UserCreateRequest{
		Name:   "Test User",
		Email:  "test@example.com",
		Status: 1,
	}

	expectedUser := &dtos.UserResponseSingle{
		UID:    "some-uid",
		Name:   input.Name,
		Email:  input.Email,
		Status: input.Status,
	}

	suite.usecase.
		EXPECT().
		CreateUser(input, gomock.Any()).
		Return(expectedUser, nil).
		Times(1)

	inputJSON, err := json.Marshal(input)
	suite.NoError(err)

	response, err := http.Post(suite.server.URL+"/users", "application/json", bytes.NewBuffer(inputJSON))
	suite.NoError(err)
	defer response.Body.Close()

	suite.Equal(http.StatusCreated, response.StatusCode)

	var responseUser dtos.UserResponseSingle
	err = json.NewDecoder(response.Body).Decode(&responseUser)
	suite.NoError(err)
	suite.Equal(expectedUser, &responseUser)
}

func (suite *UserControllerTestSuite) TestUpdateUser_Success() {
	id := "some-id"
	input := dtos.UserUpdateRequest{
		Name:  "Updated User",
		Email: "updated@example.com",
	}

	expectedUser := &dtos.UserResponseSingle{
		UID:    "some-uid",
		Name:   input.Name,
		Email:  input.Email,
		Status: 1,
	}

	suite.usecase.
		EXPECT().
		UpdateUser(id, input, gomock.Any()).
		Return(expectedUser, nil).
		Times(1)

	inputJSON, err := json.Marshal(input)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+id, bytes.NewBuffer(inputJSON))
	suite.NoError(err)

	response, err := http.DefaultClient.Do(req)
	suite.NoError(err)
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)

	var responseUser dtos.UserResponseSingle
	err = json.NewDecoder(response.Body).Decode(&responseUser)
	suite.NoError(err)
	suite.Equal(expectedUser, &responseUser)
}

func (suite *UserControllerTestSuite) TestDeleteUser_Success() {
	id := "some-id"

	suite.usecase.
		EXPECT().
		DeleteUser(id, gomock.Any()).
		Return(nil).
		Times(1)

	req, err := http.NewRequest(http.MethodDelete, suite.server.URL+"/users/"+id, nil)
	suite.NoError(err)

	response, err := http.DefaultClient.Do(req)
	suite.NoError(err)
	defer response.Body.Close()

	suite.Equal(http.StatusNoContent, response.StatusCode)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
