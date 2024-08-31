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
	GetRoleUsers(c *gin.Context)
}

type RoleUseCase interface {
	GetAllRoles(ctx context.Context) ([]*dtos.RoleResponse, *models.ErrorResponse)
	GetRoleById(id string, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	CreateRole(role dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	UpdateRole(id string, role dtos.RoleUpdateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	DeleteRole(id string, ctx context.Context) *models.ErrorResponse
	GetRoleUsers(id string, ctx context.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
}

type RoleRepository interface {
	GetAllRoles(ctx context.Context) ([]*dtos.RoleResponse, *models.ErrorResponse)
	GetRoleById(id string, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	CreateRole(role dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	UpdateRole(id string, role dtos.RoleUpdateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	DeleteRole(id string, ctx context.Context) *models.ErrorResponse
	GetRoleUsers(role *dtos.RoleResponse, ctx context.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	GetRoleByNameAndRights(role dtos.RoleCreateRequest, ctx context.Context) (*models.Role, *models.ErrorResponse)
	
}
