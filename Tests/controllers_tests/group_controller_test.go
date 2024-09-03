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
	mocks "github.com/google-run-code/Tests/Mocks"
	"github.com/stretchr/testify/suite"
)

type GroupControllerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	useCaseMock *mocks.MockGroupUseCase
}

func (suite *GroupControllerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	suite.useCaseMock = mocks.NewMockGroupUseCase(mockCtrl)
	suite.router = gin.Default()

	groupController := controllers.NewGroupController(suite.useCaseMock)
	suite.router.GET("/groups", groupController.GetAllGroups)
	suite.router.GET("/groups/:id", groupController.GetGroupById)
	suite.router.GET("/groups/:id/users", groupController.GetGroupUsers)
	suite.router.POST("/groups", groupController.CreateGroup)
	suite.router.PUT("/groups/:id", groupController.UpdateGroup)
	suite.router.DELETE("/groups/:id", groupController.DeleteGroup)
}

func (suite *GroupControllerTestSuite) TestGetAllGroups_Success() {
	expectedGroups := []*dtos.GroupResponse{{UID: "group-id", Name: "Admin"}}
	suite.useCaseMock.EXPECT().GetAllGroups(gomock.Any()).Return(expectedGroups, nil)

	req, _ := http.NewRequest("GET", "/groups", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Admin")
}

func (suite *GroupControllerTestSuite) TestGetGroupById_Success() {
	expectedGroup := &dtos.GroupResponse{UID: "group-id", Name: "Admin"}
	suite.useCaseMock.EXPECT().GetGroupById("group-id", gomock.Any()).Return(expectedGroup, nil)

	req, _ := http.NewRequest("GET", "/groups/group-id", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Admin")
}

func (suite *GroupControllerTestSuite) TestGetGroupUsers_Success() {
	expectedUsers := []dtos.UserResponse{{UID: "user-id", Name: "User1"}}
	suite.useCaseMock.EXPECT().GetGroupUsers("group-id", gomock.Any()).Return(expectedUsers, nil)

	req, _ := http.NewRequest("GET", "/groups/group-id/users", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "User1")
}

func (suite *GroupControllerTestSuite) TestCreateGroup_Success() {
	groupRequest := dtos.GroupCreateRequest{Name: "Admin"}
	groupResponse := &dtos.GroupResponse{UID: "group-id", Name: "Admin"}
	suite.useCaseMock.EXPECT().CreateGroup(groupRequest, gomock.Any()).Return(groupResponse, nil)

	groupJson, _ := json.Marshal(groupRequest)
	req, _ := http.NewRequest("POST", "/groups", bytes.NewBuffer(groupJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)
	suite.Contains(w.Body.String(), "group-id")
}

func (suite *GroupControllerTestSuite) TestUpdateGroup_Success() {
	groupRequest := dtos.GroupUpdateRequest{Name: "Admin"}
	groupResponse := &dtos.GroupResponse{UID: "group-id", Name: "Admin"}
	suite.useCaseMock.EXPECT().UpdateGroup("group-id", groupRequest, gomock.Any()).Return(groupResponse, nil)

	groupJson, _ := json.Marshal(groupRequest)
	req, _ := http.NewRequest("PUT", "/groups/group-id", bytes.NewBuffer(groupJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "group-id")
}

func (suite *GroupControllerTestSuite) TestDeleteGroup_Success() {
	suite.useCaseMock.EXPECT().DeleteGroup("group-id", gomock.Any()).Return(nil)

	req, _ := http.NewRequest("DELETE", "/groups/group-id", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func TestGroupControllerTestSuite(t *testing.T) {
	suite.Run(t, new(GroupControllerTestSuite))
}
