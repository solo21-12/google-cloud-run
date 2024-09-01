package interfaces

import (
	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	models "github.com/google-run-code/Domain/Models"
)

type UserController interface {
	GetUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	GetUsersGroup(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	AddUserToGroup(c *gin.Context)
	AddUserToRole(c *gin.Context)
}

type UserUseCase interface {
	ValidateEmail(email string) *models.ErrorResponse
	CheckEmailExists(email string, ctx *gin.Context) (*models.User, *models.ErrorResponse)
	GetAllUsers(ctx *gin.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	GetUserById(id string, ctx *gin.Context) (*dtos.UserResponseSingle, *models.ErrorResponse)
	GetUsersGroup(id string, ctx *gin.Context) ([]*dtos.GroupResponse, *models.ErrorResponse)
	SearchUsers(searchFields dtos.SearchFields, ctx *gin.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	CreateUser(user dtos.UserCreateRequest, ctx *gin.Context) (*dtos.UserResponse, *models.ErrorResponse)
	UpdateUser(id string, user dtos.UserUpdateRequest, ctx *gin.Context) (*dtos.UserResponse, *models.ErrorResponse)
	DeleteUser(id string, ctx *gin.Context) *models.ErrorResponse
	AddUserToGroup(req dtos.AddUserToGroupRequest, ctx *gin.Context) *models.ErrorResponse
	AddUserToRole(req dtos.AddUserToRoleRequest, ctx *gin.Context) *models.ErrorResponse
}

type UserRepository interface {
	GetAllUsers(ctx *gin.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	GetUserById(id string, ctx *gin.Context) (*dtos.UserResponseSingle, *models.ErrorResponse)
	GetUserByEmail(email string, ctx *gin.Context) (*models.User, *models.ErrorResponse)
	GetUsersGroups(uid string, ctx *gin.Context) ([]*dtos.GroupResponse, *models.ErrorResponse)
	SearchUsers(searchFields dtos.SearchFields, ctx *gin.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	CreateUser(user dtos.UserCreateRequest, ctx *gin.Context) (*dtos.UserResponse, *models.ErrorResponse)
	UpdateUser(id string, user *dtos.UserUpdateRequest, ctx *gin.Context) (*dtos.UserResponse, *models.ErrorResponse)
	DeleteUser(id string, ctx *gin.Context) *models.ErrorResponse
	AddUserToGroup(req dtos.AddUserToGroupRequest, ctx *gin.Context) *models.ErrorResponse
	AddUserToRole(req dtos.AddUserToRoleRequest, ctx *gin.Context) *models.ErrorResponse
}
