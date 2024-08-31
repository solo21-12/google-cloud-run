package repository

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/gorm"

	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) interfaces.RoleRepository {
	return &roleRepository{db: db}
}
func (r *roleRepository) GetAllRoles(ctx context.Context) ([]*dtos.RoleResponse, *models.ErrorResponse) {
	var roles []*models.Role
	// Preload the correct association name, which is "Users"
	if err := r.db.WithContext(ctx).Preload("Users").Find(&roles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*dtos.RoleResponse{}, nil
		}
		return nil, models.InternalServerError(err.Error())
	}

	var result []*dtos.RoleResponse

	for _, role := range roles {
		result = append(result, &dtos.RoleResponse{
			RID:    role.RID.String(),
			Name:   role.Name,
			Rights: role.Rights,
		})
	}
	return result, nil
}

func (r *roleRepository) GetRoleById(id string, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse) {
	roleId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var role models.Role
	// Preload the correct association name, which is "Users"
	if err := r.db.WithContext(ctx).Preload("Users").First(&role, roleId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Role not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.RoleResponse{
		RID:    role.RID.String(),
		Name:   role.Name,
		Rights: role.Rights,
	}, nil
}


func (r *roleRepository) CreateRole(role dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse) {

	newRole := models.Role{
		RID:    uuid.New(),
		Name:   role.Name,
		Rights: role.Rights,
	}

	if err := r.db.WithContext(ctx).Create(&newRole).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.RoleResponse{
		RID:    newRole.RID.String(),
		Name:   newRole.Name,
		Rights: newRole.Rights,
	}, nil
}

func (r *roleRepository) UpdateRole(id string, role dtos.RoleUpdateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse) {
	roleId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var existingRole models.Role
	if err := r.db.WithContext(ctx).First(&existingRole, roleId).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	existingRole.Name = role.Name
	existingRole.Rights = role.Rights

	if err := r.db.WithContext(ctx).Save(&existingRole).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.RoleResponse{
		RID:    existingRole.RID.String(),
		Name:   existingRole.Name,
		Rights: existingRole.Rights,
	}, nil
}

func (r *roleRepository) DeleteRole(id string, ctx context.Context) *models.ErrorResponse {
    roleId, err := uuid.Parse(id)
    if err != nil {
        return models.InternalServerError("Invalid UUID format")
    }

    // Explicitly set RoleID to NULL for all users associated with this role
    if err := r.db.WithContext(ctx).Model(&models.User{}).Where("role_id = ?", roleId).Update("role_id", gorm.Expr("NULL")).Error; err != nil {
        return models.InternalServerError("Failed to dissociate users from role: " + err.Error())
    }

    // Now, safely delete the role
    if err := r.db.WithContext(ctx).Delete(&models.Role{}, roleId).Error; err != nil {
        return models.InternalServerError("Failed to delete role: " + err.Error())
    }

    return nil
}



func (r *roleRepository) GetRoleUsers(role *dtos.RoleResponse, ctx context.Context) ([]*dtos.UserResponse, *models.ErrorResponse) {
	var roleModel models.Role

	if err := r.db.WithContext(ctx).First(&roleModel, "r_id = ?", role.RID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Role not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	var users []*models.User
	if err := r.db.WithContext(ctx).Model(&roleModel).Association("Users").Find(&users); err != nil {
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

func (r *roleRepository) GetRoleByNameAndRights(rol dtos.RoleCreateRequest, ctx context.Context) (*models.Role, *models.ErrorResponse) {
	var role models.Role

	rightsJSON, err := json.Marshal(rol.Rights)
	if err != nil {
		return nil, models.InternalServerError("Error marshalling rights: " + err.Error())
	}

	rightsStr := string(rightsJSON)

	if err := r.db.WithContext(ctx).
		Where("name = ? AND rights::jsonb = ?", rol.Name, rightsStr).
		First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, models.InternalServerError("Error querying role: " + err.Error())
	}

	return &role, nil
}
