package repository

import (
	"context"

	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, *models.ErrorResponse) {
	var users []*models.User
	if err := r.db.WithContext(ctx).Preload("Groups").Preload("Roles").Find(&users).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	return users, nil
}

func (r *userRepository) GetUserById(id string, ctx context.Context) (*models.User, *models.ErrorResponse) {
	uId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var user models.User
	if err := r.db.WithContext(ctx).Preload("Groups").Preload("Roles").First(&user, uId).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	return &user, nil
}

func (r *userRepository) GetUsersGroups(id string, ctx context.Context) ([]*models.Group, *models.ErrorResponse) {
	var user models.User
	uId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}

	if err := r.db.WithContext(ctx).
		Preload("Groups").
		First(&user, uId).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return user.Groups, nil
}

func (r *userRepository) SearchUsers(query string, ctx context.Context) ([]*models.User, *models.ErrorResponse) {
	var users []*models.User
	if err := r.db.WithContext(ctx).Where("name LIKE ?", "%"+query+"%").Preload("Groups").Preload("Roles").Find(&users).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	return users, nil
}

func (r *userRepository) CreateUser(user dtos.UserCreateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse) {
	newUser := models.User{
		UID:    uuid.New(),
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
	if err := r.db.WithContext(ctx).Create(&newUser).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	return &dtos.UserResponse{
		UID:    newUser.UID.String(),
		Name:   newUser.Name,
		Email:  newUser.Email,
		Status: newUser.Status,
	}, nil
}

func (r *userRepository) UpdateUser(id string, user dtos.UserUpdateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse) {
	uId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var existingUser models.User
	if err := r.db.WithContext(ctx).First(&existingUser, uId).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Status = user.Status

	if err := r.db.WithContext(ctx).Save(&existingUser).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.UserResponse{UID: existingUser.UID.String(), Name: existingUser.Name, Email: existingUser.Email, Status: existingUser.Status}, nil
}

func (r *userRepository) DeleteUser(id string, ctx context.Context) *models.ErrorResponse {
	uId, err := uuid.Parse(id)
	if err != nil {
		return models.InternalServerError("Invalid UUID format")
	}

	if err := r.db.WithContext(ctx).Delete(&models.User{}, uId).Error; err != nil {
		return models.InternalServerError(err.Error())
	}
	return nil
}
