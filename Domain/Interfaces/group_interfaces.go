package interfaces

import (
	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	models "github.com/google-run-code/Domain/Models"
)

type GroupController interface {
	GetAllGroups(c *gin.Context)
	GetGroupById(c *gin.Context)
	GetGroupUsers(c *gin.Context)
	CreateGroup(c *gin.Context)
	UpdateGroup(c *gin.Context)
	DeleteGroup(c *gin.Context)
}

type GroupUseCase interface {
	GetAllGroups(ctx *gin.Context) ([]*dtos.GroupResponse, *models.ErrorResponse)
	GetGroupById(id string, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	GetGroupUsers(id string, ctx *gin.Context) ([]dtos.UserResponse, *models.ErrorResponse)
	CreateGroup(group dtos.GroupCreateRequest, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	UpdateGroup(id string, group dtos.GroupUpdateRequest, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	DeleteGroup(id string, ctx *gin.Context) *models.ErrorResponse
}

type GroupRepository interface {
	GetAllGroups(ctx *gin.Context) ([]*dtos.GroupResponse, *models.ErrorResponse)
	GetGroupById(id string, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	GetGroupUsers(id string, ctx *gin.Context) ([]dtos.UserResponse, *models.ErrorResponse)
	GetGroupByName(name string, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	CreateGroup(group dtos.GroupCreateRequest, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	UpdateGroup(id string, group dtos.GroupUpdateRequest, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	DeleteGroup(id string, ctx *gin.Context) *models.ErrorResponse
}
