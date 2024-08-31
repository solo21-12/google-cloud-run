package interfaces

import (
	"context"

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
	GetAllUsers(ctx context.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	GetUserById(id string, ctx context.Context) (*dtos.UserResponseSingle, *models.ErrorResponse)
	GetUsersGroup(id string, ctx context.Context) ([]*dtos.GroupResponse, *models.ErrorResponse)
	SearchUsers(searchFields dtos.SearchFields, ctx context.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	CreateUser(user dtos.UserCreateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse)
	UpdateUser(id string, user dtos.UserUpdateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse)
	DeleteUser(id string, ctx context.Context) *models.ErrorResponse
	AddUserToGroup(req dtos.AddUserToGroupRequest, ctx context.Context) *models.ErrorResponse
	AddUserToRole(req dtos.AddUserToRoleRequest, ctx context.Context) *models.ErrorResponse
}

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	GetUserById(id string, ctx context.Context) (*dtos.UserResponseSingle, *models.ErrorResponse)
	GetUserByEmail(email string, ctx context.Context) (*models.User, *models.ErrorResponse)
	GetUsersGroups(uid string, ctx context.Context) ([]*dtos.GroupResponse, *models.ErrorResponse)
	SearchUsers(searchFields dtos.SearchFields, ctx context.Context) ([]*dtos.UserResponse, *models.ErrorResponse)
	CreateUser(user dtos.UserCreateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse)
	UpdateUser(id string, user *dtos.UserUpdateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse)
	DeleteUser(id string, ctx context.Context) *models.ErrorResponse
	AddUserToGroup(req dtos.AddUserToGroupRequest, ctx context.Context) *models.ErrorResponse
	AddUserToRole(req dtos.AddUserToRoleRequest, ctx context.Context) *models.ErrorResponse
}
