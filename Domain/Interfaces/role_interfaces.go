package interfaces

import (
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
	GetAllRoles(ctx *gin.Context) ([]*dtos.RoleResponse, *models.ErrorResponse)
	GetRoleById(id string, ctx *gin.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	CreateRole(role dtos.RoleCreateRequest, ctx *gin.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	UpdateRole(id string, role dtos.RoleUpdateRequest, ctx *gin.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	DeleteRole(id string, ctx *gin.Context) *models.ErrorResponse
	GetRoleUsers(id string, ctx *gin.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
}

type RoleRepository interface {
	GetAllRoles(ctx *gin.Context) ([]*dtos.RoleResponse, *models.ErrorResponse)
	GetRoleById(id string, ctx *gin.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	CreateRole(role dtos.RoleCreateRequest, ctx *gin.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	UpdateRole(id string, role dtos.RoleUpdateRequest, ctx *gin.Context) (*dtos.RoleResponse, *models.ErrorResponse)
	DeleteRole(id string, ctx *gin.Context) *models.ErrorResponse
	GetRoleUsers(role *dtos.RoleResponse, ctx *gin.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	GetRoleByNameAndRights(role dtos.RoleCreateRequest, ctx *gin.Context) (*models.Role, *models.ErrorResponse)
}
