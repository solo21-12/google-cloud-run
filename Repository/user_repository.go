package repository

import (
	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
}

func NewUserRepository() interfaces.UserRepository {
	return &userRepository{}
}

func (r *userRepository) getDB(ctx *gin.Context) (*gorm.DB, error) {
	db, exists := ctx.Get("dbClient")
	if !exists {
		return nil, models.InternalServerError("Database connection not found")
	}

	dbClient, ok := db.(*gorm.DB)
	if !ok {
		return nil, models.InternalServerError("Invalid database connection")
	}

	return dbClient, nil
}

func (r *userRepository) GetAllUsers(ctx *gin.Context) ([]*dtos.UserResponseAll, *models.ErrorResponse) {
	db, err := r.getDB(ctx)
	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	var users []*models.User
	if err := db.WithContext(ctx).Preload("Groups").Preload("Role").Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*dtos.UserResponseAll{}, nil
		}
		return nil, models.InternalServerError(err.Error())
	}

	var result []*dtos.UserResponseAll
	for _, user := range users {
		var roleResponse *dtos.RoleResponseNoRight
		if user.Role != nil {
			roleResponse = &dtos.RoleResponseNoRight{
				UID:  user.Role.UID.String(),
				Name: user.Role.Name,
			}
		}

		result = append(result, &dtos.UserResponseAll{
			UID:    user.UID.String(),
			Name:   user.Name,
			Email:  user.Email,
			Status: user.Status,
			Role:   roleResponse, // roleResponse will be nil if the user has no role
		})
	}

	return result, nil
}

func (r *userRepository) GetUserById(uid string, ctx *gin.Context) (*dtos.UserResponseSingle, *models.ErrorResponse) {
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	var user models.User
	if err := db.WithContext(ctx).
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
			UID:    user.Role.UID.String(),
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
			UID:  group.UID.String(),
			Name: group.Name,
		})
	}

	return &result, nil
}

func (r *userRepository) GetUserByEmail(email string, ctx *gin.Context) (*models.User, *models.ErrorResponse) {
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	var user models.User
	if err := db.WithContext(ctx).Preload("Groups").Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("User not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	return &user, nil
}

func (r *userRepository) GetUsersGroups(uid string, ctx *gin.Context) ([]*dtos.GroupResponse, *models.ErrorResponse) {
	var user models.User
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).
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
			UID:  group.UID.String(),
			Name: group.Name,
		})
	}

	return groups, nil
}

func (repo *userRepository) SearchUsers(searchFields dtos.SearchFields, ctx *gin.Context) ([]*dtos.UserResponseAll, *models.ErrorResponse) {
	var users []*models.User
	db, err := repo.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	query := db.Preload("Role").Where("name ILIKE ?", "%"+searchFields.Name+"%")

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

	var result []*dtos.UserResponseAll
	for _, user := range users {
		var roleResponse *dtos.RoleResponseNoRight
		if user.Role != nil {
			roleResponse = &dtos.RoleResponseNoRight{
				UID:  user.Role.UID.String(),
				Name: user.Role.Name,
			}
		}

		result = append(result, &dtos.UserResponseAll{
			UID:    user.UID.String(),
			Name:   user.Name,
			Email:  user.Email,
			Status: user.Status,
			Role:   roleResponse,
		})
	}

	return result, nil
}

func (r *userRepository) CreateUser(user dtos.UserCreateRequest, ctx *gin.Context) (*dtos.UserResponse, *models.ErrorResponse) {
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	newUser := models.User{
		UID:    uuid.New(),
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}

	if err := db.WithContext(ctx).Create(&newUser).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	return &dtos.UserResponse{
		UID:    newUser.UID.String(),
		Name:   newUser.Name,
		Email:  newUser.Email,
		Status: newUser.Status,
	}, nil
}

func (r *userRepository) UpdateUser(uid string, user *dtos.UserUpdateRequest, ctx *gin.Context) (*dtos.UserResponseSingle, *models.ErrorResponse) {
	var existingUser models.User
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	uId, err := uuid.Parse(uid)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format for UID")
	}

	if err := db.WithContext(ctx).
		Where("uid = ?", uId).
		Preload("Role").
		First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("User not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Status = user.Status

	if err := db.WithContext(ctx).Save(&existingUser).Error; err != nil {
		return nil, models.InternalServerError("Failed to update user: " + err.Error())
	}

	return &dtos.UserResponseSingle{
		UID:    existingUser.UID.String(),
		Name:   existingUser.Name,
		Email:  existingUser.Email,
		Status: existingUser.Status,
		Role: &dtos.RoleResponse{
			UID:    existingUser.Role.UID.String(),
			Name:   existingUser.Role.Name,
			Rights: existingUser.Role.Rights,
		},
	}, nil
}

func (r *userRepository) DeleteUser(uid string, ctx *gin.Context) *models.ErrorResponse {
	var existingUser models.User
	db, err := r.getDB(ctx)

	if err != nil {
		return models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).
		Where("uid = ?", uid).
		First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("User not found")
		}
		return models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).
		Model(&existingUser).
		Association("Groups").Clear(); err != nil {
		return models.InternalServerError("Failed to dissociate user from groups: " + err.Error())
	}

	// Delete the user
	if err := db.WithContext(ctx).
		Delete(&models.User{}, existingUser.ID).Error; err != nil {
		return models.InternalServerError("Failed to delete user: " + err.Error())
	}

	return nil
}

func (r *userRepository) AddUserToGroup(req dtos.AddUserToGroupRequest, ctx *gin.Context) *models.ErrorResponse {
	db, err := r.getDB(ctx)

	if err != nil {
		return models.InternalServerError(err.Error())
	}

	var user models.User
	if err := db.WithContext(ctx).
		Where("uid = ?", req.UserUID).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("User not found")
		}
		return models.InternalServerError(err.Error())
	}

	var groups []models.Group
	if err := db.WithContext(ctx).
		Where("uid IN ?", req.GroupIds).
		Find(&groups).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("One or more groups not found")
		}
		return models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).Model(&user).Association("Groups").Append(groups); err != nil {
		return models.InternalServerError(err.Error())
	}

	return nil
}

func (repo *userRepository) AddUserToRole(req dtos.AddUserToRoleRequest, ctx *gin.Context) *models.ErrorResponse {
	var user models.User
	var role models.Role

	db, err := repo.getDB(ctx)

	if err != nil {
		return models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).
		Where("uid = ?", req.RoleId).
		First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.ErrorResponse{Message: "Role not found", Code: 404}
		}
		return &models.ErrorResponse{Message: "Failed to fetch role: " + err.Error(), Code: 500}
	}

	if err := db.WithContext(ctx).
		Where("uid = ?", req.UserUID).
		First(&user).Error; err != nil {
		return &models.ErrorResponse{Message: "User not found", Code: 404}
	}

	if user.RoleID != nil && *user.RoleID == role.ID {
		return &models.ErrorResponse{Message: "User already has the specified role", Code: 400}
	}

	user.RoleID = &role.ID
	if err := db.WithContext(ctx).Save(&user).Error; err != nil {
		return &models.ErrorResponse{Message: "Failed to update user role: " + err.Error(), Code: 500}
	}

	return nil
}

func (repo *userRepository) RemoveUserFromGroups(userUID string, groupUIDs []string, ctx *gin.Context) *models.ErrorResponse {
	db, err := repo.getDB(ctx)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	var user models.User
	if err := db.WithContext(ctx).Where("uid = ?", userUID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("User not found")
		}
		return models.InternalServerError(err.Error())
	}

	var groups []models.Group
	if err := db.WithContext(ctx).Where("uid IN ?", groupUIDs).Find(&groups).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("Some of the groups were not found")
		}
		return models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).Model(&user).Association("Groups").Delete(&groups); err != nil {
		return models.InternalServerError(err.Error())
	}

	return nil
}
