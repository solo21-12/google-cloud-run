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
	if err := r.db.WithContext(ctx).Preload("Groups").Preload("Role").Find(&users).Error; err != nil {
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

func (r *userRepository) GetUserById(uid string, ctx context.Context) (*dtos.UserResponseSingle, *models.ErrorResponse) {

	var user models.User
	if err := r.db.WithContext(ctx).
		Preload("Groups").
		Preload("Role").
		Where("uid = ?", uid).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("User not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	var roleRes *dtos.RoleResponse
	if user.Role != nil {
		roleRes = &dtos.RoleResponse{
			RID:    user.Role.RID.String(),
			Name:   user.Role.Name,
			Rights: user.Role.Rights,
		}
	}

	var result dtos.UserResponseSingle
	result.UID = user.UID.String()
	result.Name = user.Name
	result.Email = user.Email
	result.Status = user.Status
	result.Role = roleRes

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
	if err := r.db.WithContext(ctx).Preload("Groups").Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("User not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	return &user, nil
}

func (r *userRepository) GetUsersGroups(uid string, ctx context.Context) ([]*dtos.GroupResponse, *models.ErrorResponse) {
	var user models.User

	if err := r.db.WithContext(ctx).
		Preload("Groups").
		Where("uid = ?", uid).
		First(&user).Error; err != nil {
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

func (r *userRepository) UpdateUser(uid string, user *dtos.UserUpdateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse) {
	var existingUser models.User

	uId, err := uuid.Parse(uid)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format for UID")
	}

	if err := r.db.WithContext(ctx).
		Where("uid = ?", uId).
		First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("User not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Status = user.Status

	if err := r.db.WithContext(ctx).Save(&existingUser).Error; err != nil {
		return nil, models.InternalServerError("Failed to update user: " + err.Error())
	}

	return &dtos.UserResponse{
		UID:    existingUser.UID.String(),
		Name:   existingUser.Name,
		Email:  existingUser.Email,
		Status: existingUser.Status,
	}, nil
}

func (r *userRepository) DeleteUser(uid string, ctx context.Context) *models.ErrorResponse {
	var existingUser models.User

	if err := r.db.WithContext(ctx).
		Where("uid = ?", uid).
		First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("User not found")
		}
		return models.InternalServerError(err.Error())
	}

	if err := r.db.WithContext(ctx).
		Model(&existingUser).
		Association("Groups").Clear(); err != nil {
		return models.InternalServerError("Failed to dissociate user from groups: " + err.Error())
	}

	// Delete the user
	if err := r.db.WithContext(ctx).
		Delete(&models.User{}, existingUser.ID).Error; err != nil {
		return models.InternalServerError("Failed to delete user: " + err.Error())
	}

	return nil
}

func (r *userRepository) AddUserToGroup(req dtos.AddUserToGroupRequest, ctx context.Context) *models.ErrorResponse {

	var user models.User
	if err := r.db.WithContext(ctx).
		Where("uid = ?", req.UserId).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("User not found")
		}
		return models.InternalServerError(err.Error())
	}

	var group models.Group
	if err := r.db.WithContext(ctx).
		Where("g_id = ?", req.GroupId).
		First(&group).Error; err != nil {
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

func (repo *userRepository) AddUserToRole(req dtos.AddUserToRoleRequest, ctx context.Context) *models.ErrorResponse {
	var user models.User
	var role models.Role

	if err := repo.db.WithContext(ctx).
		Where("r_id = ?", req.RoleId).
		First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.ErrorResponse{Message: "Role not found", Code: 404}
		}
		return &models.ErrorResponse{Message: "Failed to fetch role: " + err.Error(), Code: 500}
	}

	if err := repo.db.WithContext(ctx).
		Where("uid = ?", req.UserId). // Use uid to fetch the user
		First(&user).Error; err != nil {
		return &models.ErrorResponse{Message: "User not found", Code: 404}
	}

	// Check if the user already has the specified role
	if user.RoleID != nil && *user.RoleID == role.ID {
		return &models.ErrorResponse{Message: "User already has the specified role", Code: 400}
	}

	// Update the user's RoleID
	user.RoleID = &role.ID
	if err := repo.db.WithContext(ctx).Save(&user).Error; err != nil {
		return &models.ErrorResponse{Message: "Failed to update user role: " + err.Error(), Code: 500}
	}

	return nil
}
