package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	models "github.com/google-run-code/Domain/Models"
)

type RoleController interface {
	GetAllRoles(c *gin.Context)
	GetRoleById(c *gin.Context)
	CreateRole(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}

type RoleUseCase interface {
	GetAllRoles(ctx context.Context) ([]*models.Role, *models.ErrorResponse)
	GetRoleById(id string, ctx context.Context) (*models.Role, *models.ErrorResponse)
	CreateRole(role dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	UpdateRole(id string, role dtos.RoleUpdateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	DeleteRole(id string, ctx context.Context) *models.ErrorResponse
}

type RoleRepository interface {
	GetAllRoles(ctx context.Context) ([]*models.Role, *models.ErrorResponse)
	GetRoleById(id string, ctx context.Context) (*models.Role, *models.ErrorResponse)
	CreateRole(role dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	UpdateRole(id string, role dtos.RoleUpdateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	DeleteRole(id string, ctx context.Context) *models.ErrorResponse
}
