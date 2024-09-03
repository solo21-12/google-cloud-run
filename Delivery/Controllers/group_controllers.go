package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
)

type groupController struct {
	usecase interfaces.GroupUseCase
}

func NewGroupController(usecase interfaces.GroupUseCase) interfaces.GroupController {
	return &groupController{
		usecase: usecase,
	}
}

func (gc *groupController) GetAllGroups(c *gin.Context) {
	groups, errResp := gc.usecase.GetAllGroups(c)

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	if len(groups) == 0 {
		c.IndentedJSON(http.StatusNotFound, []string{})
		return
	}

	c.IndentedJSON(http.StatusOK, groups)
}

func (gc *groupController) GetGroupById(c *gin.Context) {
	id := c.Param("id")
	group, errResp := gc.usecase.GetGroupById(id, c)

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, group)
}

func (gc *groupController) GetGroupUsers(c *gin.Context) {
	id := c.Param("id")

	users, errResp := gc.usecase.GetGroupUsers(id, c)

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (gc *groupController) CreateGroup(c *gin.Context) {
	var group dtos.GroupCreateRequest

	if err := c.ShouldBindJSON(&group); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "one or more required fields are missing"})
		return
	}

	groupResponse, errResp := gc.usecase.CreateGroup(group, c)

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusCreated, groupResponse)
}

func (gc *groupController) UpdateGroup(c *gin.Context) {
	var group dtos.GroupUpdateRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&group); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupResponse, errResp := gc.usecase.UpdateGroup(id, group, c)

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, groupResponse)
}

func (gc *groupController) DeleteGroup(c *gin.Context) {
	id := c.Param("id")

	errResp := gc.usecase.DeleteGroup(id, c)

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{})
}
