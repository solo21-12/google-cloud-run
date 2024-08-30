package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	models "github.com/google-run-code/Domain/Models"
)

type RoleController interface {
	GetAllRoles(gin.Context)
	GetRoleById(gin.Context)
	CreateRole(gin.Context)
	UpdateRole(gin.Context)
	DeleteRole(gin.Context)
}

type RoleUseCase interface {
	GetAllRoles(ctx context.Context) ([]*models.Role, *models.ErrorResponse)
	GetRoleById(id int, ctx context.Context) (*models.Role, *models.ErrorResponse)
	CreateRole(role dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	UpdateRole(role dtos.RoleUpdateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	DeleteRole(id int, ctx context.Context) *models.ErrorResponse
}

type RoleRepository interface {
	GetAllRoles(ctx context.Context) ([]*models.Role, *models.ErrorResponse)
	GetRoleById(id int, ctx context.Context) (*models.Role, *models.ErrorResponse)
	CreateRole(role dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	UpdateRole(role dtos.RoleUpdateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	DeleteRole(id int, ctx context.Context) *models.ErrorResponse
}
