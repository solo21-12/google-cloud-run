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

func (r *userRepository) GetAllUsers(ctx context.Context) ([]*dtos.UserResponse, *models.ErrorResponse) {
	var users []*models.User
	if err := r.db.WithContext(ctx).Preload("Groups.Users").Preload("Roles").Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*dtos.UserResponse{}, nil
		}
		return nil, models.InternalServerError(err.Error())
	}

	var result []*dtos.UserResponse

	for _, user := range users {
		result = append(result, &dtos.UserResponse{
			UID:    user.UID.String(),
			Name:   user.Name,
			Email:  user.Email,
			Status: user.Status,
		})
	}
	return result, nil
}

func (r *userRepository) GetUserById(id string, ctx context.Context) (*dtos.UserResponseSingle, *models.ErrorResponse) {
	uId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var user models.User
	if err := r.db.WithContext(ctx).
		Preload("Groups.Users").
		Preload("Roles").
		First(&user, uId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("User not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	var result dtos.UserResponseSingle
	result.UID = user.UID.String()
	result.Name = user.Name
	result.Email = user.Email
	result.Status = user.Status

	for _, group := range user.Groups {
		result.Groups = append(result.Groups, dtos.GroupResponse{
			GID:  group.GID.String(),
			Name: group.Name,
		})
	}

	return &result, nil
}

func (r *userRepository) GetUserByEmail(email string, ctx context.Context) (*models.User, *models.ErrorResponse) {
	var user models.User
	if err := r.db.WithContext(ctx).Preload("Groups").Preload("Roles").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("User not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	return &user, nil
}

func (r *userRepository) GetUsersGroups(id string, ctx context.Context) ([]*dtos.GroupResponse, *models.ErrorResponse) {
	var user models.User
	uId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}

	if err := r.db.WithContext(ctx).
		Preload("Groups").
		First(&user, uId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("User not found")
		}
		return nil, models.InternalServerError(err.Error())
	}
	var groups []*dtos.GroupResponse
	for _, group := range user.Groups {

		groups = append(groups, &dtos.GroupResponse{
			GID:  group.GID.String(),
			Name: group.Name,
		})

	}

	return groups, nil
}

func (repo *userRepository) SearchUsers(searchFields dtos.SearchFields, ctx context.Context) ([]*dtos.UserResponse, *models.ErrorResponse) {
	var users []*models.User

	query := repo.db.Where("name ILIKE ?", "%"+searchFields.Search+"%")

	if searchFields.OrderBy != "" {
		query = query.Order(searchFields.OrderBy)
	}

	if searchFields.Limit > 0 {
		query = query.Limit(searchFields.Limit)
	} else {
		query = query.Limit(100)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	var result []*dtos.UserResponse

	for _, user := range users {
		result = append(result, &dtos.UserResponse{
			UID:    user.UID.String(),
			Name:   user.Name,
			Email:  user.Email,
			Status: user.Status,
		})
	}

	return result, nil
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

func (r *userRepository) UpdateUser(id string, user *dtos.UserUpdateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse) {
	uId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var existingUser models.User
	if err := r.db.WithContext(ctx).First(&existingUser, uId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("User not found")
		}
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

func (r *userRepository) AddUserToGroup(req dtos.AddUserToGroupRequest, ctx context.Context) *models.ErrorResponse {
	uId, err := uuid.Parse(req.UserId)
	if err != nil {
		return models.InternalServerError("Invalid UUID format")
	}

	gId, err := uuid.Parse(req.GroupId)
	if err != nil {
		return models.InternalServerError("Invalid UUID format")
	}

	var user models.User
	if err := r.db.WithContext(ctx).First(&user, uId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("User not found")
		}
		return models.InternalServerError(err.Error())
	}

	var group models.Group
	if err := r.db.WithContext(ctx).First(&group, gId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("Group not found")
		}
		return models.InternalServerError(err.Error())
	}

	if err := r.db.WithContext(ctx).Model(&user).Association("Groups").Append(&group); err != nil {
		return models.InternalServerError(err.Error())
	}

	return nil
}

func (r *userRepository) AddUserToRole(req dtos.AddUserToRoleRequest, ctx context.Context) *models.ErrorResponse {
	uId, err := uuid.Parse(req.UserId)
	if err != nil {
		return models.InternalServerError("Invalid UUID format")
	}

	rId, err := uuid.Parse(req.RoleId)
	if err != nil {
		return models.InternalServerError("Invalid UUID format")
	}

	var user models.User
	if err := r.db.WithContext(ctx).First(&user, uId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("User not found")
		}
		return models.InternalServerError(err.Error())
	}

	var role models.Role
	if err := r.db.WithContext(ctx).First(&role, rId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("Role not found")
		}
		return models.InternalServerError(err.Error())
	}

	if err := r.db.WithContext(ctx).Model(&user).Association("Roles").Append(&role); err != nil {
		return models.InternalServerError(err.Error())
	}

	return nil
}
