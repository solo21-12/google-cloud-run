package controllers_tests

import (
	"bytes"
	"context"
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

type RoleControllerTestSuite struct {
	suite.Suite
	roleUsecase *mocks.MockRoleUseCase
	controller  interfaces.RoleController
	server      *httptest.Server
	ctrl        *gomock.Controller
}

func (suite *RoleControllerTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.roleUsecase = mocks.NewMockRoleUseCase(suite.ctrl)
	suite.controller = controllers.NewRoleController(suite.roleUsecase)

	router := gin.Default()
	router.GET("/roles", suite.controller.GetAllRoles)
	router.GET("/roles/:id", suite.controller.GetRoleById)
	router.POST("/roles", suite.controller.CreateRole)
	router.PUT("/roles/:id", suite.controller.UpdateRole)
	router.DELETE("/roles/:id", suite.controller.DeleteRole)
	router.GET("/roles/:id/users", suite.controller.GetRoleUsers)

	suite.server = httptest.NewServer(router)
}

func (suite *RoleControllerTestSuite) TearDownSuite() {
	suite.server.Close()
	suite.ctrl.Finish()
}

func (suite *RoleControllerTestSuite) TestCreateRole_Success() {
	roleRequest := dtos.RoleCreateRequest{
		Name:   "Admin",
		Rights: json.RawMessage(`{"read": true, "write": true}`),
	}

	roleResponse := dtos.RoleResponse{
		UID:    "some-id",
		Name:   "Admin",
		Rights: json.RawMessage(`{"read": true, "write": true}`),
	}

	suite.roleUsecase.
		EXPECT().
		CreateRole(gomock.Any(), gomock.Any()).
		DoAndReturn(func(req dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, error) {
			// Compare unformatted JSON
			suite.Equal(roleRequest.Name, req.Name)
			suite.JSONEq(string(roleRequest.Rights), string(req.Rights))
			return &roleResponse, nil
		}).
		Times(1)

	roleJSON, err := json.Marshal(roleRequest)
	suite.NoError(err)

	response, err := http.Post(suite.server.URL+"/roles", "application/json", bytes.NewBuffer(roleJSON))
	suite.NoError(err)
	defer response.Body.Close()

	suite.Equal(http.StatusCreated, response.StatusCode)

	var responseBody dtos.RoleResponse
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err)

	// Compare JSON responses ignoring formatting
	suite.Equal(roleResponse.UID, responseBody.UID)
	suite.Equal(roleResponse.Name, responseBody.Name)
	suite.JSONEq(string(roleResponse.Rights), string(responseBody.Rights))
}

func (suite *RoleControllerTestSuite) TestDeleteRole_Success() {
	suite.roleUsecase.
		EXPECT().
		DeleteRole("some-id", gomock.Any()).
		Return(nil).
		Times(1)

	req, err := http.NewRequest(http.MethodDelete, suite.server.URL+"/roles/some-id", nil)
	suite.NoError(err)

	response, err := http.DefaultClient.Do(req)
	suite.NoError(err)
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)
}

func (suite *RoleControllerTestSuite) TestGetRoleUsers_Success() {
	users := []*dtos.UserResponse{
		{UID: "user1-uid", Name: "User 1", Email: "user1@example.com"},
		{UID: "user2-uid", Name: "User 2", Email: "user2@example.com"},
	}

	suite.roleUsecase.
		EXPECT().
		GetRoleUsers("some-id", gomock.Any()).
		Return(users, nil).
		Times(1)

	response, err := http.Get(suite.server.URL + "/roles/some-id/users")
	suite.NoError(err)
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)

	var responseBody []*dtos.UserResponse
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err)
	suite.Equal(users, responseBody)
}

func TestRoleControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RoleControllerTestSuite))
}
