package repository

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type roleRepository struct {
}

func NewRoleRepository() interfaces.RoleRepository {
	return &roleRepository{}
}

func (r *roleRepository) getDB(ctx *gin.Context) (*gorm.DB, error) {

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

func (r *roleRepository) GetAllRoles(ctx *gin.Context) ([]*dtos.RoleResponse, *models.ErrorResponse) {
	var roles []*models.Role
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).Preload("Users").Find(&roles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*dtos.RoleResponse{}, nil
		}
		return nil, models.InternalServerError(err.Error())
	}

	var result []*dtos.RoleResponse

	for _, role := range roles {
		result = append(result, &dtos.RoleResponse{
			UID:    role.UID.String(),
			Name:   role.Name,
			Rights: role.Rights,
		})
	}
	return result, nil
}

func (r *roleRepository) GetRoleById(uid string, ctx *gin.Context) (*dtos.RoleResponse, *models.ErrorResponse) {

	var role models.Role
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	if err := db.WithContext(ctx).
		Preload("Users").
		Where("uid = ?", uid).
		First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Role not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.RoleResponse{
		UID:    role.UID.String(),
		Name:   role.Name,
		Rights: role.Rights,
	}, nil
}

func (r *roleRepository) CreateRole(role dtos.RoleCreateRequest, ctx *gin.Context) (*dtos.RoleResponse, *models.ErrorResponse) {
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	newRole := models.Role{
		UID:    uuid.New(),
		Name:   role.Name,
		Rights: role.Rights,
	}

	if err := db.WithContext(ctx).Create(&newRole).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.RoleResponse{
		UID:    newRole.UID.String(),
		Name:   newRole.Name,
		Rights: newRole.Rights,
	}, nil
}

func (r *roleRepository) UpdateRole(uid string, role dtos.RoleUpdateRequest, ctx *gin.Context) (*dtos.RoleResponse, *models.ErrorResponse) {
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	var existingRole models.Role
	if err := db.WithContext(ctx).
		Where("uid = ?", uid).
		First(&existingRole).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Role not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	existingRole.Name = role.Name
	existingRole.Rights = role.Rights

	if err := db.WithContext(ctx).Save(&existingRole).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.RoleResponse{
		UID:    existingRole.UID.String(),
		Name:   existingRole.Name,
		Rights: existingRole.Rights,
	}, nil
}

func (r *roleRepository) DeleteRole(UID string, ctx *gin.Context) *models.ErrorResponse {
	db, err := r.getDB(ctx)

	if err != nil {
		return models.InternalServerError(err.Error())
	}
	roleUID, err := uuid.Parse(UID)
	if err != nil {
		return models.InternalServerError("Invalid UUID format")
	}

	var role models.Role
	if err := db.WithContext(ctx).
		Where("uid = ?", roleUID).
		First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("Role not found")
		}
		return models.InternalServerError("Failed to fetch role: " + err.Error())
	}

	if err := db.WithContext(ctx).Model(&models.User{}).
		Where("role_id = ?", role.ID).
		Update("role_id", gorm.Expr("NULL")).Error; err != nil {
		return models.InternalServerError("Failed to dissociate users from role: " + err.Error())
	}

	if err := db.WithContext(ctx).
		Where("uid = ?", roleUID).
		Delete(&models.Role{}).Error; err != nil {
		return models.InternalServerError("Failed to delete role: " + err.Error())
	}

	return nil
}

func (r *roleRepository) GetRoleUsers(role *dtos.RoleResponse, ctx *gin.Context) ([]*dtos.UserResponse, *models.ErrorResponse) {
	var roleModel models.Role
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).First(&roleModel, "uid = ?", role.UID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Role not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	var users []*models.User
	if err := db.WithContext(ctx).Model(&roleModel).Association("Users").Find(&users); err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	result := make([]*dtos.UserResponse, len(users))
	for i, user := range users {
		result[i] = &dtos.UserResponse{
			UID:    user.UID.String(),
			Name:   user.Name,
			Email:  user.Email,
			Status: user.Status,
		}
	}

	return result, nil
}

func (r *roleRepository) GetRoleByNameAndRights(rol dtos.RoleCreateRequest, ctx *gin.Context) (*models.Role, *models.ErrorResponse) {
	var role models.Role
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	rightsJSON, err := json.Marshal(rol.Rights)
	if err != nil {
		return nil, models.InternalServerError("Error marshalling rights: " + err.Error())
	}

	rightsStr := string(rightsJSON)

	if err := db.WithContext(ctx).
		Where("name = ? AND rights::jsonb = ?", rol.Name, rightsStr).
		First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, models.InternalServerError("Error querying role: " + err.Error())
	}

	return &role, nil
}
