package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	models "github.com/google-run-code/Domain/Models"
)

type UserController interface {
	GetAllUsers(gin.Context)
	GetUserById(gin.Context)
	GetUsersGroup(gin.Context)
	SearchUsers(gin.Context)
	CreateUser(gin.Context)
	UpdateUser(gin.Context)
	DeleteUser(gin.Context)
}

type UserUseCase interface {
	GetAllUsers(ctx context.Context) ([]*models.User, *models.ErrorResponse)
	GetUserById(id string, ctx context.Context) (*models.User, *models.ErrorResponse)
	GetUsersGroup(ctx context.Context) ([]*models.Group, *models.ErrorResponse)
	SearchUsers(query string, ctx context.Context) ([]*models.User, *models.ErrorResponse)
	CreateUser(user dtos.UserCreateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse)
	UpdateUser(user dtos.UserUpdateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse)
	DeleteUser(id string, ctx context.Context) *models.ErrorResponse
}

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*models.User, *models.ErrorResponse)
	GetUserById(id string, ctx context.Context) (*models.User, *models.ErrorResponse)
	GetUsersGroups(uid string, ctx context.Context) ([]*models.Group, *models.ErrorResponse)
	SearchUsers(query string, ctx context.Context) ([]*models.User, *models.ErrorResponse)
	CreateUser(user dtos.UserCreateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse)
	UpdateUser(id string, user dtos.UserUpdateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse)
	DeleteUser(id string, ctx context.Context) *models.ErrorResponse
}
