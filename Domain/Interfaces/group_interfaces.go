package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	models "github.com/google-run-code/Domain/Models"
)

type GroupController interface {
	GetAllGroups(gin.Context)
	GetGroupById(gin.Context)
	GetGroupUsers(gin.Context)
	CreateGroup(gin.Context)
	UpdateGroup(gin.Context)
	DeleteGroup(gin.Context)
}

type GroupUseCase interface {
	GetAllGroups() ([]*models.Group, *models.ErrorResponse)
	GetGroupById(id string, ctx context.Context) (*models.Group, *models.ErrorResponse)
	GetGroupUsers(id string, ctx context.Context) ([]*models.User, *models.ErrorResponse)
	CreateGroup(group dtos.GroupCreateRequest, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	UpdateGroup(group dtos.GroupUpdateRequest, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	DeleteGroup(id string, ctx context.Context) *models.ErrorResponse
}

type GroupRepository interface {
	GetAllGroups(ctx context.Context) ([]*models.Group, *models.ErrorResponse)
	GetGroupById(id string, ctx context.Context) (*models.Group, *models.ErrorResponse)
	GetGroupUsers(id string, ctx context.Context) ([]models.User, *models.ErrorResponse)
	CreateGroup(group dtos.GroupCreateRequest, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	UpdateGroup(id string, group dtos.GroupUpdateRequest, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse)
	DeleteGroup(id string, ctx context.Context) *models.ErrorResponse
}
